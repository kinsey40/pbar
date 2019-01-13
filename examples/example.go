package main

import (
	"time"

	"github.com/kinsey40/tqdm"
)

func main() {
	x := [5]int{1, 2, 3, 4, 5}
	t := tqdm.Tqdm(x)

	for _, value := range x {
		t.Update()
		value += 1
		time.Sleep(time.Second * 1)
	}
}
