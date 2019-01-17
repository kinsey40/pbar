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
 * The test file for the Pbar package.
 *
 */

package pbar_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestDescription(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockPbarInterface := mocks.NewMockPbarInterface(mockCtrl)
	testCases := []struct {
		description    string
		expectedReturn string
		correct        bool
	}{
		{"Testing", "Testing", true},
		{"Another Test", "Another Test", true},
		{"Test", "Incorrect return", false},
	}

	for _, testCase := range testCases {
		mockPbarInterface.EXPECT().SetDescription(testCase.description).Return().Times(1)
		mockPbarInterface.EXPECT().GetDescription().Return(testCase.description).Times(1)

		mockPbarInterface.SetDescription(testCase.description)
		desc := mockPbarInterface.GetDescription()

		if desc != testCase.expectedReturn && testCase.correct {
			t.Fail(fmt.Sprintf("Incorrect description returned: %v, expected: %v", desc, testCase.expectedReturn))
		}

		if desc == testCase.expectedReturn && !testCase.correct {
			t.Fail(fmt.Sprintf("Incorrect description returned: %v, expected: %v", desc, testCase.expectedReturn))
		}
	}
}

func TestRetain(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockPbarInterface := mocks.NewMockPbarInterface(mockCtrl)
	testCases := []struct {
		retain         bool
		expectedReturn bool
		correct        bool
	}{
		{true, true, true},
		{false, false, true},
		{true, false, false},
		{false, true, false},
	}

	for _, testCase := range testCases {
		mockPbarInterface.EXPECT().SetRetain(testCase.retain).Return().Times(1)
		mockPbarInterface.EXPECT().GetRetain().Return(testCase.retain).Times(1)

		mockPbarInterface.SetRetain(testCase.retain)
		ret := mockPbarInterface.GetRetain()

		if ret != testCase.expectedReturn && testCase.correct {
			t.Fail(fmt.Sprintf("Incorrect description returned: %v, expected: %v", ret, testCase.expectedReturn))
		}

		if ret == testCase.expectedReturn && !testCase.correct {
			t.Fail(fmt.Sprintf("Incorrect description returned: %v, expected: %v", ret, testCase.expectedReturn))
		}
	}
}
