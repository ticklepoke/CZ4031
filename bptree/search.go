package bptree

import (
	"errors"
	"fmt"
)

// search right hand side for linked list
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
		// TODO: traverse to next block to find key?
		return nil, errors.New("key not found")
	}

	r, _ := c.Pointers[i].(*Record)

	return r, nil
}

/**
* TODO: Modify this to return a range of records that return the same key
*/
func (t *Tree) FindAndPrint(key int, verbose bool) {
	r, err := t.Find(key, verbose)

	if err != nil || r == nil {
		fmt.Printf("Record not found under key %d.\n", key)
	} else {
		fmt.Printf("Record at %d -- key %d, value %s.\n", r, key, r.Value)
	}
}

/**
* TODO: Modify to return a range of records in the range
*/
func (t *Tree) FindAndPrintRange(key_start, key_end int, verbose bool) {
	var i int
	array_size := key_end - key_start + 1
	returned_keys := make([]int, array_size)
	returned_pointers := make([]interface{}, array_size)
	num_found := t.findRange(key_start, key_end, verbose, returned_keys, returned_pointers)
	if num_found == 0 {
		fmt.Println("None found,\n")
	} else {
		for i = 0; i < num_found; i++ {
			c, _ := returned_pointers[i].(*Record)
			fmt.Printf("Key: %d  Location: %d  Value: %s\n",
				returned_keys[i],
				returned_pointers[i],
				c.Value)
		}
	}
}

// implement search rhs at leaf nodes (for the duplicate keys)
func (t *Tree) findRange(key_start, key_end int, verbose bool, returned_keys []int, returned_pointers []interface{}) int {
	var i int
	num_found := 0

	n := t.findLeaf(key_start, verbose)
	if n == nil {
		return 0
	}
	for i = 0; i < n.NumKeys && n.Keys[i] < key_start; i++ {
		if i == n.NumKeys { // could be wrong
			return 0
		}
	}
	for n != nil {
		for i = i; i < n.NumKeys && n.Keys[i] <= key_end; i++ {
			returned_keys[num_found] = n.Keys[i]
			returned_pointers[num_found] = n.Pointers[i]
			num_found += 1
		}
		n, _ = n.Pointers[order-1].(*Node)
		i = 0
	}
	return num_found
}

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
				i += 1
			} else {
				break
			}
		}
		if verbose {
			fmt.Printf("%d ->\n", i)
		}
		c, _ = c.Pointers[i].(*Node)
	}
	if verbose {
		fmt.Printf("Leaf [")
		for i = 0; i < c.NumKeys-1; i++ {
			fmt.Printf("%d ", c.Keys[i])
		}
		fmt.Printf("%d] ->\n", c.Keys[i])
	}
	return c
}
