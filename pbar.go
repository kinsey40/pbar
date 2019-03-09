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
 * File:   pbar.go
 * Author: kinsey40
 *
 * Created on 13 January 2019, 11:05
 *
 * The main file for the pbar package, this enables the creation of the pbar
 * object. The user can then edit the specific variables associated with the
 * object.
 *
 */

package pbar

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/kinsey40/pbar/render"
)

// Iterate enables the progress bar execution.
// Various settings can be manipulated using the Set*() functions.
// To create a progress bar, call the Pbar() function.
// Initialize() the progress bar before the for-loop
// Update() at the end of each iteration within the for-loop.
//
// It is recommended that you do not create an Iterate object directly,
// but instead use the Pbar() function which will automatically set the variables
// correctly
type Iterate interface {
	Initialize() error
	Update() error
	SetDescription(string)
	SetFinishedIterationSymbol(string)
	SetCurrentIterationSymbol(string)
	SetRemainingIterationSymbol(string)
	SetLParen(string)
	SetRParen(string)
	SetRetain(bool)
	SetEqualTo()
	Multi()
	MultiEnd()

	progress() error
	createIteratorFromObject(interface{})
	createIteratorFromValues(...interface{})
}

// Iterator object stores the relevant parameters
// associated with the progress bar, this is returned
// by the Pbar function.
type Iterator struct {
	Values   render.Values
	Clock    render.Clock
	Settings render.Settings
	Write    render.Write
}

// makeIteratorObject creates an Iterate interface
func makeIteratorObject() Iterate {
	itr := new(Iterator)
	itr.Clock = render.NewClock()
	itr.Settings = render.NewSettings()
	itr.Values = render.NewValues()
	itr.Write = render.NewWrite()

	return itr
}

// Pbar creates a progress bar from the inputted values or object.
// The user can pass either a valid object or list of numbers
// (of type: float32, float64, int8, int16, int32, int64 or int).
// Valid objects should be passed as single values, a valid object
// (of type: array, slice, string, map or buffered channel).
func Pbar(values ...interface{}) (Iterate, error) {
	itr := makeIteratorObject()
	isObject, err := isObject(values...)
	if err != nil {
		return nil, err
	}

	if err = checkValues(isObject, values...); err != nil {
		return nil, err
	}

	if isObject {
		itr.createIteratorFromObject(values[0])
	} else {
		itr.createIteratorFromValues(values...)
	}

	return itr, err
}

// Initialize sets the internal timer to start,
// enabling output relating to the time taken for
// iterations within the progress bar.
func (itr *Iterator) Initialize() error {
	itr.Clock.SetStartTime()
	if err := itr.Settings.SetIdealLineSize(); err != nil {
		return err
	}

	return itr.Update()
}

// Update moves the iteration forward by one step. This should
// be performed at the end of the iteration sequence
// (i.e. at the end of the for-loop).
func (itr *Iterator) Update() error {
	itr.Clock.Now()
	return itr.progress()
}

// SetDescription sets the Description parameter, which causes the Pbar
// to output a String at the start of the progress bar, effectively
// enabling the progress bars to be named within the output.
//
// Default Value: ""
func (itr *Iterator) SetDescription(descrip string) {
	itr.Settings.SetDescription(descrip)
}

// SetFinishedIterationSymbol sets the FinishedIterationSymbol, which
// is the symbol within the progress bar that shows that particular
// iteration has completed it's execution.
//
// Default Value: "#"
func (itr *Iterator) SetFinishedIterationSymbol(newSymbol string) {
	itr.Settings.SetFinishedIterationSymbol(newSymbol)
}

// SetCurrentIterationSymbol sets the CurrentIterationSymbol, which
// is the symbol within the progress bar that shows the iteration
// which is currently being executed.
//
// Default Value: "#"
func (itr *Iterator) SetCurrentIterationSymbol(newSymbol string) {
	itr.Settings.SetCurrentIterationSymbol(newSymbol)
}

// SetRemainingIterationSymbol sets the RemainingIterationSymbol, which
// is the symbol within the progress bar that shows that particular
// iteration has not yet completed it's execution.
//
// Default Value: "-"
func (itr *Iterator) SetRemainingIterationSymbol(newSymbol string) {
	itr.Settings.SetRemainingIterationSymbol(newSymbol)
}

// SetLParen sets the symbol to be used to show the start
// of the progress bar.
//
// Default Value: "|"
func (itr *Iterator) SetLParen(newSymbol string) {
	itr.Settings.SetLParen(newSymbol)
}

// SetRParen sets the symbol to be used to show the end
// of the progress bar.
//
// Default Value: "|"
func (itr *Iterator) SetRParen(newSymbol string) {
	itr.Settings.SetRParen(newSymbol)
}

// SetRetain sets whether to clear the progress bar
// from the writer (false) or not (true)
//
// Default Value: true
func (itr *Iterator) SetRetain(value bool) {
	if value {
		itr.Settings.SetSuffix(render.DefaultSuffix)
	} else {
		itr.Settings.SetSuffix("\r\033[K")
	}
}

// SetEqualTo adds an extra step to the stop value
// This is to be used when the for loop uses an 'equals' value
// for the upper limit
func (itr *Iterator) SetEqualTo() {
	if itr.Values.GetIsObject() {
		panic("Cannot use Equal To when creating Pbar from an Object!")
	}

	itr.Values.SetStop(itr.Values.GetStop() + itr.Values.GetStep())
}

// Multi enables multiple progress bars to be displayed at the same time.
// It should be called before Initialize on the nested pbar object
func (itr *Iterator) Multi() {
	if err := itr.render("\n\033[K"); err != nil {
		panic(fmt.Sprintf("Error in rendering: %v", err))
	}

	itr.Settings.SetSuffix("\033[1A")
}

// MultiEnd enables you to escape nicely out of the multiple progress bars.
// If using the multiple option, this is the recommended way to finish.
// Note this should be called after the outer-most loop has completed.
func (itr *Iterator) MultiEnd() {
	if err := itr.render("\033[1B\n"); err != nil {
		panic(fmt.Sprintf("Error in rendering: %v", err))
	}
}

// progress moves the iteration sequence forward by altering the
// CurrentValue inside the iterator object
func (itr *Iterator) progress() error {
	start := itr.Values.GetStart()
	stop := itr.Values.GetStop()
	step := itr.Values.GetStep()
	current := itr.Values.GetCurrent()
	lineSize := itr.Settings.GetLineSize()

	if current < start || current > stop {
		return fmt.Errorf("Current: %f is incorrect. Start: %f; end: %f", current, start, stop)
	}

	bar := itr.formatProgressBar(start, stop, current, lineSize)
	if err := itr.render(bar); err != nil {
		return err
	}

	if current == stop {
		if err := itr.render(itr.Settings.GetSuffix()); err != nil {
			return err
		}
	}

	itr.Values.SetCurrent(current + step)

	return nil
}

// render writes the relevant string to the relevant writer
func (itr *Iterator) render(s string) error {
	if itr.Write == nil {
		return errors.New("Write is nil!")
	}

	if err := itr.Write.WriteString(fmt.Sprintf("\r%s", s)); err != nil {
		return err
	}

	return nil
}

// formatProgressBar creates the progress bar to be displayed
// by the writer. It gathers all the relevant sections from
// the other functions.
func (itr *Iterator) formatProgressBar(start, stop, current float64, lineSize int) string {
	statistics, numStepsCompleted := itr.Values.Statistics(lineSize)
	barString := itr.Settings.CreateBarString(numStepsCompleted)
	speedMeter := itr.Clock.CreateSpeedMeter(start, stop, current)
	progressBar := strings.Join([]string{barString, statistics, speedMeter}, " ")

	return progressBar
}

// createIteratorFromObject creates the iterator object from
// an object value.
func (itr *Iterator) createIteratorFromObject(object interface{}) {
	itr.Values.SetStart(0.0)
	itr.Values.SetStop(float64(reflect.ValueOf(object).Len()))
	itr.Values.SetStep(1.0)
	itr.Values.SetCurrent(0.0)
	itr.Values.SetIsObject(true)
}

// createIteratorFromValues creates an iterator object from a list of numerical
// values.
func (itr *Iterator) createIteratorFromValues(values ...interface{}) {
	floatValues := make([]float64, 0)
	for _, value := range values {
		floatValues = append(floatValues, convertToFloatValue(value))
	}

	switch len(floatValues) {
	case 1:
		itr.Values.SetStop(floatValues[0])
		itr.Values.SetStep(1.0)
	case 2:
		itr.Values.SetStart(floatValues[0])
		itr.Values.SetCurrent(floatValues[0])
		itr.Values.SetStop(floatValues[1])
		itr.Values.SetStep(1.0)
	case 3:
		itr.Values.SetStart(floatValues[0])
		itr.Values.SetCurrent(floatValues[0])
		itr.Values.SetStop(floatValues[1])
		itr.Values.SetStep(floatValues[2])
	}

	itr.Values.SetIsObject(false)
}

// convertToFloatValue converts an interface to a float using
// the reflect package
func convertToFloatValue(value interface{}) float64 {
	floatValue := reflect.Indirect(reflect.ValueOf(value)).Convert(reflect.TypeOf(*new(float64))).Float()

	return floatValue
}

// isObject examines if the interface values are indeed an Object
// of the correct type. An error is raised if the values are not
// all of the same type. A separate error is raised if the value
// is not a valid object or number.
func isObject(values ...interface{}) (bool, error) {
	var isObject bool
	var err error

	for index, v := range values {
		if isConvertibleToFloat(v) {
			if index >= 1 && isObject {
				err = errors.New("Mixed value types!")
				break
			}
			isObject = false
		} else if isValidObject(v) {
			if index >= 1 && !isObject {
				err = errors.New("Mixed value types!")
				break
			}
			isObject = true
		} else {
			err = fmt.Errorf("Type: %v is not as number or valid object!", reflect.TypeOf(v))
			break
		}
	}

	return isObject, err
}

// checkValues checks that the user-passed values are of the correct type.
func checkValues(isObject bool, values ...interface{}) error {
	if isObject && len(values) != 1 {
		return errors.New("Must only pass a single valid object!")
	}

	if !isObject && (len(values) < 1 || len(values) > 3) {
		return errors.New("Expect 1, 2 or 3 parameters (Stop); (Start, Stop) or (Start, Stop, Step)")
	}

	if !isObject && len(values) > 1 && convertToFloatValue(values[0]) > convertToFloatValue(values[1]) {
		return fmt.Errorf("Start value (%v) is greater than Stop value (%v)!",
			convertToFloatValue(values[0]),
			convertToFloatValue(values[1]),
		)
	}

	return nil
}

// isConvertibleToFloat checks to see if the interface value can
// be converted to a float64 value.
func isConvertibleToFloat(v interface{}) bool {
	return reflect.TypeOf(v).ConvertibleTo(reflect.TypeOf(*new(float64)))
}

// isValidObject assesses if the user-passed value is a valid object.
// Note that a valid object can be a populated object of type:
// Array, slice, map, Buffered channel or string
func isValidObject(v interface{}) bool {
	validObj := true
	defer func() {
		if r := recover(); r != nil {
			validObj = false
		}
	}()
	reflect.ValueOf(v).Len()

	return validObj
}
