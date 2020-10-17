package experiment

import (
	"github.com/ticklepoke/CZ4031/bptree"
)

/* Delete those movies with the attribute “averageRating” equal to 7, update the B+ tree accordingly, and report the following statistics:
- the number of times that a node is deleted (or two nodes are merged) during the process of the updating the B+ tree;
- the number nodes of the updated B+ tree;
- the height of the updated B+ tree;
- the root node and its child nodes of the updated B+ tree;
*/

func experiment5(t *bptree.Tree) {
	t.Delete(7)
	t.PrintTree()
	t.PrintHeight()
	t.PrintLeaves()
}
