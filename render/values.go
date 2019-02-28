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
 * File:   values.go
 * Author: kinsey40
 *
 * Created on 13 January 2019, 11:05
 *
 * Holds the underlying start, stop, step and current values for the progress bar
 *
 */

package render

import "fmt"

// Values holds the start, stop, step and current values for the progress bar.
// It also enables statistics to be calculated for these values relating to
// the ratio of completion.
type Values interface {
	SetStart(float64)
	SetStop(float64)
	SetStep(float64)
	SetCurrent(float64)

	GetStart() float64
	GetStop() float64
	GetStep() float64
	GetCurrent() float64

	Statistics(int) (string, int)
}

// Vals holds the Start, Stop, Step and Current values
type Vals struct {
	Start   float64
	Stop    float64
	Step    float64
	Current float64
}

// NewValues generates a NewValues interface
func NewValues() Values {
	v := new(Vals)

	return v
}

// SetStart sets the Start value
func (v *Vals) SetStart(s float64) {
	v.Start = s
}

// SetStop sets the Stop value
func (v *Vals) SetStop(s float64) {
	v.Stop = s
}

// SetStep sets the Step value
func (v *Vals) SetStep(s float64) {
	v.Step = s
}

// SetCurrent sets the Current value
func (v *Vals) SetCurrent(s float64) {
	v.Current = s
}

// GetStart gets the Start value
func (v *Vals) GetStart() float64 {
	return v.Start
}

// GetStop gets the Stop value
func (v *Vals) GetStop() float64 {
	return v.Stop
}

// GetStep gets the Step value
func (v *Vals) GetStep() float64 {
	return v.Step
}

// GetCurrent gets the Current value
func (v *Vals) GetCurrent() float64 {
	return v.Current
}

// Statistics calculates all the numerical values relating to the
// progression of the progress bar. These are then formed and returned
// in a string, alongside the number of steps that have been completed.
func (v *Vals) Statistics(linesize int) (string, int) {
	ratio := v.Current / v.Stop
	percentage := ratio * 100.0
	statistics := fmt.Sprintf("%.1f/%.1f %.1f%%", v.Current, v.Stop, percentage)
	numStepsCompleted := int(ratio * float64(linesize))

	return statistics, numStepsCompleted
}
