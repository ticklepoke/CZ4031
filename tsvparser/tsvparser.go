package tsvparser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ParseTSV - reads tsv file and returns parsed contents
func ParseTSV(file string) [][]string {
	tsv, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	scn := bufio.NewScanner(tsv)

	// raw table w/o parsing
	var table []string

	// insert row into table slice
	for scn.Scan() {
		row := scn.Text()
		table = append(table, row)
	}

	if err := scn.Err(); err != nil {
		fmt.Println(err)
		return nil
	}

	// remove header
	table = table[1:]

	// parsed table
	res := make([][]string, len(table))

	// split by tab delimitter and insert into res slice
	for i, row := range table {
		record := strings.Split(row, "\t")
		res[i] = record
	}

	return res
}
