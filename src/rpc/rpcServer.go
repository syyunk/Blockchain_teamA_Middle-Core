package main

import (
	"fmt"
	"net"
	"net/rpc"
	"src/wallet"
)

var wallets = make(map[string]*Wallet)

type RpcServer struct{}

type Args struct {
	Alias   string
	Address string
}

type Reply struct {
	Alias      string
	Address    string
	PublicKey  []byte
	PrivateKey []byte
	Check      bool
}

type Wallet struct {
	PrivateKey []byte
	PublicKey  []byte
	Address    string
	Alias      string
}

func (wRPC *RpcServer) MakeNewWallet(Alias string, reply *Reply) error {
	prvKey, pubKey := wallet.NewKeyPair()
	w := wallet.MakeWallet(&prvKey, pubKey, Alias)
	reply.Address = w.Address
	reply.PrivateKey = w.Prvkey.D.Bytes()
	fmt.Println(reply.PrivateKey, "reply.PrvKey 입니다")
	reply.PublicKey = w.Pubkey
	fmt.Println(reply.PublicKey, "reply.PubKey 입니다")
	reply.Alias = w.Alias
	return nil
}
func (wRPC *RpcServer) CheckAddress(Address string, reply *Reply) error {
	// 주소가 존재한다면
	if wallets[Address] != nil {
		reply.Check = true
	} else {
		reply.Check = false
	}
	return nil
}

func (wRPC *RpcServer) GetWallet(Address string, reply *Reply) error {

	w := wallets[Address]
	reply.PrivateKey = w.PrivateKey
	reply.PublicKey = w.PublicKey
	return nil
}

// -------------------- main ----------------------------------------------------

func main() {
	rpc.Register(new(RpcServer))

	In, err := net.Listen("tcp", ":9000")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer In.Close()

	for {
		conn, err := In.Accept()

		if err != nil {
			continue
		}
		defer conn.Close()

		go rpc.ServeConn(conn)
	}
}
