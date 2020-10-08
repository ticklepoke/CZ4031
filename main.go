package main

import (
	"fmt"
	"strconv"

	"github.com/ticklepoke/CZ4031/blockmanager"
	"github.com/ticklepoke/CZ4031/tsvparser"

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

	// t.PrintTree()

	b := blockmanager.InitializeBlockManager(100)

	// testing out tsv
	rows := tsvparser.ParseTSV("data.tsv")

	for i := 0; i < 10; i++ {
		tconts, rating, votes := rows[i][0], rows[i][1], rows[i][2]

		// TODO: insert record to bptree
		b.InsertRecord(tconts, rating, votes)
	}

	b.DisplayStatus(true)
}
