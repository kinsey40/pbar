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

package render_test

import (
	"testing"
)

// func TestMakeRenderObject(t *testing.T) {
// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockRenderInterface := mocks.NewMockRenderInterface(mockCtrl)
// 	testCases := []struct {
// 		startValue float64
// 		stopValue  float64
// 		stepValue  float64
// 	}{
// 		{float64(1), float64(3), float64(1)},
// 	}

// 	for _, testCase := range testCases {
// 		renderObj := new(render.RenderObject)
// 		renderObj.StartValue = testCase.startValue

// 		mockRenderInterface.EXPECT().MakeRenderObject(testCase.startValue, testCase.stopValue, testCase.stepValue).Return(renderObj).Times(1)
// 		returnedRenderObj := mockRenderInterface.MakeRenderObject(testCase.startValue, testCase.stopValue, testCase.stepValue)

// 		assert.Equal(t, testCase.startValue, returnedRenderObj.StartValue, fmt.Sprintf("Start values don't match, expected: %v; got: %v", testCase.startValue, returnedRenderObj.StartValue))

// 	}
// }

func TestUpdate(t *testing.T) {

}
