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
   // merkleRoot [32]byte
   version int64 // blockchain(block) version  ----> Wallet Address <--- key pair
   Nonce   int64 // Nonce from Pow
   bits    int64 // targetBytes of Pow
   Height  int64

   Timestamp []byte // local time

   Data []byte // copyrights

   //Timestamp int64 // local time
   Txs []*Tx
   //  MR        [32]byte
   //   MT        []*txID
   //  Signature []byte
}

type blocks struct {
   blockchain map[string]*Block
}


//제네시스 블록 생성
func GenesisBlock() *Block {
   b := &Block{}

   b.Height = 0
   // b.Nonce = 1234 // Pow, Hash, TargetBit : bits = 4/8/12/16/20

   loc, _ := time.LoadLocation("Asia/Seoul")
   now := time.Now()
   t := now.In(loc)
   b.Timestamp = []byte(t.String())
   b.Hash = b.setHash()
   b.Data = []byte("The Times 03/Jan/2009 Chanceller on brink of second bailout for banks")

   pow := newProofOfWork(b)
   b.Nonce, b.Pow = pow.Run()
   b.bits = int64(targetBites)
   return b
}

//새로운 블록생성
func NewBlock(PrevHash []byte ,Txs []*Tx,Height int64, s string) *Block {
   b := &Block{}

   b.PrevHash = PrevHash[:]
   b.Height = Height + 1

   b.Txs = Txs
   // fmt.Println("!!!!!!!!1",b.Txs[0].Amount)
   loc, _ := time.LoadLocation("Asia/Seoul")
   now := time.Now()
   t := now.In(loc)
   b.Timestamp = []byte(t.String())
   b.Hash = b.setHash()
   b.Data = []byte(s)

   pow := newProofOfWork(b)
   b.Nonce, b.Pow = pow.Run() //nonce & Hash
   b.bits = int64(targetBites)

   return b
}

//해시생성함수
func (b *Block) setHash() []byte {
   sha := sha256.New()
   sha.Write(b.Data)
   sha.Write(b.PrevHash)
   sha.Write([]byte(strconv.FormatInt(b.Height, 10)))
   hash := sha.Sum(nil)

   return hash
}


func (b *Block) getHeight() int64 {
   return b.Height
}

//제니시스 블록체인 생성
func NewBlockchain(b *Block) *blocks {
   bs := &blocks{}
   bs.blockchain = make(map[string]*Block)
   bs.blockchain[bytes.NewBuffer(b.Hash).String()] = b

   return bs
}


func (b *Block) getBlockID(txid []byte) ([]byte, bool) {
   if b.isExisted(txid) {
      return b.Hash, true
   } else {
      return []byte{}, false
   }
}

// ------------------------
// 1 Tx -> 1 Block
// n Txs -> 1 Block
func (b *Block) isExisted(txid []byte) bool {
   // ToDo
   // n Txs -> 1 Block
   return bytes.Equal([]byte(b.Txid), txid)
}


func (bs *blocks) addBlock(o *Block) error {
   if bytes.Equal(o.Hash, []byte{}) {
      return errors.New("유효하지 않은 블럭입니다")
   }

   // 최신 블록체인의 높이를 구한다.
   currentHeight := len(bs.blockchain) - 1

   // 최신 블록을 찾는다.
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

   //   최신 블록을 찾는다.
   var curBlockID []byte //Hash

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
   for i := 0; i < len(bs.blockchain); i++ {
      bs.findBlock(int64(i)).printBlock()
   }
}

func (b *Block) printBlock() {
   fmt.Println("===============Print Info===============")
   fmt.Printf("Hash : %x\n", b.Hash)
   fmt.Printf("PrevHash : %x\n", b.PrevHash)
   fmt.Printf("Timestamp : %x\n", b.Timestamp)
   fmt.Println(bytes.NewBuffer(b.Data).String())
   fmt.Println("Nonce : ", b.Nonce)
   fmt.Printf("========================================\n\n")
}



//-------------------중현----------------------
// blcokID값으로 블록을 뽑아오기위한 함수

func (bs *blocks) findRespBlock(blkID []byte) *Block {
   blk := &Block{}
   // curBlockID에 해당하는 블록을 받아온다.
   blk = bs.blockchain[bytes.NewBuffer(blkID).String()]

   return blk
}


func (bs *blocks) getBlockID(TxID []byte) (blkID []byte){
   tmpbls := bs.blockchain[bytes.NewBuffer(blkID).String()]
   for _, val := range tmpbls {
      if TxID == val.Txid {
         blkID := val.Hash
      }
   }
   return blkID
}

//TxID 가져오기
func (bs *blocks) GetTxID(blkID []byte) (TxID []byte) {
   blk := &Block{}
   // curBlockID에 해당하는 블록을 받아온다.
   blk = bs.blockchain[bytes.NewBuffer(blkID).String()]
   return blk.Txid
}