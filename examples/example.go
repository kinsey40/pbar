package main

import (
	"time"

	"github.com/kinsey40/pbar"
)

func iterateUsingArray() {
	x := [5]int{1, 2, 3, 4, 5}
	t, _ := pbar.Pbar(x)

	for range x {
		time.Sleep(time.Second * 1)
		t.Update()
	}
}

func iterateUsingValues() {
	t, _ := pbar.Pbar(10)
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		t.Update()
	}
}

// Need an initialize function which does not take into account the
// start, it merely displays the value to the screen.
// The timer should only start on the first Update.

// func multipleProgressBars() {
// 	t, err := pbar.Pbar(10)

// 	for i := 0; i < 10; i++ {
// 		t.Update()
// 		time.Sleep(time.Second * 1)
// 		tq, err := pbar.Pbar(5)
// 		for j := 0; j < 5; j++ {
// 			tq.Update()
// 			time.Sleep(time.Second * 1)
// 		}
// 		tq.Update()
// 	}
// 	t.Update()
// }

func main() {
	// multipleProgressBars()
	iterateUsingArray()
	iterateUsingValues()
}

/* ******

NOTES:
- Need to finish off the tests, generating the mocks etc.
- Need method by which to do multiple progress bars
- potentially look at changing what it writes to
- Look at using RETAIN as well.
- Then can release v1.0
*/
