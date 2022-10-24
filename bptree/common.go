package bptree

import (
	"errors"

	"github.com/ticklepoke/CZ4031/blockmanager"
)

var (
	err error

	// N is the maximum number of keys in a node
	N              = 4
	queue          *Node
	verbose_output = false
	version        = 0.1
)

// Tree is a B+ Tree
type Tree struct {
	Root         *Node
	BlckMngr     *blockmanager.BlockManager
	numDeletions int
}

// Record serialize and unserialize function / library
type Record struct {
	Value *[]byte
	Next  *Record
}

// Node represents a B+ tree node
type Node struct {
	Pointers     []interface{}
	TailPointers []interface{}
	Keys         []float64
	Parent       *Node
	IsLeaf       bool
	NumKeys      int
	Next         *Node
}

// NewTree Constructor with Order
func NewTree(n, blocksize int) *Tree {
	N = n
	b := blockmanager.InitializeBlockManager(blocksize)
	return &Tree{BlckMngr: &b}
}

func makeNode() (*Node, error) {
	newNode := new(Node)
	newNode.Keys = make([]float64, N-1)
	if newNode.Keys == nil {
		return nil, errors.New("error: New node keys array")
	}
	newNode.Pointers = make([]interface{}, N)
	newNode.TailPointers = make([]interface{}, N)
	if newNode.Keys == nil {
		return nil, errors.New("error: New node pointers array")
	}
	newNode.IsLeaf = false
	newNode.NumKeys = 0
	newNode.Parent = nil
	newNode.Next = nil
	return newNode, nil
}

func makeLeaf() (*Node, error) {
	leaf, err := makeNode()
	if err != nil {
		return nil, err
	}
	leaf.IsLeaf = true
	return leaf, nil
}

func (t *Tree) startNewTree(key float64, pointer *Record) error {
	t.Root, err = makeLeaf()
	if err != nil {
		return err
	}
	t.Root.Keys[0] = key
	t.Root.Pointers[0] = pointer
	t.Root.TailPointers[0] = pointer
	t.Root.Pointers[N-1] = nil
	t.Root.Parent = nil
	t.Root.NumKeys++
	return nil
}

func contains(arrSearch []float64, valSearch float64) bool {
	for _, valIter := range arrSearch {
		if valIter == valSearch {
			return true
		}
	}
	return false
}

func iterLeafLL(recordPtr *Record) []*Record {
	// iterate over the LL at the leaf node returning the list of records
	var recordsArr []*Record
	curr := recordPtr
	for curr != nil {
		recordsArr = append(recordsArr, curr)
		curr = curr.Next
	}
	return recordsArr
}
