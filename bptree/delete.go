package bptree

import (
	"fmt"
	"reflect"
)

// Delete - implement deletion logic for node w/o parent pointers
func (t *Tree) Delete(key float64) error {
	fmt.Println("Deleting node: " + fmt.Sprintf("%f", key))
	keyRecords, err := t.Find(key, false)
	if err != nil {
		return err
	}
	keyLeaf := t.findLeaf(key, false) // TODO: this becomes a slice of leaf nodes
	if len(keyRecords) != 0 && keyLeaf != nil {
		keyRecordHead := keyRecords[0]
		t.deleteEntry(keyLeaf, key, keyRecordHead)
	}
	return nil
}

func getNeighbourIndex(n *Node) int {
	var i int

	for i = 0; i <= n.Parent.NumKeys; i++ {
		if reflect.DeepEqual(n.Parent.Pointers[i], n) {
			return i - 1
		}
	}

	return i
}

func removeEntryFromNode(n *Node, key float64, pointer interface{}) *Node {
	var i, num_pointers int

	for n.Keys[i] != key {
		i++
	}
	// i is the index of the key
	for i++; i < n.NumKeys; i++ {
		n.Keys[i-1] = n.Keys[i]
	}

	if n.IsLeaf {
		num_pointers = n.NumKeys
	} else {
		num_pointers = n.NumKeys + 1
	}

	i = 0
	for n.Pointers[i] != pointer {
		i++
	}
	for i += 1; i < num_pointers; i++ {
		n.Pointers[i-1] = n.Pointers[i]
	}
	n.NumKeys--

	if n.IsLeaf {
		for i = n.NumKeys; i < N-1; i++ {
			n.Pointers[i] = nil
		}
	} else {
		for i = n.NumKeys + 1; i < N; i++ {
			n.Pointers[i] = nil
		}
	}

	return n
}

// figure out what this does
func (t *Tree) adjustRoot() {
	var new_root *Node

	if t.Root.NumKeys > 0 {
		return
	}

	if !t.Root.IsLeaf {
		new_root, _ = t.Root.Pointers[0].(*Node)
		new_root.Parent = nil
	} else {
		new_root = nil
	}
	t.Root = new_root

	return
}

func (t *Tree) coalesceNodes(n, neighbour *Node, neighbour_index int, k_prime float64) {
	// combine two nodes
	var i, j, neighbour_insertion_index, n_end int
	var tmp *Node

	if neighbour_index == -1 { // swap n and neighbour node var to standardize neighbour , n
		tmp = n
		n = neighbour
		neighbour = tmp // neighbour is the left
	}

	neighbour_insertion_index = neighbour.NumKeys

	if !n.IsLeaf {
		neighbour.Keys[neighbour_insertion_index] = k_prime
		neighbour.NumKeys++

		n_end = n.NumKeys
		i = neighbour_insertion_index + 1
		for j = 0; j < n_end; j++ {
			neighbour.Keys[i] = n.Keys[j]
			neighbour.Pointers[i] = n.Pointers[j]
			neighbour.NumKeys++
			n.NumKeys--
			i++
		}
		neighbour.Pointers[i] = n.Pointers[j]

		for i = 0; i < neighbour.NumKeys+1; i++ {
			tmp, _ = neighbour.Pointers[i].(*Node)
			tmp.Parent = neighbour
		}
	} else {
		i = neighbour_insertion_index
		for j = 0; j < n.NumKeys; j++ {
			neighbour.Keys[i] = n.Keys[j]
			neighbour.Pointers[i] = n.Pointers[j]
			neighbour.NumKeys++
			i++
		}
		neighbour.Pointers[N-1] = n.Pointers[N-1]
	}

	t.deleteEntry(n.Parent, k_prime, n)
}

func (t *Tree) redistributeNodes(n, neighbour *Node, neighbour_index, k_prime_index int, k_prime float64) {
	var i int
	var tmp *Node

	if neighbour_index != -1 {
		if !n.IsLeaf {
			n.Pointers[n.NumKeys+1] = n.Pointers[n.NumKeys]
		}
		for i = n.NumKeys; i > 0; i-- {
			n.Keys[i] = n.Keys[i-1]
			n.Pointers[i] = n.Pointers[i-1]
		}
		if !n.IsLeaf { // why the second if !n.IsLeaf
			n.Pointers[0] = neighbour.Pointers[neighbour.NumKeys]
			tmp, _ = n.Pointers[0].(*Node)
			tmp.Parent = n
			neighbour.Pointers[neighbour.NumKeys] = nil
			n.Keys[0] = k_prime
			n.Parent.Keys[k_prime_index] = neighbour.Keys[neighbour.NumKeys-1]
		} else {
			n.Pointers[0] = neighbour.Pointers[neighbour.NumKeys-1]
			neighbour.Pointers[neighbour.NumKeys-1] = nil
			n.Keys[0] = neighbour.Keys[neighbour.NumKeys-1]
			n.Parent.Keys[k_prime_index] = n.Keys[0]
		}
	} else {
		if n.IsLeaf {
			n.Keys[n.NumKeys] = neighbour.Keys[0]
			n.Pointers[n.NumKeys] = neighbour.Pointers[0]
			n.Parent.Keys[k_prime_index] = neighbour.Keys[1]
		} else {
			n.Keys[n.NumKeys] = k_prime
			n.Pointers[n.NumKeys+1] = neighbour.Pointers[0]
			tmp, _ = n.Pointers[n.NumKeys+1].(*Node)
			tmp.Parent = n
			n.Parent.Keys[k_prime_index] = neighbour.Keys[0]
		}
		for i = 0; i < neighbour.NumKeys-1; i++ {
			neighbour.Keys[i] = neighbour.Keys[i+1]
			neighbour.Pointers[i] = neighbour.Pointers[i+1]
		}
		if !n.IsLeaf {
			neighbour.Pointers[i] = neighbour.Pointers[i+1]
		}
	}
	n.NumKeys++
	neighbour.NumKeys--

	return
}

func (t *Tree) deleteEntry(n *Node, key float64, pointer interface{}) {
	var min_keys, neighbour_index, k_prime_index, capacity int
	var k_prime float64
	var neighbour *Node

	n = removeEntryFromNode(n, key, pointer)

	if n == t.Root {
		t.adjustRoot()
		return
	}

	if n.IsLeaf {
		min_keys = (N) / 2
	} else {
		min_keys = (N - 1) / 2
	}

	if n.NumKeys >= min_keys {
		return
	}

	neighbour_index = getNeighbourIndex(n) // combine to the left
	if neighbour_index == -1 {
		k_prime_index = 0
	} else {
		k_prime_index = neighbour_index
	}

	k_prime = n.Parent.Keys[k_prime_index]

	if neighbour_index == -1 {
		neighbour, _ = n.Parent.Pointers[1].(*Node) // when most LHS node, neighbour is on the right
	} else {
		neighbour, _ = n.Parent.Pointers[neighbour_index].(*Node)
	}

	capacity = N - 1

	if neighbour.NumKeys+n.NumKeys <= capacity {
		t.coalesceNodes(n, neighbour, neighbour_index, k_prime)
		return
	}
	t.redistributeNodes(n, neighbour, neighbour_index, k_prime_index, k_prime)
	return
}
