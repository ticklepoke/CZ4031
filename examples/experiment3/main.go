package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ticklepoke/CZ4031/bptree"
	"github.com/ticklepoke/CZ4031/tsvparser"
)

func main() {
	start := time.Now()
	n := 5
	t := bptree.NewTree(n)
	rows := tsvparser.ParseTSV("../../data.tsv")

	for _, s := range rows {
		tconsts, ratingString, votes := s[0], s[1], s[2]
		rating, _ := strconv.Atoi(ratingString)
		t.Insert(rating, []byte(tconsts+votes))
	}

	t.FindAndPrint(8, true)

	elapse := time.Since(start)
	fmt.Println("Experiment 3 elapsed time: ", elapse)
}
