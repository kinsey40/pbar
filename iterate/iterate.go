/*



 */

package iterate

import (
	"errors"
	"fmt"
	"math"
	"reflect"

	"github.com/kinsey40/tqdm/render"
)

var numberTypes = []reflect.Kind{
	reflect.Int8,
	reflect.Int16,
	reflect.Int32,
	reflect.Int64,
	reflect.Int,
	reflect.Float32,
	reflect.Float64,
}
var objectTypes = []reflect.Kind{
	reflect.Map,
	reflect.Array,
	reflect.Chan,
	reflect.Slice,
	reflect.String,
}
var floatType = reflect.TypeOf(float64(0))

type Iterator struct {
	start     float64
	stop      float64
	step      float64
	current   float64
	renderObj *render.RenderObject
}

func (itr *Iterator) createIteratorFromObject(object interface{}) (err error) {
	itr.start = 0.0
	itr.stop = float64(reflect.ValueOf(object).Len())
	itr.step = 1.0
	itr.current = 0.0
	itr.renderObj = render.MakeRenderObject(float64(reflect.ValueOf(object).Len()))

	return err
}

func (itr *Iterator) createIteratorFromOneValues(values ...interface{}) (err error) {
	itr.start = 0.0
	itr.step = 1.0
	itr.current = 0.0
	itr.renderObj = render.MakeRenderObject(itr.stop)

	if stopValue, err := convertToFloatValue(values[0]); err != nil {
		return err
	} else {
		itr.stop = stopValue
	}

	return nil
}

func (itr *Iterator) createIteratorFromTwoValues(values ...interface{}) (err error) {
	itr.step = 1.0
	itr.current = 0.0
	itr.renderObj = render.MakeRenderObject(itr.stop)

	if startValue, err := convertToFloatValue(values[0]); err != nil {
		return err
	} else {
		itr.start = startValue
	}

	if stopValue, err := convertToFloatValue(values[1]); err != nil {
		return err
	} else {
		itr.stop = stopValue
	}

	if incorrectValues := itr.start > itr.stop; incorrectValues {
		return errors.New("Start value is greater than stop value!")
	}

	return nil
}

func (itr *Iterator) createIteratorFromThreeValues(values ...interface{}) (err error) {
	itr.current = 0.0
	itr.renderObj = render.MakeRenderObject(itr.stop)

	if startValue, err := convertToFloatValue(values[0]); err != nil {
		return err
	} else {
		itr.start = startValue
	}

	if stopValue, err := convertToFloatValue(values[1]); err != nil {
		return err
	} else {
		itr.stop = stopValue
	}

	if stepValue, err := convertToFloatValue(values[2]); err != nil {
		return err
	} else {
		itr.step = stepValue
	}

	if incorrectValues := itr.start > itr.stop; incorrectValues {
		return errors.New("Start value is greater than stop value!")
	}

	return err
}

func checkSameTypes(values ...interface{}) error {
	err := *new(error)
	prevType := *new(reflect.Type)

	for index, value := range values {
		valueType := reflect.ValueOf(value).Type()
		if index > 1 {
			if prevType != valueType {
				err = errors.New(fmt.Sprintf("Value types are not the same: %v and %v", prevType, valueType))
			}
		}
	}

	return err
}

func checkNumbers(values ...interface{}) error {
	for _, value := range values {
		number := isNumber(value)
		if !number {
			return errors.New(fmt.Sprintf("Number is of incoorect type, value: %v, type: %v", reflect.ValueOf(value), reflect.TypeOf(value)))
		}
	}

	return nil
}

func isNumber(value interface{}) bool {
	valueIsNumber := false
	for _, numberType := range numberTypes {
		if reflect.ValueOf(value).Type().Kind() == numberType {
			valueIsNumber = true
		}
	}

	return valueIsNumber
}

func acceptableObject(value interface{}) bool {
	acceptableObj := false
	for _, objectType := range objectTypes {
		if reflect.ValueOf(value).Type().Kind() == objectType {
			acceptableObj = true
		}
	}

	return acceptableObj
}

func convertToFloatValue(value interface{}) (float64, error) {
	newValue := reflect.ValueOf(value)
	newValue = reflect.Indirect(newValue)

	if !newValue.Type().ConvertibleTo(floatType) {
		return math.NaN(), errors.New(fmt.Sprintf("Cannot convert %v to float64", newValue.Type()))
	}

	floatValue := newValue.Convert(floatType)

	return floatValue.Float(), nil
}

func CreateIterator(values ...interface{}) (*Iterator, error) {
	var itr *Iterator
	var err error

	if err = checkSameTypes(values); err != nil {
		return itr, err
	}

	switch len(values) {
	case 1:
		if isNumber(values[0]) {
			err = itr.createIteratorFromObject(values)
		} else if acceptableObject(values[0]) {
			err = itr.createIteratorFromOneValues(values)
		} else {
			err = errors.New(fmt.Sprintf("Incorrect type (%v) for parameter", reflect.TypeOf(values[0])))
		}

	case 2:
		if err = checkNumbers(values); err != nil {
			return itr, err
		}

		err = itr.createIteratorFromTwoValues(values)

	case 3:
		if err = checkNumbers(values); err != nil {
			return itr, err
		}

		err = itr.createIteratorFromThreeValues(values)

	default:
		err = errors.New("Expect 1, 2 or 3 parameters (stop); (start, stop) or (start, stop, step)")
	}

	return itr, err
}

func (itr *Iterator) Update() error {
	itr.current += itr.step

	if itr.current >= itr.stop {
		return errors.New("Stop Iteration error")
	}

	return nil
}
