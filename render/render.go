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
	"math"
	"strings"
)

type Render interface {
	Initialize(Clock, Settings)
	Update(float64) error
	render(string) error
	formatProgressBar() string
	getStatistics() (string, int)
	getBarString(numStepsCompleted int) string
	getSpeedMeter() string
}

// RenderObject is the underlying object which controls the
// various parameters relating to the rendering of the Pbar
// object.
type RenderObject struct {
	Write        Write
	Clock        Clock
	Settings     Settings
	StartValue   float64
	CurrentValue float64
	StopValue    float64
	StepValue    float64
}

// MakeRenderObject creates a RenderObject with the initial values set as the
// default variables.
func MakeRenderObject(startValue, stopValue, stepValue float64) Render {
	renderObj := new(RenderObject)
	renderObj.StartValue = startValue
	renderObj.CurrentValue = startValue
	renderObj.StepValue = stepValue
	renderObj.StopValue = stopValue

	return renderObj
}

// Initialize sets the Clock parameter within the RenderObject
// to a given Clock object.
func (r *RenderObject) Initialize(c Clock, s Settings) {
	r.Clock = c
	r.Settings = s
	// r.Write = NewWrite(s.GetWriter())
}

// Update causes the RenderObject to progress to the next step,
// returning an error if the currentValue is below the StartValue
// or above the StopValue.
func (r *RenderObject) Update(currentValue float64) error {
	if currentValue < r.StartValue || currentValue > r.StopValue {
		return fmt.Errorf(
			"Current value: %f is incorrect. Start: %f; end: %f",
			currentValue,
			r.StartValue,
			r.StopValue,
		)
	}

	r.CurrentValue = currentValue
	wholeProgressBar := r.formatProgressBar()

	if currentValue == r.StopValue {
		wholeProgressBar += "\n"
	}

	err := r.render(wholeProgressBar)

	return err
}

// render writes the relevant string to the relevant writer
func (r *RenderObject) render(s string) error {
	err := r.Write.WriteString(fmt.Sprintf("\r%s", s))
	if err != nil {
		return err
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
		r.Settings.GetLParen(),
		barString,
		r.Settings.GetRParen(),
	)

	speedMeter := r.getSpeedMeter()
	progressBar := strings.Join([]string{r.Settings.GetDescription(), bar, statistics, speedMeter}, " ")

	return progressBar
}

// getStatistics calculates all the numerical values relating to the
// progression of the progress bar. These are then formed and returned
// in a string, alongside the number of steps that have been completed.
func (r *RenderObject) getStatistics() (string, int) {
	ratio := r.CurrentValue / r.StopValue
	percentage := ratio * 100.0
	statistics := fmt.Sprintf("%.1f/%.1f %.1f%%", r.CurrentValue, r.StopValue, percentage)
	numStepsCompleted := int(ratio * float64(r.Settings.GetLineSize()))

	return statistics, numStepsCompleted
}

// getBarString creates the actual 'bar' within the progress bar
func (r *RenderObject) getBarString(numStepsCompleted int) string {
	var finString string
	var currString string
	var remString string

	remainingSymbol := r.Settings.GetRemainingIterationSymbol()
	currentSymbol := r.Settings.GetCurrentIterationSymbol()
	finishedSymbol := r.Settings.GetFinishedIterationSymbol()
	lineSize := r.Settings.GetLineSize()

	switch numStepsCompleted {
	case 0:
		remString = strings.Repeat(remainingSymbol, lineSize)
	case 1:
		currString = currentSymbol
		remString = strings.Repeat(remainingSymbol, lineSize-1)
	case lineSize:
		finString = strings.Repeat(finishedSymbol, lineSize-1)
		currString = currentSymbol
	default:
		finString = strings.Repeat(finishedSymbol, numStepsCompleted-1)
		currString = currentSymbol
		remString = strings.Repeat(remainingSymbol, lineSize-numStepsCompleted)
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
		remainingTime := r.Clock.Remaining(math.Round((r.StopValue - r.CurrentValue) / rate))

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
