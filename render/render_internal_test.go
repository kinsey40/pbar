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
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/kinsey40/pbar/mocks"
	"github.com/stretchr/testify/assert"
)

func TestRender(t *testing.T) {
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
	}

	for _, testCase := range testCases {
		r := MakeRenderObject(testCase.startVal, testCase.endVal, testCase.stepVal)
		r.W = testCase.buffer
		err := r.render(testCase.inputString)

		got := testCase.buffer.String()

		if testCase.expectError {
			assert.Error(t, err, fmt.Sprintf("Expected error not raised!"))
		} else {
			assert.NoError(t, err, fmt.Sprintf("Unexpected error raised: %v", err))
		}

		assert.Equal(
			t,
			testCase.expectedOutput,
			got,
			fmt.Sprintf("Outputted string incorrect expected: %s; got %s", testCase.expectedOutput, got),
		)
	}
}

func TestFormatProgressBar(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClock := mocks.NewMockClock(mockCtrl)
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
		r := MakeRenderObject(testCase.startVal, testCase.endVal, testCase.stepVal)
		r.Initialize(mockClock)
		r.CurrentValue = testCase.currentVal

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
		r := MakeRenderObject(testCase.startVal, testCase.endVal, testCase.stepVal)
		r.CurrentValue = testCase.currentVal
		r.Initialize(mockClock)

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
		r := MakeRenderObject(0.0, 0.0, 0.0)
		r.LineSize = testCase.lineSize
		r.FinishedIterationSymbol = testCase.finSymbol
		r.CurrentIterationSymbol = testCase.currSymbol
		r.RemainingIterationSymbol = testCase.remSymbol

		barString := r.getBarString(testCase.numStepsComplete)

		assert.Equal(t,
			testCase.expectedString,
			barString,
			fmt.Sprintf("Strings not equal, expected: %s; got: %s", testCase.expectedString, barString))
	}
}

func TestGetStatistics(t *testing.T) {
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
		r := MakeRenderObject(testCase.start, testCase.stop, testCase.step)
		r.CurrentValue = testCase.current
		r.LineSize = testCase.lineSize

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
