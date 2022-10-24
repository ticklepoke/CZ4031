package main

import (
	"os"
	"strconv"

	"github.com/ticklepoke/CZ4031/bptree"
	"github.com/ticklepoke/CZ4031/logger"

	log "github.com/sirupsen/logrus"
)

var tree *bptree.Tree

func lecExample() {
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
		{"28", "2", "3"},
		{"42", "2", "3"},
	}
	tree = runner(4, rows)
}

func lecExampleDel17() {
	lecExample()
	tree.Insert(5, "17", "2", "3")
	tree.Insert(16, "17", "2", "3")
	tree.PrintTree()
	delete(17, tree)
}

func lecExampleDelFourTwo() {
	lecExample()
	delete(17, tree)
	delete(19, tree)
	delete(4, tree)
}

func exampleStealFromNonLeaf() {
	lecExample()
	tree.Insert(43, "43", "2", "3")
	tree.Insert(44, "44", "2", "3")
	tree.Insert(5, "5", "2", "3")
	tree.Insert(16, "16", "2", "3")
	tree.Insert(45, "45", "2", "3")
	tree.Insert(46, "46", "2", "3")
	tree.Insert(47, "47", "2", "3")
	tree.Insert(48, "48", "2", "3")
	tree.Insert(49, "48", "2", "3")
	tree.Insert(50, "48", "2", "3")
	tree.PrintTree()

	delete(47, tree)
	delete(48, tree)
	delete(5, tree)
	delete(4, tree)
	delete(10, tree)
	delete(17, tree)
	delete(19, tree)
	delete(42, tree)
}

func lecExampleDelFourOne() {
	lecExample()
	delete(4, tree)
}

func delete(key float64, t *bptree.Tree) {
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
	// set logging level
	args := os.Args[1:]
	if len(args) > 0 {
		level, err := log.ParseLevel(args[0])
		if err != nil {
			log.Fatal(err)
		}
		logger.Logger.SetLevel(level)
	}

	lecExampleDel17()      // working
	lecExampleDelFourOne() // working
	lecExampleDelFourTwo() // working
	// exampleStealFromNonLeaf()
}
