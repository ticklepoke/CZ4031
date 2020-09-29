package main

import (
	"fmt"
	"strconv"

	"github.com/ticklepoke/CZ4031/bptree"
)

func main() {

	t := bptree.NewTree()

	// Insert 1, 4, 7, 10, 17, 21, 31, 25, 19, 20, 28, 42
	// lecture example

	arr := [12]int{1, 4, 7, 10, 17, 21, 31, 25, 19, 20, 28, 42}

	for _, num := range arr {
		err := t.Insert(num, []byte("hello friend"+strconv.Itoa(num)))
		if err != nil {
			fmt.Printf("error: %s\n\n", err)
		}
	}

	t.PrintTree()

	r, err := t.Find(7, true)
	if err != nil {
		fmt.Printf("error: %s\n\n", err)
	}

	fmt.Printf("%s\n\n", r.Value)

}
