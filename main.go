package main

import (
	"fmt"
	"strconv"

	"github.com/ticklepoke/CZ4031/bptree"
)

func main() {

	t := bptree.NewTree(4)
	// b := blockmanager.InitializeBlockManager(100)

	arr := []float64{1, 4, 7, 10, 17, 21, 31, 25, 19, 20, 28, 42}
	for _, num := range arr {
		temp := strconv.Itoa(int(num))
		// addr := b.InsertRecord(temp, temp, temp)
		err := t.Insert(num, temp, temp, temp)
		if err != nil {
			fmt.Printf("error: %s\n\n", err)
		}
	}

	t.FindAndPrint(7, true)
	t.BlckMngr.GetBlocksAccessed()

	// t.PrintTree()
	// b.DisplayStatus(true)
}
