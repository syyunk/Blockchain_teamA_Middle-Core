package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"

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

func makeWallet(Prvkey ecdsa.PrivateKey, Pubkey []byte, Alias string) *Wallet {
	w := &Wallet{}

	publicRIPEMD160 := HashPubKey(Pubkey)
	version := byte(0x00)

	Address := base58.CheckEncode(publicRIPEMD160, version)

	fmt.Println(Address)

	w.Prvkey = Prvkey
	w.Pubkey = Pubkey
	w.Address = Address
	w.Alias = Alias

	return w
}

func makeWallets() *Wallets {
	Wallets := &Wallets{}
	Wallets.wallets = make(map[string]*Wallet)

	return Wallets
}

func newKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	prvKey, _ := ecdsa.GenerateKey(curve, rand.Reader)
	pubKey := prvKey.PublicKey
	bpubKey := append(pubKey.X.Bytes(), pubKey.Y.Bytes()...)

	return *prvKey, bpubKey
}

func HashPubKey(pubKey []byte) []byte {
	publicSHA256 := sha256.Sum256(pubKey)
	RIPEMD160Hasher := ripemd160.New()
	RIPEMD160Hasher.Write(publicSHA256[:])
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)
	return publicRIPEMD160
}

func (Wallets *Wallets) getWallet(address string) *Wallet {
	return Wallets.wallets[address]
}

func (Wallets *Wallets) addWallet(wallet *Wallet) {
	Wallets.wallets[wallet.Address] = wallet
}

func (Wallet *Wallet) printInfo() {
	fmt.Println("Alias :", Wallet.Alias)
	fmt.Println("Address :", Wallet.Address)
	fmt.Printf("PublicKey : %x\n", Wallet.Pubkey)
	fmt.Println("PrivateKey :", Wallet.Prvkey)
}

// func main() {

// encoded := base58.Encode(pubKey)

// decoded := base58.Decode(encoded)

// 	// if bytes.Equal(pubKey, decoded) {
// 	// 	fmt.Println("Same\n")
// 	// } else {
// 	// 	fmt.Println("Not same\n")
// 	// }
// }
