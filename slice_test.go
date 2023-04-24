package slice_test

import (
	"fmt"

	"github.com/aria3ppp/slice"
)

func Example() {
	hotbar := []string{
		"Torch",
		"Rock",
		"Potion",
		"Sword",
		"Shield",
	}

	slc := slice.From(hotbar)

	index := slc.Index("Sword")
	item := slc.Pop(index)
	slc.Push(item, 0)

	for _, v := range slc.Into() {
		fmt.Println(v)
	}
	// Output:
	// Sword
	// Torch
	// Rock
	// Potion
	// Shield
}
