/*


 */

package render

import "io"

var (
	iterationFinishedSymbol  = "#"
	remainingIterationSymbol = "-"
	lineSize                 = 15
	lParen                   = "|"
	rParen                   = "|"
)

type RenderObject struct {
	w            io.Writer
	currentValue float64
	endIteration float64
}

func MakeRenderObject(endValue float64) *RenderObject {
	renderObj := new(RenderObject)
	renderObj.currentValue = 0.0
	renderObj.endIteration = endValue

	return renderObj
}

// func formatProgressBar()
