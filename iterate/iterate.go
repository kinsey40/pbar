/*



 */

package iterate

import (
	"errors"
	"reflect"

	"github.com/kinsey40/tqdm/render"
)

var StopIteration = errors.New("Stop Iteration error")

type Iterator struct {
	start     float64
	stop      float64
	step      float64
	current   float64
	renderObj *render.RenderObject
}

func IteratorFromObject(object interface{}) (itr *Iterator, err error) {
	itr.start = 0.0
	itr.stop = float64(reflect.ValueOf(object).Len())
	itr.step = 1.0
	itr.current = 0.0
	itr.renderObj, err = render.MakeRenderObject(float64(reflect.ValueOf(object).Len()))

	return itr, err
}

func IteratorFromValues(start, stop, step float64) (itr *Iterator, err error) {
	itr.start = start
	itr.stop = stop
	itr.step = step
	itr.current = start
	itr.renderObj, err = render.MakeRenderObject(stop)

	return itr, err
}

func (itr *Iterator) Update() error {
	itr.current += itr.step

	if itr.current >= itr.stop {
		return StopIteration
	}

	return nil
}
