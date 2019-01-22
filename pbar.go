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

	"github.com/kinsey40/pbar/render"
)

type iteratorInterface interface {
	createIteratorFromObject(...interface{}) error
	createIteratorFromValues(...interface{}) error
	Update() error
}

type iterator struct {
	start        float64
	stop         float64
	step         float64
	current      float64
	rendered     bool
	renderObject *render.RenderObject
}

func Pbar(values ...interface{}) (*iterator, error) {
	var itr *iterator
	isObject, err := isObject(values...)

	if err != nil {
		return nil, err
	}

	if err := checkValues(isObject, values); err != nil {
		return nil, err
	}

	if isObject {
		itr = createIteratorFromObject(values[0])
	} else {
		itr = createIteratorFromValues(values...)
	}

	itr.Update()

	return itr, err
}

func (itr *iterator) Update() error {
	itr.renderObject.Update(itr.current, itr.rendered)

	if itr.current == itr.start && !itr.rendered {
		itr.rendered = true
	}

	itr.current += itr.step
	if itr.current > itr.stop {
		return errors.New("Stop Iteration error")
	}

	return nil
}

func (itr *iterator) SetDescription(descrip string) {
	itr.renderObject.Description = descrip + ": "
}

func (itr *iterator) GetDescription() string {
	return itr.renderObject.Description
}

func (itr *iterator) SetIterationFinishedSymbol(newSymbol string) {
	itr.renderObject.IterationFinishedSymbol = newSymbol
}

func (itr *iterator) GetIterationFinishedSymbol() string {
	return itr.renderObject.IterationFinishedSymbol
}

func (itr *iterator) SetCurrentIterationSymbol(newSymbol string) {
	itr.renderObject.CurrentIterationSymbol = newSymbol
}

func (itr *iterator) GetCurrentIterationSymbol() string {
	return itr.renderObject.CurrentIterationSymbol
}

func (itr *iterator) SetRemainingIterationSymbol(newSymbol string) {
	itr.renderObject.RemainingIterationSymbol = newSymbol
}

func (itr *iterator) GetRemainingIterationSymbol() string {
	return itr.renderObject.RemainingIterationSymbol
}

func (itr *iterator) SetLParen(newSymbol string) {
	itr.renderObject.LParen = newSymbol
}

func (itr *iterator) GetLParen() string {
	return itr.renderObject.LParen
}

func (itr *iterator) SetRParen(newSymbol string) {
	itr.renderObject.RParen = newSymbol
}

func (itr *iterator) GetRParen() string {
	return itr.renderObject.RParen
}
