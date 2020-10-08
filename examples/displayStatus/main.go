package main

import (
	"strconv"

	"github.com/ticklepoke/CZ4031/blockmanager"
)

func main() {
	b := blockmanager.InitializeBlockManager(100)

	// multiple blocks
	for i := 0; i < 101; i++ {
		b.InsertRecord("tt"+strconv.Itoa(i), "7.7", strconv.Itoa(i))
	}
	b.DisplayStatus(true)
}
