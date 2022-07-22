package main

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Hash     []byte // BlkID O
	PrevHash []byte // O
	Pow      []byte // Hash from Pow
	Txid     []byte // O
	//	// merkleRoot [32]byte
	version int64 // blockchain(block) version 000000 ---> Wallet Address <--- key pair
	Nonce   int64 // Nonce from Pow
	bits    int64 // targetBytes of Pow
	Height  int64

	Timestamp []byte // local time

	// Data []byte // copyrights

	//Timestamp int64 // local time
	Txs *transactions
	//  MR        [32]byte
	//	MT        []*txID
	//  Signature []byte
}

// ----------비트코인 / GenesisBlock data -----
// "The Times 03/Jan/2009 Chanceller
// on brink of second bailout for banks"

func GenesisBlock() *Block {
	return NewBlock([]byte{}, &transactions{}, -1)
}

func NewBlock(PrevHash []byte, txs *transactions, Height int64) *Block {
	b := &Block{}

	b.PrevHash = PrevHash[:]
	b.Height = Height + 1

	loc, _ := time.LoadLocation("Asia/Seoul")
	now := time.Now()
	t := now.In(loc)
	b.Timestamp = []byte(t.String())
	b.Hash = b.setHash()
	b.Txs = txs

	fmt.Println("--------------- Mining Start ---------------")

	pow := newProofOfWork(b)
	b.Nonce, b.Pow = pow.Run()
	b.bits = int64(targetBites)

	fmt.Println("Produced Nonce Value :", b.Nonce)
	fmt.Println("Present bits Value :", b.bits)
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

//func (b *Block) getBlockID() []byte {
//	return b.Hash
//}

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

//------------------------
// 1 Tx -> 1 Block
// n Txs -> 1 Block
func (b *Block) isExisted(txid []byte) bool {
	// ToDo
	// n Txs -> 1 Block
	return bytes.Equal(b.Txid, txid)
}

func (b *Block) setHash() []byte {
	sha := sha256.New()
	sha.Write(b.PrevHash)
	sha.Write([]byte(strconv.FormatInt(b.Height, 10)))
	hash := sha.Sum(nil)

	return hash
}

func (bs *blocks) addBlock(o *Block) error {
	if bytes.Equal(o.Hash, []byte{}) {
		return errors.New("유효하지 않은 블럭입니다")
	}

	// 최신 블록체인의 높이를 구한다.
	currentHeight := len(bs.blockchain) - 1

	//	최신 블록을 찾는다.
	var curBlockID []byte

	for _, v := range bs.blockchain {
		if v.Height == int64(currentHeight) {
			curBlockID = v.Hash[:]
		}
	}

	o.PrevHash = bs.blockchain[bytes.NewBuffer(curBlockID).String()].Hash

	bs.blockchain[bytes.NewBuffer(o.Hash).String()] = o

	return nil
}

func (bs *blocks) getBlock(blkID []byte) *Block {
	return bs.blockchain[bytes.NewBuffer(blkID).String()]
}

func (bs *blocks) findBlock(Height int64) *Block {
	// 최신 블록체인의 높이를 구한다.
	currentHeight := len(bs.blockchain) - 1

	//	최신 블록을 찾는다.
	var curBlockID []byte

	for _, v := range bs.blockchain {
		if v.Height == int64(currentHeight) {
			curBlockID = v.Hash[:]
		}
	}

	for {
		// curBlockID에 해당하는 블록을 받아온다.
		blk := bs.blockchain[bytes.NewBuffer(curBlockID).String()]

		// 블록의 높이를 계산한다.
		if blk.Height == Height {
			// 같으면 반환한다.
			return blk
		} else {
			// 높이가 다르면 다음 블록이 마지막(제네시스)블록인지 확인한다.
			if bytes.Equal(blk.PrevHash, []byte{}) {
				return nil
			}

			// 다시 순회하도록 (PrevHash) 셋팅한다.
			curBlockID = blk.PrevHash[:]
		}
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

func main() {
	firstBlock := GenesisBlock()

	ws := makeWallets()
	bs := NewBlockchain(firstBlock)

	prvKey, pubKey := newKeyPair()
	prvKey2, pubKey2 := newKeyPair()

	w1 := makeWallet(prvKey, pubKey, "JamesBond")
	w2 := makeWallet(prvKey2, pubKey2, "Company")

	ws.wallets[w1.Address] = w1
	ws.wallets[w2.Address] = w2

	tx := NewTx([]byte(w1.Address), []byte(w2.Address), 10000)
	txs := NewTransactions(tx)

	for i := 1; i < 100; i++ {
		prevBlock := bs.findBlock(int64(len(bs.blockchain) - 1))
		temp := NewBlock(prevBlock.Hash, txs, prevBlock.Height)
		bs.blockchain[bytes.NewBuffer(temp.Hash).String()] = temp
	}

	bs.printOrder()

	w1.printInfo()
}
