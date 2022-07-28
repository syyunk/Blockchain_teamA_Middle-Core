package block

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"src/util"
	"strconv"
)

type Block struct {
	Hash      []byte // BlkID O
	PrevHash  []byte // O
	Pow       []byte // Hash from Pow
	Data      []byte // copyrights
	Timestamp []byte // local time
	Txid      []byte // transaction id

	version byte // blockchain(block) version 000000 ---> Wallet Address <--- key pair

	Height int64
	Nonce  int64 // Nonce from Pow
	bits   int64 // targetBytes of Pow
}

func GenesisBlock() *Block {
	return NewBlock(
		[]byte{},
		[]byte{},
		-1,
		[]byte("The Times 03/Jan/2009 Chanceller on brink of second bailout for banks"),
	)
}

func NewBlock(prevHash []byte, txid []byte, height int64, data []byte) *Block {
	b := &Block{
		[]byte{},
		prevHash,
		[]byte{},
		data,
		[]byte{},
		txid,
		byte(0x00),
		height + 1,
		0,
		0,
	}

	b.Hash = b.setHash()

	fmt.Println(b.Hash)

	b.Timestamp = []byte(util.GetTimestamp().String())

	b.setPowInfo()

	return b
}

func (b *Block) setHash() []byte {
	sha := sha256.New()

	sha.Write(b.PrevHash)
	sha.Write([]byte(strconv.FormatInt(b.Height, 10)))
	hash := sha.Sum(nil)

	return hash
}

func (b *Block) setPowInfo() {
	pow := newProofOfWork(b)
	b.Nonce, b.Pow = pow.Run()
	b.bits = int64(targetBites)

	fmt.Printf("-------------------------------Mining Info-------------------------------\n")
	fmt.Printf("pow : %x\n", b.Pow)
	fmt.Printf("Present bits : %d\n", b.bits)
	fmt.Printf("Produced Nonce : %d\n", b.Nonce)
	fmt.Printf("-------------------------------------------------------------------------\n\n")

}

func (b *Block) getHeight() int64 {
	return b.Height
}

func (b *Block) getBlockID(txid []byte) ([]byte, bool) {
	if bytes.Equal(b.Txid, txid) {
		return b.Hash, true
	}

	return []byte{}, false
}

func (b *Block) isExisted(txid []byte) bool {
	return bytes.Equal(b.Txid, txid)
}

func (b *Block) printBlock() {
	fmt.Println("==============================================Block Info==============================================")

	fmt.Printf("Hash		: %x\n", b.Hash)
	fmt.Printf("PrevHash	: %x\n", b.PrevHash)
	fmt.Printf("Timestamp	: %x\n", b.Timestamp)
	fmt.Printf("Saved Data	: %s\n", bytes.NewBuffer(b.Data).String())

	fmt.Printf("====================================================================================================\n\n")
}
