package main

import (
   "bytes"
   "crypto/sha256"
   "fmt"
   "strconv"
   "time"
)

type Tx struct {
   Txid      []byte
   Wallet    *Wallet
   CWallet   *Wallet
   //Wallet    []byte
   //CWallet   []byte
   Timestamp []byte
   Amount    int64
}

func NewTx(amount int64) *Tx {
   tx := &Tx{}

   loc, _ := time.LoadLocation("Asia/Seoul")
   now := time.Now()
   time := now.In(loc)
   
   tx.Wallet = NewWallet()
   tx.CWallet = NewWallet()
   tx.Amount = amount
   tx.Timestamp = []byte(time.String())
   tx.Txid = tx.setHash()

   return tx
}

func (tx *Tx) printBlock() {
   fmt.Println("-------------------------------Transaction Info---------------------------------")
   fmt.Printf("Hash : %x\n", tx.Txid)
   fmt.Printf("Timestamp : %x\n", tx.Timestamp)
   fmt.Printf("Amount : %d\n", tx.Amount)
   fmt.Printf("-------------------------------------------------------------------------------\n\n")
}

type Txs struct {
   Txs []*Tx
}

func NewTxs(t *Tx) *Txs {
   if t == nil {
      return nil
   }

   txs := &Txs{}
   txs.Txs = []*Tx{}
   txs.Txs = append(txs.Txs, t)

   return txs
}

//------------------------
// 1 Tx -> 1 Block
// n Txs -> 1 Block
func (t *Tx) isExisted(txid []byte) bool {
   // ToDo
   // n Txs -> 1 Block
   return bytes.Equal(t.Txid, txid)
}

func (tx *Tx) setHash() []byte {
   sha := sha256.New()
   sha.Write(tx.Timestamp)
   sha.Write([]byte(strconv.FormatInt(tx.Amount, 10)))
   hash := sha.Sum(nil)

   return hash
}

func (txs *Txs) addBlock(t *Tx) error {
   // 최신 블록체인의 높이를 구한다.
   currentHeight := len(txs.Txs) - 1

   txs.Txs[currentHeight+1] = t

   return nil
}

func (txs *Txs) getTransaction(txid []byte) *Tx {
   // 최신 블록체인의 높이를 구한다.
   currentHeight := len(txs.Txs) - 1

   for {
      if bytes.Equal(txs.Txs[currentHeight].Txid, txid) {
         return txs.Txs[currentHeight]
      }

      if currentHeight == 0 {
         return nil
      } else {
         currentHeight -= 1
      }
   }
}

func (txs *Txs) printTxs() {
   for _, v := range txs.Txs {
      v.printBlock()
   }
}