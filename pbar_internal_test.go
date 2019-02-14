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
 * The test file for pbar_interal.go
 *
 */

package pbar

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeIteratorObject(t *testing.T) {
	itr := MakeIteratorObject()

	assert.Zero(t, itr.Start, fmt.Sprintf("Start not zero value: %v", itr.Start))
	assert.Zero(t, itr.Stop, fmt.Sprintf("Stop not zero value: %v", itr.Start))
	assert.Zero(t, itr.Step, fmt.Sprintf("Step not zero value: %v", itr.Start))
	assert.Zero(t, itr.Current, fmt.Sprintf("Current not zero value: %v", itr.Start))
	assert.NotNil(t, itr, fmt.Sprintf("Iterator is nil!"))
	assert.NotNil(t, itr.RenderObject, fmt.Sprintf("renderObject is nil"))
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
	testCases := []struct {
		start            float64
		stop             float64
		step             float64
		current          float64
		expectedEndValue float64
		expectedError    bool
	}{
		{0.0, 5.0, 1.0, 0.0, 2.0, false},
		{0.0, 5.0, 1.0, 4.0, 5.0, false},
		{0.0, 5.0, 1.0, 5.0, 6.0, true},
	}

	for _, testCase := range testCases {
		itr := new(iterator)
		itr.Start = testCase.start
		itr.Stop = testCase.stop
		itr.Step = testCase.step
		itr.Current = testCase.current

		err := itr.progress()
		if testCase.expectedError {
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
		obj := createIteratorFromValues(testCase.values...)

		assert.NotNil(
			t,
			obj.RenderObject,
			fmt.Sprintf("The render object is nil!"),
		)

		assert.Equal(
			t,
			testCase.expectedStart,
			obj.Start,
			fmt.Sprintf("Start Value incorrect expected: %f; got: %f", testCase.expectedStart, obj.Start),
		)

		assert.Equal(
			t,
			testCase.expectedStop,
			obj.Stop,
			fmt.Sprintf("Stop Value incorrect expected: %f; got: %f", testCase.expectedStop, obj.Stop),
		)

		assert.Equal(
			t,
			testCase.expectedStep,
			obj.Step,
			fmt.Sprintf("Step Value incorrect expected: %f; got: %f", testCase.expectedStep, obj.Step),
		)

		assert.Equal(
			t,
			testCase.expectedCurrent,
			obj.Current,
			fmt.Sprintf("Current Value incorrect expected: %f; got: %f", testCase.expectedCurrent, obj.Current),
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
		obj := createIteratorFromObject(testCase.object)

		assert.NotNil(
			t,
			obj.RenderObject,
			fmt.Sprintf("The render object is nil!"),
		)

		assert.Equal(
			t,
			testCase.expectedStart,
			obj.Start,
			fmt.Sprintf("Start Value incorrect expected: %f; got: %f", testCase.expectedStart, obj.Start),
		)

		assert.Equal(
			t,
			testCase.expectedStop,
			obj.Stop,
			fmt.Sprintf("Stop Value incorrect expected: %f; got: %f", testCase.expectedStop, obj.Stop),
		)

		assert.Equal(
			t,
			testCase.expectedStep,
			obj.Step,
			fmt.Sprintf("Step Value incorrect expected: %f; got: %f", testCase.expectedStep, obj.Step),
		)

		assert.Equal(
			t,
			testCase.expectedCurrent,
			obj.Current,
			fmt.Sprintf("Current Value incorrect expected: %f; got: %f", testCase.expectedCurrent, obj.Current),
		)
	}
}
