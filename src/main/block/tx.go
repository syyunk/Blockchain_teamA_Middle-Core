package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"time"
)

type Tx struct {
   Txid        []byte
   From        []byte   //구매자
   To          []byte
   Timestamp   []byte
   Amount      int64
}

type Txs struct {
   Txs []*Tx
}

//새로운 트렌젝션 생성
func NewTx(ToWallet *Wallet, FromWallet *Wallet) *Tx {
   t := &Tx{}

   t.From = []byte(FromWallet.Address)
   t.To = []byte(ToWallet.Address)
   
   loc, _ := time.LoadLocation("Asia/Seoul")
   now := time.Now()
   time := now.In(loc)
   t.Timestamp = []byte(time.String())
   t.Txid = t.setTxidHash()

   t.Amount = int64(rand.Intn(1000))

   return t
}

//해시값 생성
func (t *Tx) setTxidHash() []byte {
   sha := sha256.New()
   sha.Write(t.From)
   sha.Write(t.Timestamp)
   hash := sha.Sum(nil)

   return hash
}

//트렌젝션을 트렌젝션 슬라이스에 추가
func AddTxToTxs(t *Tx) *Txs {
   if t == nil {
      return nil
   }

   Txs := &Txs{}
   Txs.Txs = []*Tx{}
   Txs.Txs = append(Txs.Txs, t)

   return Txs
}

func (t *Tx) printTx() {
   fmt.Println("===============Print Info===============")
   fmt.Println("From : ", bytes.NewBuffer(t.From).String())
   fmt.Println("To : ",bytes.NewBuffer(t.To).String())
   fmt.Printf("Amount : %x원\n", t.Amount)
   fmt.Printf("Timestamp : %x\n", t.Timestamp)
   fmt.Println("========================================")
}




// import "bytes"
func (t *Tx) getBlockTxID(txid []byte) ([]byte, bool) {
   if t.isTxExisted(txid) {
      return t.Txid, true
   } else {
      return []byte{}, false
   }
}

//------------------------
// 1 Tx -> 1 Block
// n Txs -> 1 Block
func (t *Tx) isTxExisted(txid []byte) bool {
   // ToDo
   // n Txs -> 1 Block
   return bytes.Equal(t.Txid, txid)
}

//-----------중현 코딩--------------
func (Txs *Txs) GetTxs(txid []byte) *Tx {
   for i := 0; i < len(Txs); i++ {
      if txid == Txs.Tx.Txid{
         return Txs.Tx
      }
   }
}

// 판매자의 구매내역의 모든 TX를 뽑아오는 함수
func getTxsHistory(wallet []byte) *Txs {
   txs := *Txs{}     // 실제 담겨있는 데이터 
   retxs := &Txs{}   // 응답하기위한 데이터
   for i := 0; i < len(txs); i++ {
      if txs.Tx.From == wallet { //판매자 지갑주소와 같은 지갑주소 모두 담는다.
         retxs = txs.Tx
      }
   }
   return retxs
}