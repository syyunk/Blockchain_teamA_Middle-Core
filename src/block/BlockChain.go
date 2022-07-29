package block

import (
	"bytes"
	"fmt"
)

var Blockchain = make(map[string]*Block)

func NewBlockchain(b *Block) {
	Blockchain[bytes.NewBuffer(b.Hash).String()] = b
}

func AddBlock(b *Block) {
	curBlockID := GetCurrentBlockId()

	b.PrevHash = Blockchain[bytes.NewBuffer(curBlockID).String()].Hash

	fmt.Printf("%x\n", b.Hash)
	fmt.Printf("%x", b.PrevHash)

	Blockchain[bytes.NewBuffer(b.Hash).String()] = b
}

func GetCurrentBlockId() []byte {
	var curBlockID []byte

	currentHeight := len(Blockchain) - 1

	for _, v := range Blockchain {
		if v.Height == int64(currentHeight) {
			curBlockID = v.Hash[:]
		}
	}

	return curBlockID
}

func getBlock(blkID []byte) *Block {
	return Blockchain[bytes.NewBuffer(blkID).String()]
}

func findBlock(Height int64) *Block {
	curBlockID := GetCurrentBlockId()

	for {
		blk := Blockchain[bytes.NewBuffer(curBlockID).String()]

		if blk.Height == Height {
			return blk
		}

		if bytes.Equal(blk.PrevHash, []byte{}) {
			return nil
		}

		curBlockID = blk.PrevHash[:]

	}
}

func printOrder() {
	Height := len(Blockchain) - 1

	for {
		block := findBlock(int64(Height))

		if bytes.Equal(block.PrevHash, []byte{}) || block == nil {
			break
		}

		block.printBlock()

		Height -= 1
	}
}
