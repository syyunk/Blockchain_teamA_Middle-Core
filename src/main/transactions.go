package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Tx struct {
	Txid      []byte // O
	From      []byte
	To        []byte
	Timestamp []byte
	Amount    int64
}

func NewTx(from []byte, to []byte, amount int64) *Tx {
	tx := &Tx{[]byte{}, from, to, []byte{}, amount}

	loc, _ := time.LoadLocation("Asia/Seoul")
	now := time.Now()
	t := now.In(loc)

	tx.Txid = tx.setHash()
	tx.Timestamp = []byte(t.String())

	return tx
}

func (t *Tx) printBlock() {
	fmt.Println("----------------------------------------Transaction Info-------------------------------------------")
	fmt.Printf("Hash : %x\n", t.Txid)
	fmt.Printf("from :%s\n", bytes.NewBuffer(t.From).String())
	fmt.Printf("to :%s\n", bytes.NewBuffer(t.To).String())
	fmt.Printf("Timestamp : %x\n", t.Timestamp)
	fmt.Printf("Amount : %d\n", t.Amount)
	fmt.Printf("----------------------------------------------------------------------------------------------------\n\n")
}

type transactions struct {
	Txs []*Tx
}

func NewTransactions(t *Tx) *transactions {
	if t == nil {
		return nil
	}

	txs := &transactions{}
	txs.Txs = []*Tx{}
	txs.Txs = append(txs.Txs, t)

	return txs
}

func (txs *transactions) isExisted(txid []byte) bool {
	if txs.getTransaction(txid) != nil {
		return true
	}
	return false
}

func (t *Tx) setHash() []byte {
	sha := sha256.New()

	sha.Write(t.From)
	sha.Write(t.To)
	sha.Write(t.Timestamp)
	sha.Write([]byte(strconv.FormatInt(t.Amount, 10)))

	hash := sha.Sum(nil)

	return hash
}

func (txs *transactions) addTx(t *Tx) error {
	// 최신 블록체인의 높이를 구한다.
	currentHeight := len(txs.Txs) - 1

	txs.Txs[currentHeight+1] = t

	return nil
}

func (txs *transactions) getTransaction(txid []byte) *Tx {
	// 최신 블록체인의 높이를 구한다.
	currentIndex := len(txs.Txs) - 1

	for {
		if bytes.Equal(txs.Txs[currentIndex].Txid, txid) {
			return txs.Txs[currentIndex]
		} else if currentIndex == 0 {
			return nil
		} else {
			currentIndex -= 1
		}
	}
}

func (txs *transactions) printTxs() {
	for _, v := range txs.Txs {
		v.printBlock()
	}
}
