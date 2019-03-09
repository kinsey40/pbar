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
 * The test file for clock.go
 *
 */

package render_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/kinsey40/pbar/render"
	"github.com/stretchr/testify/assert"
)

func TestNewClock(t *testing.T) {
	c := render.NewClock()

	message := fmt.Sprintf("NewClock type does not match expected: %v; got: %v", (*render.Clock)(nil), c)
	assert.Implements(t, (*render.Clock)(nil), c, message)
}

func TestSubtract(t *testing.T) {
	testCases := []struct {
		startTime time.Time
		inputTime time.Time
		duration  string
	}{
		{time.Unix(10, 0), time.Unix(20, 0), "10s"},
	}

	for _, testCase := range testCases {
		c := &render.ClockVal{
			StartTime:   testCase.startTime,
			CurrentTime: testCase.inputTime,
		}

		expectedDuration, err := time.ParseDuration(testCase.duration)
		if err != nil {
			t.Errorf("Error parsing duration: %v", err)
		}

		output := c.Subtract()
		message := fmt.Sprintf("Subtract output unequal expected: %v; got: %v", expectedDuration, output)

		assert.Equal(t, expectedDuration, output, message)
	}
}

func TestNow(t *testing.T) {
	testCases := []struct {
		timeVal time.Time
	}{
		{time.Now()},
	}

	for _, testCase := range testCases {
		render.NowTime = func() time.Time { return testCase.timeVal }
		c := render.ClockVal{}
		c.Now()

		message := fmt.Sprintf("Current time incorrect expected: %v; got: %v", testCase.timeVal, c.CurrentTime)
		assert.Equal(t, testCase.timeVal, c.CurrentTime, message)

	}
}

func TestStart(t *testing.T) {
	testCases := []struct {
		timeVal time.Time
	}{
		{time.Now()},
	}

	for _, testCase := range testCases {
		c := &render.ClockVal{
			StartTime: testCase.timeVal,
		}

		message := fmt.Sprintf("Start Time unequal expected: %v; got: %v", testCase.timeVal, c.Start())
		assert.Equal(t, testCase.timeVal, c.Start(), message)
	}
}

func TestSetStartTime(t *testing.T) {
	testCases := []struct {
		timeVal time.Time
	}{
		{time.Now()},
	}

	for _, testCase := range testCases {
		render.NowTime = func() time.Time { return testCase.timeVal }
		c := &render.ClockVal{}
		c.SetStartTime()

		message := fmt.Sprintf("Times are not equal expected: %v; got: %v", testCase.timeVal, c.StartTime)
		assert.Equal(t, testCase.timeVal, c.StartTime, message)
	}
}

func TestSeconds(t *testing.T) {
	testCases := []struct {
		input          time.Duration
		expectedOutput float64
	}{
		{time.Duration(10) * time.Second, 10.0},
	}

	for _, testCase := range testCases {
		c := &render.ClockVal{}
		secs := c.Seconds(testCase.input)
		message := fmt.Sprintf("Seconds incorrect expected: %v; got: %v", testCase.expectedOutput, secs)

		assert.Equal(t, testCase.expectedOutput, secs, message)
	}
}

func TestRemaining(t *testing.T) {
	testCases := []struct {
		input    float64
		duration string
	}{
		{10.0, "10s"},
	}

	for _, testCase := range testCases {
		c := &render.ClockVal{}
		expectedDuration, err := time.ParseDuration(testCase.duration)
		if err != nil {
			t.Errorf("Error parsing duration: %v", err)
		}

		output := c.Remaining(testCase.input)
		message := fmt.Sprintf("Remaining output unequal expected: %v; got: %v", expectedDuration, output)

		assert.Equal(t, expectedDuration, output, message)
	}
}

func TestIsStartTimeSet(t *testing.T) {
	testCases := []struct {
		input       time.Time
		expectError bool
	}{
		{time.Time{}, true},
		{time.Now(), false},
	}

	for _, testCase := range testCases {
		c := &render.ClockVal{
			StartTime: testCase.input,
		}

		err := c.IsStartTimeSet()
		if testCase.expectError {
			assert.Error(t, err, fmt.Sprintf("Expected error not returned"))
		} else {
			assert.NoError(t, err, fmt.Sprintf("Unexpected error returned: %v", err))
		}
	}
}

func TestFormat(t *testing.T) {
	testCases := []struct {
		timeValue      time.Duration
		expectedString string
	}{
		{time.Duration(10) * time.Second, "00m:10s"},
		{time.Duration(10000) * time.Second, "02h:46m:40s"},
	}

	for _, testCase := range testCases {
		c := &render.ClockVal{}
		returnedString := c.Format(testCase.timeValue)
		message := fmt.Sprintf("Time string incorrect, expected: %v; got: %v", testCase.expectedString, returnedString)
		assert.Equal(t, testCase.expectedString, returnedString, message)
	}
}

func TestCreateSpeedMeter(t *testing.T) {
	testCases := []struct {
		start           float64
		stop            float64
		current         float64
		elapsedSecs     int64
		elapsedNanoSecs int64
		expectedOutput  string
	}{
		{0.0, 5.0, 0.0, 0, 0, "[elapsed: 00m:00s, left: N/A, N/A iters/sec]"},
		{0.0, 5.0, 1.0, 1, 0, "[elapsed: 00m:01s, left: 00m:04s, 1.00 iters/sec]"},
	}

	for _, testCase := range testCases {
		c := render.ClockVal{
			StartTime:   time.Unix(0, 0),
			CurrentTime: time.Unix(testCase.elapsedSecs, testCase.elapsedNanoSecs),
		}

		speedMeter := c.CreateSpeedMeter(testCase.start, testCase.stop, testCase.current)

		assert.Equal(
			t,
			testCase.expectedOutput,
			speedMeter,
			fmt.Sprintf("Speed Meter Incorrect expected: %v; got: %v", testCase.expectedOutput, speedMeter),
		)
	}
}
