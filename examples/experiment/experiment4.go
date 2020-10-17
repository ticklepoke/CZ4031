package experiment

import "github.com/ticklepoke/CZ4031/bptree"

func experiment4(t *bptree.Tree) {
	t.FindAndPrintRange(7, 9, true)
}
