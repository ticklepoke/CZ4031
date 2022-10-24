package bptree

import (
	"fmt"
)

// Insert - implement duplicate key insertion functionality
// create new node w/o parent pointer (point to the left)
func (t *Tree) Insert(key float64, tconsts, rating, votes string) error {
	var pointer *Record
	var leaf *Node

	// Inserting a new key
	addr := t.BlckMngr.InsertRecord(tconsts, rating, votes)
	pointer = &Record{Value: addr}
	// pointer, err := makeRecord(value)
	// if err != nil {
	// 	return err
	// }

	if t.Root == nil {
		return t.startNewTree(key, pointer)
	}

	leaf, _ = t.findLeaf(key, false)

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

func findInsertionIndex(leaf *Node, key float64) int {
	var insertionPoint int
	for insertionPoint < leaf.NumKeys && leaf.Keys[insertionPoint] < key {
		insertionPoint++
	}
	return insertionPoint
}

// insert record in to leaf node
func insertIntoLeaf(leaf *Node, key float64, pointer *Record) {
	var i int

	var insertionPoint int = findInsertionIndex(leaf, key)

	// key found in the leaf node
	if leaf.Keys[insertionPoint] == key {
		curr := leaf.Pointers[insertionPoint]
		if curr == nil {
			fmt.Println("no record")
			// no records yet
			leaf.Pointers[insertionPoint] = pointer
			leaf.TailPointers[insertionPoint] = pointer
			return
		} else {
			prev := leaf.TailPointers[insertionPoint].(*Record)
			prev.Next = pointer
			leaf.TailPointers[insertionPoint] = pointer
			return
		}
	}

	// leaf not found in leaf node keys, move everything right and insert
	for i = leaf.NumKeys; i > insertionPoint; i-- {
		leaf.Keys[i] = leaf.Keys[i-1]
		leaf.Pointers[i] = leaf.Pointers[i-1]
		leaf.TailPointers[i] = leaf.TailPointers[i-1]
	}
	leaf.Keys[insertionPoint] = key
	leaf.Pointers[insertionPoint] = pointer
	leaf.TailPointers[insertionPoint] = pointer
	leaf.NumKeys++
}

func (t *Tree) insertIntoLeafAfterSplitting(leaf *Node, key float64, pointer *Record) error {
	var newLeaf *Node
	var split, i, j int
	var newKey float64
	var err error

	newLeaf, err = makeLeaf()
	if err != nil {
		return nil
	}

	tempKeys := make([]float64, N)
	tempPointers := make([]interface{}, N)
	tempTailPointers := make([]interface{}, N)
	insertionIndex := findInsertionIndex(leaf, key)

	// copy the array and insert at insertion point
	for i = 0; i < leaf.NumKeys; i++ {
		// skips the space of the insertion index
		if j == insertionIndex {
			j++
		}
		tempKeys[j] = leaf.Keys[i]
		tempPointers[j] = leaf.Pointers[i]
		tempTailPointers[j] = leaf.TailPointers[i]
		j++
	}

	tempKeys[insertionIndex] = key
	tempPointers[insertionIndex] = pointer
	for pointer.Next != nil {
		pointer = pointer.Next
	}
	tempTailPointers[insertionIndex] = pointer

	leaf.NumKeys = 0
	split = findMidPoint(N)
	for i = 0; i < split; i++ {
		leaf.Pointers[i] = tempPointers[i]
		leaf.TailPointers[i] = tempTailPointers[i]
		leaf.Keys[i] = tempKeys[i]
		leaf.NumKeys++
	}

	j = 0
	for i = split; i < N; i++ {
		newLeaf.Pointers[j] = tempPointers[i]
		newLeaf.TailPointers[j] = tempTailPointers[i]
		newLeaf.Keys[j] = tempKeys[i]
		newLeaf.NumKeys++
		j++
	}

	newLeaf.Pointers[N-1] = leaf.Pointers[N-1]
	leaf.Pointers[N-1] = newLeaf
	// set the indices after insertion point to nil
	for i = leaf.NumKeys; i < N-1; i++ {
		leaf.Pointers[i] = nil
		leaf.TailPointers[i] = nil
	}
	for i = newLeaf.NumKeys; i < N-1; i++ {
		newLeaf.Pointers[i] = nil
		newLeaf.TailPointers[i] = nil
	}

	// point to the left for dup keys
	newLeaf.Parent = leaf.Parent
	newKey = newLeaf.Keys[0]

	return t.insertIntoParent(leaf, newKey, newLeaf)
}

func insertIntoNode(n *Node, leftIndex int, key float64, right *Node) {
	var i int
	for i = n.NumKeys; i > leftIndex; i-- {
		n.Pointers[i+1] = n.Pointers[i]
		n.Keys[i] = n.Keys[i-1]
	}
	n.Pointers[leftIndex+1] = right
	n.Keys[leftIndex] = key
	n.NumKeys++
}

// insertion into internal node
func (t *Tree) insertIntoNodeAfterSplitting(oldNode *Node, leftIndex int, key float64, right *Node) error {
	var i, j, split int
	var kPrime float64
	var newNode, child *Node
	var err error

	tempPointers := make([]interface{}, N+1)
	tempKeys := make([]float64, N)

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

	// Kprime is the key that will be inserted into the parent
	kPrime = tempKeys[split]

	j = 0
	for i++; i < N; i++ {
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

	// adjust parent pointer of the two nodes
	return t.insertIntoParent(oldNode, kPrime, newNode)
}

func (t *Tree) insertIntoParent(left *Node, key float64, right *Node) error {
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

func (t *Tree) insertIntoNewRoot(left *Node, key float64, right *Node) error {
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
