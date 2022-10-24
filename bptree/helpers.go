package bptree

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"

	"github.com/ticklepoke/CZ4031/logger"
)

func (t *Tree) PrintTree() {

	// create temp buffer to store traversal results
	// helps formatting in the logger
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	var n *Node
	numberOfNodes := 0
	i := 0
	rank := 0
	new_rank := 0

	if t.Root == nil {
		fmt.Printf("Empty tree.\n")
		return
	}
	queue = nil
	enqueue(t.Root)
	for queue != nil {
		n = dequeue()
		numberOfNodes++
		if n != nil {
			if n.Parent != nil && n == n.Parent.Pointers[0] {
				new_rank = t.pathToRoot(n)
				if new_rank != rank {
					fmt.Printf("\n\n")
					rank = new_rank
				}
			}
			if verbose_output {
				fmt.Printf("(%v)", n)
			}
			for i = 0; i < n.NumKeys; i++ {
				if verbose_output {
					fmt.Printf("%d ", n.Pointers[i])
				}
				fmt.Printf("%.1f ", n.Keys[i])
			}
			if !n.IsLeaf {
				for i = 0; i <= n.NumKeys; i++ {
					c, _ := n.Pointers[i].(*Node)
					enqueue(c)
				}
			}
			if verbose_output {
				if n.IsLeaf {
					fmt.Printf("%d ", n.Pointers[N-1])
				} else {
					fmt.Printf("%d ", n.Pointers[n.NumKeys])
				}
			}
			fmt.Printf(" | ")
		}
	}
	logger.Logger.Printf("B+ Tree Number of Nodes: %v", numberOfNodes)

	// close temp buffer and output result to log
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = stdout

	logger.Logger.Printf("\n\n================\nBPTREE STRUCTURE\n================\n\n%s\n\n", out)
}

// PrintHeight prints the height of the tree
func (t *Tree) PrintHeight() {
	fmt.Printf("Height: %v\n", strconv.Itoa(t.Height()))
}

// PrintLeaves print leaves
func (t *Tree) PrintLeaves() {
	if t.Root == nil {
		fmt.Printf("Empty tree.\n")
		return
	}

	fmt.Print("Leave nodes: ")

	var i int
	c := t.Root
	for !c.IsLeaf {
		c, _ = c.Pointers[0].(*Node)
	}

	for {
		for i = 0; i < c.NumKeys; i++ {
			if verbose_output {
				fmt.Printf("%d ", c.Pointers[i])
			}
			fmt.Printf("%.1f ", c.Keys[i])
		}
		if verbose_output {
			fmt.Printf("%d ", c.Pointers[N-1])
		}
		if c.Pointers[N-1] != nil {
			fmt.Printf(" | ")
			c, _ = c.Pointers[N-1].(*Node)
		} else {
			break
		}
	}
	fmt.Printf("\n")
}

// used in printing
func enqueue(new_node *Node) {
	var c *Node
	if queue == nil {
		queue = new_node
		queue.Next = nil
	} else {
		c = queue
		for c.Next != nil {
			c = c.Next
		}
		c.Next = new_node
		new_node.Next = nil
	}
}

// used in printing
func dequeue() *Node {
	n := queue
	queue = queue.Next
	n.Next = nil
	return n
}

func (t *Tree) Height() int {
	h := 0
	c := t.Root
	for !c.IsLeaf {
		c, _ = c.Pointers[0].(*Node)
		h++
	}
	return h
}

func (t *Tree) pathToRoot(child *Node) int {
	length := 0
	c := child
	for c != t.Root {
		c = c.Parent
		length += 1
	}
	return length
}

// findMidPoint - return num keys needed during insertion and deletion
func findMidPoint(length int) int {
	return int(math.Ceil(float64(length) / 2))
}

// FindNumDeletions - returns the number of nodes deleted
func (t Tree) FindNumDeletions() {
	logger.Logger.Printf("Number of nodes deleted: %d\n", t.numDeletions)
}
