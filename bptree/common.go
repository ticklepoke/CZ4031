package bptree

import "errors"

var (
	err error

	// defaultOrder = 4
	// minOrder     = 3
	// maxOrder     = 20

	// order          = defaultOrder
	N              = 4
	queue          *Node
	verbose_output = false
	version        = 0.1
)

type Tree struct {
	Root *Node
}

// serialize and unserialize function / library
type Record struct {
	Value []byte
	Next  *Record
}

// block manager struct
// allocate new block
// keep track of blocks with free space
// delete record

type Node struct {
	Pointers []*Record
	Keys     []int
	Parent   *Node
	IsLeaf   bool
	NumKeys  int
	Next     *Node
}

func NewTree() *Tree {
	return &Tree{}
}

// call block manager
// allocate space to the record
func makeRecord(value []byte) (*Record, error) {
	new_record := new(Record)
	if new_record == nil {
		return nil, errors.New("Error: Record creation.")
	} else {
		new_record.Value = value
	}
	return new_record, nil
}

func makeNode() (*Node, error) {
	new_node := new(Node)
	if new_node == nil {
		return nil, errors.New("Error: Node creation.")
	}
	new_node.Keys = make([]int, N-1)
	if new_node.Keys == nil {
		return nil, errors.New("Error: New node keys array.")
	}
	new_node.Pointers = make([]*Record, order)
	if new_node.Keys == nil {
		return nil, errors.New("Error: New node pointers array.")
	}
	new_node.IsLeaf = false
	new_node.NumKeys = 0
	new_node.Parent = nil
	new_node.Next = nil
	return new_node, nil
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

func (t *Tree) iterLeafLL(recordPtr *Record) []*Record {
	// iterate over the LL at the leaf node returning the list of records
	var recordsArr []*Record
	curr := recordPtr
	for curr != nil {
		recordsArr = append(recordsArr, curr)
		curr = curr.Next
	}
	return recordsArr
}
