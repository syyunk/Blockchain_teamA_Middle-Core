package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
)

type Tx struct {
	Txid      []byte // O
	From      []byte
	To        []byte
	Timestamp []byte
	Amount    int64
}

func NewTransaction(from []byte, to []byte, amount int64) *Tx {
	tx := &Tx{}

	tx.Txid = tx.setHash()

	tx.From = from
	tx.To = to

	tx.Amount = amount
	tx.Timestamp = []byte(getTimestamp().String())

	return tx
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

func (t *Tx) setHash() []byte {
	sha := sha256.New()
	sha.Write(t.From)
	sha.Write(t.To)
	sha.Write(t.Timestamp)
	sha.Write([]byte(strconv.FormatInt(t.Amount, 10)))
	hash := sha.Sum(nil)

	return hash
}

func (t *Tx) printTransaction() {
	fmt.Println("----------------------------------------Transaction Info-------------------------------------------")
	fmt.Printf("Hash		: %x\n", t.Txid)
	fmt.Printf("from		: %s\n", bytes.NewBuffer(t.From).String())
	fmt.Printf("to		: %s\n", bytes.NewBuffer(t.To).String())
	fmt.Printf("Timestamp	: %x\n", t.Timestamp)
	fmt.Printf("Amount		: %d\n", t.Amount)
	fmt.Printf("----------------------------------------------------------------------------------------------------\n")
}

func (txs *transactions) addBlock(t *Tx) error {
	currentHeight := len(txs.Txs) - 1

	txs.Txs[currentHeight+1] = t

	return nil
}

func (txs *transactions) getTransaction(txid []byte) *Tx {
	currentIndex := len(txs.Txs) - 1

	for currentIndex != 0 {
		if bytes.Equal(txs.Txs[currentIndex].Txid, txid) {
			return txs.Txs[currentIndex]
		}

		currentIndex -= 1
	}

	return nil
}

func (txs *transactions) isExisted(txid []byte) bool {
	return txs.getTransaction(txid) != nil
}

func (txs *transactions) JoinAllHash() []byte {
	sha := sha256.New()

	for _, v := range txs.Txs {
		sha.Write(v.Txid)
	}

	hash := sha.Sum(nil)

	return hash
}

func (txs *transactions) printTxs() {
	for _, v := range txs.Txs {
		v.printTransaction()
	}
}
