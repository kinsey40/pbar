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
 * File:   doc.go
 * Author: kinsey40
 *
 * Created on 13 January 2019, 11:05
 *
 * Pbar is a simple, flexible, easy-to-use progress bar for the Go language.
 * To create and use a progress bar you first call the Pbar() function, passing
 * in the relevant values (either a valid object or numeric values).
 *
 * Extra settings can then be easily set (such as setting a description to allow
 * you to label the progress bar).
 *
 * Immediately before the progress bar, the Initialize function on the progress bar
 * should be called.

 * The Update function should be called at the end of the for-loop. An example is
 * provided below, many more examples can be found in the examples/example.go file
 * in the root of the project.
 *
 * p, err := Pbar(0, 10, 1)		// start = 0; stop = 10; step = 1
 * if err != nil {
 *     panic(err)
 * }
 *
 * p.SetDescription("Pbar One")
 * p.Initialize()
 * for i := 0; i < 10; i++ {
 *     time.Sleep(time.Second * 1)
 *     p.Update()
 * }
 *
 */

package pbar
