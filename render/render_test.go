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
	"fmt"
	"testing"

	"github.com/kinsey40/pbar/render"
	"github.com/stretchr/testify/assert"
)

func TestMakeRenderObject(t *testing.T) {
	testCases := []struct {
		start float64
		stop  float64
		step  float64
	}{
		{0.0, 0.0, 0.0},
	}

	for _, testCase := range testCases {
		renderObj := render.MakeRenderObject(testCase.start, testCase.stop, testCase.step)
		assert.Equal(t, testCase.start, renderObj.StartValue, fmt.Sprintf(""))
		assert.Equal(t, testCase.stop, renderObj.EndValue, fmt.Sprintf(""))
		assert.Equal(t, testCase.step, renderObj.StepValue, fmt.Sprintf(""))
		assert.Equal(t, testCase.start, renderObj.CurrentValue, fmt.Sprintf(""))

		assert.Equal(t, render.DefaultFinishedIterationSymbol, renderObj.FinishedIterationSymbol, fmt.Sprintf(""))
		assert.Equal(t, render.DefaultCurrentIterationSymbol, renderObj.CurrentIterationSymbol, fmt.Sprintf(""))
		assert.Equal(t, render.DefaultRemainingIterationSymbol, renderObj.RemainingIterationSymbol, fmt.Sprintf(""))
		assert.Equal(t, render.DefaultLParen, renderObj.LParen, fmt.Sprintf(""))
		assert.Equal(t, render.DefaultRParen, renderObj.RParen, fmt.Sprintf(""))
		assert.Equal(t, render.DefaultMaxLineSize, renderObj.MaxLineSize, fmt.Sprintf(""))
		assert.Equal(t, render.DefaultLineSize, renderObj.LineSize, fmt.Sprintf(""))

		assert.Zero(t, renderObj.Description, fmt.Sprintf(""))
	}
}

func TestUpdate(t *testing.T) {
	// REQUIRES MOCKING
}

func TestInitialize(t *testing.T) {
	// mockCtrl := gomock.NewController(t)
	// defer mockCtrl.Finish()

	// mockRenderInterface := mocks.NewMockRenderInterface(mockCtrl)
	// renderObj := render.MakeRenderObject(0.0, 0.0, 0.0)
	// testCases := []struct {
	// 	timeValue    time.Time
	// 	updateReturn error
	// 	startValue   float64
	// }{
	// 	{time.Now(), nil, 0.0},
	// }

	// for _, testCase := range testCases {
	// 	mockRenderInterface.EXPECT().Update(testCase.startValue).Return(testCase.updateReturn).Times(1)
	// 	err := renderObj.Initialize(testCase.timeValue)

	// 	// err := mockRenderInterface.Initialize(testCase.timeValue)
	// 	// underlyingStruct := mockRenderInterface.(*render.RenderObject)

	// 	assert.Equal(t,
	// 		testCase.updateReturn,
	// 		err,
	// 		fmt.Sprintf("Errors not equal, expected: %v; got: %v", testCase.updateReturn, err))

	// 	assert.Equal(t,
	// 		testCase.timeValue,
	// 		renderObj.StartTime,
	// 		fmt.Sprintf("Times not equal, expected: %v; got: %v", testCase.timeValue, renderObj.StartTime))

	// }
	// // This is fine, but need to mock out the call to Update, as this may
	// // or may not return an error and we need to check that if it does,
	// // then the initialize function will also return that same error.

	// // r := render.MakeRenderObject(0.0, 0.0, 0.0)
	// // for i := 0; i < 10; i++ {
	// // 	timeVal := time.Now()
	// // 	err := r.Initialize(timeVal)

	// // 	assert.Equal(t,
	// // 		timeVal,
	// // 		r.StartTime,
	// // 		fmt.Sprintf("Times not equal, expected: %v; got: %v", timeVal, r.StartTime))
	// // }
}

func TestDescription(t *testing.T) {
	// 	mockCtrl := gomock.NewController(t)
	// 	defer mockCtrl.Finish()

	// 	mockPbarInterface := mocks.NewMockPbarInterface(mockCtrl)
	// 	testCases := []struct {
	// 		description    string
	// 		expectedReturn string
	// 		correct        bool
	// 	}{
	// 		{"Testing", "Testing", true},
	// 		{"Another Test", "Another Test", true},
	// 		{"Test", "Incorrect return", false},
	// 	}

	// 	for _, testCase := range testCases {
	// 		mockPbarInterface.EXPECT().SetDescription(testCase.description).Return().Times(1)
	// 		mockPbarInterface.EXPECT().GetDescription().Return(testCase.description).Times(1)

	// 		mockPbarInterface.SetDescription(testCase.description)
	// 		desc := mockPbarInterface.GetDescription()

	// 		if desc != testCase.expectedReturn && testCase.correct {
	// 			t.Fail(fmt.Sprintf("Incorrect description returned: %v, expected: %v", desc, testCase.expectedReturn))
	// 		}

	// 		if desc == testCase.expectedReturn && !testCase.correct {
	// 			t.Fail(fmt.Sprintf("Incorrect description returned: %v, expected: %v", desc, testCase.expectedReturn))
	// 		}
	// 	}
	// }
}
