package main

import (
	"bytes"
	"strconv"
)

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

	tx := NewTransaction([]byte(w1.Address), []byte(w2.Address), 10000)
	txs := NewTransactions(tx)

	for i := 1; i < 5; i++ {
		prevBlock := bs.findBlock(int64(len(bs.blockchain) - 1))
		temp := NewBlock(prevBlock.Hash, txs, prevBlock.Height, strconv.FormatInt(int64(i), 10)+"번째 블록")
		bs.blockchain[bytes.NewBuffer(temp.Hash).String()] = temp
	}

	bs.printOrder()
}
