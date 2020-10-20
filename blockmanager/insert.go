package blockmanager

// InsertRecord - insert to markedDeleted records if available otherwise
// insert Record to current block
func (b *BlockManager) InsertRecord(tconst string, avgRating string, numVotes string) *[]byte {
	newRecord := makeRecord(tconst, avgRating, numVotes)
	del := (b.markedDeleted)
	var addr *[]byte

	if len(del) > 0 {
		addr, b.markedDeleted = del[len(del)-1], del[:len(del)-1]
		copy(*addr, newRecord)
	} else {
		addr = b.insertToBlock(newRecord)
	}

	b.numRecords++
	b.blockSet[addr] = b.blocks[len(b.blocks)-1]
	return addr
}

// insert Record into current block if there is sufficient capacity otherwise
// create and insert into new block
func (b *BlockManager) insertToBlock(newRecord Record) *[]byte {
	var target []byte

	if !b.hasCapacity {
		target = b.createBlock()
	} else {
		target = (*b.blocks[b.numBlocks-1])[b.currentCount*RECORDSIZE:]
		if b.currentCount*RECORDSIZE >= b.BLOCKSIZE-RECORDSIZE {
			b.hasCapacity = false
		}
	}
	copy(target, newRecord)
	b.currentCount++
	return &target
}

// serialize fields with fixed offsets
func makeRecord(tconst string, avgRating string, numVotes string) Record {
	val := make([]byte, RECORDSIZE, RECORDSIZE)
	copy(val, []byte(tconst))
	copy(val[RATINGOFFSET:], []byte(avgRating))
	copy(val[VOTESOFFSET:], []byte(numVotes))

	return val
}
