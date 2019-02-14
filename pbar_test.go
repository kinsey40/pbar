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
 * File:   pbar_test.go
 * Author: kinsey40
 *
 * Created on 13 January 2019, 11:05
 *
 * The test file for pbar.go
 *
 */

package pbar_test

import (
	"fmt"
	"testing"

	"github.com/kinsey40/pbar"
	"github.com/stretchr/testify/assert"
)

func TestMakeIteratorObject(t *testing.T) {
	itr := pbar.MakeIteratorObject()

	assert.Equal(t, 0.0, itr.Start, "Start value is not 0!")
	assert.Equal(t, 0.0, itr.Stop, "Stop value is not 0!")
	assert.Equal(t, 0.0, itr.Step, "Step value is not 0!")
	assert.Equal(t, 0.0, itr.Current, "Current value is not 0!")
	assert.NotNil(t, itr.RenderObject, "Render Object is nil!")
}

func TestInitialize(t *testing.T) {
	// Requires mocking of the time module
}

// func TestUpdate(t *testing.T) {
// 	Requires mocking of the time module
// 	testCases := []struct {
// 		startVal             float64
// 		stopVal              float64
// 		stepVal              float64
// 		currentVal           float64
// 		buffer               *bytes.Buffer
// 		expectedOutput       string
// 		expectedCurrentValue float64
// 	}{
// 		{0.0, 5.0, 1.0, 1.0, new(bytes.Buffer), "", 2.0},
// 	}

// 	for _, testCase := range testCases {
// 		r := render.MakeRenderObject(testCase.startVal, testCase.stopVal, testCase.stepVal)
// 		r.W = testCase.buffer
// 		r.Initialize(time.Now())
// 		r.Update(testCase.currentVal)

// 		got := testCase.buffer.String()

// 		assert.Equal(t, testCase.expectedOutput, got, fmt.Sprintf(""))
// 		assert.Equal(t, testCase.expectedCurrentValue, r.CurrentValue, fmt.Sprintf(""))
// 	}
// }

func TestPbar(t *testing.T) {
	testCases := []struct {
		values      []interface{}
		expectError bool
	}{
		{[]interface{}{float64(1), []int{}}, true},
		{[]interface{}{[]int{}, float64(1)}, true},
		{[]interface{}{complex128(1)}, true},
		{[]interface{}{}, true},
		{[]interface{}{float64(2), float64(1), float64(1)}, true},
		{[]interface{}{float64(1), float64(2), float64(100)}, false},
		{[]interface{}{float64(1), float64(10), float64(1)}, false},
		{[]interface{}{[]int{1, 2, 3}}, false},
		{[]interface{}{"Hello!"}, false},
		{[]interface{}{map[string]int{"1": 1, "2": 2}}, false},
	}

	for _, testCase := range testCases {
		itr, err := pbar.Pbar(testCase.values...)
		if testCase.expectError {
			assert.Error(t, err, fmt.Sprintf("Expected error was not raised!"))
			assert.Nil(t, itr, fmt.Sprintf("Iterator is not nil (%v)", itr))
		} else {
			assert.NoError(t, err, fmt.Sprintf("Unexpected error(%v) was raised!", err))
			assert.NotNil(t, itr, fmt.Sprintf("Iterator is nil!"))
		}
	}
}

func TestSetDescription(t *testing.T) {
	itr := pbar.MakeIteratorObject()
	testCases := []struct {
		desc           string
		expectedPrefix string
	}{
		{"Hello", "Hello: "},
		{"World", "World: "},
	}

	for _, testCase := range testCases {
		itr.SetDescription(testCase.desc)
		returnedDesc := itr.RenderObject.Description

		assert.Equal(t,
			testCase.expectedPrefix,
			returnedDesc,
			fmt.Sprintf("Descriptions not equal; expected: %s, got: %s", testCase.expectedPrefix, returnedDesc))
	}
}

func TestSetFinishedIterationSymbol(t *testing.T) {
	itr := pbar.MakeIteratorObject()
	testCases := []struct {
		symbol string
	}{
		{"Hello"},
		{"World"},
	}

	for _, testCase := range testCases {
		itr.SetFinishedIterationSymbol(testCase.symbol)
		returnedSymbol := itr.RenderObject.FinishedIterationSymbol

		assert.Equal(t,
			testCase.symbol,
			returnedSymbol,
			fmt.Sprintf("Descriptions not equal; expected: %s, got: %s", testCase.symbol, returnedSymbol))
	}
}

func TestSetCurrentIterationSymbol(t *testing.T) {
	itr := pbar.MakeIteratorObject()
	testCases := []struct {
		symbol string
	}{
		{"Hello"},
		{"World"},
	}

	for _, testCase := range testCases {
		itr.SetCurrentIterationSymbol(testCase.symbol)
		returnedSymbol := itr.RenderObject.CurrentIterationSymbol

		assert.Equal(t,
			testCase.symbol,
			returnedSymbol,
			fmt.Sprintf("Descriptions not equal; expected: %s, got: %s", testCase.symbol, returnedSymbol))
	}
}

func TestSetRemainingIterationSymbol(t *testing.T) {
	itr := pbar.MakeIteratorObject()
	testCases := []struct {
		symbol string
	}{
		{"Hello"},
		{"World"},
	}

	for _, testCase := range testCases {
		itr.SetRemainingIterationSymbol(testCase.symbol)
		returnedSymbol := itr.RenderObject.RemainingIterationSymbol

		assert.Equal(t,
			testCase.symbol,
			returnedSymbol,
			fmt.Sprintf("Descriptions not equal; expected: %s, got: %s", testCase.symbol, returnedSymbol))
	}
}

func TestSetLParenSymbol(t *testing.T) {
	itr := pbar.MakeIteratorObject()
	testCases := []struct {
		symbol string
	}{
		{"Hello"},
		{"World"},
	}

	for _, testCase := range testCases {
		itr.SetLParen(testCase.symbol)
		returnedSymbol := itr.RenderObject.LParen

		assert.Equal(t,
			testCase.symbol,
			returnedSymbol,
			fmt.Sprintf("Descriptions not equal; expected: %s, got: %s", testCase.symbol, returnedSymbol))
	}
}

func TestSetRParenSymbol(t *testing.T) {
	itr := pbar.MakeIteratorObject()
	testCases := []struct {
		symbol string
	}{
		{"Hello"},
		{"World"},
	}

	for _, testCase := range testCases {
		itr.SetRParen(testCase.symbol)
		returnedSymbol := itr.RenderObject.RParen

		assert.Equal(t,
			testCase.symbol,
			returnedSymbol,
			fmt.Sprintf("Descriptions not equal; expected: %s, got: %s", testCase.symbol, returnedSymbol))
	}
}
