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
	parseRecords(*b.blocks[b.numBlocks-1], b.currentCount)
}

func parseRecords(block []byte, currCount int) {
	for i := 0; i < currCount; i++ {
		offset := i * RECORDSIZE
		tconst := string(block[offset : offset+10])
		rating := string(block[offset+10 : offset+13])
		votes := string(block[offset+13 : offset+20])

		if bytes.Compare(block[offset:offset+20], make([]byte, RECORDSIZE, RECORDSIZE)) == 0 {
			continue
		}

		fmt.Printf("id: %s, average rating: %s, number of votes: %s \n", tconst, rating, votes)
	}
	fmt.Println()
}
