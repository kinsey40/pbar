/*


 */

package tqdm

import (
	"github.com/kinsey40/tqdm/iterate"
)

type tqdmObject struct {
	iterator    *iterate.Iterator
	description string
	retain      bool
}

func Tqdm(values ...interface{}) *tqdmObject {
	tqdmObj := new(tqdmObject)
	tqdmObj.description = ""
	tqdmObj.retain = false

	if itr, err := iterate.CreateIterator(values); err != nil {
		panic(err)
	} else {
		tqdmObj.iterator = itr
	}

	return tqdmObj
}

func (tqdmObj *tqdmObject) Update() {
	tqdmObj.iterator.Update()
}

func (tqdmObj *tqdmObject) SetDescrition(description string) {
	tqdmObj.description = description
}

func (tqdmObj *tqdmObject) GetDescrition() string {
	return tqdmObj.description
}

func (tqdmObj *tqdmObject) SetRetain(retain bool) {
	tqdmObj.retain = retain
}

func (tqdmObj *tqdmObject) GetRetain() bool {
	return tqdmObj.retain
}
