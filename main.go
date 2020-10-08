package main

import (
	"fmt"
	"strconv"

	"github.com/ticklepoke/CZ4031/bptree"
)

func main() {

	t := bptree.NewTree(4)

	// Insertion example
	// 1, 4, 7, 10, 17, 21, 31, 25, 19, 20, 28, 42
	// 20  |
	// 7 17  | 25  |
	// 1 4 5  | 7 10  | 17 19  | 20 21  | 25 31  |

	// Deletion example
	// build btree with the following
	// 1, 4, 7, 10, 17, 21, 31, 25, 19, 20, 5
	// delete 5
	// 20  |
	// 7 17  | 25  |
	// 1 4  | 7 10  | 17 19  | 20 21  | 25 31  |

	arr := []int{1, 4, 7, 10, 17, 21, 31, 25, 19, 20, 28, 42, 42, 42, 42}
	for _, num := range arr {
		err := t.Insert(num, []byte("hello friend"+strconv.Itoa(num)))
		if err != nil {
			fmt.Printf("error: %s\n\n", err)
		}
	}

	t.PrintTree()
	t.FindAndPrint(42, false)
	t.FindAndPrint(4, false)
	t.FindAndPrint(21, false)

	t.Delete(5)
	t.Delete(42)
	t.PrintTree()

	// b := blockmanager.InitializeBlockManager(100)

	// testing out tsv
	// rows := tsvparser.ParseTSV("data.tsv")

	// for i := 0; i < 10; i++ {
	// 	tconts, rating, votes := rows[i][0], rows[i][1], rows[i][2]

	// 	// TODO: insert record to bptree
	// 	b.InsertRecord(tconts, rating, votes)
	// }

	// b.DisplayStatus(true)
}
