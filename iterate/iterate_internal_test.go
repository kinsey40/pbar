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
 * File:   iterate_test.go
 * Author: kinsey40
 *
 * Created on 13 January 2019, 11:05
 *
 * The export file used for testing.
 *
 */

package iterate

import (
	"fmt"
	"testing"
)

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
	}

	for _, testCase := range testCases {
		returnedValue := isConvertibleToFloat(testCase.value)
		if returnedValue != testCase.expectedResult {
			t.Errorf(
				fmt.Sprintf(
					"Conversion value not correct returned: %v, expected: %v for: %v",
					returnedValue,
					testCase.expectedResult,
					testCase.value))
		}
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
		if floatValue != testCase.output {
			t.Errorf(fmt.Sprintf("Incorrect float conversion: floatValue: %v, expected: %v", floatValue, testCase.output))
		}
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
		if returnedValue != testCase.expectedResult {
			t.Errorf(fmt.Sprintf("Is Valid Object incorrect returned: %v, expected: %v", returnedValue, testCase.expectedResult))
		}
	}
}

func TestObjectOrNumber(t *testing.T) {
	testCases := []struct {
		values         []interface{}
		expectedResult bool
		errorRaised    bool
	}{
		{[]interface{}{float64(1), float64(1), float64(1)}, true, false},
		{[]interface{}{float32(1), float32(1), float32(1)}, true, false},
		{[]interface{}{int8(1), int8(1), int8(1)}, true, false},
		{[]interface{}{int16(1), int16(1), int16(1)}, true, false},
		{[]interface{}{int32(1), int32(1), int32(1)}, true, false},
		{[]interface{}{int64(1), int64(1), int64(1)}, true, false},
		{[]interface{}{uint8(1), uint8(1), uint8(1)}, true, false},
		{[]interface{}{uint16(1), uint16(1), uint16(1)}, true, false},
		{[]interface{}{make(map[int]string, 1)}, false, false},
		{[]interface{}{string("Hello")}, false, false},
		{[]interface{}{make(chan int, 1)}, false, false},
		{[]interface{}{[1]int{1}}, false, false},
		{[]interface{}{[]int{}}, false, false},
		{[]interface{}{float64(1), []int{}}, true, true},
		{[]interface{}{[]int{}, float64(1)}, false, true},
		{[]interface{}{complex128(1)}, false, true},
	}

	for _, testCase := range testCases {
		isNumber, err := objectOrNumber(testCase.values...)
		if isNumber != testCase.expectedResult {
			t.Errorf(
				fmt.Sprintf(
					"Incorrect isNumber; expected: %v; got: %v for values: %v",
					testCase.expectedResult,
					isNumber,
					testCase.values))
		}

		if err != nil && !testCase.errorRaised {
			t.Errorf(fmt.Sprintf("Error (%v) incorrectly raised", err))
		}

		if err == nil && testCase.errorRaised {
			t.Errorf(fmt.Sprintf("Expected error not raised"))
		}
	}
}

func TestCheckSize(t *testing.T) {
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
	}

	for _, testCase := range testCases {
		err := checkSize(testCase.isObject, testCase.values...)
		if err != nil && !testCase.errorRaised {
			t.Errorf(
				fmt.Sprintf(
					"Unexpected error raised, for values: %v; with isObject: %v",
					testCase.values,
					testCase.isObject))
		}

		if err == nil && testCase.errorRaised {
			t.Errorf(
				fmt.Sprintf(
					"Expected error not raised, for values: %v; with isObject: %v",
					testCase.values,
					testCase.isObject))
		}
	}
}
