package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Hash      []byte // BlkID O
	PrevHash  []byte // O
	Timestamp []byte // local time
	Pow       []byte // Hash from Pow

	Version byte // blockchain(block) version 000000 ---> Wallet Address <--- key pair

	Height int64
	Nonce  int64 // Nonce from Pow
	Bits   int64 // targetBytes of Pow

	Txs *transactions
}

func GenesisBlock() *Block {
	return NewBlock([]byte{}, &transactions{}, -1)
}

func NewBlock(PrevHash []byte, txs *transactions, Height int64) *Block {
	b := &Block{
		[]byte{},
		PrevHash,
		[]byte{},
		[]byte{},
		byte(0x00),
		Height + 1,
		0,
		0,
		txs,
	}

	loc, _ := time.LoadLocation("Asia/Seoul")
	now := time.Now()
	t := now.In(loc)

	b.Hash = b.setHash()
	b.Timestamp = []byte(t.String())

	fmt.Println("--------------- Mining Start ---------------")

	pow := newProofOfWork(b)
	b.Nonce, b.Pow = pow.Run()
	b.Bits = int64(targetBites)

	fmt.Println("Produced Nonce Value :", b.Nonce)
	fmt.Println("Present bits Value :", b.Bits)
	fmt.Printf("---------------- Mining End ----------------\n\n")

	return b
}

func (b *Block) printBlock() {
	fmt.Println("-------------------------------------------Block Info----------------------------------------------")
	fmt.Printf("Hash : %x\n", b.Hash)
	fmt.Printf("PrevHash : %x\n", b.PrevHash)
	fmt.Printf("Timestamp : %x\n", b.Timestamp)
	b.Txs.printTxs()
	fmt.Printf("----------------------------------------------------------------------------------------------------\n\n")
}

func (b *Block) getHeight() int64 {
	return b.Height
}

type blocks struct {
	blockchain map[string]*Block
}

func NewBlockchain(b *Block) *blocks {
	if b == nil {
		return nil
	}

	bs := &blocks{}
	bs.blockchain = make(map[string]*Block)
	bs.blockchain[bytes.NewBuffer(b.Hash).String()] = b

	return bs
}

// import "bytes"
func (b *Block) getBlockID(txid []byte) ([]byte, bool) {
	if b.isExisted(txid) {
		return b.Hash, true
	} else {
		return []byte{}, false
	}
}

func (b *Block) isExisted(txid []byte) bool {
	if b.Txs.getTransaction(txid) != nil {
		return true
	}

	return false
}

func (b *Block) setHash() []byte {
	sha := sha256.New()
	sha.Write(b.PrevHash)
	sha.Write([]byte(strconv.FormatInt(b.Height, 10)))
	hash := sha.Sum(nil)

	return hash
}

func (bs *blocks) addBlock(o *Block) {
	o.PrevHash = bs.blockchain[bytes.NewBuffer(bs.getCurrentBlockID()).String()].Hash
	bs.blockchain[bytes.NewBuffer(o.Hash).String()] = o
}

func (bs *blocks) getBlock(blkID []byte) *Block {
	return bs.blockchain[bytes.NewBuffer(blkID).String()]
}

func (bs *blocks) getCurrentBlockID() []byte {
	var curBlockID []byte

	currentHeight := len(bs.blockchain) - 1

	for _, v := range bs.blockchain {
		if v.Height == int64(currentHeight) {
			curBlockID = v.Hash[:]
		}
	}

	return curBlockID
}

func (bs *blocks) findBlock(Height int64) *Block {
	curBlockID := bs.getCurrentBlockID()

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
