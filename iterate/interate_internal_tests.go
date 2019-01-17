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
		{new(float64), true},
		{new(float32), true},
		{new(int8), true},
		{new(int16), true},
		{new(int32), true},
		{new(int64), true},
		{new(uint8), true},
		{new(uint16), true},
		{new(int), true},
		{new(map[int]int), false},
		{new([]string), false},
	}

	for _, testCase := range testCases {
		returnedValue := isConvertibleToFloat(testCase.value)
		if returnedValue != testCase.expectedResult {
			t.Errorf(fmt.Sprintf("Conversion value not correct returned: %v, expected: %v", returnedValue, testCase.expectedResult))
		}
	}
}

func TestConvertToFloat(t *testing.T) {
	testCases := []struct {
		input  interface{}
		output float64
	}{
		{new(float64), 3.0},
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
		{new([]int), true},
		{new(complex128), false},
	}

	for _, testCase := range testCases {
		returnedValue := isValidObject(testCase.object)
		if returnedValue != testCase.expectedResult {
			t.Errorf(fmt.Sprintf("Is Valid Object incorrect returned: %v, expected: %v", returnedValue, testCase.expectedResult))
		}
	}
}

func TestObjectOrNumber(t *testing.T) {

}

func TestCheckSize(t *testing.T) {

}
