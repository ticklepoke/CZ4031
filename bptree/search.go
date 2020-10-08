package bptree

import (
	"errors"
	"fmt"
)

// Find returns slice of records, error
// TODO: search right hand side for linked list
func (t *Tree) Find(key int, verbose bool) (*Record, error) {
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

	r, _ := c.Pointers[i].(*Record)

	return r, nil
}

// FindAndPrint returns void
func (t *Tree) FindAndPrint(key int, verbose bool) {
	r, err := t.Find(key, verbose)

	// TODO: have to traverse linked list and print out
	if err != nil || r == nil {
		fmt.Printf("Record not found under key %d.\n", key)
	} else {
		fmt.Printf("Record at %d -- key %d, value %s.\n", r, key, r.Value)
	}
}

// FindAndPrintRange returns void
func (t *Tree) FindAndPrintRange(keyStart, keyEnd int, verbose bool) {
	var i int
	arraySize := keyEnd - keyStart + 1
	returnedKeys := make([]int, arraySize)
	returnedPointers := make([]interface{}, arraySize)
	numFound := t.findRange(keyStart, keyEnd, verbose, returnedKeys, returnedPointers)
	if numFound == 0 {
		fmt.Println("none found")
	} else {
		for i = 0; i < numFound; i++ {
			c, _ := returnedPointers[i].(*Record)
			fmt.Printf("Key: %d  Location: %d  Value: %s\n",
				returnedKeys[i],
				returnedPointers[i],
				c.Value)
		}
	}
}

/* ============================ Private Methods ============================*/

// implement search rhs at leaf nodes (for the duplicate keys)
func (t *Tree) findRange(keyStart, keyEnd int, verbose bool, returnedKeys []int, returnedPointers []interface{}) int {
	var i int
	numFound := 0

	n := t.findLeaf(keyStart, verbose) // TODO: n is now an array of nodes, need to refactor to handle this
	if n == nil {
		return 0
	}
	for i = 0; i < n.NumKeys && n.Keys[i] < keyStart; i++ {
		if i == n.NumKeys { // could be wrong
			return 0
		}
	}
	for n != nil {
		// for i = i; i < n.NumKeys && n.Keys[i] <= keyEnd; i++ {
		for i = 0; i < n.NumKeys && n.Keys[i] <= keyEnd; i++ {
			returnedKeys[numFound] = n.Keys[i]
			returnedPointers[numFound] = n.Pointers[i]
			numFound++
		}
		n, _ = n.Pointers[order-1].(*Node)
		i = 0
	}
	return numFound
}

// TODO: modify to traverse and find all the same keys
func (t *Tree) findLeaf(key int, verbose bool) *Node {
	i := 0
	c := t.Root

	if c == nil {
		if verbose {
			fmt.Printf("Empty tree.\n")
		}
		return c
	}

	for !c.IsLeaf {
		if verbose {
			fmt.Printf("[")
			for i = 0; i < c.NumKeys-1; i++ {
				fmt.Printf("%d ", c.Keys[i])
			}
			fmt.Printf("%d]", c.Keys[i])
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
			fmt.Printf("%d ", c.Keys[i])
		}
		fmt.Printf("%d] ->\n", c.Keys[i])
	}
	return c
}
