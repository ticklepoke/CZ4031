package main

import (
	"fmt"
	"strconv"

	"github.com/ticklepoke/CZ4031/bptree"
)

func main() {
	// key := 1
	// value := []byte("hello friend")

	t := bptree.NewTree()

	// err := t.Insert(key, value)
	// if err != nil {
	// 	fmt.Printf("error: %s\n\n", err)
	// }

	// r, err := t.Find(key, true)
	// if err != nil {
	// 	fmt.Printf("error: %s\n\n", err)
	// }

	// fmt.Printf("%s\n\n", r.Value)
	// Insert 1, 4, 7, 10, 17, 21, 31, 25, 19, 20, 28, 42

	arr := [12]int{1, 4, 7, 10, 17, 21, 31, 25, 19, 20, 28, 42}

	for _, num := range arr {
		err := t.Insert(num, []byte("hello friend"+strconv.Itoa(num)))
		if err != nil {
			fmt.Printf("error: %s\n\n", err)
		}
	}

	t.PrintTree()

}
