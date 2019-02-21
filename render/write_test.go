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
 * File:   write_test.go
 * Author: kinsey40
 *
 * Created on 13 January 2019, 11:05
 *
 * Test file for write.go
 *
 */

package render_test

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/kinsey40/pbar/render"
	"github.com/stretchr/testify/assert"
)

func TestNewWrite(t *testing.T) {
	w := render.NewWrite(os.Stdout)

	message := fmt.Sprintf("NewWrite type does not match expected: %v; got: %v", (*render.Write)(nil), w)
	assert.Implements(t, (*render.Write)(nil), w, message)
}

func TestWriteString(t *testing.T) {
	testCases := []struct {
		buffer *bytes.Buffer
		str    string
	}{
		{new(bytes.Buffer), "Hello"},
	}

	for _, testCase := range testCases {
		w := &render.Writing{W: testCase.buffer}
		w.WriteString(testCase.str)
		got := testCase.buffer.String()

		message := fmt.Sprintf("Output and input strings are different expected: %v; got: %v", testCase.str, got)
		assert.Equal(t, testCase.str, got, message)
	}

}
