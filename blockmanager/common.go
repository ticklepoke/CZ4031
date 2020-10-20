package blockmanager

const (
	// RECORDSIZE - number of bytes in a Record
	// longest character length for the following fields
	// tconst:		tt11285516 - 10
	// avgRating:	9.9		   - 3
	// numVotes		2279223    - 7
	RECORDSIZE = 20

	// RATINGOFFSET - byte slice index for ratings
	RATINGOFFSET = 10
	// VOTESOFFSET - byte slice indes for votes
	VOTESOFFSET = 13
)

// BlockManager - handles the logic for the creation and
// maintenance of blocks
type BlockManager struct {
	numBlocks    int
	numRecords   int
	blocks       []*[]byte
	blockSet     map[*[]byte]*[]byte
	visited      map[*[]byte]bool
	hasCapacity  bool
	currentCount int

	// contains the addreses of deleted records
	markedDeleted []*[]byte

	// BLOCKSIZE (100 or 500 bytes)
	BLOCKSIZE int
}

// Record - a byte slice that is of the size RECORDSIZE
type (
	Record []byte
)

// InitializeBlockManager - create new blockmanager instance
// with specified block size
func InitializeBlockManager(size int) BlockManager {
	b := BlockManager{numBlocks: 0, BLOCKSIZE: size, blockSet: make(map[*[]byte]*[]byte), visited: make(map[*[]byte]bool)}
	return b
}

func (b *BlockManager) createBlock() []byte {
	b.numBlocks++
	newBlock := make([]byte, b.BLOCKSIZE, b.BLOCKSIZE)
	b.blocks = append(b.blocks, &newBlock)
	b.blockSet[&newBlock] = &newBlock
	b.visited[&newBlock] = false
	b.hasCapacity = true
	b.currentCount = 0

	return newBlock
}
