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
 * File:   iterate.go
 * Author: kinsey40
 *
 * Created on 13 January 2019, 11:05
 *
 * The iterate package enables the user to create an iterate object which will
 * call updates to the render object.
 *
 */

package iterate

import (
	"errors"
	"fmt"
	"math"
	"reflect"

	"github.com/kinsey40/tqdm/render"
)

var numberTypes = []reflect.Kind{
	reflect.Int8,
	reflect.Int16,
	reflect.Int32,
	reflect.Int64,
	reflect.Int,
	reflect.Float32,
	reflect.Float64,
}
var objectTypes = []reflect.Kind{
	reflect.Map,
	reflect.Array,
	reflect.Chan,
	reflect.Slice,
	reflect.String,
}
var floatType = reflect.TypeOf(float64(0))
var StopIterationError = errors.New("Stop Iteration error")

type IteratorInterface interface {
	createIteratorFromObject(...interface{}) error
	createIteratorFromValues(...interface{}) error
	Update() error
}

type Iterator struct {
	start     float64
	stop      float64
	step      float64
	current   float64
	renderObj *render.RenderObject
}

func (itr *Iterator) createIteratorFromObject(values ...interface{}) (err error) {
	object := values[0]

	itr.start = 0.0
	itr.stop = float64(reflect.ValueOf(object).Len())
	itr.step = 1.0
	itr.current = 0.0
	itr.renderObj = render.MakeRenderObject(itr.start, itr.stop, itr.step)

	return err
}

func (itr *Iterator) createIteratorFromValues(values ...interface{}) (err error) {
	itr.step = 1.0
	floatValues := make([]float64, 0)

	for _, value := range values {
		if floatValue, err := convertToFloatValue(value); err != nil {
			return err
		} else {
			floatValues = append(floatValues, floatValue)
		}
	}

	switch len(floatValues) {
	case 1:
		itr.stop = floatValues[0]
	case 2:
		itr.start = floatValues[0]
		itr.stop = floatValues[1]
	case 3:
		itr.start = floatValues[0]
		itr.stop = floatValues[1]
		itr.step = floatValues[2]
	default:
		return errors.New(fmt.Sprintf("Values have incorrect length: %d, expect length of 1, 2, or 3", len(values)))
	}
	itr.renderObj = render.MakeRenderObject(itr.start, itr.stop, itr.step)

	return nil
}

func checkTypes(values ...interface{}) (reflect.Type, error) {
	err := *new(error)
	prevType := *new(reflect.Type)

	for index, value := range values {
		valueType := reflect.ValueOf(value).Type()
		if index > 1 {
			if prevType != valueType {
				err = errors.New(fmt.Sprintf("Value types are not the same: %v and %v", prevType, valueType))
			}
		}
		prevType = valueType
	}

	return prevType, err
}

func convertToFloatValue(value interface{}) (float64, error) {
	newValue := reflect.ValueOf(value)
	newValue = reflect.Indirect(newValue)

	if !newValue.Type().ConvertibleTo(floatType) {
		return math.NaN(), errors.New(fmt.Sprintf("Cannot convert %v to float64", newValue.Type()))
	}

	floatValue := newValue.Convert(floatType)

	return floatValue.Float(), nil
}

func isNumber(type_ reflect.Type) bool {
	isNumber := false
	for _, numberType := range numberTypes {
		if type_.Kind() == numberType {
			isNumber = true
		}
	}

	return isNumber
}

func isObject(type_ reflect.Type) bool {
	isObj := false
	for _, objectType := range objectTypes {
		if type_.Kind() == objectType {
			isObj = true
		}
	}

	return isObj
}

func objectOrNumber(type_ reflect.Type, values ...interface{}) (bool, error) {
	if isNum := isNumber(type_); isNum {
		err := checkSize(false, values...)
		return true, err
	}

	if isObj := isObject(type_); isObj {
		err := checkSize(true, values...)
		return false, err
	}

	return false, errors.New(fmt.Sprintf("Type: %v is not as expected!", type_))
}

func checkSize(isObject bool, values ...interface{}) error {
	if isObject && len(values) != 1 {
		return errors.New("Must only pass a single valid object!")
	}

	if !isObject && (len(values) < 1 || len(values) > 3) {
		return errors.New("Expect 1, 2 or 3 parameters (stop); (start, stop) or (start, stop, step)")
	}

	return nil
}

func CreateIterator(values ...interface{}) (*Iterator, error) {
	var err error
	var type_ reflect.Type
	itr := new(Iterator)

	if type_, err = checkTypes(values...); err != nil {
		return itr, err
	}

	if num, err := objectOrNumber(type_, values...); err != nil {
		return itr, err
	} else if num {
		itr.createIteratorFromValues(values...)
	} else {
		itr.createIteratorFromObject(values...)
	}

	if itr.start > itr.stop {
		return itr, errors.New(fmt.Sprintf("Start value (%v) is less than stop value (%v)!", itr.start, itr.stop))
	}

	return itr, err
}

func (itr *Iterator) Update() error {
	itr.current += itr.step
	itr.renderObj.Update(itr.current)

	if itr.current > itr.stop {
		return StopIterationError
	}

	return nil
}
