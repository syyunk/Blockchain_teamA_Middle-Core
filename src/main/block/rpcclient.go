package main

import(
	"fmt"
	"net/rpc"
)

// 매개변수
//type Args struct {
//	A, B int
//}

// 리턴값
type Reply struct {
	W *Wallet
}

func main(){
	client, err := rpc.Dial("tcp", "127.0.0.1:6000")	// RPC 서버에 연결
	if err != nil{
		fmt.Println(err)
		return
	}
	defer client.Close()	// main 함수가 끝나기 직전에 RPC 연결을 닫음

	//동기 호출
	reply := new(Reply)
	client.Call("RPCWallet.MkRWallet", []int{} , reply)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(reply.W)



	MkRWalletCall := client.Go("RPCWallet.MkRWallet", []int{}, reply, nil)
	<-MkRWalletCall.Done // 함수가 끝날 때까지 대기
	fmt.Println(reply.W)
}