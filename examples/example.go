package main

import (
	"fmt"
	"time"

	"github.com/kinsey40/pbar"
)

func iterateUsingArray() {
	x := [5]int{1, 2, 3, 4, 5}
	t, err := pbar.Pbar(x)
	if err != nil {
		panic(err)
	}

	t.SetDescription("Array")
	t.Initialize()
	for range x {
		time.Sleep(time.Second * 1)
		t.Update()
	}
}

func iterateUsingString() {
	x := "abcde"
	t, err := pbar.Pbar(x)
	if err != nil {
		panic(err)
	}

	t.SetDescription("String")
	t.Initialize()
	for range x {
		time.Sleep(time.Second * 1)
		t.Update()
	}
}

func iterateUsingSlice() {
	x := []int{1, 2, 3, 4, 5}
	t, err := pbar.Pbar(x)
	if err != nil {
		panic(err)
	}

	t.SetDescription("Slice")
	t.Initialize()
	for range x[:] {
		time.Sleep(time.Second * 1)
		if t != nil {
			t.Update()
		}
	}
}

func iterateUsingChannel() {
	size := 5
	x := make(chan int, size)
	for index := 0; index < size; index++ {
		x <- index
	}

	close(x)
	t, err := pbar.Pbar(x)
	if err != nil {
		panic(err)
	}

	t.SetDescription("Channel")
	t.Initialize()
	for range x {
		time.Sleep(time.Second * 1)
		t.Update()
	}
}

func iterateUsingMap() {
	x := map[string]string{"1": "a", "2": "b", "3": "c", "4": "d", "5": "e"}
	t, err := pbar.Pbar(x)
	if err != nil {
		panic(err)
	}

	t.SetDescription("Map")
	t.Initialize()
	for range x {
		time.Sleep(time.Second * 1)
		t.Update()
	}
}

func iterateUsingValues() {
	t, err := pbar.Pbar(10)
	if err != nil {
		panic(err)
	}

	t.SetDescription("Values")
	t.Initialize()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		t.Update()
	}
}

func multipleIterators() {
	t, err := pbar.Pbar(10)
	p, err := pbar.Pbar(2)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		if t != nil {
			t.Update()
		}

		// for j := 0; j < 2; i++ {
		// 	if p != nil {
		// 		time.Sleep(time.Second * 1)
		// 		p.Update()
		// 	}
		// }
	}

	time.Sleep(time.Second * 1)
	p.Update()
	time.Sleep(time.Second * 1)
	p.Update()
}

func main() {
	// multipleIterators()

	iterateUsingArray()
	iterateUsingString()
	iterateUsingMap()
	iterateUsingChannel()
	iterateUsingSlice()
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
