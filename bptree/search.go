package bptree

import (
	"errors"
	"fmt"

	"github.com/ticklepoke/CZ4031/blockmanager"
)

// Find the records for a given key
func (t *Tree) Find(key float64, verbose bool) ([]*Record, error) {
	i := 0
	c := t.findLeaf(key, verbose)
	if c == nil {
		return nil, errors.New("key not found")
	}

	for i = 0; i < c.NumKeys; i++ {
		if c.Keys[i] == key {
			break
		}
	}

	if i == c.NumKeys {
		return nil, errors.New("key not found")
	}

	curr := c.Pointers[i].(*Record)
	var recordsArr []*Record = iterLeafLL(curr)
	return recordsArr, nil
}

// FindAndPrint returns void
func (t *Tree) FindAndPrint(key float64, verbose bool) {
	r, err := t.Find(key, verbose)

	// TODO: have to traverse linked list and print out
	if err != nil || r == nil {
		fmt.Printf("Record not found under key %f.\n", key)
	} else {
		for _, recordPtr := range r {
			fmt.Printf("Record at %p -- key %f, value %s.\n", recordPtr, key, recordPtr.Value)
			blockmanager.PrintRecord(recordPtr.Value)
			fmt.Println()
			t.BlckMngr.SetBlocksAccessed(recordPtr.Value)
		}
	}
}

// FindAndPrintRange returns void
func (t *Tree) FindAndPrintRange(keyStart, keyEnd float64, verbose bool) {
	var i int
	returnedKeys := make([]float64, 0)
	returnedPointers := make([]interface{}, 0)
	numFound := t.findRange(keyStart, keyEnd, verbose, &returnedKeys, &returnedPointers)
	if numFound == 0 {
		fmt.Println("none found")
	} else {
		for i = 0; i < numFound; i++ {
			c, _ := returnedPointers[i].(*Record)
			fmt.Printf("Key: %f  Location: %p ",
				returnedKeys[i],
				returnedPointers[i])
			blockmanager.PrintRecord(c.Value)
			t.BlckMngr.SetBlocksAccessed(c.Value)
		}
	}
}

/* ============================ Private Methods ============================*/

func (t *Tree) findRange(keyStart, keyEnd float64, verbose bool, returnedKeys *[]float64, returnedPointers *[]interface{}) int {
	var i, left_bound int
	numFound := 0

	n := t.findLeaf(keyStart, verbose)
	if n == nil {
		return 0
	}
	for left_bound = 0; left_bound < n.NumKeys && n.Keys[left_bound] < keyStart; left_bound++ {
		if left_bound == n.NumKeys { // could be wrong
			return 0
		}
	}
	t.PrintTree()
	for n != nil { // traverse right
		for i = left_bound; i < n.NumKeys && n.Keys[i] <= keyEnd; i++ {
			curr := n.Pointers[i]
			var recordPtrs []*Record = iterLeafLL(curr.(*Record))

			for _, recPtr := range recordPtrs {
				// returnedPointers[numFound] = recPtr
				*returnedPointers = append(*returnedPointers, recPtr)
				*returnedKeys = append(*returnedKeys, n.Keys[i])
				numFound++
			}
		}
		// n = n.Next //go to the next leaf node
		n, _ = n.Pointers[N-1].(*Node)
		i = 0
	}
	return numFound
}

// TODO: modify to traverse and find all the same keys
func (t *Tree) findLeaf(key float64, verbose bool) *Node {
	i := 0
	c := t.Root

	if c == nil {
		if verbose {
			fmt.Printf("Empty tree.\n")
		}
		return c
	}

	// traverse down the tree till reach leaf node
	for !c.IsLeaf {
		if verbose {
			fmt.Printf("[")
			for i = 0; i < c.NumKeys-1; i++ {
				fmt.Printf("%f ", c.Keys[i])
			}
			fmt.Printf("%f]", c.Keys[i])
		}
		i = 0
		for i < c.NumKeys {
			if key >= c.Keys[i] {
				i++
			} else {
				break
			}
		}
		if verbose {
			fmt.Printf("%d ->\n", i)
		}
		c, _ = c.Pointers[i].(*Node)
	}

	// TODO: modify c to factor in slice
	if verbose {
		fmt.Printf("Leaf [")
		for i = 0; i < c.NumKeys-1; i++ {
			fmt.Printf("%f ", c.Keys[i])
		}
		fmt.Printf("%f] ->\n", c.Keys[i])
	}
	return c
}
