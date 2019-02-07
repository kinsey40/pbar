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

package render

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRender(t *testing.T) {
	// mockCtrl := gomock.NewController(t)
	// defer mockCtrl.Finish()

	// mockRenderInterface := mocks.NewMockRenderInterface(mockCtrl)

}

func TestFormatProgressBar(t *testing.T) {
	// Testing this will require use of MOCKS!
}

func TestFormatSpeedMeter(t *testing.T) {
	// Need some way of mocking time.Now() or setting it to a set value
	// so that when it gets called, it will give an answer I already
	// know
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

func TestFormatTime(t *testing.T) {
	testCases := []struct {
		timeValue      time.Duration
		expectedString string
	}{
		{time.Duration(10) * time.Second, "00m:10s"},
	}

	for _, testCase := range testCases {
		returnedString := formatTime(testCase.timeValue)
		message := fmt.Sprintf("Time string incorrect, expected: %v; got: %v", testCase.expectedString, returnedString)
		assert.Equal(t, testCase.expectedString, returnedString, message)
	}
}
