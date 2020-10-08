package main

import (
	"strconv"

	"github.com/ticklepoke/CZ4031/bptree"
	"github.com/ticklepoke/CZ4031/tsvparser"
)

/* Delete those movies with the attribute “averageRating” equal to 7, update the B+ tree accordingly, and report the following statistics:
- the number of times that a node is deleted (or two nodes are merged) during the process of the updating the B+ tree;
- the number nodes of the updated B+ tree;
- the height of the updated B+ tree;
- the root node and its child nodes of the updated B+ tree;
*/

func main() {
	n := 5
	t := bptree.NewTree(n)
	rows := tsvparser.ParseTSV("../data.tsv")

	for _, s := range rows {
		tconsts, ratingString, votes := s[0], s[1], s[2]
		rating, _ := strconv.Atoi(ratingString)
		t.Insert(rating, []byte(tconsts+votes))
	}
	t.Delete(7)
	t.PrintTree()
	t.PrintHeight()
	t.PrintLeaves()
}
