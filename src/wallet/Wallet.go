package wallet

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

func makeWallet(prvkey ecdsa.PrivateKey, pubkey []byte, alias string) *Wallet {
	w := &Wallet{}

	publicRIPEMD160 := HashPubKey(pubkey)
	version := byte(0x00)

	Address := base58.CheckEncode(publicRIPEMD160, version)

	w.Prvkey = prvkey
	w.Pubkey = pubkey
	w.Address = Address
	w.Alias = alias

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

func (Wallets *Wallets) addWallet(wallet *Wallet) {
	Wallets.wallets[wallet.Address] = wallet
}

func (Wallets *Wallets) getWallet(address string) *Wallet {
	return Wallets.wallets[address]
}

func (Wallet *Wallet) printInfo() {
	fmt.Printf("Alias : %s\n", Wallet.Alias)
	fmt.Printf("Address : %s\n", Wallet.Address)
	fmt.Printf("PublicKey : %x\n", Wallet.Pubkey)
	fmt.Printf("PrivateKey : %s\n", Wallet.Prvkey)
}
