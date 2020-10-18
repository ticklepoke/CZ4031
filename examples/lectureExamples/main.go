package main

import (
	"fmt"
	"strconv"

	"github.com/ticklepoke/CZ4031/blockmanager"
	"github.com/ticklepoke/CZ4031/bptree"
)

func buildTree(t *bptree.Tree, b blockmanager.BlockManager, arr []float64, target float64, v2 bool) {
	for _, num := range arr {
		temp := strconv.Itoa(int(num))
		addr := b.InsertRecord(temp, temp, temp)
		err := t.Insert(num, addr)
		if err != nil {
			fmt.Printf("error: %s\n\n", err)
		}
	}

	// idk how to do it without deletion
	if v2 {
		t.Delete(17)
		t.Delete(19)
	}

	fmt.Printf("====================\n DELETION EXAMPLE %d\n====================\n", int(target))

	t.PrintTree()
	t.Delete(target)
	t.PrintTree()
}

func insert() {
	t := bptree.NewTree(4)
	b := blockmanager.InitializeBlockManager(100)
	arr := []float64{1, 4, 7, 10, 17, 21, 31, 25, 19, 20, 28, 42}
	buildTree(t, b, arr, 5, false)
}

func del5() {
	t := bptree.NewTree(4)
	b := blockmanager.InitializeBlockManager(100)
	arr := []float64{1, 4, 7, 10, 17, 21, 31, 25, 19, 20, 5}
	buildTree(t, b, arr, 5, false)

	// 20.0  |
	// 7.0 17.0  | 25.0  |
	// 1.0 4.0  | 7.0 10.0  | 17.0 19.0  | 20.0 21.0  | 25.0 31.0  |
}

func del17() {
	t := bptree.NewTree(4)
	b := blockmanager.InitializeBlockManager(100)
	arr := []float64{1, 4, 7, 10, 17, 21, 31, 25, 19, 20, 5, 16}
	buildTree(t, b, arr, 17, false)

	// 20.0  |
	// 7.0  | 25.0  |
	// 1.0 4.0  | 7.0 10.0 19.0  | 20.0 21.0  | 25.0 31.0  |
}

func del4() {
	t := bptree.NewTree(4)
	b := blockmanager.InitializeBlockManager(100)
	arr := []float64{1, 4, 7, 10, 17, 21, 31, 25, 19, 20}
	buildTree(t, b, arr, 4, false)

	// 20.0  |
	// 17.0  | 25.0  |
	// 1.0 7.0 10.0  | 17.0 19.0  | 20.0 21.0  | 25.0 31.0  |
}

func del4v2() {
	t := bptree.NewTree(4)
	b := blockmanager.InitializeBlockManager(100)
	arr := []float64{1, 4, 7, 10, 17, 21, 31, 25, 19, 20}
	buildTree(t, b, arr, 4, true)

	// 20.0 25.0  |
	// 1.0 7.0 10.0  | 20.0 21.0  | 25.0 31.0  |
}

func main() {
	// insert()
	del5()
	del17()
	del4()
	del4v2()
}
