package main

import (
	"fmt"
	"strconv"
	"time"
	"unsafe"

	"github.com/ticklepoke/CZ4031/blockmanager"
	"github.com/ticklepoke/CZ4031/bptree"
	"github.com/ticklepoke/CZ4031/tsvparser"
)

func experiment1() {
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
	t := bptree.NewTree(n)
	rows := tsvparser.ParseTSV("./data.tsv")

	b := blockmanager.InitializeBlockManager(100)
	i := 0
	for _, s := range rows {
		if i == 1000 {
			break
		}
		tconsts, rating, votes := s[0], s[1], s[2]
		key, _ := strconv.ParseFloat(rating, 64)
		addr := b.InsertRecord(tconsts, rating, votes)
		fmt.Println(tconsts, key)
		t.Insert(int(key*10), addr)
		i++
	}

	// t.PrintOrder()
	// t.PrintLeaves()
	// t.PrintHeight()
	t.PrintTree()
	b.DisplayStatus(false)
	return b, t
}

func experiment3(t *bptree.Tree) {
	t.FindAndPrint(8, true)
}

func experiment4(t *bptree.Tree) {
	t.FindAndPrintRange(7, 9, true)
}

func experiment5(b blockmanager.BlockManager, t *bptree.Tree) {
	start := time.Now()
	recPtrs, _ := t.Find(70, false)
	t.PrintTree()
	t.Delete(70)
	t.PrintTree()
	t.PrintHeight()
	t.PrintLeaves()

	for _, recPtr := range recPtrs {
		b.DeleteRecord((*[]byte)(unsafe.Pointer(recPtr)))
	}

	b.DisplayStatus(true)
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
