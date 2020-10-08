package bptree

import (
	"fmt"
)

func (t *Tree) PrintTree() {
	var n *Node
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
		if n != nil {
			if n.Parent != nil && n == n.Parent.Pointers[0] {
				new_rank = t.pathToRoot(n)
				if new_rank != rank {
					fmt.Printf("\n")
					rank = new_rank
				}
			}
			if verbose_output {
				fmt.Printf("(%d)", n)
			}
			for i = 0; i < n.NumKeys; i++ {
				if verbose_output {
					fmt.Printf("%d ", n.Pointers[i])
				}
				fmt.Printf("%d ", n.Keys[i])
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
	fmt.Printf("\n")
}

func (t *Tree) PrintLeaves() {
	if t.Root == nil {
		fmt.Printf("Empty tree.\n")
		return
	}

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
			fmt.Printf("%d ", c.Keys[i])
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

func (t *Tree) height() int {
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

// cut - return num keys needed during insertion and deletion
func cut(length int) int {
	if length%2 == 0 {
		return length / 2
	}

	return length/2 + 1
}
