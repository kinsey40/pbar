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

type RenderObject struct {
	w                        io.Writer
	startValue               float64
	currentValue             float64
	endValue                 float64
	stepValue                float64
	startTime                time.Time
	prefix                   string
	iterationFinishedSymbol  string
	currentIterationSymbol   string
	remainingIterationSymbol string
	lineSize                 int
	lParen                   string
	rParen                   string
}

func MakeRenderObject(startValue, endValue, stepValue float64) *RenderObject {
	renderObj := new(RenderObject)
	renderObj.w = os.Stdout
	renderObj.startValue = startValue
	renderObj.currentValue = startValue
	renderObj.stepValue = stepValue
	renderObj.endValue = endValue
	renderObj.iterationFinishedSymbol = "="
	renderObj.currentIterationSymbol = ">"
	renderObj.remainingIterationSymbol = "-"
	renderObj.lParen = "|"
	renderObj.rParen = "|"

	if difference := int(endValue - startValue); difference < 10 {
		renderObj.lineSize = difference
	} else {
		renderObj.lineSize = 10
	}

	return renderObj
}

func (r *RenderObject) Initialize() {
	r.startTime = time.Now()
}

func (r *RenderObject) Update(currentValue float64) error {
	r.currentValue = currentValue
	barString := r.formatProgressBar()

	if currentValue == r.endValue {
		barString += "\n"
	}

	if err := r.render(barString); err != nil {
		return err
	}

	return nil
}

func (r *RenderObject) SetDescription(description string) {
	prefix := description + ":"
	r.prefix = prefix
}

func (r *RenderObject) SetIterationFinishedSymbol(newSymbol string) {
	r.iterationFinishedSymbol = newSymbol
}

func (r *RenderObject) SetRemainingIterationSymbol(newSymbol string) {
	r.remainingIterationSymbol = newSymbol
}

func (r *RenderObject) SetLParen(newSymbol string) {
	r.lParen = newSymbol
}

func (r *RenderObject) SetRParen(newSymbol string) {
	r.rParen = newSymbol
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
	ratio := r.currentValue / r.endValue
	percentage := ratio * 100.0

	numStepsComplete := int(ratio * float64(r.lineSize))

	bar := fmt.Sprintf("%s%s%s%s%s",
		r.lParen,
		strings.Repeat(r.iterationFinishedSymbol, numStepsComplete),
		strings.Repeat(r.currentIterationSymbol, 1),
		strings.Repeat(r.remainingIterationSymbol, int(r.endValue)-numStepsComplete),
		r.rParen)

	statistics := fmt.Sprintf("%.1f/%.1f %.1f%%", r.currentValue, r.endValue, percentage)
	speedMeter := r.formatSpeedMeter()
	progressBar := strings.Join([]string{r.prefix, bar, statistics, speedMeter}, " ")

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
