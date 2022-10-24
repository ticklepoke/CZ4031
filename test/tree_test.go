package main

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/ticklepoke/CZ4031/bptree"
	"github.com/ticklepoke/CZ4031/logger"
)

func getLevelOrder(tree *bptree.Tree) [][]float64 {
	// traverse the tree in level order and append keys to a slice
	var n *bptree.Node
	var levelOrder [][]float64
	queue := []*bptree.Node{tree.Root}
	for len(queue) > 0 {
		n = queue[0]
		queue = queue[1:]
		if n != nil {
			var keys []float64
			for i := 0; i < n.NumKeys; i++ {
				keys = append(keys, n.Keys[i])
			}
			levelOrder = append(levelOrder, keys)
			if !n.IsLeaf {
				for i := 0; i <= n.NumKeys; i++ {
					c, _ := n.Pointers[i].(*bptree.Node)
					queue = append(queue, c)
				}
			}
		}
	}

	return levelOrder
}

// init
func init() {
	logger.Logger.SetLevel(log.WarnLevel)
}

func compare(expected [][]float64, result [][]float64) (bool, int) {
	if len(expected) != len(result) {
		return false, -1
	}

	// compare the actual and result
	for i := 0; i < len(expected); i++ {
		for j := 0; j < len(expected[i]); j++ {
			if expected[i][j] != result[i][j] {
				return false, i
			}
		}
	}
	return true, -1
}

func TestInsertion(t *testing.T) {
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
	tree := runner(4, rows)
	result := getLevelOrder(tree)

	expected := [][]float64{
		{20}, {7, 17}, {25}, {1, 4}, {7, 10}, {17, 19}, {20, 21}, {25, 31},
	}

	valid, index := compare(expected, result)
	if !valid {
		if index == -1 {
			t.Errorf("The length of the tree is not correct")
		} else {
			t.Errorf("The tree is not correct at index %d", index)
		}
		t.Errorf("Expected: %v", expected)
		t.Errorf("Result: %v", result)
	}
}

func TestDeletion17(t *testing.T) {
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
	tree := runner(4, rows)
	tree.Insert(5, "17", "2", "3")
	tree.Insert(16, "17", "2", "3")
	delete(17, tree)
	result := getLevelOrder(tree)

	expected := [][]float64{
		{20}, {7, 16}, {25}, {1, 4, 5}, {7, 10}, {16, 19}, {20, 21}, {25, 31},
	}

	valid, index := compare(expected, result)
	if !valid {
		if index == -1 {
			t.Errorf("The length of the tree is not correct")
		} else {
			t.Errorf("The tree is not correct at index %d", index)
		}
		t.Errorf("Expected: %v", expected)
		t.Errorf("Result: %v", result)
	}
}

func TestDeletion4_1(t *testing.T) {
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
	delete(4, tree)

	result := getLevelOrder(tree)
	expected := [][]float64{
		{20}, {17}, {25, 31}, {1, 7, 10}, {17, 19}, {20, 21}, {25, 28}, {31, 42},
	}

	valid, index := compare(expected, result)
	if !valid {
		if index == -1 {
			t.Errorf("The length of the tree is not correct")
		} else {
			t.Errorf("The tree is not correct at index %d", index)
		}
		t.Errorf("Expected: %v", expected)
		t.Errorf("Result: %v", result)
	}
}

func TestDeletion4_2(t *testing.T) {
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

	result := getLevelOrder(tree)
	expected := [][]float64{
		{20, 25}, {1, 7, 10}, {20, 21}, {25, 31},
	}

	valid, index := compare(expected, result)
	if !valid {
		if index == -1 {
			t.Errorf("The length of the tree is not correct")
		} else {
			t.Errorf("The tree is not correct at index %d", index)
		}
		t.Errorf("Expected: %v", expected)
		t.Errorf("Result: %v", result)
	}
}
