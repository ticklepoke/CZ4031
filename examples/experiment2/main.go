package main

import (
	"fmt"
	"strconv"

	"github.com/ticklepoke/CZ4031/blockmanager"

	"github.com/ticklepoke/CZ4031/bptree"
	"github.com/ticklepoke/CZ4031/tsvparser"
)

func main() {
	n := 5
	t := bptree.NewTree(n)
	rows := tsvparser.ParseTSV("data.tsv")

	b := blockmanager.InitializeBlockManager(100)

	for _, s := range rows[:100] {
		tconsts, rating, votes := s[0], s[1], s[2]
		key, _ := strconv.ParseFloat(rating, 64)
		addr := b.InsertRecord(tconsts, rating, votes)
		fmt.Println(tconsts)
		t.Insert(int(key*10), addr)
	}

	// t.PrintOrder()
	// t.PrintLeaves()
	// t.PrintHeight()
	t.PrintTree()
	b.DisplayStatus(false)
}
