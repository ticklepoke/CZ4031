package bptree

import (
	"fmt"
	"reflect"

	"github.com/ticklepoke/CZ4031/logger"
)

// Delete - implement deletion logic for node w/o parent pointers
func (t *Tree) Delete(key float64) error {
	logger.Logger.Println("Deleting node: " + fmt.Sprintf("%f", key))
	keyRecords, err := t.Find(key, false)
	if err != nil {
		return err
	}
	keyLeaf, _ := t.findLeaf(key, false)
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

	// if last key, set to 0
	if i == n.NumKeys-1 {
		n.Keys[i] = 0
	}

	// move all keys after selected to the left
	for i++; i < n.NumKeys; i++ {
		n.Keys[i-1] = n.Keys[i]
	}

	if n.IsLeaf {
		num_pointers = n.NumKeys
	} else {
		num_pointers = n.NumKeys + 1
	}

	i = 0
	// find pointer to remove
	for n.Pointers[i] != pointer {
		i++
	}

	// move all pointers after selected to the left
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

// adjust root if underflowed
func (t *Tree) adjustRoot() {
	var new_root *Node

	// root must have at least one key
	if t.Root.NumKeys > 0 {
		return
	}

	// if empty root has a child, promote first child as new root
	if !t.Root.IsLeaf {
		new_root, _ = t.Root.Pointers[0].(*Node)
		new_root.Parent = nil
	} else {
		new_root = nil
	}
	t.Root = new_root
}

// merge two nodes
func (t *Tree) coalesceNodes(right, left *Node, neighbour_index int, k_prime float64) {
	var i, j, insertion_index, n_end int
	var tmp *Node

	// swap left and right if node in prev step was on the left
	if neighbour_index == -1 {
		tmp = right
		right = left
		left = tmp
	}

	logger.Logger.Debugln("Coalescing nodes", left.Keys[:left.NumKeys], right.Keys[:right.NumKeys])
	insertion_index = left.NumKeys

	if !right.IsLeaf {
		// smallest key in right subtree is k_prime
		left.Keys[insertion_index] = k_prime
		left.NumKeys++

		n_end = right.NumKeys
		i = insertion_index + 1
		for j = 0; j < n_end; j++ {
			left.Keys[i] = right.Keys[j]
			left.Pointers[i] = right.Pointers[j]
			left.NumKeys++
			right.NumKeys--
			i++
		}
		left.Pointers[i] = right.Pointers[j]

		// update parent pointer for new children
		for i = 0; i < left.NumKeys+1; i++ {
			tmp, _ = left.Pointers[i].(*Node)
			tmp.Parent = left
		}
	} else {
		i = insertion_index
		for j = 0; j < right.NumKeys; j++ {
			left.Keys[i] = right.Keys[j]
			left.Pointers[i] = right.Pointers[j]
			left.NumKeys++
			i++
		}
		left.Pointers[N-1] = right.Pointers[N-1]
	}

	logger.Logger.Debugln("Finished node", left.Keys[:left.NumKeys])
	logger.Logger.Debugf("Removing %f from parent %v\n", k_prime, right.Parent.Keys[:right.Parent.NumKeys])
	t.deleteEntry(right.Parent, k_prime, right)
}

func (t *Tree) redistributeNodes(right, left *Node, neighbour_index, k_prime_index int, k_prime float64) {
	var i int
	var tmp *Node

	logger.Logger.Debugln("Redistributing nodes", right.Keys[:right.NumKeys], left.Keys[:left.NumKeys])
	if neighbour_index != -1 {
		if !right.IsLeaf {
			right.Pointers[right.NumKeys+1] = right.Pointers[right.NumKeys]
		}
		for i = right.NumKeys; i > 0; i-- {
			right.Keys[i] = right.Keys[i-1]
			right.Pointers[i] = right.Pointers[i-1]
		}
		if !right.IsLeaf {
			right.Pointers[0] = left.Pointers[left.NumKeys]
			tmp, _ = right.Pointers[0].(*Node)
			tmp.Parent = right
			left.Pointers[left.NumKeys] = nil
			right.Keys[0] = k_prime
			right.Parent.Keys[k_prime_index] = left.Keys[left.NumKeys-1]
		} else {
			right.Pointers[0] = left.Pointers[left.NumKeys-1]
			left.Pointers[left.NumKeys-1] = nil
			right.Keys[0] = left.Keys[left.NumKeys-1]
			right.Parent.Keys[k_prime_index] = right.Keys[0]
		}
	} else {
		if right.IsLeaf {
			right.Keys[right.NumKeys] = left.Keys[0]
			right.Pointers[right.NumKeys] = left.Pointers[0]
			right.Parent.Keys[k_prime_index] = left.Keys[1]
		} else {
			right.Keys[right.NumKeys] = k_prime
			right.Pointers[right.NumKeys+1] = left.Pointers[0]
			tmp, _ = right.Pointers[right.NumKeys+1].(*Node)
			tmp.Parent = right
			right.Parent.Keys[k_prime_index] = left.Keys[0]
		}
		for i = 0; i < left.NumKeys-1; i++ {
			left.Keys[i] = left.Keys[i+1]
			left.Pointers[i] = left.Pointers[i+1]
		}
		if !right.IsLeaf {
			left.Pointers[i] = left.Pointers[i+1]
		}
	}
	right.NumKeys++
	left.NumKeys--
}

// delete key and pointer from node
func (t *Tree) deleteEntry(n *Node, key float64, pointer interface{}) {
	var min_keys, neighbour_index, k_prime_index, capacity int
	var k_prime float64
	var neighbour *Node

	logger.Logger.Println("Removing key", key, "from node", n.Keys[:n.NumKeys])
	n = removeEntryFromNode(n, key, pointer)

	if n == t.Root {
		logger.Logger.Info("Adjusting root")
		t.adjustRoot()
		return
	}

	if n.IsLeaf {
		min_keys = (N) / 2
	} else {
		min_keys = (N - 1) / 2
	}

	if n.NumKeys >= min_keys {
		t.adjustParentKeys(n, n.Keys[0])
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

	if n.IsLeaf {
		capacity = N - 1
	} else {
		// one key will be inherited from parent
		capacity = N - 2
	}
	// capacity = N - 1

	logger.Logger.Debugf("Max capacity for coalescing %d, current capacity %d\n", capacity, neighbour.NumKeys+n.NumKeys)
	if neighbour.NumKeys+n.NumKeys <= capacity {
		t.numDeletions++
		t.coalesceNodes(n, neighbour, neighbour_index, k_prime)
	} else {
		t.redistributeNodes(n, neighbour, neighbour_index, k_prime_index, k_prime)
	}

	if neighbour_index != -1 {
		n = neighbour
	}

	t.adjustParentKeys(n, n.Keys[0])
}

// ensure that smallest value in right subtree is promoted to parent
func (t *Tree) adjustParentKeys(n *Node, small float64) {
	// return
	if !n.IsLeaf {
		return
	}

	for n.Parent != nil {
		if n.Parent.Pointers[0] == n {
			n = n.Parent
		} else {
			break
		}
	}

	// reached root as leftmost subtree
	// no changes needed
	if t.Root == n {
		return
	}

	for i := 0; i < n.Parent.NumKeys+1; i++ {
		if n.Parent.Pointers[i] == n {
			n.Parent.Keys[i-1] = small
			return
		}
	}
}
