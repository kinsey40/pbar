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
		{[]interface{}{float64(1), float64(2), float64(100)}, true},
		{[]interface{}{float64(1), float64(10), float64(1)}, false},
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

func TestUpdate(t *testing.T) {
	// testCases := []struct {
	// 	values []interface{}
	// 	expectError	bool
	// }

}

func TestSetGetDescription(t *testing.T) {
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
		returnedDesc := itr.GetDescription()

		assert.Equal(t,
			testCase.expectedPrefix,
			returnedDesc,
			fmt.Sprintf("Descriptions not equal; expected: %s, got: %s", testCase.expectedPrefix, returnedDesc))
	}
}

func TestSetGetFinishedIterationSymbol(t *testing.T) {
	itr := pbar.MakeIteratorObject()
	testCases := []struct {
		symbol string
	}{
		{"Hello"},
		{"World"},
	}

	for _, testCase := range testCases {
		itr.SetFinishedIterationSymbol(testCase.symbol)
		returnedSymbol := itr.GetFinishedIterationSymbol()

		assert.Equal(t,
			testCase.symbol,
			returnedSymbol,
			fmt.Sprintf("Descriptions not equal; expected: %s, got: %s", testCase.symbol, returnedSymbol))
	}
}

func TestSetGetCurrentIterationSymbol(t *testing.T) {
	itr := pbar.MakeIteratorObject()
	testCases := []struct {
		symbol string
	}{
		{"Hello"},
		{"World"},
	}

	for _, testCase := range testCases {
		itr.SetCurrentIterationSymbol(testCase.symbol)
		returnedSymbol := itr.GetCurrentIterationSymbol()

		assert.Equal(t,
			testCase.symbol,
			returnedSymbol,
			fmt.Sprintf("Descriptions not equal; expected: %s, got: %s", testCase.symbol, returnedSymbol))
	}
}

func TestSetGetRemainingIterationSymbol(t *testing.T) {
	itr := pbar.MakeIteratorObject()
	testCases := []struct {
		symbol string
	}{
		{"Hello"},
		{"World"},
	}

	for _, testCase := range testCases {
		itr.SetRemainingIterationSymbol(testCase.symbol)
		returnedSymbol := itr.GetRemainingIterationSymbol()

		assert.Equal(t,
			testCase.symbol,
			returnedSymbol,
			fmt.Sprintf("Descriptions not equal; expected: %s, got: %s", testCase.symbol, returnedSymbol))
	}
}

func TestSetGetLParenSymbol(t *testing.T) {
	itr := pbar.MakeIteratorObject()
	testCases := []struct {
		symbol string
	}{
		{"Hello"},
		{"World"},
	}

	for _, testCase := range testCases {
		itr.SetLParen(testCase.symbol)
		returnedSymbol := itr.GetLParen()

		assert.Equal(t,
			testCase.symbol,
			returnedSymbol,
			fmt.Sprintf("Descriptions not equal; expected: %s, got: %s", testCase.symbol, returnedSymbol))
	}
}

func TestSetGetRParenSymbol(t *testing.T) {
	itr := pbar.MakeIteratorObject()
	testCases := []struct {
		symbol string
	}{
		{"Hello"},
		{"World"},
	}

	for _, testCase := range testCases {
		itr.SetRParen(testCase.symbol)
		returnedSymbol := itr.GetRParen()

		assert.Equal(t,
			testCase.symbol,
			returnedSymbol,
			fmt.Sprintf("Descriptions not equal; expected: %s, got: %s", testCase.symbol, returnedSymbol))
	}
}

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// )

// func TestDescription(t *testing.T) {
// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockPbarInterface := mocks.NewMockPbarInterface(mockCtrl)
// 	testCases := []struct {
// 		description    string
// 		expectedReturn string
// 		correct        bool
// 	}{
// 		{"Testing", "Testing", true},
// 		{"Another Test", "Another Test", true},
// 		{"Test", "Incorrect return", false},
// 	}

// 	for _, testCase := range testCases {
// 		mockPbarInterface.EXPECT().SetDescription(testCase.description).Return().Times(1)
// 		mockPbarInterface.EXPECT().GetDescription().Return(testCase.description).Times(1)

// 		mockPbarInterface.SetDescription(testCase.description)
// 		desc := mockPbarInterface.GetDescription()

// 		if desc != testCase.expectedReturn && testCase.correct {
// 			t.Fail(fmt.Sprintf("Incorrect description returned: %v, expected: %v", desc, testCase.expectedReturn))
// 		}

// 		if desc == testCase.expectedReturn && !testCase.correct {
// 			t.Fail(fmt.Sprintf("Incorrect description returned: %v, expected: %v", desc, testCase.expectedReturn))
// 		}
// 	}
// }

// func TestRetain(t *testing.T) {
// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockPbarInterface := mocks.NewMockPbarInterface(mockCtrl)
// 	testCases := []struct {
// 		retain         bool
// 		expectedReturn bool
// 		correct        bool
// 	}{
// 		{true, true, true},
// 		{false, false, true},
// 		{true, false, false},
// 		{false, true, false},
// 	}

// 	for _, testCase := range testCases {
// 		mockPbarInterface.EXPECT().SetRetain(testCase.retain).Return().Times(1)
// 		mockPbarInterface.EXPECT().GetRetain().Return(testCase.retain).Times(1)

// 		mockPbarInterface.SetRetain(testCase.retain)
// 		ret := mockPbarInterface.GetRetain()

// 		if ret != testCase.expectedReturn && testCase.correct {
// 			t.Fail(fmt.Sprintf("Incorrect description returned: %v, expected: %v", ret, testCase.expectedReturn))
// 		}

// 		if ret == testCase.expectedReturn && !testCase.correct {
// 			t.Fail(fmt.Sprintf("Incorrect description returned: %v, expected: %v", ret, testCase.expectedReturn))
// 		}
// 	}
// }
