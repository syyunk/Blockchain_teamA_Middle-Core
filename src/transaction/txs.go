package transaction

//import (
//	"bytes"
//	"crypto/sha256"
//	"github.com/lecture"
//)
//
//type Transactions struct {
//	Txs []*lecture.Tx
//}
//
//func NewTransactions(t *lecture.Tx) *Transactions {
//	if t == nil {
//		return nil
//	}
//
//	txs := &Transactions{}
//	txs.Txs = []*lecture.Tx{}
//	txs.Txs = append(txs.Txs, t)
//
//	return txs
//}
//
//func (txs *Transactions) addBlock(t *lecture.Tx) error {
//	currentHeight := len(txs.Txs) - 1
//
//	txs.Txs[currentHeight+1] = t
//
//	return nil
//}
//
//func (txs *Transactions) getTransaction(txid []byte) *lecture.Tx {
//	currentIndex := len(txs.Txs) - 1
//
//	for currentIndex != 0 {
//		if bytes.Equal(txs.Txs[currentIndex].Txid, txid) {
//			return txs.Txs[currentIndex]
//		}
//
//		currentIndex -= 1
//	}
//
//	return nil
//}
//
//func (txs *Transactions) isExisted(txid []byte) bool {
//	return txs.getTransaction(txid) != nil
//}
//
//func (txs *Transactions) JoinAllHash() []byte {
//	sha := sha256.New()
//
//	for _, v := range txs.Txs {
//		sha.Write(v.Txid)
//	}
//
//	hash := sha.Sum(nil)
//
//	return hash
//}
//
//func (txs *Transactions) printTxs() {
//	for _, v := range txs.Txs {
//		v.printTransaction()
//	}
//}
