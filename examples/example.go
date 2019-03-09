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
 * File:   example.go
 * Author: kinsey40
 *
 * Created on 13 January 2019, 11:05
 *
 * This file contains lots of examples for how you use the Pbar package in
 * a variety of scenarios. The user is encouraged to copy these examples!
 * Note that this file does not contain functionality present within the Pbar
 * module and is included for user convience.
 *
 * Depending on the desired behaviour, the user can check the errors individually
 * outputted by the pbar operations. Alternatively, as Pbar is generally not used
 * as the main part of an application, errors can be left to silently fail.
 *
 */

package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/kinsey40/pbar"
)

// Create a Pbar object for iteration over an array
func iterateUsingArray() {
	x := [5]int{1, 2, 3, 4, 5}
	p, err := pbar.Pbar(x)
	if err != nil {
		panic(err)
	}

	p.SetDescription("Array")
	p.Initialize()
	for range x {
		time.Sleep(time.Millisecond * 500)
		p.Update()
	}
}

// Create a Pbar object for iteration over a string
func iterateUsingString() {
	x := "abcde"
	p, err := pbar.Pbar(x)
	if err != nil {
		panic(err)
	}

	p.SetDescription("String")
	p.Initialize()
	for range x {
		time.Sleep(time.Millisecond * 500)
		p.Update()
	}
}

// Create a Pbar object for iteration over a slice
func iterateUsingSlice() {
	x := []int{1, 2, 3, 4, 5}
	p, err := pbar.Pbar(x)
	if err != nil {
		panic(err)
	}

	p.SetDescription("Slice")
	p.Initialize()
	for range x[:] {
		time.Sleep(time.Millisecond * 500)
		p.Update()
	}
}

// Create a Pbar object for iteration over a buffered channel
func iterateUsingChannel() {
	size := 5
	x := make(chan int, size)
	for index := 0; index < size; index++ {
		x <- index
	}

	close(x)
	p, err := pbar.Pbar(x)
	if err != nil {
		panic(err)
	}

	p.SetDescription("Channel")
	p.Initialize()
	for range x {
		time.Sleep(time.Millisecond * 500)
		p.Update()
	}
}

// Create a Pbar object for iteration over a map
func iterateUsingMap() {
	x := map[string]string{"1": "a", "2": "b", "3": "c", "4": "d", "5": "e"}
	p, err := pbar.Pbar(x)
	if err != nil {
		panic(err)
	}

	p.SetDescription("Map")
	p.Initialize()
	for range x {
		time.Sleep(time.Millisecond * 500)
		p.Update()
	}
}

// Create a Pbar object for iteration over values
func iterateUsingValues() {
	p, err := pbar.Pbar(5)
	if err != nil {
		panic(err)
	}

	p.SetDescription("Values")
	p.Initialize()
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 500)
		p.Update()
	}
}

func multipleProgressBars() {
	x := []int{1, 2, 3}
	y := []int{1, 2, 3}
	z := []int{1, 2, 3}

	p, _ := pbar.Pbar(x)
	p.SetDescription("First")
	p.Initialize()
	for range x {
		pb, _ := pbar.Pbar(z)
		pb.SetDescription("Second")
		pb.Multi()
		pb.Initialize()
		for range z {
			pba, _ := pbar.Pbar(y)
			pba.SetDescription("Third")
			pba.Multi()
			pba.Initialize()
			for range y {
				time.Sleep(time.Millisecond * 500)
				pba.Update()
			}
			pb.Update()
		}
		p.Update()
	}

	p.MultiEnd()
}

// Threaded bars currently does not work correctly. See the Issues board.
func threadedBars() {
	var wg sync.WaitGroup
	x := []int{1, 2, 3, 4, 5}

	wg.Add(1)
	go func() {
		defer wg.Done()
		p, _ := pbar.Pbar(x)
		p.Initialize()
		for range x {
			time.Sleep(time.Millisecond * 500)
			p.Update()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		p, _ := pbar.Pbar(x)
		p.Initialize()
		for range x {
			time.Sleep(time.Millisecond * 2)
			p.Update()
		}
	}()

	wg.Wait()
}

func main() {
	iterateUsingArray()
	iterateUsingString()
	iterateUsingMap()
	iterateUsingChannel()
	iterateUsingSlice()
	iterateUsingValues()

	fmt.Println("\nUsing Multiple Progress Bars:")
	multipleProgressBars()
	// threadedBars()
}
