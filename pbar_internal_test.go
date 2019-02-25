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
 * File:   pbar_internal_test.go
 * Author: kinsey40
 *
 * Created on 13 January 2019, 11:05
 *
 * The internal test file for pbar.go
 *
 */

package pbar

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kinsey40/pbar/mocks"
	"github.com/kinsey40/pbar/render"
	"github.com/stretchr/testify/assert"
)

func TestMakeIteratorObject(t *testing.T) {
	itr := makeIteratorObject().(*Iterator)

	assert.NotNil(t, itr, fmt.Sprintf("Iterator is nil!"))
	assert.NotNil(t, itr.Values, fmt.Sprintf("Values is nil"))
	assert.NotNil(t, itr.Clock, fmt.Sprintf("Clock is nil"))
	assert.NotNil(t, itr.Settings, fmt.Sprintf("Settings is nil"))
	assert.NotNil(t, itr.Write, fmt.Sprintf("Write is nil"))
}

func TestIsConvertibleToFloat(t *testing.T) {
	testCases := []struct {
		value          interface{}
		expectedResult bool
	}{
		{float64(1), true},
		{float32(1), true},
		{int8(1), true},
		{int16(1), true},
		{int32(1), true},
		{int64(1), true},
		{uint8(1), true},
		{uint16(1), true},
		{int(1), true},
		{make(map[int]int), false},
		{make([]string, 1), false},
		{complex128(1), false},
	}

	for _, testCase := range testCases {
		returnedValue := isConvertibleToFloat(testCase.value)
		message := fmt.Sprintf("Conversion value not correct expected: %v, returned: %v for: %v",
			testCase.expectedResult,
			returnedValue,
			testCase.value)

		assert.Equal(t, testCase.expectedResult, returnedValue, message)
	}
}

func TestConvertToFloat(t *testing.T) {
	testCases := []struct {
		input  interface{}
		output float64
	}{
		{float64(3.0), 3.0},
		{float32(3.0), 3.0},
		{int8(3.0), 3.0},
		{int16(3.0), 3.0},
		{int32(3.0), 3.0},
		{int64(3.0), 3.0},
		{uint8(3.0), 3.0},
		{uint16(3.0), 3.0},
		{int(3.0), 3.0},
	}

	for _, testCase := range testCases {
		floatValue := convertToFloatValue(testCase.input)
		message := fmt.Sprintf("Incorrect float conversion: expected: %v, returned: %v",
			testCase.output,
			floatValue)

		assert.Equal(t, testCase.output, floatValue, message)
	}
}

func TestIsValidObject(t *testing.T) {
	testCases := []struct {
		object         interface{}
		expectedResult bool
	}{
		{make([]int, 1), true},
		{complex128(1), false},
	}

	for _, testCase := range testCases {
		returnedValue := isValidObject(testCase.object)
		message := fmt.Sprintf("Is Valid Object incorrect expected: %v, returned: %v",
			testCase.expectedResult,
			returnedValue)

		assert.Equal(t, testCase.expectedResult, returnedValue, message)
	}
}

func TestIsObject(t *testing.T) {
	testCases := []struct {
		values         []interface{}
		expectedResult bool
		errorRaised    bool
	}{
		{[]interface{}{float64(1), float64(1), float64(1)}, false, false},
		{[]interface{}{float32(1), float32(1), float32(1)}, false, false},
		{[]interface{}{int8(1), int8(1), int8(1)}, false, false},
		{[]interface{}{int16(1), int16(1), int16(1)}, false, false},
		{[]interface{}{int32(1), int32(1), int32(1)}, false, false},
		{[]interface{}{int64(1), int64(1), int64(1)}, false, false},
		{[]interface{}{uint8(1), uint8(1), uint8(1)}, false, false},
		{[]interface{}{uint16(1), uint16(1), uint16(1)}, false, false},
		{[]interface{}{make(map[int]string, 1)}, true, false},
		{[]interface{}{string("Hello")}, true, false},
		{[]interface{}{make(chan int, 1)}, true, false},
		{[]interface{}{[1]int{1}}, true, false},
		{[]interface{}{[]int{}}, true, false},
		{[]interface{}{float64(1), []int{}}, false, true},
		{[]interface{}{[]int{}, float64(1)}, true, true},
		{[]interface{}{complex128(1)}, false, true},
	}

	for _, testCase := range testCases {
		isObject, err := isObject(testCase.values...)
		isObjectMessage := fmt.Sprintf(
			"Incorrect isObject; expected: %v; returned: %v for values: %v",
			testCase.expectedResult,
			isObject,
			testCase.values)

		assert.Equal(t, testCase.expectedResult, isObject, isObjectMessage)

		if testCase.errorRaised {
			message := fmt.Sprintf("Expected error not raised")
			assert.Error(t, err, message)
		} else {
			message := fmt.Sprintf("Error (%v) incorrectly raised", err)
			assert.NoError(t, err, message)
		}
	}
}

func TestCheckValues(t *testing.T) {
	testCases := []struct {
		isObject    bool
		values      []interface{}
		errorRaised bool
	}{
		{true, []interface{}{make(map[int]int, 1)}, false},
		{true, []interface{}{make(map[int]int, 1), make(map[int]int, 1)}, true},
		{true, []interface{}{}, true},
		{false, []interface{}{float64(1)}, false},
		{false, []interface{}{float64(1), float64(1)}, false},
		{false, []interface{}{float64(1), float64(1), float64(1)}, false},
		{false, []interface{}{}, true},
		{false, []interface{}{float64(1), float64(1), float64(1), float64(1)}, true},
		{false, []interface{}{float64(2), float64(1), float64(1)}, true},
	}

	for _, testCase := range testCases {
		err := checkValues(testCase.isObject, testCase.values...)
		if testCase.errorRaised {
			message := fmt.Sprintf(
				"Expected error not raised, for values: %v; with isObject: %v",
				testCase.values,
				testCase.isObject)

			assert.Error(t, err, message)
		} else {
			message := fmt.Sprintf(
				"Unexpected error raised, for values: %v; with isObject: %v",
				testCase.values,
				testCase.isObject)

			assert.NoError(t, err, message)
		}
	}
}

func TestProgress(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClock := mocks.NewMockClock(mockCtrl)
	mockSettings := mocks.NewMockSettings(mockCtrl)
	mockWrite := mocks.NewMockWrite(mockCtrl)
	mockValues := mocks.NewMockValues(mockCtrl)
	testCases := []struct {
		startVal    float64
		stopVal     float64
		stepVal     float64
		currentVal  float64
		lineSize    int
		numSteps    int
		barString   string
		stats       string
		speedMeter  string
		expectError bool
		writeError  error
	}{
		{0.0, 5.0, 1.0, 1.0, 10, 2, "|##--------|", "1.0/5.0 20.0%", "[elapsed: 00m:05s, left: 00m:20s, 0.20 iters/sec]", false, nil},
		{2.0, 5.0, 1.0, 1.0, 10, 2, "", "", "", true, nil},
		{0.0, 5.0, 1.0, 1.0, 10, 2, "|##--------|", "1.0/5.0 20.0%", "[elapsed: 00m:05s, left: 00m:20s, 0.20 iters/sec]", true, errors.New("An error")},
	}

	for _, testCase := range testCases {
		calls := []*gomock.Call{
			mockValues.EXPECT().GetStart().Return(testCase.startVal),
			mockValues.EXPECT().GetStop().Return(testCase.stopVal),
			mockValues.EXPECT().GetStep().Return(testCase.stepVal),
			mockValues.EXPECT().GetCurrent().Return(testCase.currentVal),
			mockSettings.EXPECT().GetLineSize().Return(testCase.lineSize),
		}

		if testCase.currentVal > testCase.startVal && testCase.currentVal < testCase.stopVal {
			newCalls := []*gomock.Call{
				mockValues.EXPECT().Statistics(testCase.lineSize).Return(testCase.stats, testCase.numSteps),
				mockSettings.EXPECT().CreateBarString(testCase.numSteps).Return(testCase.barString),
				mockClock.EXPECT().CreateSpeedMeter(testCase.startVal, testCase.stopVal, testCase.currentVal).Return(testCase.speedMeter),
				mockWrite.EXPECT().WriteString(gomock.Any()).Return(testCase.writeError),
			}

			if testCase.writeError == nil {
				newCalls = append(newCalls, mockValues.EXPECT().SetCurrent(gomock.Any()))
			}

			calls = append(calls, newCalls...)
		}

		gomock.InOrder(calls...)

		itr := &Iterator{
			Values:   mockValues,
			Settings: mockSettings,
			Clock:    mockClock,
			Write:    mockWrite,
		}

		err := itr.progress()
		if testCase.expectError {
			assert.Error(t, err, fmt.Sprintf("Expected error not raised!"))
		} else {
			assert.NoError(t, err, fmt.Sprintf("Unexpected error raised!"))
		}

	}
}

func TestCreateIteratorFromValues(t *testing.T) {
	testCases := []struct {
		values          []interface{}
		expectedStart   float64
		expectedStop    float64
		expectedStep    float64
		expectedCurrent float64
	}{
		{[]interface{}{float64(1.0), float64(5.0), float64(1.0)}, 1.0, 5.0, 1.0, 1.0},
		{[]interface{}{float64(1.0)}, 0.0, 1.0, 1.0, 0.0},
		{[]interface{}{float64(1.0), float64(5.0)}, 1.0, 5.0, 1.0, 1.0},
	}

	for _, testCase := range testCases {
		itr := new(Iterator)
		itr.Values = &render.Vals{}
		itr.createIteratorFromValues(testCase.values...)

		assert.Equal(
			t,
			testCase.expectedStart,
			itr.Values.GetStart(),
			fmt.Sprintf("Start Value incorrect expected: %f; got: %f", testCase.expectedStart, itr.Values.GetStart()),
		)

		assert.Equal(
			t,
			testCase.expectedStop,
			itr.Values.GetStop(),
			fmt.Sprintf("Stop Value incorrect expected: %f; got: %f", testCase.expectedStop, itr.Values.GetStop()),
		)

		assert.Equal(
			t,
			testCase.expectedStep,
			itr.Values.GetStep(),
			fmt.Sprintf("Step Value incorrect expected: %f; got: %f", testCase.expectedStep, itr.Values.GetStep()),
		)

		assert.Equal(
			t,
			testCase.expectedCurrent,
			itr.Values.GetCurrent(),
			fmt.Sprintf("Current Value incorrect expected: %f; got: %f", testCase.expectedCurrent, itr.Values.GetCurrent()),
		)
	}
}

func TestCreateIteratorFromObject(t *testing.T) {
	testCases := []struct {
		object          interface{}
		expectedStart   float64
		expectedStop    float64
		expectedStep    float64
		expectedCurrent float64
	}{
		{[...]float64{1.0, 2.0, 3.0}, 0.0, 3.0, 1.0, 0.0},
	}

	for _, testCase := range testCases {
		itr := new(Iterator)
		itr.Values = &render.Vals{}
		itr.createIteratorFromObject(testCase.object)

		assert.Equal(
			t,
			testCase.expectedStart,
			itr.Values.GetStart(),
			fmt.Sprintf("Start Value incorrect expected: %f; got: %f", testCase.expectedStart, itr.Values.GetStart()),
		)

		assert.Equal(
			t,
			testCase.expectedStop,
			itr.Values.GetStop(),
			fmt.Sprintf("Stop Value incorrect expected: %f; got: %f", testCase.expectedStop, itr.Values.GetStop()),
		)

		assert.Equal(
			t,
			testCase.expectedStep,
			itr.Values.GetStep(),
			fmt.Sprintf("Step Value incorrect expected: %f; got: %f", testCase.expectedStep, itr.Values.GetStep()),
		)

		assert.Equal(
			t,
			testCase.expectedCurrent,
			itr.Values.GetCurrent(),
			fmt.Sprintf("Current Value incorrect expected: %f; got: %f", testCase.expectedCurrent, itr.Values.GetCurrent()),
		)
	}
}

func TestFormatProgressBar(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClock := mocks.NewMockClock(mockCtrl)
	mockSettings := mocks.NewMockSettings(mockCtrl)
	mockValues := mocks.NewMockValues(mockCtrl)
	testCases := []struct {
		startVal       float64
		currentVal     float64
		endVal         float64
		lineSize       int
		numSteps       int
		barString      string
		stats          string
		speedMeter     string
		expectedOutput string
	}{
		{0.0, 1.0, 5.0, 10, 1, "|##--------|", "1.0/5.0 20.0%", "[elapsed: 00m:05s, left: 00m:20s, 0.20 iters/sec]", "|##--------| 1.0/5.0 20.0% [elapsed: 00m:05s, left: 00m:20s, 0.20 iters/sec]"},
	}

	for _, testCase := range testCases {
		itr := &Iterator{
			Values:   mockValues,
			Settings: mockSettings,
			Clock:    mockClock,
		}

		gomock.InOrder(
			mockValues.EXPECT().Statistics(testCase.lineSize).Return(testCase.stats, testCase.numSteps),
			mockSettings.EXPECT().CreateBarString(testCase.numSteps).Return(testCase.barString),
			mockClock.EXPECT().CreateSpeedMeter(testCase.startVal, testCase.endVal, testCase.currentVal).Return(testCase.speedMeter),
		)

		output := itr.formatProgressBar(testCase.startVal, testCase.endVal, testCase.currentVal, testCase.lineSize)
		message := fmt.Sprintf(
			"Progress bar incorrect, expected: %s; got %s",
			testCase.expectedOutput,
			output,
		)

		assert.Equal(t, testCase.expectedOutput, output, message)
	}
}

func TestRenderError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockWrite := mocks.NewMockWrite(mockCtrl)
	testCases := []struct {
		input         string
		expectedError bool
	}{
		{"Hello", false},
		{"Hello", true},
	}

	for _, testCase := range testCases {
		itr := &Iterator{
			Write: mockWrite,
		}

		if testCase.expectedError {
			mockWrite.EXPECT().WriteString(gomock.Any()).Return(errors.New("An error"))
			err := itr.render(testCase.input)
			message := fmt.Sprintf("")

			assert.Error(t, err, message)
		} else {
			mockWrite.EXPECT().WriteString(gomock.Any()).Return(nil)
			err := itr.render(testCase.input)
			message := fmt.Sprintf("")

			assert.NoError(t, err, message)
		}
	}
}

func TestRender(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput string
		buffer         *bytes.Buffer
	}{
		{"Hello", "\rHello", new(bytes.Buffer)},
	}

	for _, testCase := range testCases {
		w := render.NewWrite()
		w.SetWriter(testCase.buffer)
		itr := &Iterator{
			Write: w,
		}

		itr.render(testCase.input)
		got := testCase.buffer.String()

		message := fmt.Sprintf("Input string incorrect expected: %s; got: %s", testCase.expectedOutput, got)
		assert.Equal(t, testCase.expectedOutput, got, message)
	}
}
