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

type RenderInterface interface {
	MakeRenderObject(float64, float64, float64) interface{}
	Update(float64) error
	render(string) error
	formatProgressBar() string
	formatSpeedMeter() error
}

type RenderObject struct {
	w                        io.Writer
	startValue               float64
	currentValue             float64
	endValue                 float64
	stepValue                float64
	startTime                time.Time
	Description              string
	IterationFinishedSymbol  string
	CurrentIterationSymbol   string
	RemainingIterationSymbol string
	LineSize                 int
	MaxLineSize              int
	LParen                   string
	RParen                   string
}

func MakeRenderObject(startValue, endValue, stepValue float64) *RenderObject {
	renderObj := new(RenderObject)
	renderObj.w = os.Stdout
	renderObj.startValue = startValue
	renderObj.currentValue = startValue
	renderObj.stepValue = stepValue
	renderObj.endValue = endValue
	renderObj.IterationFinishedSymbol = "#"
	renderObj.CurrentIterationSymbol = "#"
	renderObj.RemainingIterationSymbol = "-"
	renderObj.LParen = "|"
	renderObj.RParen = "|"
	renderObj.MaxLineSize = 80
	renderObj.LineSize = 10

	return renderObj
}

func (r *RenderObject) Update(currentValue float64, rendered bool) error {
	r.currentValue = currentValue
	barString := r.formatProgressBar()

	if currentValue == r.endValue {
		barString += "\n"
	}

	if err := r.render(barString); err != nil {
		return err
	}

	if !rendered {
		r.startTime = time.Now()
	}

	return nil
}

func (r *RenderObject) render(s string) error {
	stringToWrite := fmt.Sprintf("\r%s", s)
	_, err := io.WriteString(r.w, stringToWrite)

	if err != nil {
		return err
	}

	if f, ok := r.w.(*os.File); ok {
		f.Sync()
	}

	return nil
}

func (r *RenderObject) formatProgressBar() string {
	var itrFinishedString string
	var itrCurrentString string
	var itrRemainingString string

	ratio := r.currentValue / r.endValue
	percentage := ratio * 100.0

	numStepsComplete := int(ratio * float64(r.LineSize))
	if numStepsComplete == 0 {
		itrFinishedString = ""
		itrCurrentString = ""
		itrRemainingString = strings.Repeat(r.RemainingIterationSymbol, r.LineSize)
	} else if numStepsComplete == r.LineSize {
		itrFinishedString = strings.Repeat(r.IterationFinishedSymbol, r.LineSize)
		itrCurrentString = r.CurrentIterationSymbol
		itrRemainingString = ""
	} else if numStepsComplete == 1.0 {
		itrFinishedString = ""
		itrCurrentString = r.CurrentIterationSymbol
		itrRemainingString = strings.Repeat(r.RemainingIterationSymbol, r.LineSize-1)
	} else {
		itrFinishedString = strings.Repeat(r.IterationFinishedSymbol, numStepsComplete-1)
		itrCurrentString = r.CurrentIterationSymbol
		itrRemainingString = strings.Repeat(r.RemainingIterationSymbol, r.LineSize-numStepsComplete)
	}

	bar := fmt.Sprintf("%s%s%s%s%s",
		r.LParen,
		itrFinishedString,
		itrCurrentString,
		itrRemainingString,
		r.RParen)

	statistics := fmt.Sprintf("%.1f/%.1f %.1f%%", r.currentValue, r.endValue, percentage)
	speedMeter := r.formatSpeedMeter()
	progressBar := strings.Join([]string{r.Description, bar, statistics, speedMeter}, " ")

	return progressBar
}

func (r *RenderObject) formatSpeedMeter() string {
	if r.currentValue > r.startValue {
		elapsed := time.Now().Sub(r.startTime)
		rate := (r.currentValue - r.startValue) / elapsed.Seconds()
		remainingTime := time.Duration(math.Round((r.endValue-r.currentValue)*rate)) * time.Second

		return fmt.Sprintf("[elapsed: %s, left: %s, %.2f iters/sec]",
			formatTime(elapsed),
			formatTime(remainingTime),
			rate)

	} else {
		return fmt.Sprintf("[elapsed: %s, left: %s, %s iters/sec]",
			"00:00s",
			"N/A",
			"N/A")
	}
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
