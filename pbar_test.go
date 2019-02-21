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
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/kinsey40/pbar"
	"github.com/kinsey40/pbar/mocks"
	"github.com/kinsey40/pbar/render"
	"github.com/stretchr/testify/assert"
)

func TestInitialize(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClock := mocks.NewMockClock(mockCtrl)
	testCases := []struct {
		startVal           float64
		stopVal            float64
		stepVal            float64
		currentVal         float64
		expectedCurrentVal float64
		currentTime        time.Time
		buffer             *bytes.Buffer
		expectError        bool
	}{
		{0.0, 5.0, 1.0, 0.0, 1.0, time.Unix(10, 0), new(bytes.Buffer), false},
		{1.0, 5.0, 1.0, 0.0, 1.0, time.Unix(10, 0), new(bytes.Buffer), true},
	}

	for _, testCase := range testCases {
		gomock.InOrder(
			mockClock.EXPECT().Now().Return(testCase.currentTime),
			mockClock.EXPECT().SetStart(testCase.currentTime),
		)

		itr := &pbar.Iterator{
			Start:        testCase.startVal,
			Stop:         testCase.stopVal,
			Step:         testCase.stepVal,
			Current:      testCase.currentVal,
			Timer:        mockClock,
			RenderObject: render.MakeRenderObject(testCase.startVal, testCase.stopVal, testCase.stepVal),
		}

		itr.RenderObject.Write = render.NewWrite(testCase.buffer)
		err := itr.Initialize()

		assert.Equal(t, itr.RenderObject.Clock, mockClock, fmt.Sprintf("Iterator clock and render clock not equal!"))
		if testCase.expectError {
			assert.Error(t, err, fmt.Sprintf("Expected Error not raised"))
		} else {
			assert.NoError(t, err, fmt.Sprintf("Unexpected error raised: %v", err))
			assert.Equal(t,
				testCase.expectedCurrentVal,
				itr.Current,
				fmt.Sprintf("Itr current value expected: %v; got: %v", testCase.expectedCurrentVal, itr.Current))
		}
	}
}

func TestUpdate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClock := mocks.NewMockClock(mockCtrl)
	testCases := []struct {
		startVal        float64
		stopVal         float64
		stepVal         float64
		currentVal      float64
		buffer          *bytes.Buffer
		elapsed         string
		remaining       string
		formatElapsed   string
		formatRemaining string
		expectError     bool
		expectedOutput  string
	}{
		{0.0, 5.0, 1.0, 0.0, new(bytes.Buffer), "2s", "5s", "00m:02s", "00m:05s", false, "\r |----------| 0.0/5.0 0.0% [elapsed: 00m:00s, left: N/A, N/A iters/sec]"},
		{0.0, 5.0, 1.0, 1.0, new(bytes.Buffer), "2s", "5s", "00m:02s", "00m:05s", false, "\r |##--------| 1.0/5.0 20.0% [elapsed: 00m:02s, left: 00m:05s, 0.50 iters/sec]"},
		{0.0, 5.0, 1.0, 5.0, new(bytes.Buffer), "4s", "1s", "00m:04s", "00m:01s", true, "\r |##########| 5.0/5.0 100.0% [elapsed: 00m:04s, left: 00m:01s, 1.25 iters/sec]\n"},
		{1.0, 5.0, 1.0, 0.0, new(bytes.Buffer), "1s", "1s", "00m:01s", "00m:01s", true, ""},
		{1.0, 5.0, 1.0, 6.0, new(bytes.Buffer), "1s", "1s", "00m:01s", "00m:01s", true, ""},
	}

	for _, testCase := range testCases {
		itr := pbar.Iterator{
			Start:        testCase.startVal,
			Stop:         testCase.stopVal,
			Step:         testCase.stepVal,
			Current:      testCase.currentVal,
			RenderObject: render.MakeRenderObject(testCase.startVal, testCase.stopVal, testCase.stepVal),
		}
		itr.RenderObject.Clock = mockClock
		itr.RenderObject.Write = render.NewWrite(testCase.buffer)

		if testCase.elapsed != "" && testCase.remaining != "" {
			elapsedDur, err := time.ParseDuration(testCase.elapsed)
			if err != nil {
				t.Errorf("Error raised in parsing elapsed: %v", elapsedDur)
			}

			remainingDur, err := time.ParseDuration(testCase.remaining)
			if err != nil {
				t.Errorf("Error raised in parsing remaining: %v", remainingDur)
			}

			if testCase.currentVal > testCase.startVal && testCase.currentVal <= testCase.stopVal {
				gomock.InOrder(
					mockClock.EXPECT().Now(),
					mockClock.EXPECT().Subtract(gomock.Any()).Return(elapsedDur),
					mockClock.EXPECT().Remaining(gomock.Any()).Return(remainingDur),
					mockClock.EXPECT().Format(gomock.Any()).Return(testCase.formatElapsed),
					mockClock.EXPECT().Format(gomock.Any()).Return(testCase.formatRemaining),
				)
			}
		}

		err := itr.Update()
		got := testCase.buffer.String()

		assert.Equal(t,
			testCase.expectedOutput,
			got,
			fmt.Sprintf("Output string incorrect expected: %v; got: %v", testCase.expectedOutput, got))

		if testCase.expectError {
			assert.Error(t, err, fmt.Sprintf("Unexpected error raised: %v", err))
		} else {
			assert.NoError(t, err, fmt.Sprintf("Expected Error not raised"))
		}
	}
}

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
	itr := pbar.Iterator{RenderObject: render.MakeRenderObject(0.0, 0.0, 0.0)}
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
	itr := pbar.Iterator{RenderObject: render.MakeRenderObject(0.0, 0.0, 0.0)}
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
	itr := pbar.Iterator{RenderObject: render.MakeRenderObject(0.0, 0.0, 0.0)}
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
	itr := pbar.Iterator{RenderObject: render.MakeRenderObject(0.0, 0.0, 0.0)}
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
	itr := pbar.Iterator{RenderObject: render.MakeRenderObject(0.0, 0.0, 0.0)}
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
	itr := pbar.Iterator{RenderObject: render.MakeRenderObject(0.0, 0.0, 0.0)}
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
