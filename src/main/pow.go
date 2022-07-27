package main

import (
	"bytes"
	"crypto/sha256"
	"math"
	"math/big"
)
​
var (
	maxNonce = math.MaxInt64
)
​
const targetBites = 16
​
type ProofOfWork struct {
	block  *Block
	target *big.Int
}
​
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join([][]byte{
		pow.block.PrevBlockHash,
		pow.block.Data,
		IntToHex(pow.block.Timestamp),
		IntToHex(int64(targetBites)),
		IntToHex(int64(nonce)),
	}, []byte{})
	return data
}
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
​
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	return nonce, hash[:]
}
​
func newProofOfWork(block *Block) *ProofOfWork {
​
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBites))
	pow := &ProofOfWork{block, target}
​
	return pow
}