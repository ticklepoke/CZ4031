package main

import (
	"github.com/ticklepoke/CZ4031/blockmanager"
)

func main() {
	b := blockmanager.InitializeBlockManager(100)

	b.InsertRecord("tt0848228", "7.7", "9987")
	recAddr := b.InsertRecord("tt0848224", "7.2", "99890")
	b.DisplayStatus(false)

	// delete record by address
	b.DeleteRecord(recAddr)
	b.DisplayStatus(false)

	// show that marked address is overwritten by new blocks
	b.InsertRecord("tt081", "9", "98")
	b.DisplayStatus(true)
}
