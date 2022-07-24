package main

import (
	"bytes"
	"crypto/sha256"
)

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
