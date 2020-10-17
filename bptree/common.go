package bptree

import "errors"

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
	Root *Node
}

// Record serialize and unserialize function / library
type Record struct {
	Value *[]byte
	Next  *Record
}

// block manager struct
// allocate new block
// keep track of blocks with free space
// delete record

// Node represents a B+ tree node
type Node struct {
	Pointers     []interface{}
	TailPointers []interface{}
	Keys         []int
	Parent       *Node
	IsLeaf       bool
	NumKeys      int
	Next         *Node
}

// NewTree Constructor with Order
func NewTree(n int) *Tree {
	N = n
	return &Tree{}
}

// call block manager
// allocate space to the record
// func makeRecord(value []byte) (*Record, error) {
// 	newRecord := new(Record)
// 	if newRecord == nil {
// 		return nil, errors.New("Error: Record creation")
// 	} else {
// 		newRecord.Value = value
// 	}
// 	return newRecord, nil
// }

func makeNode() (*Node, error) {
	newNode := new(Node)
	if newNode == nil {
		return nil, errors.New("Error: Node creation")
	}
	newNode.Keys = make([]int, N-1)
	if newNode.Keys == nil {
		return nil, errors.New("Error: New node keys array")
	}
	newNode.Pointers = make([]interface{}, N)
	newNode.TailPointers = make([]interface{}, N)
	if newNode.Keys == nil {
		return nil, errors.New("Error: New node pointers array")
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

func (t *Tree) startNewTree(key int, pointer *Record) error {
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

func contains(arrSearch []int, valSearch int) bool {
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
