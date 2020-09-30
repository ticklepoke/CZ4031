package blockmanager

import (
	"fmt"

	"github.com/ticklepoke/CZ4031/bptree"
)

const (
	// RECORDSIZE - number of bytes in a record
	RECORDSIZE = 19
)

// BlockManager - handles the logic for the creation and
// maintenance of blocks
type BlockManager struct {
	numBlocks    int
	numRecords   int
	blocks       []*[]byte
	hasCapacity  bool
	currentCount int

	// contains the addreses of deleted records
	markedDeleted []*record

	// BLOCKSIZE (100 or 500 bytes)
	BLOCKSIZE int
}

type (
	record []byte
)

// InitializeBlockManager - create new blockmanager instance
// with specified block size
func InitializeBlockManager(size int) BlockManager {
	b := BlockManager{numBlocks: 0, BLOCKSIZE: size}
	return b
}

func (b *BlockManager) createBlock() []byte {
	b.numBlocks++
	newBlock := make([]byte, b.BLOCKSIZE, b.BLOCKSIZE)
	b.blocks = append(b.blocks, &newBlock)
	b.hasCapacity = true
	b.currentCount = 0

	return newBlock
}

// DeleteRecord - find element to delete and add() address to markedDeleted
// to indicate free space in a block
func (b *BlockManager) DeleteRecord(recAddr *bptree.Record) {
	b.numRecords--
}

// InsertRecord - insert to markedDeleted records if available otherwise
// insert record to current block
func (b *BlockManager) InsertRecord(tconst string, avgRating string, numVotes string) {
	newRecord := makeRecord(tconst, avgRating, numVotes)
	del := b.markedDeleted

	if len(del) > 0 {
		var addr *record
		addr, del = del[len(del)-1], del[:len(del)-1]
		*addr = newRecord

	} else {
		b.copyToMemory(newRecord)
	}

	b.numRecords++
}

// insert new record to current block ifthere is sufficient capacity otherwise
// insert into a new block
func (b *BlockManager) copyToMemory(newRecord record) {
	if b.hasCapacity {
		offset := (*b.blocks[b.numBlocks-1])[b.currentCount*RECORDSIZE:]
		copy(offset, newRecord)

		if b.currentCount*RECORDSIZE >= b.BLOCKSIZE-RECORDSIZE {
			b.hasCapacity = false
		}
	} else {
		newBlock := b.createBlock()
		copy(newBlock, newRecord)
	}

	b.currentCount++
}

// DisplayStatus - get status of blockmanager and print the state of the
// current block
func (b BlockManager) DisplayStatus() {
	fmt.Println(b)
	parseRecords(*b.blocks[b.numBlocks-1], b.currentCount)
}

func parseRecords(block []byte, currCount int) {
	for i := 0; i < currCount; i++ {
		offset := i * 19
		tconst := string(block[offset : offset+9])
		rating := string(block[offset+9 : offset+12])
		votes := string(block[offset+12 : offset+19])

		fmt.Printf("id: %s, average rating: %s, number of votes: %s \n", tconst, rating, votes)
	}
}

// serialize fields with fixed offsets
func makeRecord(tconst string, avgRating string, numVotes string) []byte {
	val := make([]byte, 19, 19)
	copy(val, []byte(tconst))
	copy(val[9:], []byte(avgRating))
	copy(val[12:], []byte(numVotes))

	return val
}
