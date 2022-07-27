package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"

	"github.com/btcsuite/btcutil/base58"

	"golang.org/x/crypto/ripemd160"
)

type Wallet struct {
    Prvkey  ecdsa.PrivateKey
    Pubkey  []byte
    Address string
    Alias   string
}

type Wallets struct {
    wallets map[string]*Wallet
}

//지갑생성
func makeWallet() *Wallet {

    w := &Wallet{}
    w.Prvkey , w.Pubkey = newKeyPair()

	publicRIPEMD160 := HashPubKey(w.Pubkey)
    version := byte(0x00)

    w.Address = base58.CheckEncode(publicRIPEMD160, version)
    w.Alias = "지갑"

    return w
}

//지갑맵 생성
func makeWallets() *Wallets {
    Wallets := &Wallets{}
    Wallets.wallets = make(map[string]*Wallet)
    return Wallets
}

//지갑맵 get
func getWallets() *Wallets {
	ws := makeWallets()
	return ws
}

//Key값 암호화 함수
func newKeyPair() (ecdsa.PrivateKey, []byte) {
    curve := elliptic.P256()
    prvKey, _ := ecdsa.GenerateKey(curve, rand.Reader)
    pubKey := prvKey.PublicKey
    bpubKey := append(pubKey.X.Bytes(), pubKey.Y.Bytes()...)
    return *prvKey, bpubKey
}

//publicRIPEMD160
//임의의 길이의 입력 값을 160비트로 압축하는 암호화 해시함수
func HashPubKey(pubKey []byte) []byte {
    publicSHA256 := sha256.Sum256(pubKey)
    RIPEMD160Hasher := ripemd160.New()
    RIPEMD160Hasher.Write(publicSHA256[:])
    publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)
    return publicRIPEMD160
}