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
 * File:   render_test.go
 * Author: kinsey40
 *
 * Created on 13 January 2019, 11:05
 *
 * The test file for the render package.
 *
 */

package render

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRender(t *testing.T) {
	// mockCtrl := gomock.NewController(t)
	// defer mockCtrl.Finish()

	// mockRenderInterface := mocks.NewMockRenderInterface(mockCtrl)

}

func TestFormatProgressBar(t *testing.T) {

}

func TestFormatSpeedMeter(t *testing.T) {
	// mockCtrl := gomock.NewController(t)
	// defer mockCtrl.Finish()

	// mockRenderInterface := mocks.NewMockRenderInterface(mockCtrl)
	// testCases := []struct {
	// 	progressBar string
	// }{
	// 	{""},
	// }

	// for _, testCase := range testCases {
	// 	mockRenderInterface.EXPECT().formatProgressBar().Return(testCase.progressBar).Times(1)
	// 	message := fmt.Sprintf("Progress bar not expected, expected: %v; got: %v", testCase.progressBar, speedMeter)
	// 	speedMeter := mockRenderInterface.formatProgressBar()
	// 	assert.Equal(t, testCase.progressBar, speedMeter, message)
	// }
}

func TestFormatTime(t *testing.T) {
	testCases := []struct {
		timeValue      time.Duration
		expectedString string
	}{
		{time.Duration(10) * time.Second, "00m:10s"},
	}

	for _, testCase := range testCases {
		returnedString := formatTime(testCase.timeValue)
		message := fmt.Sprintf("Time string incorrect, expected: %v; got: %v", testCase.expectedString, returnedString)
		assert.Equal(t, testCase.expectedString, returnedString, message)
	}
}
