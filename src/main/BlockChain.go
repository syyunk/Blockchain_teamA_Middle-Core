package main

import (
	"bytes"
)

type blocks struct {
	blockchain map[string]*Block
}

func NewBlockchain(b *Block) *blocks {
	bs := &blocks{}
	bs.blockchain = make(map[string]*Block)
	bs.blockchain[bytes.NewBuffer(b.Hash).String()] = b

	return bs
}

func (bs *blocks) addBlock(b *Block) {
	curBlockID := bs.getCurrentBlockId()

	b.PrevHash = bs.blockchain[bytes.NewBuffer(curBlockID).String()].Hash

	bs.blockchain[bytes.NewBuffer(b.Hash).String()] = b
}

func (bs *blocks) getCurrentBlockId() []byte {
	var curBlockID []byte

	currentHeight := len(bs.blockchain) - 1

	for _, v := range bs.blockchain {
		if v.Height == int64(currentHeight) {
			curBlockID = v.Hash[:]
		}
	}

	return curBlockID
}

func (bs *blocks) getBlock(blkID []byte) *Block {
	return bs.blockchain[bytes.NewBuffer(blkID).String()]
}

func (bs *blocks) findBlock(Height int64) *Block {
	curBlockID := bs.getCurrentBlockId()

	for {
		blk := bs.blockchain[bytes.NewBuffer(curBlockID).String()]

		if blk.Height == Height {
			return blk
		}

		if bytes.Equal(blk.PrevHash, []byte{}) {
			return nil
		}

		curBlockID = blk.PrevHash[:]

	}
}

func (bs *blocks) printOrder() {
	Height := len(bs.blockchain) - 1

	for {
		block := bs.findBlock(int64(Height))

		if bytes.Equal(block.PrevHash, []byte{}) || block == nil {
			break
		}

		block.printBlock()

		Height -= 1
	}
}
