package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"net"
	"net/rpc"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

/*================================================================================*/
// RPC 서버 관련 코드
/*================================================================================*/
// 맵 자료형으로 지갑들이 들어갈 공간을 전역 변수로 선언
// 서버가 가동되는 동안에는 데이터가 유지됨
var wallets = make(map[string]*Wallet)

// Rpc 함수를 담을 구조체 변수 선언
type WalletRpc struct{}

// 요청 구조체(매개변수)
type WalletArgs struct {
	Alias   string
	Address string
}

// 응답 구조체
type WalletReply struct {
	Prvkey  []byte
	Pubkey  []byte
	Address string
	Alias   string
}

// 지갑 생성 요청을 처리하는 함수
func (wRPC *WalletRpc) MakeNewWallet(arg WalletArgs, reply *WalletReply) error {
	prvkey, pubkey := NewKeyPair()

	w := MakeWallet(&prvkey, pubkey, arg.Alias)

	reply.Address = w.Address

	fmt.Println(len(wallets))

	return nil
}

func (wRPC *WalletRpc) GetWalletInfo(arg WalletArgs, reply *WalletReply) error {
	w := wallets[arg.Address]

	reply.Alias = w.Alias
	reply.Address = w.Address
	reply.Pubkey = w.Pubkey
	reply.Prvkey = w.Prvkey.D.Bytes()

	return nil
}

/*================================================================================*/
// 지갑 관련 코드
/*================================================================================*/
// 지갑 구조체
type Wallet struct {
	Prvkey  *ecdsa.PrivateKey
	Pubkey  []byte
	Address string
	Alias   string
}

// 지갑 생성 함수
func MakeWallet(prvkey *ecdsa.PrivateKey, pubkey []byte, alias string) *Wallet {
	w := &Wallet{}

	// sign - 서명
	publicRIPEMD160 := HashPubKey(pubkey)
	version := byte(0x00)

	Address := base58.CheckEncode(publicRIPEMD160, version)

	w.Prvkey = prvkey
	w.Pubkey = pubkey
	w.Address = Address
	w.Alias = alias

	wallets[w.Address] = w

	return w
}

// 공개키, 비공개키 추출 함수
func NewKeyPair() (ecdsa.PrivateKey, []byte) {
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

//func makeWallets() *Wallets {
//	Wallets := &Wallets{}
//	Wallets.wallets = make(map[string]*Wallet)
//
//	return Wallets
//}

//func (Wallets *Wallets) addWallet(wallet *Wallet) {
//	Wallets.wallets[wallet.Address] = wallet
//}
//
//func (Wallets *Wallets) getWallet(address string) *Wallet {
//	return Wallets.wallets[address]
//}

func (Wallet *Wallet) printInfo() {
	fmt.Printf("Alias : %s\n", Wallet.Alias)
	fmt.Printf("Address : %s\n", Wallet.Address)
	fmt.Printf("PublicKey : %x\n", Wallet.Pubkey)
	fmt.Printf("PrivateKey : %s\n", Wallet.Prvkey)
}

func main() {
	rpc.Register(new(WalletRpc))

	listener, err := net.Listen("tcp", ":6000")

	if err != nil {
		return
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			continue
		}

		go rpc.ServeConn(conn)
	}
}
