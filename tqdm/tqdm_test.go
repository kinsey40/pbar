/*


 */

package tqdm_test

import (
	"testing"

	. "github.com/kinsey40/tqdm/tqdm"
	"github.com/stretchr/testify/assert"
)

func TestAddValue(t *testing.T) {
	testCases := []struct {
		valueOne      int
		valueTwo      int
		expectedValue int
		shouldBeEqual bool
	}{
		{1, 3, 4, true},
		{10, 5, 15, true},
		{10, 6, 21, false},
	}

	for _, tCase := range testCases {
		returnValue := AddValues(tCase.valueOne, tCase.valueTwo)
		if tCase.shouldBeEqual {
			assert.Equal(t, returnValue, tCase.expectedValue)
		} else {
			assert.NotEqual(t, returnValue, tCase.expectedValue)
		}
	}
}
