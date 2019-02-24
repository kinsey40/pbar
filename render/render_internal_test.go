/*
 * The MIT License
 *
 * Copyright 2018 kinsey40.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 *
 * File:   render_internal_test.go
 * Author: kinsey40
 *
 * Created on 13 January 2019, 11:05
 *
 * The internal test file for render.go
 *
 */

package render

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/kinsey40/pbar/mocks"
	"github.com/stretchr/testify/assert"
)

func TestRender(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockWriter := mocks.NewMockWrite(mockCtrl)
	testCases := []struct {
		startVal       float64
		endVal         float64
		currentVal     float64
		stepVal        float64
		buffer         *bytes.Buffer
		inputString    string
		expectedOutput string
		expectError    bool
	}{
		{0.0, 5.0, 1.0, 1.0, new(bytes.Buffer), "Hello", "\rHello", false},
		{0.0, 5.0, 1.0, 1.0, nil, "Hello", "\rHello", true},
	}

	for _, testCase := range testCases {
		r := &RenderObject{
			StartValue: testCase.startVal,
			StopValue:  testCase.endVal,
			StepValue:  testCase.stepVal,
			Write:      mockWriter,
		}

		if testCase.expectError {
			mockWriter.EXPECT().WriteString(gomock.Any()).Return(errors.New("An error"))
			err := r.render(testCase.inputString)
			assert.Error(t, err, fmt.Sprintf("Expected error not raised!"))
		} else {
			// r.Write = NewWrite(testCase.buffer)
			err := r.render(testCase.inputString)
			got := testCase.buffer.String()

			assert.NoError(t, err, fmt.Sprintf("Unexpected error raised: %v", err))
			assert.Equal(
				t,
				testCase.expectedOutput,
				got,
				fmt.Sprintf("Outputted string incorrect expected: %s; got %s", testCase.expectedOutput, got),
			)
		}
	}
}

func TestFormatProgressBar(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClock := mocks.NewMockClock(mockCtrl)
	mockSettings := mocks.NewMockSettings(mockCtrl)
	testCases := []struct {
		startVal        float64
		currentVal      float64
		endVal          float64
		stepVal         float64
		description     string
		elapsed         string
		remaining       string
		formatElapsed   string
		formatRemaining string
		lineSize        int
		lParen          string
		rParen          string
		remIterSymbol   string
		finIterSymbol   string
		curIterSymbol   string
		expectedOutput  string
	}{
		{0.0, 1.0, 5.0, 1.0, "", "5s", "20s", "00m:05s", "00m:20s", 10, "|", "|", "-", "#", "#", " |##--------| 1.0/5.0 20.0% [elapsed: 00m:05s, left: 00m:20s, 0.20 iters/sec]"},
	}

	for _, testCase := range testCases {
		r := &RenderObject{
			StartValue:   testCase.startVal,
			StopValue:    testCase.endVal,
			StepValue:    testCase.stepVal,
			CurrentValue: testCase.currentVal,
		}

		// mockSettings.EXPECT().GetWriter().Return(new(bytes.Buffer))
		r.Initialize(mockClock, mockSettings)

		if testCase.elapsed != "" && testCase.remaining != "" {
			elapsedDur, err := time.ParseDuration(testCase.elapsed)
			if err != nil {
				t.Errorf("Error raised in parsing elapsed: %v", elapsedDur)
			}

			remainingDur, err := time.ParseDuration(testCase.remaining)
			if err != nil {
				t.Errorf("Error raised in parsing remaining: %v", remainingDur)
			}

			gomock.InOrder(
				mockSettings.EXPECT().GetLineSize().Return(DefaultLineSize),
				mockSettings.EXPECT().GetRemainingIterationSymbol().Return(DefaultRemainingIterationSymbol),
				mockSettings.EXPECT().GetCurrentIterationSymbol().Return(DefaultCurrentIterationSymbol),
				mockSettings.EXPECT().GetFinishedIterationSymbol().Return(DefaultFinishedIterationSymbol),
				mockSettings.EXPECT().GetLineSize().Return(DefaultLineSize),
				mockSettings.EXPECT().GetLParen().Return(DefaultLParen),
				mockSettings.EXPECT().GetRParen().Return(DefaultRParen),

				mockClock.EXPECT().Now(),
				mockClock.EXPECT().Subtract(gomock.Any()).Return(elapsedDur),
				mockClock.EXPECT().Remaining(gomock.Any()).Return(remainingDur),
				mockClock.EXPECT().Format(gomock.Any()).Return(testCase.formatElapsed),
				mockClock.EXPECT().Format(gomock.Any()).Return(testCase.formatRemaining),

				mockSettings.EXPECT().GetDescription().Return(DefaultDescription),
			)
		}

		pbar := r.formatProgressBar()
		message := fmt.Sprintf(
			"Progress bar incorrect, expected: %s; got %s",
			testCase.expectedOutput,
			pbar,
		)

		assert.Equal(t, testCase.expectedOutput, pbar, message)
	}
}

func TestGetSpeedMeter(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClock := mocks.NewMockClock(mockCtrl)
	mockSettings := mocks.NewMockSettings(mockCtrl)
	testCases := []struct {
		startVal        float64
		currentVal      float64
		endVal          float64
		stepVal         float64
		elapsed         string
		remaining       string
		formatElapsed   string
		formatRemaining string
		expectedOutput  string
	}{
		{0.0, 0.0, 5.0, 1.0, "", "", "00m:00s", "00m:05s", "[elapsed: 00m:00s, left: N/A, N/A iters/sec]"},
		{0.0, 1.0, 5.0, 1.0, "5s", "20s", "00m:05s", "00m:20s", "[elapsed: 00m:05s, left: 00m:20s, 0.20 iters/sec]"},
	}

	for _, testCase := range testCases {
		r := &RenderObject{
			StartValue:   testCase.startVal,
			StopValue:    testCase.endVal,
			StepValue:    testCase.stepVal,
			CurrentValue: testCase.currentVal,
		}

		// mockSettings.EXPECT().GetWriter().Return(new(bytes.Buffer))
		r.Initialize(mockClock, mockSettings)

		if testCase.elapsed != "" && testCase.remaining != "" {
			elapsedDur, err := time.ParseDuration(testCase.elapsed)
			if err != nil {
				t.Errorf("Error raised in parsing elapsed: %v", elapsedDur)
			}

			remainingDur, err := time.ParseDuration(testCase.remaining)
			if err != nil {
				t.Errorf("Error raised in parsing remaining: %v", remainingDur)
			}

			gomock.InOrder(
				mockClock.EXPECT().Now(),
				mockClock.EXPECT().Subtract(gomock.Any()).Return(elapsedDur),
				mockClock.EXPECT().Remaining(gomock.Any()).Return(remainingDur),
				mockClock.EXPECT().Format(gomock.Any()).Return(testCase.formatElapsed),
				mockClock.EXPECT().Format(gomock.Any()).Return(testCase.formatRemaining),
			)
		}

		speedMeter := r.getSpeedMeter()
		message := fmt.Sprintf(
			"Speed Meter incorrect expected: %s; got: %s",
			testCase.expectedOutput,
			speedMeter,
		)

		assert.Equal(t, testCase.expectedOutput, speedMeter, message)
	}
}

func TestGetBarString(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSettings := mocks.NewMockSettings(mockCtrl)
	testCases := []struct {
		numStepsComplete int
		lineSize         int
		finSymbol        string
		currSymbol       string
		remSymbol        string
		expectedString   string
	}{
		{0, 10, "#", "#", "-", "----------"},
		{1, 10, "#", "#", "-", "#---------"},
		{2, 10, "#", "#", "-", "##--------"},
		{10, 10, "#", "#", "-", "##########"},
	}

	for _, testCase := range testCases {
		r := &RenderObject{
			StartValue: 0.0,
			StopValue:  0.0,
			StepValue:  0.0,
			Settings:   mockSettings,
		}

		gomock.InOrder(
			mockSettings.EXPECT().GetRemainingIterationSymbol().Return(testCase.remSymbol),
			mockSettings.EXPECT().GetCurrentIterationSymbol().Return(testCase.currSymbol),
			mockSettings.EXPECT().GetFinishedIterationSymbol().Return(testCase.finSymbol),
			mockSettings.EXPECT().GetLineSize().Return(testCase.lineSize),
		)

		barString := r.getBarString(testCase.numStepsComplete)
		assert.Equal(t,
			testCase.expectedString,
			barString,
			fmt.Sprintf("Strings not equal, expected: %s; got: %s", testCase.expectedString, barString))
	}
}

func TestGetStatistics(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSettings := mocks.NewMockSettings(mockCtrl)
	testCases := []struct {
		start                     float64
		stop                      float64
		step                      float64
		current                   float64
		lineSize                  int
		expectedStatistics        string
		expectedNumStepsCompleted int
	}{
		{0.0, 5.0, 1.0, 3.0, 10, "3.0/5.0 60.0%", 6},
	}

	for _, testCase := range testCases {
		mockSettings.EXPECT().GetLineSize().Return(testCase.lineSize)
		r := &RenderObject{
			StartValue:   testCase.start,
			StopValue:    testCase.stop,
			CurrentValue: testCase.current,
			Settings:     mockSettings,
		}

		stats, numSteps := r.getStatistics()
		assert.Equal(t,
			testCase.expectedStatistics,
			stats,
			fmt.Sprintf("Stats string not equal, expected: %s; got: %s", testCase.expectedStatistics, stats))
		assert.Equal(t,
			testCase.expectedNumStepsCompleted,
			numSteps,
			fmt.Sprintf("Num steps complete not equal, expected: %d; got: %d", testCase.expectedNumStepsCompleted, numSteps))
	}
}
