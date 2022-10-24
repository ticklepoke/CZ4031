package main

import (
	"os"
	"strconv"

	"github.com/ticklepoke/CZ4031/bptree"
	"github.com/ticklepoke/CZ4031/logger"

	log "github.com/sirupsen/logrus"
)

var tree *bptree.Tree

func del17() {
	rows := [][]string{
		{"1", "2", "3"},
		{"4", "2", "3"},
		{"7", "2", "3"},
		{"10", "2", "3"},
		{"17", "2", "3"},
		{"21", "2", "3"},
		{"31", "2", "3"},
		{"25", "2", "3"},
		{"19", "2", "3"},
		{"20", "2", "3"},
	}
	tree = runner(4, rows)
	tree.Insert(5, "17", "2", "3")
	tree.Insert(16, "17", "2", "3")
	delete(17, tree)
}

func delFourOne() {
	rows := [][]string{
		{"1", "2", "3"},
		{"4", "2", "3"},
		{"7", "2", "3"},
		{"10", "2", "3"},
		{"17", "2", "3"},
		{"21", "2", "3"},
		{"31", "2", "3"},
		{"25", "2", "3"},
		{"19", "2", "3"},
		{"20", "2", "3"},
	}
	tree = runner(4, rows)
	delete(4, tree)
}

func delFourTwo() {
	rows := [][]string{
		{"1", "2", "3"},
		{"4", "2", "3"},
		{"7", "2", "3"},
		{"10", "2", "3"},
		{"17", "2", "3"},
		{"21", "2", "3"},
		{"31", "2", "3"},
		{"25", "2", "3"},
		{"19", "2", "3"},
		{"20", "2", "3"},
	}
	tree = runner(4, rows)
	tree.PrintTree()
	tree.Delete(17)
	tree.Delete(19)
	delete(4, tree)
}

func delete(key float64, t *bptree.Tree) {
	t.PrintTree()
	t.Delete(key)
	t.PrintTree()
}

func runner(n int, rows [][]string) *bptree.Tree {
	blockSize := 100
	t := bptree.NewTree(n, blockSize)
	i := 0

	for _, s := range rows {
		tconst, rating, votes := s[0], s[1], s[2]
		key, _ := strconv.ParseFloat(tconst, 64)
		t.Insert(key, tconst, rating, votes)
		i++
	}

	return t
}

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		// set log level
		level, err := log.ParseLevel(args[0])
		if err != nil {
			log.Fatal(err)
		}
		logger.Logger.SetLevel(level)

		// choose test case
		example := args[1]
		switch example {
		case "1":
			del17()
		case "2":
			delFourOne()
		case "3":
			delFourTwo()
		}
	} else {
		// no args provided, exit
		log.Fatal("Please provide a logging level and example number [1, 2, 3]")
		os.Exit(1)
	}
}
