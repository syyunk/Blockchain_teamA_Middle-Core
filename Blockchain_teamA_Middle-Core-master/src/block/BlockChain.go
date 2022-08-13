package block

import (
	"bytes"
	"encoding/hex"
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

func GetBlockCh(blkID []byte) *Block { 
	// return Blockchain[hex.EncodeToString(blkID)] //16진수 이노딩을 리턴 
	return Blockchain[bytes.NewBuffer(blkID).String()]
 }

// func getBlock(blkID []byte) *Block {
// 	return Blockchain[bytes.NewBuffer(blkID).String()]
// }

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

//block 조회 - send : hash, data, timestamp, txid(확인용)
//반복문으로 서비스에서 받아온 txid가 from과 일치하는 블록 데이터들을 block 구조체에 담아 가져온다.
func GetBlockOne(txid []byte) *Block {
	fmt.Println("#################4444")
	block := &Block{}

	fmt.Println(len(Blockchain))

	for _, v := range Blockchain {
		if bytes.Equal(v.Txid, txid) {
			blkID := v.Hash

			fmt.Println("블럭 해쉬값:", hex.EncodeToString(v.Hash))
			// fmt.Println("11111111111")
			block = GetBlockCh(blkID)
			fmt.Println("22222222222")
			// fmt.Println("getBlock으로 찾은 블록의 txid 값 :: ", hex.EncodeToString(block.Txid))
			return block
		}
		// fmt.Println("block구조체 담긴 값 :: ", v)
	}
	return nil

}