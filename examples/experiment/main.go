package main

import (
	"fmt"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/ticklepoke/CZ4031/logger"

	"github.com/ticklepoke/CZ4031/bptree"
	"github.com/ticklepoke/CZ4031/tsvparser"
)

var (
	blockSize int
)

func experiment1And2(n, size int) *bptree.Tree {
	blockSize = size
	fmt.Printf("Running experiment 1 %dB\n", blockSize)
	t := bptree.NewTree(n, blockSize)
	rows := tsvparser.ParseTSV("data.tsv")
	loggername := "experiment1_" + strconv.Itoa(blockSize)
	logger.InitFileLogger(loggername, log.WarnLevel)

	i := 0
	for _, s := range rows {
		tconsts, rating, votes := s[0], s[1], s[2]
		key, _ := strconv.ParseFloat(rating, 64)
		t.Insert(key, tconsts, rating, votes)
		i++
	}

	t.BlckMngr.DisplayStatus(false)
	fmt.Printf("Running experiment 2 %dB\n", blockSize)
	loggername = "experiment2_" + strconv.Itoa(blockSize)
	logger.InitFileLogger(loggername, log.WarnLevel)
	logger.Logger.Println("B+ tree has parameter n of", n)
	logger.Logger.Println("B+ tree has height of", t.Height())
	logger.Logger.Println("Printing B+ tree structure")
	t.PrintTree()
	return t
}

func experiment3(t *bptree.Tree) {
	fmt.Printf("Running experiment 3 %dB\n", blockSize)
	loggername := "experiment3_" + strconv.Itoa(blockSize)
	logger.InitFileLogger(loggername, log.WarnLevel)
	t.BlckMngr.ResetBlocksAccessed()
	t.FindAndPrint(8.0, true)

	t.BlckMngr.GetBlocksAccessed()
}

func experiment4(t *bptree.Tree) {
	fmt.Printf("Running experiment 4 %dB\n", blockSize)
	loggername := "experiment4_" + strconv.Itoa(blockSize)
	logger.InitFileLogger(loggername, log.WarnLevel)
	t.BlckMngr.ResetBlocksAccessed()
	t.FindAndPrintRange(7.0, 9.0, true)

	t.BlckMngr.GetBlocksAccessed()
}

func experiment5(t *bptree.Tree) {
	fmt.Printf("Running experiment 5 %dB\n", blockSize)
	loggername := "experiment5_" + strconv.Itoa(blockSize)
	logger.InitFileLogger(loggername, log.WarnLevel)
	start := time.Now()
	recPtrs, _ := t.Find(7.0, false)
	t.PrintTree()
	t.Delete(7.0)
	t.FindNumDeletions()
	logger.Logger.Println("B+ tree has height of", t.Height())
	logger.Logger.Println("Printing B+ tree structure")
	logger.Logger.Println()
	t.PrintTree()

	for _, recPtr := range recPtrs {
		t.BlckMngr.DeleteRecord(recPtr.Value)
	}
	t.BlckMngr.DisplayStatus(false)
	elapse := time.Since(start)
	fmt.Println("Experiment 5 elapsed time: ", elapse)
}

func main() {
	n := 5
	t := experiment1And2(n, 100)
	experiment3(t)
	experiment4(t)
	experiment5(t)

	c := experiment1And2(n, 500)
	experiment3(c)
	experiment4(c)
	experiment5(c)
}
