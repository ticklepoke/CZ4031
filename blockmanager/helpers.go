package blockmanager

import (
	"bytes"
	"fmt"
	"strconv"
	"unsafe"
)

// DisplayStatus - get status of blockmanager and print the state of the
// current block
func (b BlockManager) DisplayStatus(verbose bool) {
	numBlocks, numRecords := b.numBlocks, b.numRecords
	fmt.Printf("num blocks:\t%s\n", strconv.Itoa(numBlocks))
	fmt.Printf("num records:\t%s\n", strconv.Itoa(numRecords))

	if verbose {
		fmt.Printf("current count:\t%s\n", strconv.Itoa(b.currentCount))
		fmt.Printf("has capacity:\t%s\n", strconv.FormatBool(b.hasCapacity))
		fmt.Printf("block size:\t%s\n\n", strconv.Itoa(b.BLOCKSIZE))

		fmt.Println("deleted records:", b.markedDeleted)
		// fmt.Println("byte block:", b.blocks[b.numBlocks-1])
	}
	fmt.Println()
	b.printRecords(verbose)
}

// PrintRecord - parses and prints bytes slice at record address
func PrintRecord(recAddr *[]byte) {
	tconst := string((*recAddr)[:RATINGOFFSET])
	rating := string((*recAddr)[RATINGOFFSET:VOTESOFFSET])
	votes := string((*recAddr)[VOTESOFFSET:RECORDSIZE])

	fmt.Printf("tconst: %s, average rating: %s, number of votes: %s \n", tconst, rating, votes)
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

// SetBlocksAccessed - keeps track of the blocks that have been accessed during deletion and record retrieval
func (b *BlockManager) SetBlocksAccessed(addr *[]byte) {
	start := unsafe.Pointer(addr)
	size := b.BLOCKSIZE

	// use offsets to check which block the record belongs to
	// this works as the records have a fixed size
	for i := 0; i < size/RECORDSIZE; i++ {
		item := (unsafe.Pointer(uintptr(start) - uintptr(i*32)))
		if visited, found := b.blockSet[item]; found && !visited {
			b.blockSet[item] = true
			break
		}
	}
}

// GetBlocksAccessed - prints the blocks accessed during deletion and record retrieval
func (b BlockManager) GetBlocksAccessed() {
	var count int
	fmt.Println("=================\n blocks accessed\n=================")
	for block, visited := range b.blockSet {
		if visited {
			count++
			block := *(*[]byte)(block)
			for j := 0; j < b.BLOCKSIZE/RECORDSIZE; j++ {
				offset := j * RECORDSIZE

				if bytes.Compare(block[offset:offset+RECORDSIZE], make([]byte, RECORDSIZE, RECORDSIZE)) == 0 {
					continue
				}

				record := block[offset:]
				PrintRecord(&record)
			}
			fmt.Println()
		}
	}
	fmt.Printf("Num blocks accessed: %d \n", count)
}

// ResetBlocksAccessed - resets accessed blocks map
func (b *BlockManager) ResetBlocksAccessed() {
	for k := range b.blockSet {
		b.blockSet[k] = false
	}
}
