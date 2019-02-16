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
 * File:   render_test.go
 * Author: kinsey40
 *
 * Created on 13 January 2019, 11:05
 *
 * The test file for the render package.
 *
 */

package render_test

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/kinsey40/pbar/mocks"
	"github.com/kinsey40/pbar/render"
	"github.com/stretchr/testify/assert"
)

func TestMakeRenderObject(t *testing.T) {
	testCases := []struct {
		start float64
		stop  float64
		step  float64
	}{
		{0.0, 0.0, 0.0},
	}

	for _, testCase := range testCases {
		r := render.MakeRenderObject(testCase.start, testCase.stop, testCase.step)
		assert.Equal(
			t,
			testCase.start,
			r.StartValue,
			fmt.Sprintf("Start values unequal expected: %v; got %v", testCase.start, r.StartValue),
		)

		assert.Equal(
			t,
			testCase.stop,
			r.EndValue,
			fmt.Sprintf("End values unqueal expected: %v; got: %v", testCase.stop, r.EndValue),
		)

		assert.Equal(
			t,
			testCase.step,
			r.StepValue,
			fmt.Sprintf("Step values unqueal expected: %v; got: %v", testCase.step, r.StepValue),
		)

		assert.Equal(
			t,
			testCase.start,
			r.CurrentValue,
			fmt.Sprintf("Current values unqueal expected: %v; got: %v", testCase.start, r.CurrentValue),
		)

		assert.Equal(
			t,
			render.DefaultFinishedIterationSymbol,
			r.FinishedIterationSymbol,
			fmt.Sprintf("Finished Iteration Symbol unqueal expected: %v; got: %v", render.DefaultFinishedIterationSymbol, r.FinishedIterationSymbol),
		)

		assert.Equal(
			t,
			render.DefaultCurrentIterationSymbol,
			r.CurrentIterationSymbol,
			fmt.Sprintf("Current Iteration Symbol unqueal expected: %v; got: %v", render.DefaultCurrentIterationSymbol, r.CurrentIterationSymbol),
		)

		assert.Equal(
			t,
			render.DefaultRemainingIterationSymbol,
			r.RemainingIterationSymbol,
			fmt.Sprintf("Remaining Iteration Symbol unqueal expected: %v; got: %v", render.DefaultRemainingIterationSymbol, r.RemainingIterationSymbol),
		)

		assert.Equal(
			t,
			render.DefaultLParen,
			r.LParen,
			fmt.Sprintf("LParen Symbol unequal expected: %v; got: %v", render.DefaultLParen, r.LParen),
		)

		assert.Equal(
			t,
			render.DefaultRParen,
			r.RParen,
			fmt.Sprintf("RParen Symbol unqueal expected: %v; got: %v", render.DefaultRParen, r.RParen),
		)

		assert.Equal(
			t,
			render.DefaultMaxLineSize,
			r.MaxLineSize,
			fmt.Sprintf("Max Line Size unqueal expected: %v; got: %v", render.DefaultMaxLineSize, r.MaxLineSize),
		)

		assert.Equal(
			t,
			render.DefaultLineSize,
			r.LineSize,
			fmt.Sprintf("Line Size unqueal expected: %v; got: %v", render.DefaultLineSize, r.LineSize),
		)

		assert.Zero(
			t,
			r.Description,
			fmt.Sprintf("End values unqueal expected: %v; got: %v", "", r.Description),
		)
	}
}

func TestUpdate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClock := mocks.NewMockClock(mockCtrl)
	testCases := []struct {
		startVal        float64
		stopVal         float64
		stepVal         float64
		currentVal      float64
		elapsed         string
		remaining       string
		formatElapsed   string
		formatRemaining string
		buffer          *bytes.Buffer
		expectError     bool
		expectedOutput  string
	}{
		{0.0, 5.0, 1.0, 5.0, "5s", "0s", "00m:05s", "00m:00s", new(bytes.Buffer), false, "\r |##########| 5.0/5.0 100.0% [elapsed: 00m:05s, left: 00m:00s, 1.00 iters/sec]\n"},
		{0.0, 5.0, 1.0, 1.0, "5s", "20s", "00m:05s", "00m:20s", new(bytes.Buffer), false, "\r |##--------| 1.0/5.0 20.0% [elapsed: 00m:05s, left: 00m:20s, 0.20 iters/sec]"},
		{1.0, 5.0, 1.0, 0.0, "", "", "", "", new(bytes.Buffer), true, ""},
	}

	for _, testCase := range testCases {
		r := render.MakeRenderObject(testCase.startVal, testCase.stopVal, testCase.stepVal)
		r.W = testCase.buffer
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

		err := r.Update(testCase.currentVal)
		got := testCase.buffer.String()

		if testCase.expectError {
			assert.Error(t, err, fmt.Sprintf("Expected Error not raised"))
		} else {
			assert.NoError(t, err, fmt.Sprintf("Unexpected Error raised"))
		}

		assert.Equal(
			t,
			testCase.expectedOutput,
			got,
			fmt.Sprintf("String not equal, expected: %s; got: %s", testCase.expectedOutput, got),
		)
	}
}

func TestInitialize(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClock := mocks.NewMockClock(mockCtrl)
	testCases := []struct {
		startVal float64
		endVal   float64
		stepVal  float64
	}{
		{0.0, 5.0, 1.0},
	}

	for _, testCase := range testCases {
		r := render.MakeRenderObject(testCase.startVal, testCase.endVal, testCase.stepVal)
		r.Initialize(mockClock)

		message := fmt.Sprintf(
			"Clocks not equal expected: %v; got %v",
			mockClock,
			r.Clock,
		)
		assert.Equal(t, mockClock, r.Clock, message)
	}
}
