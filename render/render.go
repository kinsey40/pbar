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
 * File:   tqdm.go
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
	"time"
)

var DefaultFinishedIterationSymbol = "#"
var DefaultCurrentIterationSymbol = "#"
var DefaultRemainingIterationSymbol = "-"
var DefaultLParen = "|"
var DefaultRParen = "|"
var DefaultMaxLineSize = 80
var DefaultLineSize = 10

type RenderObject struct {
	W                        io.Writer
	StartValue               float64
	CurrentValue             float64
	EndValue                 float64
	StepValue                float64
	StartTime                time.Time
	Description              string
	FinishedIterationSymbol  string
	CurrentIterationSymbol   string
	RemainingIterationSymbol string
	LineSize                 int
	MaxLineSize              int
	LParen                   string
	RParen                   string
}

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

func (r *RenderObject) Initialize(timeVal time.Time) {
	r.StartTime = timeVal
}

func (r *RenderObject) Update(currentValue float64, timeVal time.Time) error {
	r.CurrentValue = currentValue
	wholeProgressBar := r.formatProgressBar(timeVal)

	if currentValue == r.EndValue {
		wholeProgressBar += "\n"
	}
	err := r.render(wholeProgressBar)

	return err
}

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

func (r *RenderObject) formatProgressBar(timeVal time.Time) string {
	statistics, numStepsCompleted := r.getStatistics()
	barString := r.getBarString(numStepsCompleted)
	bar := fmt.Sprintf("%s%s%s",
		r.LParen,
		barString,
		r.RParen)

	speedMeter := r.getSpeedMeter(timeVal)
	progressBar := strings.Join([]string{r.Description, bar, statistics, speedMeter}, " ")

	return progressBar
}

func (r *RenderObject) getStatistics() (string, int) {
	ratio := r.CurrentValue / r.EndValue
	percentage := ratio * 100.0
	statistics := fmt.Sprintf("%.1f/%.1f %.1f%%", r.CurrentValue, r.EndValue, percentage)
	numStepsCompleted := int(ratio * float64(r.LineSize))

	return statistics, numStepsCompleted
}

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

func (r *RenderObject) getSpeedMeter(timeVal time.Time) string {
	if r.CurrentValue > r.StartValue {
		elapsed := timeVal.Sub(r.StartTime)
		rate := (r.CurrentValue - r.StartValue) / elapsed.Seconds()
		remainingTime := time.Duration(math.Round((r.EndValue-r.CurrentValue)/rate)) * time.Second

		return fmt.Sprintf("[elapsed: %s, left: %s, %.2f iters/sec]",
			formatTime(elapsed),
			formatTime(remainingTime),
			rate,
		)
	}

	return fmt.Sprintf("[elapsed: %s, left: %s, %s iters/sec]",
		"00m:00s",
		"N/A",
		"N/A",
	)
}

func formatTime(d time.Duration) string {
	secs := (d % time.Minute) / time.Second
	mins := (d % time.Hour) / time.Minute
	hours := d / time.Hour

	if hours == 0 {
		return fmt.Sprintf("%02dm:%02ds", mins, secs)
	}

	return fmt.Sprintf("%02dh:%02dm:%02ds", hours, mins, secs)
}
