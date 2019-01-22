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
 * File:   pbar_internal.go
 * Author: kinsey40
 *
 * Created on 19 January 2019, 00:04
 *
 * An internals file, which contains a set of useful functions needed for
 * checking the various values passed by the user.
 *
 */

package pbar

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/kinsey40/pbar/render"
)

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
			err = errors.New(fmt.Sprintf("Type: %v is not as number or valid object!", reflect.TypeOf(v)))
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
		return errors.New("Expect 1, 2 or 3 parameters (stop); (start, stop) or (start, stop, step)")
	}

	if !isObject && len(values) > 1 && convertToFloatValue(values[0]) < convertToFloatValue(values[1]) {
		return errors.New(
			fmt.Sprintf("Start value (%v) is less than stop value (%v)!",
				convertToFloatValue(values[0]),
				convertToFloatValue(values[1])))
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
	itr.start = 0.0
	itr.stop = float64(reflect.ValueOf(object).Len())
	itr.step = 1.0
	itr.current = 0.0
	itr.renderObject = render.MakeRenderObject(itr.start, itr.stop, itr.step)

	return itr
}

func createIteratorFromValues(values ...interface{}) *iterator {
	itr := new(iterator)
	itr.step = 1.0
	floatValues := make([]float64, 0)

	for _, value := range values {
		floatValues = append(floatValues, convertToFloatValue(value))
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
	}

	itr.renderObject = render.MakeRenderObject(itr.start, itr.stop, itr.step)

	return itr
}
