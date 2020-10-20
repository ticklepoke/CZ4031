package blockmanager

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/ticklepoke/CZ4031/logger"
)

// DisplayStatus - get status of blockmanager and print the state of the
// current block
func (b BlockManager) DisplayStatus(verbose bool) {
	numBlocks, numRecords := b.numBlocks, b.numRecords
	logger.Logger.Printf("num blocks:\t%s\n", strconv.Itoa(numBlocks))
	logger.Logger.Printf("num records:\t%s\n", strconv.Itoa(numRecords))

	if verbose {
		fmt.Printf("current count:\t%s\n", strconv.Itoa(b.currentCount))
		fmt.Printf("has capacity:\t%s\n", strconv.FormatBool(b.hasCapacity))
		fmt.Printf("block size:\t%s\n\n", strconv.Itoa(b.BLOCKSIZE))

		fmt.Println("deleted records:", b.markedDeleted)
		// fmt.Println("byte block:", b.blocks[b.numBlocks-1])
	}
	logger.Logger.Println()
	b.printRecords(verbose)
}

// PrintRecord - parses and prints bytes slice at record address
func PrintRecord(recAddr *[]byte) {
	tconst := string(bytes.Trim((*recAddr)[:RATINGOFFSET], "\x00"))
	rating := string(bytes.Trim((*recAddr)[RATINGOFFSET:VOTESOFFSET], "\x00"))
	votes := string(bytes.Trim((*recAddr)[VOTESOFFSET:RECORDSIZE], "\x00"))

	logger.Logger.Printf("tconst: %s, average rating: %s, number of votes: %s \n", tconst, rating, votes)
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
	logger.Logger.Println()
}

// SetBlocksAccessed - keeps track of the blocks that have been accessed during deletion and record retrieval
func (b *BlockManager) SetBlocksAccessed(addr *[]byte) {
	// start := unsafe.Pointer(addr)
	// size := b.BLOCKSIZE

	// use offsets to check which block the record belongs to
	// this works as the records have a fixed size
	// for i := 0; i < size/RECORDSIZE; i++ {
	// 	item := (unsafe.Pointer(uintptr(start) - uintptr(i*32)))
	if block, found := b.blockSet[addr]; found && !b.visited[block] {
		b.visited[block] = true
	}
	// }
}

// GetBlocksAccessed - prints the blocks accessed during deletion and record retrieval
func (b BlockManager) GetBlocksAccessed() {
	var count int
	logger.Logger.Println("=================")
	logger.Logger.Println(" blocks accessed ")
	logger.Logger.Println("=================")
	for block, visited := range b.visited {
		if visited {
			count++
			// block := *(*[]byte)(block)
			for j := 0; j < b.BLOCKSIZE/RECORDSIZE; j++ {
				offset := j * RECORDSIZE

				if bytes.Compare((*block)[offset:offset+RECORDSIZE], make([]byte, RECORDSIZE, RECORDSIZE)) == 0 {
					continue
				}

				record := (*block)[offset:]
				PrintRecord(&record)
			}
			logger.Logger.Println()
		}
	}
	logger.Logger.Printf("Num blocks accessed: %d \n", count)
}

// ResetBlocksAccessed - resets accessed blocks map
func (b *BlockManager) ResetBlocksAccessed() {
	for k := range b.blockSet {
		b.visited[k] = false
	}
}
