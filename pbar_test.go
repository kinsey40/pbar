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

	"github.com/kinsey40/pbar"
	"github.com/kinsey40/pbar/render"
	"github.com/stretchr/testify/assert"
)

func TestInitialize(t *testing.T) {
	testCases := []struct {
		startVal                 float64
		stopVal                  float64
		stepVal                  float64
		currentVal               float64
		description              string
		finishedIterationSymbol  string
		currentIterationSymbol   string
		remainingIterationSymbol string
		lineSize                 int
		maxLineSize              int
		lParen                   string
		rParen                   string
		buffer                   *bytes.Buffer
		expectError              bool
		expectedOutput           string
	}{
		{0.0, 5.0, 1.0, 0.0, "", "#", "#", "-", 10, 80, "|", "|", new(bytes.Buffer), false, "\r|----------| 0.0/5.0 0.0% [elapsed: 00m:00s, left: N/A, N/A iters/sec]"},
		{1.0, 5.0, 1.0, 0.0, "", "#", "#", "-", 10, 80, "|", "|", new(bytes.Buffer), true, ""},
	}

	for _, testCase := range testCases {
		v := &render.Vals{
			Start:   testCase.startVal,
			Stop:    testCase.stopVal,
			Step:    testCase.stepVal,
			Current: testCase.currentVal,
		}

		w := &render.Writing{
			W: testCase.buffer,
		}

		s := &render.Set{
			Description:              testCase.description,
			FinishedIterationSymbol:  testCase.finishedIterationSymbol,
			CurrentIterationSymbol:   testCase.currentIterationSymbol,
			RemainingIterationSymbol: testCase.remainingIterationSymbol,
			LineSize:                 testCase.lineSize,
			MaxLineSize:              testCase.maxLineSize,
			LParen:                   testCase.lParen,
			RParen:                   testCase.rParen,
		}

		c := &render.ClockVal{}

		itr := &pbar.Iterator{
			Clock:    c,
			Settings: s,
			Write:    w,
			Values:   v,
		}

		err := itr.Initialize()
		got := testCase.buffer.String()

		assert.NotNil(t, c.StartTime, fmt.Sprintf("StartTime is nil!"))
		if testCase.expectError {
			assert.Error(t, err, fmt.Sprintf("Expected Error not raised"))
		} else {
			assert.NoError(t, err, fmt.Sprintf("Unexpected error raised: %v", err))
			assert.Equal(
				t,
				testCase.expectedOutput,
				got,
				fmt.Sprintf("Itr output incorrect expected: %v; got: %v", testCase.expectedOutput, got))
		}
	}
}

func TestUpdate(t *testing.T) {
	testCases := []struct {
		startVal                 float64
		stopVal                  float64
		stepVal                  float64
		currentVal               float64
		description              string
		finishedIterationSymbol  string
		currentIterationSymbol   string
		remainingIterationSymbol string
		lineSize                 int
		maxLineSize              int
		lParen                   string
		rParen                   string
		retain                   bool
		elapsedSecs              int64
		elapsedNanoSecs          int64
		buffer                   *bytes.Buffer
		expectError              bool
		expectedEndCurrentVal    float64
		expectedOutput           string
	}{
		{0.0, 5.0, 1.0, 0.0, "", "#", "#", "-", 10, 80, "|", "|", true, 0, 0, new(bytes.Buffer), false, 1.0, "\r|----------| 0.0/5.0 0.0% [elapsed: 00m:00s, left: N/A, N/A iters/sec]"},
		{0.0, 5.0, 1.0, 1.0, "", "#", "#", "-", 10, 80, "|", "|", true, 2, 0, new(bytes.Buffer), false, 2.0, "\r|##--------| 1.0/5.0 20.0% [elapsed: 00m:02s, left: 00m:08s, 0.50 iters/sec]"},
		{0.0, 5.0, 1.0, 5.0, "", "#", "#", "-", 10, 80, "|", "|", true, 4, 0, new(bytes.Buffer), false, 6.0, "\r|##########| 5.0/5.0 100.0% [elapsed: 00m:04s, left: 00m:00s, 1.25 iters/sec]\r\n"},
		{0.0, 5.0, 1.0, 5.0, "", "#", "#", "-", 10, 80, "|", "|", false, 4, 0, new(bytes.Buffer), false, 6.0, "\r|##########| 5.0/5.0 100.0% [elapsed: 00m:04s, left: 00m:00s, 1.25 iters/sec]\r\r                                                                             "},
		{1.0, 5.0, 1.0, 0.0, "", "#", "#", "-", 10, 80, "|", "|", true, 0, 0, new(bytes.Buffer), true, 1.0, ""},
		{1.0, 5.0, 1.0, 6.0, "", "#", "#", "-", 10, 80, "|", "|", true, 0, 0, new(bytes.Buffer), true, 1.0, ""},
	}

	for _, testCase := range testCases {
		render.NowTime = func() time.Time { return time.Unix(testCase.elapsedSecs, testCase.elapsedNanoSecs) }
		c := &render.ClockVal{
			StartTime: time.Unix(0, 0),
		}

		v := &render.Vals{
			Start:   testCase.startVal,
			Stop:    testCase.stopVal,
			Step:    testCase.stepVal,
			Current: testCase.currentVal,
		}

		w := &render.Writing{
			W: testCase.buffer,
		}

		s := &render.Set{
			Description:              testCase.description,
			FinishedIterationSymbol:  testCase.finishedIterationSymbol,
			CurrentIterationSymbol:   testCase.currentIterationSymbol,
			RemainingIterationSymbol: testCase.remainingIterationSymbol,
			LineSize:                 testCase.lineSize,
			MaxLineSize:              testCase.maxLineSize,
			LParen:                   testCase.lParen,
			RParen:                   testCase.rParen,
			Retain:                   testCase.retain,
		}

		itr := pbar.Iterator{
			Values:   v,
			Settings: s,
			Clock:    c,
			Write:    w,
		}

		err := itr.Update()
		got := testCase.buffer.String()

		assert.Equal(t,
			testCase.expectedOutput,
			got,
			fmt.Sprintf("Output string incorrect expected: %v; got: %v", testCase.expectedOutput, got),
		)

		if testCase.expectError {
			assert.Error(t, err, fmt.Sprintf("Unexpected error raised: %v", err))
		} else {
			assert.NoError(t, err, fmt.Sprintf("Expected Error not raised"))
			assert.Equal(
				t,
				testCase.expectedEndCurrentVal,
				itr.Values.GetCurrent(),
				fmt.Sprintf("Current Value incorrect expected: %v; got: %v", testCase.expectedEndCurrentVal, itr.Values.GetCurrent()),
			)
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
			assert.Implements(t, (*pbar.Iterate)(nil), itr, fmt.Sprintf("Itr does not implement Iterate!"))
		}
	}
}

func TestSetDescription(t *testing.T) {
	itr := &pbar.Iterator{}
	testCases := []struct {
		desc           string
		expectedPrefix string
	}{
		{"Hello", "Hello:"},
		{"World", "World:"},
	}

	for _, testCase := range testCases {
		itr.Settings = &render.Set{}
		itr.SetDescription(testCase.desc)
		message := fmt.Sprintf("Descriptions not equal; expected: %s, got: %s", testCase.expectedPrefix, itr.Settings.GetDescription())

		assert.Equal(
			t,
			testCase.expectedPrefix,
			itr.Settings.GetDescription(),
			message,
		)
	}
}

func TestSetFinishedIterationSymbol(t *testing.T) {
	itr := &pbar.Iterator{}
	testCases := []struct {
		symbol string
	}{
		{"Hello"},
		{"World"},
	}

	for _, testCase := range testCases {
		itr.Settings = &render.Set{}
		itr.SetFinishedIterationSymbol(testCase.symbol)
		message := fmt.Sprintf("FinishedIterationSymbols not equal; expected: %s, got: %s", testCase.symbol, itr.Settings.GetFinishedIterationSymbol())

		assert.Equal(
			t,
			testCase.symbol,
			itr.Settings.GetFinishedIterationSymbol(),
			message,
		)
	}
}

func TestSetCurrentIterationSymbol(t *testing.T) {
	itr := &pbar.Iterator{}
	testCases := []struct {
		symbol string
	}{
		{"Hello"},
		{"World"},
	}

	for _, testCase := range testCases {
		itr.Settings = &render.Set{}
		itr.SetCurrentIterationSymbol(testCase.symbol)
		message := fmt.Sprintf("CurrentInterationSymbols not equal; expected: %s, got: %s", testCase.symbol, itr.Settings.GetCurrentIterationSymbol())

		assert.Equal(
			t,
			testCase.symbol,
			itr.Settings.GetCurrentIterationSymbol(),
			message,
		)
	}
}

func TestSetRemainingIterationSymbol(t *testing.T) {
	itr := &pbar.Iterator{}
	testCases := []struct {
		symbol string
	}{
		{"Hello"},
		{"World"},
	}

	for _, testCase := range testCases {
		itr.Settings = &render.Set{}
		itr.SetRemainingIterationSymbol(testCase.symbol)
		message := fmt.Sprintf("RemainingIterationSymbols not equal; expected: %s, got: %s", testCase.symbol, itr.Settings.GetRemainingIterationSymbol())

		assert.Equal(
			t,
			testCase.symbol,
			itr.Settings.GetRemainingIterationSymbol(),
			message,
		)
	}
}

func TestSetLParenSymbol(t *testing.T) {
	itr := &pbar.Iterator{}
	testCases := []struct {
		symbol string
	}{
		{"Hello"},
		{"World"},
	}

	for _, testCase := range testCases {
		itr.Settings = &render.Set{}
		itr.SetLParen(testCase.symbol)
		message := fmt.Sprintf("LParens not equal; expected: %s, got: %s", testCase.symbol, itr.Settings.GetLParen())

		assert.Equal(
			t,
			testCase.symbol,
			itr.Settings.GetLParen(),
			message,
		)
	}
}

func TestSetRParenSymbol(t *testing.T) {
	itr := &pbar.Iterator{}
	testCases := []struct {
		symbol string
	}{
		{"Hello"},
		{"World"},
	}

	for _, testCase := range testCases {
		itr.Settings = &render.Set{}
		itr.SetRParen(testCase.symbol)
		message := fmt.Sprintf("RParens not equal; expected: %s, got: %s", testCase.symbol, itr.Settings.GetRParen())

		assert.Equal(
			t,
			testCase.symbol,
			itr.Settings.GetRParen(),
			message,
		)
	}
}

func TestSetRetain(t *testing.T) {
	itr := &pbar.Iterator{}
	testCases := []struct {
		value bool
	}{
		{true},
		{false},
	}

	for _, testCase := range testCases {
		itr.Settings = &render.Set{}
		itr.SetRetain(testCase.value)
		message := fmt.Sprintf("Retains not equal; expected: %v, got: %v", testCase.value, itr.Settings.GetRetain())

		assert.Equal(
			t,
			testCase.value,
			itr.Settings.GetRetain(),
			message,
		)
	}
}

func TestSetEqualTo(t *testing.T) {
	itr := &pbar.Iterator{}
	testCases := []struct {
		start        float64
		step         float64
		stop         float64
		isObject     bool
		expectedStop float64
		expectPanic  bool
	}{
		{0.0, 1.0, 5.0, false, 6.0, false},
		{0.0, 1.0, 5.0, true, 6.0, true},
	}

	for _, testCase := range testCases {
		itr.Values = &render.Vals{
			Start:    testCase.start,
			Stop:     testCase.stop,
			Step:     testCase.step,
			IsObject: testCase.isObject,
		}

		if testCase.expectPanic {
			assert.Panics(t, func() { itr.SetEqualTo() }, fmt.Sprintf("Panic not raised!"))
		} else {
			assert.NotPanics(t, func() { itr.SetEqualTo() }, fmt.Sprintf("Unexpected Panic"))
			assert.Equal(
				t,
				testCase.expectedStop,
				itr.Values.GetStop(),
				fmt.Sprintf("Stop value incorrect expected: %v; got: %v", testCase.expectedStop, itr.Values.GetStop()),
			)
		}

	}
}
