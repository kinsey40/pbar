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
 * File:   write.go
 * Author: kinsey40
 *
 * Created on 13 January 2019, 11:05
 *
 * Render performs the acutual rendering of the progress bar onto the
 * terminal display.
 *
 */

package render

import (
	"io"
	"os"
)

// Write wraps the WriteString method
type Write interface {
	WriteString(string) error
}

// writing is struct holding an io.Writer
type Writing struct {
	W io.Writer
}

// NewWrite creates a new Write interface object
func NewWrite(writer io.Writer) Write {
	w := new(Writing)
	w.W = writer

	return w
}

// WriteString writes the given string to the underlying
// io.Writer object
func (w *Writing) WriteString(s string) error {
	_, err := io.WriteString(w.W, s)

	if f, ok := w.W.(*os.File); ok {
		f.Sync()
	}

	return err
}
