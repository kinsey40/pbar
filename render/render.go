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
 * File:   render.go
 * Author: kinsey40
 *
 * Created on 13 January 2019, 11:05
 *
 * Render performs the acutual rendering of the progress bar onto the
 * terminal display.
 *
 */

package render

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"

	"github.com/kinsey40/pbar/clock"
)

var DefaultFinishedIterationSymbol = "#"
var DefaultCurrentIterationSymbol = "#"
var DefaultRemainingIterationSymbol = "-"
var DefaultLParen = "|"
var DefaultRParen = "|"
var DefaultMaxLineSize = 80
var DefaultLineSize = 10

// RenderObject is the underlying object which controls the
// various parameters relating to the rendering of the Pbar
// object.
type RenderObject struct {
	W                        io.Writer
	Clock                    clock.Clock
	StartValue               float64
	CurrentValue             float64
	EndValue                 float64
	StepValue                float64
	Description              string
	FinishedIterationSymbol  string
	CurrentIterationSymbol   string
	RemainingIterationSymbol string
	LineSize                 int
	MaxLineSize              int
	LParen                   string
	RParen                   string
}

// MakeRenderObject creates a RenderObject with the initial values set as the
// default variables.
func MakeRenderObject(startValue, endValue, stepValue float64) *RenderObject {
	renderObj := new(RenderObject)
	renderObj.W = os.Stdout
	renderObj.StartValue = startValue
	renderObj.CurrentValue = startValue
	renderObj.StepValue = stepValue
	renderObj.EndValue = endValue
	renderObj.FinishedIterationSymbol = DefaultFinishedIterationSymbol
	renderObj.CurrentIterationSymbol = DefaultCurrentIterationSymbol
	renderObj.RemainingIterationSymbol = DefaultRemainingIterationSymbol
	renderObj.LParen = DefaultLParen
	renderObj.RParen = DefaultRParen
	renderObj.MaxLineSize = DefaultMaxLineSize
	renderObj.LineSize = DefaultLineSize

	return renderObj
}

// Initialize sets the Clock parameter within the RenderObject
// to a given Clock object.
func (r *RenderObject) Initialize(c clock.Clock) {
	r.Clock = c
}

// Update causes the RenderObject to progress to the next step,
// returning an error if the currentValue is below the StartValue
// or above the EndValue.
func (r *RenderObject) Update(currentValue float64) error {
	if currentValue < r.StartValue || currentValue > r.EndValue {
		return fmt.Errorf(
			"Current value: %f is incorrect. Start: %f; end: %f",
			currentValue,
			r.StartValue,
			r.EndValue)
	}

	r.CurrentValue = currentValue
	wholeProgressBar := r.formatProgressBar()

	if currentValue == r.EndValue {
		wholeProgressBar += "\n"
	}

	err := r.render(wholeProgressBar)

	return err
}

// render writes the relevant string to the relevant writer
func (r *RenderObject) render(s string) error {
	stringToWrite := fmt.Sprintf("\r%s", s)
	_, err := io.WriteString(r.W, stringToWrite)

	if err != nil {
		return err
	}

	if f, ok := r.W.(*os.File); ok {
		f.Sync()
	}

	return nil
}

// formatProgressBar creates the progress bar to be displayed
// by the writer. It gathers all the relevant sections from
// the other functions.
func (r *RenderObject) formatProgressBar() string {
	statistics, numStepsCompleted := r.getStatistics()
	barString := r.getBarString(numStepsCompleted)
	bar := fmt.Sprintf("%s%s%s",
		r.LParen,
		barString,
		r.RParen)

	speedMeter := r.getSpeedMeter()
	progressBar := strings.Join([]string{r.Description, bar, statistics, speedMeter}, " ")

	return progressBar
}

// getStatistics calculates all the numerical values relating to the
// progression of the progress bar. These are then formed and returned
// in a string, alongside the number of steps that have been completed.
func (r *RenderObject) getStatistics() (string, int) {
	ratio := r.CurrentValue / r.EndValue
	percentage := ratio * 100.0
	statistics := fmt.Sprintf("%.1f/%.1f %.1f%%", r.CurrentValue, r.EndValue, percentage)
	numStepsCompleted := int(ratio * float64(r.LineSize))

	return statistics, numStepsCompleted
}

// getBarString creates the actual 'bar' within the progress bar
func (r *RenderObject) getBarString(numStepsCompleted int) string {
	var finString string
	var currString string
	var remString string

	switch numStepsCompleted {
	case 0:
		remString = strings.Repeat(r.RemainingIterationSymbol, r.LineSize)
	case 1:
		currString = r.CurrentIterationSymbol
		remString = strings.Repeat(r.RemainingIterationSymbol, r.LineSize-1)
	case r.LineSize:
		finString = strings.Repeat(r.FinishedIterationSymbol, r.LineSize-1)
		currString = r.CurrentIterationSymbol
	default:
		finString = strings.Repeat(r.FinishedIterationSymbol, numStepsCompleted-1)
		currString = r.CurrentIterationSymbol
		remString = strings.Repeat(r.RemainingIterationSymbol, r.LineSize-numStepsCompleted)
	}

	barString := fmt.Sprintf("%s%s%s", finString, currString, remString)

	return barString
}

// getSpeedMeter forms the part of the progress bar relating
// to the elapsed and remaining time, as well as the rate of
// iterations per second.
func (r *RenderObject) getSpeedMeter() string {
	if r.CurrentValue > r.StartValue {
		elapsed := r.Clock.Subtract(r.Clock.Now())
		rate := (r.CurrentValue - r.StartValue) / elapsed.Seconds()
		remainingTime := r.Clock.Remaining(math.Round((r.EndValue - r.CurrentValue) / rate))

		return fmt.Sprintf("[elapsed: %s, left: %s, %.2f iters/sec]",
			r.Clock.Format(elapsed),
			r.Clock.Format(remainingTime),
			rate,
		)
	}

	return fmt.Sprintf("[elapsed: %s, left: %s, %s iters/sec]",
		"00m:00s",
		"N/A",
		"N/A",
	)
}
