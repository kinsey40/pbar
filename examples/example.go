package main

import (
	"fmt"
	"time"

	"github.com/kinsey40/pbar"
)

func iterateUsingArray() {
	x := [5]int{1, 2, 3, 4, 5}
	t, err := pbar.Pbar(x)

	for range x {
		if err == nil {
			t.Update()
		}

		time.Sleep(time.Second * 1)
	}
	t.Update()
}

func iterateUsingValues() {
	t, _ := pbar.Pbar(10)
	// t.Initialize()
	t.Update()
	for i := 0; i < 10; i++ {
		t.Update()
		time.Sleep(time.Second * 1)
	}
	fmt.Println("Finished!")
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
	// iterateUsingArray()
	iterateUsingValues()
}

/* ******

NOTES:
- Can i use embedded interfaces to prevent having to have
multiple Set and Get functions. i.e. the SET and GET functions
for customizables can be set in the render object and use the
interface to inherit those up to pbar
YES can do this

- With embedded interfaces, I need to be able to override the
functions
YES can do this

- Need to testify package and to use assert in the tests
YES can do this
*/
