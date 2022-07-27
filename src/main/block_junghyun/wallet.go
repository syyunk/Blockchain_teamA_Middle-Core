package main
import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "github.com/btcsuite/btcutil/base58"
    "crypto/sha256"
    "golang.org/x/crypto/ripemd160"
)

//지갑 구조체
type Wallet struct {
    prvKey  ecdsa.PrivateKey
    pubKey  []byte
    address  string
    alias   string
}

//지갑들 구조체
type Wallets struct {
    Wallets []string
}

//공개키를 통해서 지갑주소 생성
func NewWallet() *Wallet {
    wl := &Wallet{}
    
    prvKey, pubKey, address, alias := getParam()

    wl.address = address
    wl.pubKey = pubKey
    wl.prvKey = prvKey
    wl.alias = alias
    
    return wl
}


//공개키와 개인키를 만드는 함수
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
func getParam() (ecdsa.PrivateKey, []byte, string, string){
    //개인키, 공개키
    prvKey, pubKey := newKeyPair()
    // address 를 구하기위한 2줄 코드
    publicRIPEMD160 := HashPubKey(pubKey)
    version := byte(0x00)

    //주소값
    address := base58.CheckEncode(publicRIPEMD160, version)
    //구매자ID?? 
    alias := "test"

    return prvKey, pubKey, address, alias
}
