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
 * File:   values_test.go
 * Author: kinsey40
 *
 * Created on 13 January 2019, 11:05
 *
 * The test file for values.go
 *
 */

package render_test

import (
	"fmt"
	"testing"

	"github.com/kinsey40/pbar/render"
	"github.com/stretchr/testify/assert"
)

func TestNewValues(t *testing.T) {
	v := render.NewValues()
	vals := v.(*render.Vals)

	message := fmt.Sprintf("NewValues type does not match expected: %v; got: %v", (*render.Values)(nil), v)
	assert.Implements(t, (*render.Values)(nil), v, message)
	assert.Equal(t, 0.0, vals.Start, fmt.Sprintf("Start is not 0: expected: %v; got %v", 0.0, vals.Start))
	assert.Equal(t, 0.0, vals.Stop, fmt.Sprintf("Stop is not 0: expected: %v; got %v", 0.0, vals.Stop))
	assert.Equal(t, 0.0, vals.Step, fmt.Sprintf("Step is not 0: expected: %v; got %v", 0.0, vals.Step))
	assert.Equal(t, 0.0, vals.Current, fmt.Sprintf("Current is not 0: expected: %v; got %v", 0.0, vals.Current))
}

func TestSetStart(t *testing.T) {
	testCases := []struct {
		input float64
	}{
		{5.0},
	}

	for _, testCase := range testCases {
		v := &render.Vals{}
		v.SetStart(testCase.input)
		message := fmt.Sprintf("Set Start incorrect expected: %v; got: %v", testCase.input, v.Start)

		assert.Equal(t, testCase.input, v.Start, message)
	}
}

func TestSetStop(t *testing.T) {
	testCases := []struct {
		input float64
	}{
		{5.0},
	}

	for _, testCase := range testCases {
		v := &render.Vals{}
		v.SetStop(testCase.input)
		message := fmt.Sprintf("Set Stop incorrect expected: %v; got: %v", testCase.input, v.Stop)

		assert.Equal(t, testCase.input, v.Stop, message)
	}
}

func TestSetStep(t *testing.T) {
	testCases := []struct {
		input float64
	}{
		{5.0},
	}

	for _, testCase := range testCases {
		v := &render.Vals{}
		v.SetStep(testCase.input)
		message := fmt.Sprintf("Set Step incorrect expected: %v; got: %v", testCase.input, v.Step)

		assert.Equal(t, testCase.input, v.Step, message)
	}
}

func TestSetCurrent(t *testing.T) {
	testCases := []struct {
		input float64
	}{
		{5.0},
	}

	for _, testCase := range testCases {
		v := &render.Vals{}
		v.SetCurrent(testCase.input)
		message := fmt.Sprintf("Set Current incorrect expected: %v; got: %v", testCase.input, v.Current)

		assert.Equal(t, testCase.input, v.Current, message)
	}
}

func TestSetIsObject(t *testing.T) {
	testCases := []struct {
		input bool
	}{
		{false},
		{true},
	}

	for _, testCase := range testCases {
		v := &render.Vals{}
		v.SetIsObject(testCase.input)
		message := fmt.Sprintf("Set IsObject incorrect expected: %v; got: %v", testCase.input, v.IsObject)

		assert.Equal(t, testCase.input, v.IsObject, message)
	}
}

func TestGetStart(t *testing.T) {
	testCases := []struct {
		input float64
	}{
		{5.0},
	}

	for _, testCase := range testCases {
		v := &render.Vals{
			Start: testCase.input,
		}
		output := v.GetStart()
		message := fmt.Sprintf("Get Start incorrect expected: %v; got: %v", testCase.input, output)

		assert.Equal(t, testCase.input, output, message)
	}
}

func TestGetStop(t *testing.T) {
	testCases := []struct {
		input float64
	}{
		{5.0},
	}

	for _, testCase := range testCases {
		v := &render.Vals{
			Stop: testCase.input,
		}
		output := v.GetStop()
		message := fmt.Sprintf("Get Stop incorrect expected: %v; got: %v", testCase.input, output)

		assert.Equal(t, testCase.input, output, message)
	}
}

func TestGetStep(t *testing.T) {
	testCases := []struct {
		input float64
	}{
		{5.0},
	}

	for _, testCase := range testCases {
		v := &render.Vals{
			Step: testCase.input,
		}
		output := v.GetStep()
		message := fmt.Sprintf("Get Step incorrect expected: %v; got: %v", testCase.input, output)

		assert.Equal(t, testCase.input, output, message)
	}
}

func TestGetCurrent(t *testing.T) {
	testCases := []struct {
		input float64
	}{
		{5.0},
	}

	for _, testCase := range testCases {
		v := &render.Vals{
			Current: testCase.input,
		}
		output := v.GetCurrent()
		message := fmt.Sprintf("Get Current incorrect expected: %v; got: %v", testCase.input, output)

		assert.Equal(t, testCase.input, output, message)
	}
}

func TestGetIsObject(t *testing.T) {
	testCases := []struct {
		input bool
	}{
		{false},
		{true},
	}

	for _, testCase := range testCases {
		v := &render.Vals{
			IsObject: testCase.input,
		}
		output := v.GetIsObject()
		message := fmt.Sprintf("Get IsObject incorrect expected: %v; got: %v", testCase.input, output)

		assert.Equal(t, testCase.input, output, message)
	}
}

func TestStatistics(t *testing.T) {
	testCases := []struct {
		linesize                  int
		current                   float64
		stop                      float64
		expectedNumStepsCompleted int
		expectedStats             string
	}{
		{10, 1.0, 5.0, 2, "1.0/5.0 20.0%"},
	}

	for _, testCase := range testCases {
		v := &render.Vals{
			Stop:    testCase.stop,
			Current: testCase.current,
		}

		stats, numStepsCompleted := v.Statistics(testCase.linesize)

		statsMessage := fmt.Sprintf("Stats incorrect expected: %v; got: %v", testCase.expectedStats, stats)
		numStepsCompletedMessage := fmt.Sprintf("NumStepsCompleted incorrect expected: %v; got: %v", testCase.expectedNumStepsCompleted, numStepsCompleted)

		assert.Equal(t, testCase.expectedStats, stats, statsMessage)
		assert.Equal(t, testCase.expectedNumStepsCompleted, numStepsCompleted, numStepsCompletedMessage)
	}
}
