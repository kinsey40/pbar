package main

import (
	"time"

	"github.com/kinsey40/tqdm"
)

func iterateUsingArray() {
	x := [5]int{1, 2, 3, 4, 5}
	t := tqdm.Tqdm(x)

	for range x {
		t.Update()
		time.Sleep(time.Second * 1)
	}
}

func iterateUsingValues() {
	t := tqdm.Tqdm(10)

	for i := 0; i < 10; i++ {
		t.Update()
		time.Sleep(time.Second * 1)
	}
}

func main() {
	iterateUsingArray()
	// iterateUsingValues()
}
