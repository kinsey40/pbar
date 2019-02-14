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
	"time"

	"github.com/kinsey40/pbar/render"
)

type iterator struct {
	Start        float64
	Stop         float64
	Step         float64
	Current      float64
	RenderObject *render.RenderObject
}

func MakeIteratorObject() *iterator {
	itr := new(iterator)
	itr.RenderObject = render.MakeRenderObject(itr.Start, itr.Stop, itr.Step)

	return itr
}

func Pbar(values ...interface{}) (*iterator, error) {
	itr := MakeIteratorObject()
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

func (itr *iterator) Initialize() error {
	itr.RenderObject.Initialize(time.Now())
	if err := itr.Update(); err != nil {
		return err
	}

	return nil
}

func (itr *iterator) Update() error {
	var err error
	if err = itr.RenderObject.Update(itr.Current, time.Now()); err != nil {
		return err
	}

	err = itr.progress()
	return err
}

func (itr *iterator) SetDescription(descrip string) {
	itr.RenderObject.Description = descrip + ": "
}

func (itr *iterator) SetFinishedIterationSymbol(newSymbol string) {
	itr.RenderObject.FinishedIterationSymbol = newSymbol
}

func (itr *iterator) SetCurrentIterationSymbol(newSymbol string) {
	itr.RenderObject.CurrentIterationSymbol = newSymbol
}

func (itr *iterator) SetRemainingIterationSymbol(newSymbol string) {
	itr.RenderObject.RemainingIterationSymbol = newSymbol
}

func (itr *iterator) SetLParen(newSymbol string) {
	itr.RenderObject.LParen = newSymbol
}

func (itr *iterator) SetRParen(newSymbol string) {
	itr.RenderObject.RParen = newSymbol
}

func (itr *iterator) progress() error {
	itr.Current += itr.Step
	if itr.Current > itr.Stop {
		return errors.New("Stop Iteration error")
	}

	return nil
}

func convertToFloatValue(value interface{}) float64 {
	floatValue := reflect.Indirect(reflect.ValueOf(value)).Convert(reflect.TypeOf(*new(float64))).Float()

	return floatValue
}

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

func isConvertibleToFloat(v interface{}) bool {
	return reflect.TypeOf(v).ConvertibleTo(reflect.TypeOf(*new(float64)))
}

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

func createIteratorFromObject(object interface{}) *iterator {
	itr := new(iterator)
	itr.Start = 0.0
	itr.Stop = float64(reflect.ValueOf(object).Len())
	itr.Step = 1.0
	itr.Current = 0.0
	itr.RenderObject = render.MakeRenderObject(itr.Start, itr.Stop, itr.Step)

	return itr
}

func createIteratorFromValues(values ...interface{}) *iterator {
	itr := new(iterator)
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
