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
 * File:   settings_test.go
 * Author: kinsey40
 *
 * Created on 13 January 2019, 11:05
 *
 * Render performs the acutual rendering of the progress bar onto the
 * terminal display.
 *
 */

package render_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/kinsey40/pbar/render"
	"github.com/stretchr/testify/assert"
)

type MockTerminal struct {
	toSend       []byte
	bytesPerRead int
	received     []byte
}

func TestNewSettings(t *testing.T) {
	s := render.NewSettings()
	set := s.(*render.Set)

	message := fmt.Sprintf("NewSettings type does not match expected: %v; got: %v", (*render.Settings)(nil), s)
	assert.Implements(t, (*render.Settings)(nil), s, message)
	assert.Equal(
		t,
		render.DefaultDescription,
		set.Description,
		fmt.Sprintf("Description incorrect expected: %v; got: %v", render.DefaultDescription, set.Description),
	)

	assert.Equal(
		t,
		render.DefaultFinishedIterationSymbol,
		set.FinishedIterationSymbol,
		fmt.Sprintf("Finished Iteration symbol incorrect expected: %v; got %v", render.DefaultFinishedIterationSymbol, set.FinishedIterationSymbol),
	)

	assert.Equal(
		t,
		render.DefaultCurrentIterationSymbol,
		set.CurrentIterationSymbol,
		fmt.Sprintf("Current Iteration symbol incorrect expected: %v; got %v", render.DefaultCurrentIterationSymbol, set.CurrentIterationSymbol),
	)

	assert.Equal(
		t,
		render.DefaultRemainingIterationSymbol,
		set.RemainingIterationSymbol,
		fmt.Sprintf("Reminaing Iteration symbol incorrect expected: %v; got %v", render.DefaultRemainingIterationSymbol, set.RemainingIterationSymbol),
	)

	assert.Equal(
		t,
		render.DefaultLParen,
		set.LParen,
		fmt.Sprintf("LParen symbol incorrect expected: %v; got %v", render.DefaultLParen, set.LParen),
	)

	assert.Equal(
		t,
		render.DefaultRParen,
		set.RParen,
		fmt.Sprintf("RParen symbol incorrect expected: %v; got %v", render.DefaultRParen, set.RParen),
	)

	assert.Equal(
		t,
		render.DefaultLineSize,
		set.LineSize,
		fmt.Sprintf("LineSize incorred expected: %v; got: %v", render.DefaultLineSize, set.LineSize),
	)

	assert.Equal(
		t,
		render.DefaultMaxLineSize,
		set.MaxLineSize,
		fmt.Sprintf("MaxLineSize incorred expected: %v; got: %v", render.DefaultMaxLineSize, set.MaxLineSize),
	)
}

func TestSetDescription(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput string
	}{
		{"Hello", "Hello:"},
		{render.DefaultDescription, render.DefaultDescription},
	}

	for _, testCase := range testCases {
		s := &render.Set{}
		s.SetDescription(testCase.input)
		message := fmt.Sprintf("Description incorrectly set expected: %v; got %v", testCase.expectedOutput, s.Description)

		assert.Equal(t, testCase.expectedOutput, s.Description, message)
	}
}

func TestSetFinishedIterationSymbol(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput string
	}{
		{"Hello", "Hello"},
	}

	for _, testCase := range testCases {
		s := &render.Set{}
		s.SetFinishedIterationSymbol(testCase.input)
		message := fmt.Sprintf("FinishedIterationSymbol incorrectly set expected: %v; got %v", testCase.expectedOutput, s.FinishedIterationSymbol)

		assert.Equal(t, testCase.expectedOutput, s.FinishedIterationSymbol, message)
	}
}

func TestSetCurrentIterationSymbol(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput string
	}{
		{"Hello", "Hello"},
	}

	for _, testCase := range testCases {
		s := &render.Set{}
		s.SetCurrentIterationSymbol(testCase.input)
		message := fmt.Sprintf("CurrentIterationSymbol incorrectly set expected: %v; got %v", testCase.expectedOutput, s.CurrentIterationSymbol)

		assert.Equal(t, testCase.expectedOutput, s.CurrentIterationSymbol, message)
	}
}

func TestSetRemainingIterationSymbol(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput string
	}{
		{"Hello", "Hello"},
	}

	for _, testCase := range testCases {
		s := &render.Set{}
		s.SetRemainingIterationSymbol(testCase.input)
		message := fmt.Sprintf("RemainingIterationSymbol incorrectly set expected: %v; got %v", testCase.expectedOutput, s.RemainingIterationSymbol)

		assert.Equal(t, testCase.expectedOutput, s.RemainingIterationSymbol, message)
	}
}

func TestSetLineSize(t *testing.T) {
	testCases := []struct {
		input          int
		maxLineSize    int
		expectedOutput int
	}{
		{5, 10, 5},
		{10, 5, 5},
	}

	for _, testCase := range testCases {
		s := &render.Set{
			MaxLineSize: testCase.maxLineSize,
		}

		s.SetLineSize(testCase.input)
		message := fmt.Sprintf("LineSize incorrectly set expected: %v; got %v", testCase.expectedOutput, s.LineSize)

		assert.Equal(t, testCase.expectedOutput, s.LineSize, message)
	}
}

func TestSetMaxLineSize(t *testing.T) {
	testCases := []struct {
		input          int
		expectedOutput int
	}{
		{5, 5},
		{10, 10},
		{15, 15},
	}

	for _, testCase := range testCases {
		s := &render.Set{}
		s.SetMaxLineSize(testCase.input)
		message := fmt.Sprintf("MaxLineSize incorrectly set expected: %v; got %v", testCase.expectedOutput, s.MaxLineSize)

		assert.Equal(t, testCase.expectedOutput, s.MaxLineSize, message)
	}
}

func TestSetLParen(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput string
	}{
		{"Hello", "Hello"},
	}

	for _, testCase := range testCases {
		s := &render.Set{}
		s.SetLParen(testCase.input)
		message := fmt.Sprintf("LParen incorrectly set expected: %v; got %v", testCase.expectedOutput, s.LParen)

		assert.Equal(t, testCase.expectedOutput, s.LParen, message)
	}
}

func TestSetRParen(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput string
	}{
		{"Hello", "Hello"},
	}

	for _, testCase := range testCases {
		s := &render.Set{}
		s.SetRParen(testCase.input)
		message := fmt.Sprintf("RParen incorrectly set expected: %v; got %v", testCase.expectedOutput, s.RParen)

		assert.Equal(t, testCase.expectedOutput, s.RParen, message)
	}
}

func TestSetSuffix(t *testing.T) {
	testCases := []struct {
		suffix         string
		input          string
		expectedOutput string
	}{
		{"Hello", "Hello", "Hello"},
		{"Hello[1A", "Hello", "HelloHello[1A"},
		{"Hello[K", "Hello", "Hello[KHello"},
	}

	for _, testCase := range testCases {
		s := &render.Set{
			Suffix: testCase.suffix,
		}
		s.SetSuffix(testCase.input)
		message := fmt.Sprintf("Suffix incorrectly set expected: %v; got %v", testCase.expectedOutput, s.Suffix)

		assert.Equal(t, testCase.expectedOutput, s.Suffix, message)
	}
}

func TestGetDescription(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput string
	}{
		{"Hello", "Hello"},
	}

	for _, testCase := range testCases {
		s := &render.Set{}
		s.Description = testCase.input
		output := s.GetDescription()
		message := fmt.Sprintf("Description incorrect get expected: %v, got: %v", testCase.expectedOutput, output)

		assert.Equal(t, testCase.expectedOutput, output, message)
	}
}

func TestGetFinishedIterationSymbol(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput string
	}{
		{"Hello", "Hello"},
	}

	for _, testCase := range testCases {
		s := &render.Set{}
		s.FinishedIterationSymbol = testCase.input
		output := s.GetFinishedIterationSymbol()
		message := fmt.Sprintf("FinishedIterationSymbol incorrect get expected: %v, got: %v", testCase.expectedOutput, output)

		assert.Equal(t, testCase.expectedOutput, output, message)
	}
}

func TestGetCurrentIterationSymbol(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput string
	}{
		{"Hello", "Hello"},
	}

	for _, testCase := range testCases {
		s := &render.Set{}
		s.CurrentIterationSymbol = testCase.input
		output := s.GetCurrentIterationSymbol()
		message := fmt.Sprintf("CurrentIterationSymbol incorrect get expected: %v, got: %v", testCase.expectedOutput, output)

		assert.Equal(t, testCase.expectedOutput, output, message)
	}
}

func TestGetRemainingIterationSymbol(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput string
	}{
		{"Hello", "Hello"},
	}

	for _, testCase := range testCases {
		s := &render.Set{}
		s.RemainingIterationSymbol = testCase.input
		output := s.GetRemainingIterationSymbol()
		message := fmt.Sprintf("RemainingIterationSymbol incorrect get expected: %v, got: %v", testCase.expectedOutput, output)

		assert.Equal(t, testCase.expectedOutput, output, message)
	}
}

func TestGetLineSize(t *testing.T) {
	testCases := []struct {
		input          int
		expectedOutput int
	}{
		{5, 5},
		{10, 10},
		{15, 15},
		{20, 20},
	}

	for _, testCase := range testCases {
		s := &render.Set{}
		s.LineSize = testCase.input
		output := s.GetLineSize()
		message := fmt.Sprintf("LineSize incorrect get expected: %v, got: %v", testCase.expectedOutput, output)

		assert.Equal(t, testCase.expectedOutput, output, message)
	}
}

func TestGetMaxLineSize(t *testing.T) {
	testCases := []struct {
		input          int
		expectedOutput int
	}{
		{5, 5},
		{10, 10},
		{15, 15},
		{20, 20},
	}

	for _, testCase := range testCases {
		s := &render.Set{}
		s.MaxLineSize = testCase.input
		output := s.GetMaxLineSize()
		message := fmt.Sprintf("MaxLineSize incorrect get expected: %v, got: %v", testCase.expectedOutput, output)

		assert.Equal(t, testCase.expectedOutput, output, message)
	}
}

func TestGetLParen(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput string
	}{
		{"Hello", "Hello"},
	}

	for _, testCase := range testCases {
		s := &render.Set{}
		s.LParen = testCase.input
		output := s.GetLParen()
		message := fmt.Sprintf("LParen incorrect get expected: %v, got: %v", testCase.expectedOutput, output)

		assert.Equal(t, testCase.expectedOutput, output, message)
	}
}

func TestGetRParen(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput string
	}{
		{"Hello", "Hello"},
	}

	for _, testCase := range testCases {
		s := &render.Set{}
		s.RParen = testCase.input
		output := s.GetRParen()
		message := fmt.Sprintf("RParen incorrect get expected: %v, got: %v", testCase.expectedOutput, output)

		assert.Equal(t, testCase.expectedOutput, output, message)
	}
}

func TestGetSuffix(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput string
	}{
		{"Hello", "Hello"},
	}

	for _, testCase := range testCases {
		s := &render.Set{}
		s.Suffix = testCase.input
		output := s.GetSuffix()
		message := fmt.Sprintf("Suffix incorrect get expected: %v, got: %v", testCase.expectedOutput, output)

		assert.Equal(t, testCase.expectedOutput, output, message)
	}
}

func TestSetIdealLineSize(t *testing.T) {
	testCases := []struct {
		description      string
		lParen           string
		rParen           string
		width            int
		terminalErr      error
		expectedLineSize int
	}{
		{"", "|", "|", 80, nil, 80 - render.NumberOfCharacters - render.NumberOfCharactersBuffer},
		{"", "|", "|", 80, errors.New("An error"), 80 - render.NumberOfCharacters - render.NumberOfCharactersBuffer},
	}

	for _, testCase := range testCases {
		s := &render.Set{}
		render.GetTerminal = func() uintptr { return 0 }
		render.TerminalSize = func(_ int) (int, int, error) { return testCase.width, 0, testCase.terminalErr }

		err := s.SetIdealLineSize()
		if testCase.terminalErr != nil {
			assert.Error(t, err, fmt.Sprintf("Expected error not raised"))
		} else {
			message := fmt.Sprintf("LineSize incorrectly set through IdealLineSize expected: %v; got %v", testCase.expectedLineSize, s.LineSize)
			assert.NoError(t, err, fmt.Sprintf("Unexpected error raised!"))
			assert.Equal(t, testCase.expectedLineSize, s.LineSize, message)
		}
	}
}

func TestCreateBarString(t *testing.T) {
	testCases := []struct {
		numStepsCompleted        int
		lineSize                 int
		finishedIterationSymbol  string
		currentIterationSymbol   string
		remainingIterationSymbol string
		description              string
		lParen                   string
		rParen                   string
		expectedOutput           string
	}{
		{0, 10, "#", "#", "-", "", "|", "|", "|----------|"},
		{0, 10, "#", "#", "-", "", "|", "|", "|----------|"},
		{1, 10, "#", "#", "-", "", "|", "|", "|#---------|"},
		{1, 10, "=", ">", "-", "", "|", "|", "|>---------|"},
		{2, 10, "#", "#", "-", "", "|", "|", "|##--------|"},
		{2, 10, "=", ">", "-", "", "|", "|", "|=>--------|"},
		{10, 10, "=", ">", "-", "", "|", "|", "|=========>|"},
		{10, 10, "#", "#", "-", "", "|", "|", "|##########|"},
		{10, 10, "#", "#", "-", "Hello:", "|", "|", "Hello: |##########|"},
		{10, 10, "#", "#", "-", "World:", "|", "|", "World: |##########|"},
	}

	for _, testCase := range testCases {
		s := &render.Set{
			FinishedIterationSymbol:  testCase.finishedIterationSymbol,
			CurrentIterationSymbol:   testCase.currentIterationSymbol,
			RemainingIterationSymbol: testCase.remainingIterationSymbol,
			LineSize:                 testCase.lineSize,
			Description:              testCase.description,
			LParen:                   testCase.lParen,
			RParen:                   testCase.rParen,
		}

		output := s.CreateBarString(testCase.numStepsCompleted)
		message := fmt.Sprintf("Output incorrect expected: %v; got: %v", testCase.expectedOutput, output)

		assert.Equal(t, testCase.expectedOutput, output, message)
	}
}
