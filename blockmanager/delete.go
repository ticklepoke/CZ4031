package blockmanager

// DeleteRecord - find record to delete and add address to markedDeleted
// to indicate free space in a block
func (b *BlockManager) DeleteRecord(recAddr *[]byte) {
	b.numRecords--
	b.markedDeleted = append(b.markedDeleted, recAddr)
	copy(*recAddr, make([]byte, RECORDSIZE, RECORDSIZE))
}
