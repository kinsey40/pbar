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

var StopIterationError = errors.New("Stop Iteration error")

type IteratorInterface interface {
	createIteratorFromObject(...interface{}) error
	createIteratorFromValues(...interface{}) error
	Update() error
}

type Iterator struct {
	Start   float64
	Stop    float64
	Step    float64
	Current float64
	*render.RenderObject
}

func (itr *Iterator) createIteratorFromObject(object interface{}) (err error) {
	itr.Start = 0.0
	itr.Stop = float64(reflect.ValueOf(object).Len())
	itr.Step = 1.0
	itr.Current = 0.0

	return err
}

func (itr *Iterator) createIteratorFromValues(values ...interface{}) (err error) {
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
	default:
		return errors.New(fmt.Sprintf("Values have incorrect length: %d, expect length of 1, 2, or 3", len(values)))
	}

	return nil
}

func (itr *Iterator) Update() error {
	itr.RenderObject.Update(itr.Current)
	itr.Current += itr.Step

	if itr.Current > itr.Stop {
		return StopIterationError
	}

	return nil
}

func Pbar(values ...interface{}) (*Iterator, error) {
	itr := new(Iterator)
	isObject, err := isObject(values...)

	if err != nil {
		return itr, err
	}

	if err := checkSize(isObject, values); err != nil {
		return itr, err
	}

	if isObject {
		itr.createIteratorFromObject(values[0])
	} else {
		itr.createIteratorFromValues(values...)
	}

	itr.RenderObject = render.MakeRenderObject(itr.Start, itr.Stop, itr.Step)
	if itr.Start > itr.Stop {
		return itr, errors.New(fmt.Sprintf("Start value (%v) is less than stop value (%v)!", itr.Start, itr.Stop))
	}

	return itr, err
}
