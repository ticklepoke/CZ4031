package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ticklepoke/CZ4031/blockmanager"
	"github.com/ticklepoke/CZ4031/bptree"
	"github.com/ticklepoke/CZ4031/tsvparser"
)

func experiment1() {
	fmt.Println("================= Experiment 1 =================")
	b := blockmanager.InitializeBlockManager(100)

	rows := tsvparser.ParseTSV("../../data.tsv")

	for _, row := range rows {
		tconts, rating, votes := row[0], row[1], row[2]

		// TODO: insert record to bptree
		b.InsertRecord(tconts, rating, votes)
	}

	b.DisplayStatus(false)
}

func experiment2(n int) (blockmanager.BlockManager, *bptree.Tree) {
	fmt.Println("================= Experiment 2 =================")
	t := bptree.NewTree(n)
	rows := tsvparser.ParseTSV("../../data.tsv")

	b := blockmanager.InitializeBlockManager(100)

	for _, s := range rows {
		tconsts, rating, votes := s[0], s[1], s[2]
		key, _ := strconv.ParseFloat(rating, 64)
		addr := b.InsertRecord(tconsts, rating, votes)
		t.Insert(int(key*10), addr)
	}

	// t.PrintOrder()
	// t.PrintLeaves()
	// t.PrintHeight()
	t.PrintTree()
	b.DisplayStatus(false)
	return b, t
}

func experiment3(t *bptree.Tree) {
	fmt.Println("================= Experiment 3 =================")
	t.FindAndPrint(80, true)
}

func experiment4(t *bptree.Tree) {
	fmt.Println("================= Experiment 4 =================")
	t.FindAndPrintRange(70, 90, true)
}

func experiment5(b blockmanager.BlockManager, t *bptree.Tree) {
	fmt.Println("================= Experiment 5 =================")
	start := time.Now()
	t.Delete(7)
	t.PrintTree()
	t.PrintHeight()
	t.PrintLeaves()

	elapse := time.Since(start)
	fmt.Println("Experiment 5 elapsed time: ", elapse)
}

func main() {
	n := 5
	// b := experiment1()
	b, t := experiment2(n)
	experiment3(t)
	experiment4(t)
	experiment5(b, t)
}
