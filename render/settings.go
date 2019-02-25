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
 * File:   settings.go
 * Author: kinsey40
 *
 * Created on 13 January 2019, 11:05
 *
 * Settings holds the various progress bar settings
 *
 */

package render

import (
	"fmt"
	"strings"
)

var DefaultDescription = ""
var DefaultFinishedIterationSymbol = "#"
var DefaultCurrentIterationSymbol = "#"
var DefaultRemainingIterationSymbol = "-"
var DefaultLParen = "|"
var DefaultRParen = "|"
var DefaultMaxLineSize = 80
var DefaultLineSize = 10

// Settings enables the setting and getting the setting parameters
// for the progress bar. It also enables the creation of the bar
// string
type Settings interface {
	SetDescription(string)
	SetFinishedIterationSymbol(string)
	SetCurrentIterationSymbol(string)
	SetRemainingIterationSymbol(string)
	SetLineSize(int)
	SetMaxLineSize(int)
	SetLParen(string)
	SetRParen(string)

	GetDescription() string
	GetFinishedIterationSymbol() string
	GetCurrentIterationSymbol() string
	GetRemainingIterationSymbol() string
	GetLineSize() int
	GetMaxLineSize() int
	GetLParen() string
	GetRParen() string

	CreateBarString(int) string
}

// Set holds the setting parameters
type Set struct {
	Description              string
	FinishedIterationSymbol  string
	CurrentIterationSymbol   string
	RemainingIterationSymbol string
	LineSize                 int
	MaxLineSize              int
	LParen                   string
	RParen                   string
}

// NewSettings creates a Settings interface
func NewSettings() Settings {
	s := new(Set)
	s.Description = DefaultDescription
	s.FinishedIterationSymbol = DefaultFinishedIterationSymbol
	s.CurrentIterationSymbol = DefaultCurrentIterationSymbol
	s.RemainingIterationSymbol = DefaultRemainingIterationSymbol
	s.LineSize = DefaultLineSize
	s.MaxLineSize = DefaultMaxLineSize
	s.LParen = DefaultLParen
	s.RParen = DefaultRParen

	return s
}

// SetDescription sets the Description value
func (s *Set) SetDescription(str string) {
	if str != DefaultDescription {
		s.Description = str + ":"
	} else {
		s.Description = str
	}
}

// SetFinishedIterationSymbol sets the FinishedIterationSymbol value
func (s *Set) SetFinishedIterationSymbol(str string) {
	s.FinishedIterationSymbol = str
}

// SetCurrentIterationSymbol sets the CurrentIterationSymbol value
func (s *Set) SetCurrentIterationSymbol(str string) {
	s.CurrentIterationSymbol = str
}

// SetRRemainingIterationSymbol sets the RemainingIterationSymbol value
func (s *Set) SetRemainingIterationSymbol(str string) {
	s.RemainingIterationSymbol = str
}

// SetLineSize sets the LineSize value
func (s *Set) SetLineSize(i int) {
	if i > s.MaxLineSize {
		s.LineSize = s.MaxLineSize
	} else {
		s.LineSize = i
	}
}

// SetMaxLineSize sets the MaxLineSize value
func (s *Set) SetMaxLineSize(i int) {
	s.MaxLineSize = i
}

// SetLParen sets the LParen value
func (s *Set) SetLParen(str string) {
	s.LParen = str
}

// SetRParem sets the RParen value
func (s *Set) SetRParen(str string) {
	s.RParen = str
}

// GetDescription gets the Description value
func (s *Set) GetDescription() string {
	return s.Description
}

// GetFinishedIterationSymbol gets the FinishedIterationSymbol value
func (s *Set) GetFinishedIterationSymbol() string {
	return s.FinishedIterationSymbol
}

// GetCurrentIterationSymbol gets the CurrentIterationSymbol value
func (s *Set) GetCurrentIterationSymbol() string {
	return s.CurrentIterationSymbol
}

// GetRemainingIterationSymbol gets the RemainingIterationSymbol value
func (s *Set) GetRemainingIterationSymbol() string {
	return s.RemainingIterationSymbol
}

// GetLineSize gets the LineSize value
func (s *Set) GetLineSize() int {
	return s.LineSize
}

// GetMaxLineSize gets the MaxLineSize value
func (s *Set) GetMaxLineSize() int {
	return s.MaxLineSize
}

// GetLParen gets the LParen value
func (s *Set) GetLParen() string {
	return s.LParen
}

// GetRParen gets the RParen value
func (s *Set) GetRParen() string {
	return s.RParen
}

// getBarString creates the actual 'bar' within the progress bar
func (s *Set) CreateBarString(numStepsCompleted int) string {
	var finString string
	var currString string
	var remString string

	switch numStepsCompleted {
	case 0:
		remString = strings.Repeat(s.RemainingIterationSymbol, s.LineSize)
	case 1:
		currString = s.CurrentIterationSymbol
		remString = strings.Repeat(s.RemainingIterationSymbol, s.LineSize-1)
	case s.LineSize:
		finString = strings.Repeat(s.FinishedIterationSymbol, s.LineSize-1)
		currString = s.CurrentIterationSymbol
	default:
		finString = strings.Repeat(s.FinishedIterationSymbol, numStepsCompleted-1)
		currString = s.CurrentIterationSymbol
		remString = strings.Repeat(s.RemainingIterationSymbol, s.LineSize-numStepsCompleted)
	}

	barString := fmt.Sprintf("%s%s%s%s%s", s.LParen, finString, currString, remString, s.RParen)
	if s.Description != DefaultDescription {
		barString = strings.Join([]string{s.Description, barString}, " ")
	}

	return barString
}
