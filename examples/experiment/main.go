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

	rows := tsvparser.ParseTSV("data.tsv")

	for _, row := range rows {
		tconts, rating, votes := row[0], row[1], row[2]

		// TODO: insert record to bptree
		b.InsertRecord(tconts, rating, votes)
	}

	b.DisplayStatus(false)
}

func experiment2(n int) *bptree.Tree {
	fmt.Println("================= Experiment 2 =================")
	t := bptree.NewTree(n)
	rows := tsvparser.ParseTSV("data.tsv")

	i := 0
	for _, s := range rows {
		if i == 1000 {
			break
		}
		tconsts, rating, votes := s[0], s[1], s[2]
		key, _ := strconv.ParseFloat(rating, 64)
		t.Insert(key, tconsts, rating, votes)
		i++
	}

	// t.PrintOrder()
	// t.PrintLeaves()
	// t.PrintHeight()
	t.PrintTree()
	t.BlckMngr.DisplayStatus(false)
	return t
}

func experiment3(t *bptree.Tree) {
	fmt.Println("================= Experiment 3 =================")
	t.BlckMngr.ResetBlocksAccessed()
	t.FindAndPrint(8.0, true)

	t.BlckMngr.GetBlocksAccessed()
}

func experiment4(t *bptree.Tree) {
	fmt.Println("================= Experiment 4 =================")
	t.BlckMngr.ResetBlocksAccessed()
	t.FindAndPrintRange(7.0, 9.0, true)

	t.BlckMngr.GetBlocksAccessed()
}

func experiment5(t *bptree.Tree) {
	fmt.Println("================= Experiment 5 =================")
	start := time.Now()
	recPtrs, _ := t.Find(7.0, false)
	t.PrintTree()
	t.Delete(7.0)
	t.PrintTree()
	t.PrintHeight()
	t.PrintLeaves()

	// TODO: add this into the DeleteRecord func
	for _, recPtr := range recPtrs {
		t.BlckMngr.DeleteRecord(recPtr.Value)
	}
	t.BlckMngr.DisplayStatus(false)
	elapse := time.Since(start)
	fmt.Println("Experiment 5 elapsed time: ", elapse)
}

func main() {
	n := 5
	// b := experiment1()
	t := experiment2(n)
	experiment3(t)
	experiment4(t)
	experiment5(t)
}
