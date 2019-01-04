/*


 */

package tqdm

import (
	"errors"

	"github.com/kinsey40/tqdm/iterate"
)

var IncorrectNumberOfParameters = errors.New("Expect 1, 2 or 3 parameters (stop); (start, stop) or (start, stop, step)")
var IncorrectValueType := errors.New("Value is not in the expected acceptable types!")
var acceptableTypes = []reflect.Type{reflect.Int8, reflect.Int16} // Need to put in the other types

type Tqdm struct {
	iterator *iterate.Iterator
}

type startStopStepConversion struct {
	start	interface{}
	stop	interface{}
	step	interface{}
	startConverted	float64
	stopConverted	float64
	stepConverted	float64
}

func TqdmFromObject(obj interface{}) (tqdm *Tqdm, err error) {
	tqdm.iterator, err = iterate.IteratorFromObject(obj)

	return tqdm, err
}

func TqdmFromValues(values ...interface{}) (tqdm *Tqdm, err error) {
	startStopStepValues := new(startStopStepConversion)
	
	switch len(values) {
	case 1:
		startStopStepValues.stop = values[0]
	case 2:
		startStopStepValues.start = values[0]
		startStopStepValues.stop = values[1]
	case 3:
		startStopStepValues.start = values[0]
		startStopStepValues.stop = values[1]
		startStopStepValues.step = values[2]
	default:
		err = IncorrectNumberOfParameters
	}

	err = startStopStepValues.convertToFloatValues()
	tqdm.iterator, err = iterate.IteratorFromValues(startStopStepValues.start, startStopStepValues.stop, startStopStepValues.step)

	return tqdm, err
}

func (conversionValues *startStopStepConversion) convertToFloatValues() (err error) {
	if conversionValues.start != nil {
		conversionValues.startConverted, err = convertFloatValue(interfaces.start)
	} else {
		conversionValues.startConverted = 0.0
	}

	if conversionValues.step != nil {
		conversionValues.stepConverted, err = convertFloatValue(interfaces.step)
	} else {
		conversionValues.stepConverted = 1.0
	}

	if conversionValues.stop != nil {
		conversionValues.stopConverted, err = convertFloatValue(interfaces.stop)
	} else {
		// Raise an error
	}

	return err
}

func convertFloatValue(value interface{}) (float64, error) {
	var p *float64
	
	for _, acceptableType := range acceptableTypes {
		if reflect.ValueOf(value).Type().Kind() == acceptableType {
			newValue = float64(reflect.ValueOf(value))
			p = &newValue
		}
	}

	if p != nil {
		return *p, nil
	} else {
		return 0.0, IncorrectValueType
	}
}
