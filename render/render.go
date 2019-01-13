/*


 */

package render

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type RenderObject struct {
	w                        io.Writer
	startValue               float64
	currentValue             float64
	endValue                 float64
	startTime                time.Time
	prefix                   string
	iterationFinishedSymbol  string
	remainingIterationSymbol string
	lineSize                 int
	lParen                   string
	rParen                   string
}

func MakeRenderObject(startValue, endValue float64) *RenderObject {
	renderObj := new(RenderObject)
	renderObj.w = os.Stdout
	renderObj.startValue = startValue
	renderObj.currentValue = startValue
	renderObj.endValue = endValue
	renderObj.iterationFinishedSymbol = "#"
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

func (r *RenderObject) Update(currentValue float64) error {
	if currentValue == r.startValue {
		r.startTime = time.Now()
	}

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

func (r *RenderObject) SetPrefix(description string) {
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
	bar := fmt.Sprintf("%s%s%s%s",
		r.lParen,
		strings.Repeat(r.iterationFinishedSymbol, numStepsComplete),
		strings.Repeat(r.remainingIterationSymbol, r.lineSize-numStepsComplete),
		r.rParen)

	statistics := fmt.Sprintf("%.1f/%.1f %5.2f%%", r.currentValue, r.endValue, percentage)
	speedMeter := r.formatSpeedMeter()
	progressBar := strings.Join([]string{r.prefix, bar, statistics, speedMeter}, " ")

	return progressBar
}

func (r *RenderObject) formatSpeedMeter() string {
	var rate float64
	elapsed := time.Now().Sub(r.startTime)
	remainingTime := time.Duration((elapsed.Seconds()/r.currentValue-r.startValue)*(r.endValue-(r.currentValue-r.startValue))) * time.Second

	if r.currentValue != r.startValue {
		rate = float64(r.endValue) / elapsed.Seconds()
	} else {
		rate = 0.0
	}

	return fmt.Sprintf("[elapsed: %s, left: %s, %5.2f iters/sec]",
		formatTime(elapsed),
		formatTime(remainingTime),
		rate)
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
