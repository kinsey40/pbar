package main

import (
	"time"

	"github.com/kinsey40/tqdm"
)

func iterateUsingArray() {
	x := [5]int{1, 2, 3, 4, 5}
	t, err := tqdm.Tqdm(x)

	for range x {
		if err == nil {
			t.Update()
		}

		time.Sleep(time.Second * 1)
	}
	t.Update()
}

func iterateUsingValues() {
	t, err := tqdm.Tqdm(10)

	for i := 0; i < 10; i++ {
		t.Update()
		time.Sleep(time.Second * 1)
	}
	t.Update()
}

func multipleProgressBars() {
	t := tqdm.Tqdm(10)

	for i := 0; i < 10; i++ {
		t.Update()
		time.Sleep(time.Second * 1)
		tq := tqdm.Tqdm(5)
		for j := 0; j < 5; j++ {
			tq.Update()
			time.Sleep(time.Second * 1)
		}
		tq.Update()
	}
	t.Update()
}

func main() {
	// multipleProgressBars()
	iterateUsingArray()
	iterateUsingValues()
}
