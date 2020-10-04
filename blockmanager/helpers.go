package blockmanager

import (
	"bytes"
	"fmt"
)

// DisplayStatus - get status of blockmanager and print the state of the
// current block
func (b BlockManager) DisplayStatus(verbose bool) {
	if verbose {
		fmt.Println(b)
		fmt.Println("deleted records", b.markedDeleted)
		fmt.Println("byte block", *b.blocks[b.numBlocks-1])
	}
	b.printRecords(true)
}

// PrintRecord - parses and prints bytes slice at record address
func PrintRecord(recAddr *[]byte) {
	tconst := string((*recAddr)[:RATINGOFFSET])
	rating := string((*recAddr)[RATINGOFFSET:VOTESOFFSET])
	votes := string((*recAddr)[VOTESOFFSET:RECORDSIZE])

	fmt.Printf("id: %s, average rating: %s, number of votes: %s \n", tconst, rating, votes)
}

// printRecords - prints all records in the database
// setting all to false prints records in current block only
func (b BlockManager) printRecords(all bool) {
	start := 0
	if all == false {
		start = b.numBlocks - 1
	}

	for i := start; i < b.numBlocks; i++ {
		block := *b.blocks[i]
		for j := 0; j < b.BLOCKSIZE/RECORDSIZE; j++ {
			offset := j * RECORDSIZE

			if bytes.Compare(block[offset:offset+RECORDSIZE], make([]byte, RECORDSIZE, RECORDSIZE)) == 0 {
				continue
			}

			record := block[offset:]
			PrintRecord(&record)
		}
	}
	fmt.Println()
}
