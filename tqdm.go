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
 * The main file for the tqdm package, this enables the creation of the tqdm
 * object. The user can then edit the specific variables associated with the
 * object.
 *
 */

package tqdm

import (
	"github.com/kinsey40/tqdm/iterate"
)

type TqdmInterface interface {
	Update()
	SetDescription(string)
	GetDescription() string
	SetRetain(bool)
	GetRetain() bool
}

type TqdmSettings struct {
	Iterator    *iterate.Iterator
	Description string
	Retain      bool
}

func Tqdm(values ...interface{}) *TqdmSettings {
	tqdmObj := new(TqdmSettings)
	tqdmObj.Description = ""
	tqdmObj.Retain = false

	if itr, err := iterate.CreateIterator(values...); err != nil {
		panic(err)
	} else {
		tqdmObj.Iterator = itr
	}

	return tqdmObj
}

func (tqdmObj *TqdmSettings) Update() {
	err := tqdmObj.Iterator.Update()

	if err != nil && err != iterate.StopIterationError {
		panic(err)
	}
}

func (tqdmObj *TqdmSettings) SetDescription(description string) {
	tqdmObj.Description = description
}

func (tqdmObj *TqdmSettings) GetDescription() string {
	return tqdmObj.Description
}

func (tqdmObj *TqdmSettings) SetRetain(retain bool) {
	tqdmObj.Retain = retain
}

func (tqdmObj *TqdmSettings) GetRetain() bool {
	return tqdmObj.Retain
}
