package bptree

import (
	"errors"
	"fmt"

	"github.com/ticklepoke/CZ4031/logger"

	"github.com/ticklepoke/CZ4031/blockmanager"
)

// Find the records for a given key
func (t *Tree) Find(key float64, verbose bool) ([]*Record, error) {
	i := 0
	c, numIndexNodes := t.findLeaf(key, verbose)
	fmt.Println("Number of Index Nodes Accessed: ", numIndexNodes)
	fmt.Println()
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

	logger.Logger.Println("Printing the attribute tconst of the records that are returned")
	// TODO: have to traverse linked list and print out
	if err != nil || r == nil {
		fmt.Printf("Record not found under key %f.\n", key)
	} else {
		for _, recordPtr := range r {
			logger.Logger.Printf("Record -- key %f, ", key)
			blockmanager.PrintRecord(recordPtr.Value)
			logger.Logger.Println()
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
	fmt.Println("Printing the attribute tconst of the records that are returned")
	if numFound == 0 {
		fmt.Println("none found")
	} else {
		for i = 0; i < numFound; i++ {
			c, _ := returnedPointers[i].(*Record)
			fmt.Printf("Key: %f ",
				returnedKeys[i])
			blockmanager.PrintRecord(c.Value)
			t.BlckMngr.SetBlocksAccessed(c.Value)
		}
	}
}

/* ============================ Private Methods ============================*/

func (t *Tree) findRange(keyStart, keyEnd float64, verbose bool, returnedKeys *[]float64, returnedPointers *[]interface{}) int {
	var i, left_bound, numIndexNodes int
	numFound := 0

	n, numIndexNodes := t.findLeaf(keyStart, verbose)
	if n == nil {
		return 0
	}
	for left_bound = 0; left_bound < n.NumKeys && n.Keys[left_bound] < keyStart; left_bound++ {
		if left_bound == n.NumKeys { // could be wrong
			return 0
		}
	}
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
		numIndexNodes++
		if n != nil {
			fmt.Printf("Leaf Node %d %f\n", numIndexNodes, n.Keys[:n.NumKeys])
		}
	}
	numIndexNodes-- //prevent double counting of the first leaf node
	fmt.Printf("Number of Index Nodes Accessed: %v\n\n", numIndexNodes)
	return numFound
}

// TODO: modify to traverse and find all the same keys
func (t *Tree) findLeaf(key float64, verbose bool) (*Node, int) {
	i := 0
	c := t.Root

	if c == nil {
		if verbose {
			fmt.Printf("Empty tree.\n")
		}
		return c, 0
	}
	noOfIndexNodes := 0
	// traverse down the tree till reach leaf node
	for !c.IsLeaf {

		noOfIndexNodes++
		if verbose {
			// TODO
			// create temp buffer to store results
			// helps formatting in the logger
			// stdout := os.Stdout
			// r, w, _ := os.Pipe()
			// os.Stdout = w

			fmt.Printf("Index node %d keys [", noOfIndexNodes)
			for i = 0; i < c.NumKeys-1; i++ {
				fmt.Printf("%f ", c.Keys[i])
			}
			fmt.Printf("%f]", c.Keys[i])

			// close temp buffer and output result to log
			// w.Close()
			// out, _ := ioutil.ReadAll(r)
			// os.Stdout = stdout
			// logger.Logger.Println(out)
		}
		i = 0
		for i < c.NumKeys {
			if key >= c.Keys[i] {
				i++
			} else {
				break
			}
		}
		c, _ = c.Pointers[i].(*Node)
	}

	noOfIndexNodes++ // add one for child node
	// TODO: modify c to factor in slice
	if verbose {
		fmt.Printf("Leaf Node %d [", noOfIndexNodes)
		for i = 0; i < c.NumKeys-1; i++ {
			fmt.Printf("%f ", c.Keys[i])
		}
		fmt.Printf("%f]\n", c.Keys[i])
	}
	return c, noOfIndexNodes
}
