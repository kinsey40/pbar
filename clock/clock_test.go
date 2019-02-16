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
 * File:   clock_test.go
 * Author: kinsey40
 *
 * Created on 13 January 2019, 11:05
 *
 *
 */

package clock_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/kinsey40/pbar/clock"
	"github.com/stretchr/testify/assert"
)

func TestNewClock(t *testing.T) {
	c := clock.NewClock()

	message := fmt.Sprintf("NewClock type does not match expected: %v; got: %v", (*clock.Clock)(nil), c)
	assert.Implements(t, (*clock.Clock)(nil), c, message)
}

func TestSubtract(t *testing.T) {
	c := clock.NewClock()
	testCases := []struct {
		startTime time.Time
		inputTime time.Time
		duration  string
	}{
		{time.Unix(10, 0), time.Unix(20, 0), "10s"},
	}

	for _, testCase := range testCases {
		c.SetStart(testCase.startTime)

		expectedDuration, err := time.ParseDuration(testCase.duration)
		if err != nil {
			t.Errorf("Error parsing duration: %v", err)
		}

		output := c.Subtract(testCase.inputTime)
		message := fmt.Sprintf("Subtract output unequal expected: %v; got: %v", expectedDuration, output)

		assert.Equal(t, expectedDuration, output, message)
	}
}

func TestStart(t *testing.T) {
	c := clock.NewClock()
	testCases := []struct {
		timeVal time.Time
	}{
		{time.Now()},
	}

	for _, testCase := range testCases {
		c.SetStart(testCase.timeVal)
		message := fmt.Sprintf("Start Time unequal expected: %v; got: %v", testCase.timeVal, c.Start())
		assert.Equal(t, testCase.timeVal, c.Start(), message)
	}
}

func TestSeconds(t *testing.T) {
	c := clock.NewClock()
	testCases := []struct {
		input          time.Duration
		expectedOutput float64
	}{
		{time.Duration(10) * time.Second, 10.0},
	}

	for _, testCase := range testCases {
		secs := c.Seconds(testCase.input)
		message := fmt.Sprintf("Seconds incorrect expected: %v; got: %v", testCase.expectedOutput, secs)

		assert.Equal(t, testCase.expectedOutput, secs, message)
	}
}

func TestRemaining(t *testing.T) {
	c := clock.NewClock()
	testCases := []struct {
		input    float64
		duration string
	}{
		{10.0, "10s"},
	}

	for _, testCase := range testCases {
		expectedDuration, err := time.ParseDuration(testCase.duration)
		if err != nil {
			t.Errorf("Error parsing duration: %v", err)
		}

		output := c.Remaining(testCase.input)
		message := fmt.Sprintf("Remaining output unequal expected: %v; got: %v", expectedDuration, output)

		assert.Equal(t, expectedDuration, output, message)
	}
}

func TestFormat(t *testing.T) {
	c := clock.NewClock()
	testCases := []struct {
		timeValue      time.Duration
		expectedString string
	}{
		{time.Duration(10) * time.Second, "00m:10s"},
		{time.Duration(10000) * time.Second, "02h:46m:40s"},
	}

	for _, testCase := range testCases {
		returnedString := c.Format(testCase.timeValue)
		message := fmt.Sprintf("Time string incorrect, expected: %v; got: %v", testCase.expectedString, returnedString)
		assert.Equal(t, testCase.expectedString, returnedString, message)
	}
}
