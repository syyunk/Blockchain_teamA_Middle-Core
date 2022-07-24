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
