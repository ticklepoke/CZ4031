package main

import (
	"github.com/ticklepoke/CZ4031/blockmanager"
	"github.com/ticklepoke/CZ4031/tsvparser"
)

func main() {
	b := blockmanager.InitializeBlockManager(100)

	rows := tsvparser.ParseTSV("data.tsv")

	for _, row := range rows {
		tconts, rating, votes := row[0], row[1], row[2]

		// TODO: insert record to bptree
		b.InsertRecord(tconts, rating, votes)
	}

	b.DisplayStatus(true)
}
