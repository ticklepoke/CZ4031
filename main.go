package main

import (
	"fmt"
	"strconv"

	"github.com/ticklepoke/CZ4031/blockmanager"
	"github.com/ticklepoke/CZ4031/bptree"
)

func main() {

	t := bptree.NewTree(4)
	b := blockmanager.InitializeBlockManager(100)

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

	// delete 4
	// 20  |
	// 17  | 25  |
	// 1 7 10  | 17 19  | 20 21  | 25 31  |

	arr := []float64{1, 4, 7, 10, 17, 21, 31, 25, 19, 20, 5}
	for _, num := range arr {
		temp := strconv.Itoa(num)
		addr := b.InsertRecord(temp, temp, temp)
		err := t.Insert(num, addr)
		if err != nil {
			fmt.Printf("error: %s\n\n", err)
		}
	}

	t.PrintTree()

	fmt.Println("Deleting 17")
	t.Delete(17)
	t.PrintTree()

	fmt.Println("Deleting 19")
	t.Delete(19)
	t.PrintTree()

	fmt.Println("Deleting 19")
	t.Delete(7)
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
