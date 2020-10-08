package bptree

import (
	"errors"
)

// Insert - implement duplicate key insertion functionality
// create new node w/o parent pointer (point to the left)
func (t *Tree) Insert(key int, value []byte) error {
	var pointer *Record
	var leaf *Node

	// Inserting a new key
	pointer, err := makeRecord(value)
	if err != nil {
		return err
	}

	if t.Root == nil {
		return t.startNewTree(key, pointer)
	}

	leaf = t.findLeaf(key, false)

	// if leaf node contains the key we want we can just insert to leaf node
	if leaf.NumKeys < N-1 || contains(leaf.Keys, key) {
		insertIntoLeaf(leaf, key, pointer)
		return nil
	}

	// split the leaf node if it is already full
	return t.insertIntoLeafAfterSplitting(leaf, key, pointer)
}

/* ============================ Private Methods ============================*/

// implement binsearch
func getLeftIndex(parent, left *Node) int {
	leftIndex := 0
	for leftIndex <= parent.NumKeys && parent.Pointers[leftIndex] != left {
		leftIndex++
	}
	return leftIndex
}

func findInsertionIndex(leaf *Node, key int) int {
	var insertionPoint int
	for insertionPoint < leaf.NumKeys && leaf.Keys[insertionPoint] < key {
		insertionPoint++
	}
	return insertionPoint
}

// insert record in to leaf node
func insertIntoLeaf(leaf *Node, key int, pointer *Record) {
	var i int

	var insertionPoint int = findInsertionIndex(leaf, key)

	// key found in the leaf node
	if leaf.Keys[insertionPoint] == key {
		curr := leaf.Pointers[insertionPoint]
		if curr == nil {
			// no records yet
			leaf.Pointers[insertionPoint] = pointer
			return
		} else {
			// get last leaf node
			var nodes []*Record = iterLeafLL(curr.(*Record))
			lastRecNode := nodes[len(nodes)-1]
			lastRecNode.Next = pointer
			return
		}
	}

	// leaf not found in leaf node keys, move everything right and insert
	for i = leaf.NumKeys; i > insertionPoint; i-- {
		leaf.Keys[i] = leaf.Keys[i-1]
		leaf.Pointers[i] = leaf.Pointers[i-1]
	}
	leaf.Keys[insertionPoint] = key
	leaf.Pointers[insertionPoint] = pointer
	leaf.NumKeys++
	return
}

// implement binsearch
func (t *Tree) insertIntoLeafAfterSplitting(leaf *Node, key int, pointer *Record) error {
	var newLeaf *Node
	var split, newKey, i, j int
	var err error

	newLeaf, err = makeLeaf()
	if err != nil {
		return nil
	}

	tempKeys := make([]int, N)
	// if tempKeys == nil {
	// 	return errors.New("error: Temporary keys array")
	// }

	tempPointers := make([]interface{}, N)
	// if tempPointers == nil {
	// 	return errors.New("error: Temporary pointers array")
	// }

	var insertionIndex = findInsertionIndex(leaf, key)

	// copy the array and insert at insertion point
	for i = 0; i < leaf.NumKeys; i++ {
		// skips the space of the insertion index
		if j == insertionIndex {
			j++
		}
		tempKeys[j] = leaf.Keys[i]
		tempPointers[j] = leaf.Pointers[i]
		j++
	}

	tempKeys[insertionIndex] = key
	tempPointers[insertionIndex] = pointer

	leaf.NumKeys = 0

	split = findMidPoint(N - 1)

	// could just make use of copy to do this
	for i = 0; i < split; i++ {
		leaf.Pointers[i] = tempPointers[i]
		leaf.Keys[i] = tempKeys[i]
		leaf.NumKeys++
	}

	j = 0
	for i = split; i < N; i++ {
		newLeaf.Pointers[j] = tempPointers[i]
		newLeaf.Keys[j] = tempKeys[i]
		newLeaf.NumKeys++
		j++
	}

	// newLeaf.Pointers[N-1] = leaf.Pointers[N-1]
	// leaf.Pointers[N-1] = []*Record{newLeaf}
	newLeaf.Next = leaf.Next
	leaf.Next = newLeaf.Next

  // set the indices after insertion point to nil
	for i = leaf.NumKeys; i < N-1; i++ {
		leaf.Pointers[i] = nil
	}
	for i = newLeaf.NumKeys; i < N-1; i++ {
		newLeaf.Pointers[i] = nil
	}

	// point to the left for dup keys
	// CHECK this might be wrong
	newLeaf.Parent = leaf.Parent
	newKey = newLeaf.Keys[0]

	return t.insertIntoParent(leaf, newKey, newLeaf)
}

func insertIntoNode(n *Node, leftIndex, key int, right *Node) {
	var i int
	for i = n.NumKeys; i > leftIndex; i-- {
		n.Pointers[i+1] = n.Pointers[i]
		n.Keys[i] = n.Keys[i-1]
	}
	n.Pointers[leftIndex+1] = right
	n.Keys[leftIndex] = key
	n.NumKeys++
}

// implement binsearch
func (t *Tree) insertIntoNodeAfterSplitting(oldNode *Node, leftIndex, key int, right *Node) error {
	var i, j, split, kPrime int
	var newNode, child *Node
	var tempKeys []int
	var tempPointers []interface{}
	var err error

	tempPointers = make([]interface{}, N+1)
	if tempPointers == nil {
		return errors.New("error: unable to make temporary pointers array for splitting nodes")
	}

	tempKeys = make([]int, N)
	if tempKeys == nil {
		return errors.New("error: unable to make temporary keys array for splitting nodes")
	}

	for i = 0; i < oldNode.NumKeys+1; i++ {
		if j == leftIndex+1 {
			j++
		}
		tempPointers[j] = oldNode.Pointers[i]
		j++
	}

	j = 0
	for i = 0; i < oldNode.NumKeys; i++ {
		if j == leftIndex {
			j++
		}
		tempKeys[j] = oldNode.Keys[i]
		j++
	}

	tempPointers[leftIndex+1] = right

	tempKeys[leftIndex] = key

	split = findMidPoint(N)
	newNode, err = makeNode()
	if err != nil {
		return err
	}
	oldNode.NumKeys = 0
	for i = 0; i < split; i++ {
		oldNode.Pointers[i] = tempPointers[i]
		oldNode.Keys[i] = tempKeys[i]
		oldNode.NumKeys++
	}
	oldNode.Pointers[i] = tempPointers[i]

	kPrime = tempKeys[split]

	j = 0
	// for i += 1; i < N; i++ {
	for i += 1; i < N; i++ {
		newNode.Pointers[j] = tempPointers[i]
		newNode.Keys[j] = tempKeys[i]
		newNode.NumKeys++
		j++
	}

	newNode.Pointers[j] = tempPointers[i]
	newNode.Parent = oldNode.Parent
	for i = 0; i <= newNode.NumKeys; i++ {
		child, _ = newNode.Pointers[i].(*Node)
		child.Parent = newNode
	}

	for i = 0; i <= oldNode.NumKeys; i++ {
		child, _ = oldNode.Pointers[i].(*Node)
		fmt.Println(child.Keys)
	}

	// adjust parent pointer of the two nodes
	return t.insertIntoParent(oldNode, kPrime, newNode)
}

func (t *Tree) insertIntoParent(left *Node, key int, right *Node) error {
	var leftIndex int
	parent := left.Parent

	if parent == nil {
		return t.insertIntoNewRoot(left, key, right)
	}

	leftIndex = getLeftIndex(parent, left)

	// if there is space, insert directly

	if parent.NumKeys < N-1 {
		insertIntoNode(parent, leftIndex, key, right)
		return nil
	}

	return t.insertIntoNodeAfterSplitting(parent, leftIndex, key, right)
}

func (t *Tree) insertIntoNewRoot(left *Node, key int, right *Node) error {
	t.Root, err = makeNode()
	if err != nil {
		return err
	}
	t.Root.Keys[0] = key
	t.Root.Pointers[0] = left
	t.Root.Pointers[1] = right
	t.Root.NumKeys++
	t.Root.Parent = nil
	left.Parent = t.Root
	right.Parent = t.Root
	return nil
}
