package transaction

import (
	"bytes"
)

var Txs = []*Tx{}

func AddTx(t *Tx) {
	Txs = append(Txs, t)
}

func GetTransaction(txid []byte) *Tx {
	currentIndex := len(Txs) - 1

	for currentIndex > -1 {
		if bytes.Equal(Txs[currentIndex].Txid, txid) {
			return Txs[currentIndex]
		}

		currentIndex -= 1
	}

	return nil
}

func isExisted(txid []byte) bool {
	return GetTransaction(txid) != nil
}

func printTxs() {
	for _, v := range Txs {
		v.printTransaction()
	}
}
