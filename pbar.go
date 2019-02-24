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

	"github.com/kinsey40/pbar/render"
)

type Iterate interface {
	Initialize() error
	Update() error
	SetDescription(string)
	SetFinishedIterationSymbol(string)
	SetCurrentIterationSymbol(string)
	SetRemainingIterationSymbol(string)
	SetLParen(string)
	SetRParen(string)
	progress() error
}

// iterator object stores the relevant parameters
// associated with the progress bar, this is returned
// by the Pbar function.
type Iterator struct {
	Start        float64
	Stop         float64
	Step         float64
	Current      float64
	Timer        render.Clock
	Settings     render.Settings
	RenderObject render.Render
}

// makeIteratorObject creates an Iterate interface
func makeIteratorObject() *Iterator {
	itr := new(Iterator)
	itr.RenderObject = render.MakeRenderObject(itr.Start, itr.Stop, itr.Step)

	return itr
}

// Pbar creates a progress bar from the inputted values or object.
// The user can pass either a valid object or list of numbers
// (of type: float32, float64, int8, int16, int32, int64 or int).
// Valid objects should be passed as single values, a valid object
// (of type: array, slice, string, map or buffered channel).
func Pbar(values ...interface{}) (Iterate, error) {
	itr := makeIteratorObject()
	itr.Timer = render.NewClock()
	itr.Settings = render.NewSettings()

	isObject, err := isObject(values...)
	if err != nil {
		return nil, err
	}

	if err = checkValues(isObject, values...); err != nil {
		return nil, err
	}

	if isObject {
		itr = createIteratorFromObject(values[0])
	} else {
		itr = createIteratorFromValues(values...)
	}

	return itr, err
}

// Initialize sets the internal timer to start,
// enabling output relating to the time taken for
// iterations within the progress bar.
func (itr *Iterator) Initialize() error {
	itr.Timer.SetStart(itr.Timer.Now())
	itr.RenderObject.Initialize(itr.Timer, itr.Settings)
	if err := itr.Update(); err != nil {
		return err
	}

	return nil
}

// Update moves the iteration forward by one step. This should
// be performed at the end of the iteration sequence
// (i.e. at the end of the for-loop).
func (itr *Iterator) Update() error {
	var err error
	if err = itr.RenderObject.Update(itr.Current); err != nil {
		return err
	}

	err = itr.progress()
	return err
}

// SetDescription sets the Description parameter, which causes the Pbar
// to output a String at the start of the progress bar, effectively
// enabling the progress bars to be named within the output.
// Default Value: ""
func (itr *Iterator) SetDescription(descrip string) {
	itr.Settings.SetDescription(descrip)
}

// SetFinishedIterationSymbol sets the FinishedIterationSymbol, which
// is the symbol within the progress bar that shows that particular
// iteration has completed it's execution.
// Default Value: "#"
func (itr *Iterator) SetFinishedIterationSymbol(newSymbol string) {
	itr.Settings.SetFinishedIterationSymbol(newSymbol)
}

// SetCurrentIterationSymbol sets the CurrentIterationSymbol, which
// is the symbol within the progress bar that shows the iteration
// which is currently being executed.
// Default Value: "#"
func (itr *Iterator) SetCurrentIterationSymbol(newSymbol string) {
	itr.Settings.SetCurrentIterationSymbol(newSymbol)
}

// SetRemainingIterationSymbol sets the RemainingIterationSymbol, which
// is the symbol within the progress bar that shows that particular
// iteration has not yet completed it's execution.
// Default Value: "-"
func (itr *Iterator) SetRemainingIterationSymbol(newSymbol string) {
	itr.Settings.SetRemainingIterationSymbol(newSymbol)
}

// SetLParen sets the symbol to be used to show the start
// of the progress bar.
// Default Value: "|"
func (itr *Iterator) SetLParen(newSymbol string) {
	itr.Settings.SetLParen(newSymbol)
}

// SetRParen sets the symbol to be used to show the end
// of the progress bar.
// Default Value: "|"
func (itr *Iterator) SetRParen(newSymbol string) {
	itr.Settings.SetRParen(newSymbol)
}

// progress moves the iteration sequence forward by altering the
// CurrentValue inside the iterator object
func (itr *Iterator) progress() error {
	itr.Current += itr.Step
	if itr.Current > itr.Stop {
		return errors.New("Stop Iteration error")
	}

	return nil
}

// convertToFloatValue converts an interface to a float using
// the reflect package
func convertToFloatValue(value interface{}) float64 {
	floatValue := reflect.Indirect(reflect.ValueOf(value)).Convert(reflect.TypeOf(*new(float64))).Float()

	return floatValue
}

// isObject examines if the interface values are indeed an Object
// of the correct type. An error is raised if the values are not
// all of the same type. A seperate error is raised if the value
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

// createIteratorFromObject creates the iterator object from
// an object value.
func createIteratorFromObject(object interface{}) *Iterator {
	itr := new(Iterator)
	itr.Start = 0.0
	itr.Stop = float64(reflect.ValueOf(object).Len())
	itr.Step = 1.0
	itr.Current = 0.0
	itr.RenderObject = render.MakeRenderObject(itr.Start, itr.Stop, itr.Step)

	return itr
}

// createIteratorFromValues creates an iterator object from a list of numerical
// values.
func createIteratorFromValues(values ...interface{}) *Iterator {
	itr := new(Iterator)
	itr.Step = 1.0
	floatValues := make([]float64, 0)

	for _, value := range values {
		floatValues = append(floatValues, convertToFloatValue(value))
	}

	switch len(floatValues) {
	case 1:
		itr.Stop = floatValues[0]
	case 2:
		itr.Start = floatValues[0]
		itr.Stop = floatValues[1]
	case 3:
		itr.Start = floatValues[0]
		itr.Stop = floatValues[1]
		itr.Step = floatValues[2]
	}

	itr.Current = itr.Start
	itr.RenderObject = render.MakeRenderObject(itr.Start, itr.Stop, itr.Step)

	return itr
}
