package restAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/rpc"
)

type WalletArgs struct {
	Alias   string
	Address string
}

type WalletReply struct {
	Alias   string
	Address string
	PrvKey  []byte
	PubKey  []byte
}

func GetRpcConnection() *rpc.Client {
	client, err := rpc.Dial("tcp", "127.0.0.1:6000")

	if err != nil {
		panic(err)
	}

	return client
}

func sendToRpcMakeWallet(client *rpc.Client, args *WalletArgs) *WalletReply {
	reply := new(WalletReply)

	err := client.Call("WalletRpc.MakeNewWallet", args, reply)

	if err != nil {
		panic(err)
	}

	return reply
}

func sendToRpcGetInfo(client *rpc.Client, args *WalletArgs) *WalletReply {
	reply := new(WalletReply)

	err := client.Call("WalletRpc.GetWalletInfo", args, reply)

	if err != nil {
		panic(err)
	}

	return reply
}

func MakeWallet(rw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	b, _ := ioutil.ReadAll(req.Body)

	args := &WalletArgs{}

	err := json.Unmarshal(b, &args)

	if err != nil {
		panic(err)
	}

	client := GetRpcConnection()

	reply := sendToRpcMakeWallet(client, args)

	response, _ := json.Marshal(reply)

	rw.Write(response)
}

func GetWalletInfo(rw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	b, _ := ioutil.ReadAll(req.Body)

	args := &WalletArgs{}

	err := json.Unmarshal(b, &args)

	if err != nil {
		panic(err)
	}

	fmt.Println(args.Address)

	client := GetRpcConnection()

	reply := sendToRpcGetInfo(client, args)

	response, _ := json.Marshal(reply)

	rw.Write(response)
}
