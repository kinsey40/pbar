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

import "github.com/kinsey40/pbar/iterate"

type PbarInterface interface {
	Update()
	SetDescription(string)
	GetDescription() string
	SetRetain(bool)
	GetRetain() bool
}

type PbarSettings struct {
	Iterator *iterate.Iterator
	// Description string
	// Size        int
	// Retain      bool
}

func Pbar(values ...interface{}) (*PbarSettings, error) {
	var err error
	pbarObj := new(PbarSettings)
	// pbarObj.Description = ""
	// pbarObj.Retain = false

	if itr, err := iterate.CreateIterator(values...); err != nil {
		return nil, err
	} else {
		pbarObj.Iterator = itr
	}

	return pbarObj, err
}

func (pbarObj *PbarSettings) Update() {
	err := pbarObj.Iterator.Update()

	if err != nil && err != iterate.StopIterationError {
		panic(err)
	}
}

// func (pbarObj *PbarSettings) SetDescription(description string) {
// 	pbarObj.Description = description
// 	renderObj, err := pbarObj.Iterator.GetRenderObject()

// 	if err != nil {
// 		panic(err)
// 	}

// 	renderObj.SetPrefix(pbarObj.Description)
// }

// func (pbarObj *PbarSettings) GetDescription() string {
// 	return pbarObj.Description
// }

// func (pbarObj *PbarSettings) SetRetain(retain bool) {
// 	pbarObj.Retain = retain
// }

// func (pbarObj *PbarSettings) GetRetain() bool {
// 	return pbarObj.Retain
// }
