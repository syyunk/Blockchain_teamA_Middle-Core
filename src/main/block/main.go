package main

import (
	"bytes"
	"fmt"
	"strconv"
	"net"
	"net/rpc"
)

type RPCWallet Wallet	//PRC 서버에 등록하기 위해 지갑타입

func (RW *RPCWallet) MkRWallet() *Wallet {
	return makeWallet()
}

func main() {

	//지갑생성
	w1 := makeWallet()
	w2 := makeWallet()

	//생성된 지갑을 슬라이스에 추가
	ws := makeWallets()
	ws.wallets[w1.Address] = w1
	ws.wallets[w2.Address] = w2

	//제네시스 블록 및 블록체인 생성
	firstBlock := GenesisBlock()
	bs := NewBlockchain(firstBlock)

	//제네시스 블록 출력
	bs.printOrder()

	//블록생성함수
	blockLength := 5
	for i := 0; i < blockLength; i++ {
		//트렌젝션 생성
		t := NewTx(w1, w2)
		//트렌젠션(Tx)을 트렌젝션슬라이스(Txs)에 추가
		Txs := AddTxToTxs(t)

		if len(bs.blockchain) == 0 {
			fmt.Println("제네시스 블록이 없습니다")

		} else {
			//블록삽입
			prevBlock := bs.findBlock(int64(len(bs.blockchain) - 1))
			temp := NewBlock(prevBlock.Hash, Txs.Txs, prevBlock.Height, strconv.FormatInt(int64(i+1), 10)+"번째 블록")
			bs.addBlock(temp)

			//블록출력
			fmt.Println("===============Print Info===============")
			fmt.Println(bytes.NewBuffer(temp.Data).String())
			fmt.Printf("Hash : %x\n", temp.Hash)
			fmt.Printf("PrevHash : %x\n", temp.PrevHash)
			fmt.Printf("Timestamp : %x\n", temp.Timestamp)
			fmt.Println("Nonce : ", temp.Nonce)
			fmt.Printf("---transaction---\n")
			fmt.Println("From : ", bytes.NewBuffer(temp.Txs[0].From).String())
			fmt.Println("To : ", bytes.NewBuffer(temp.Txs[0].To).String())
			fmt.Printf("Amount : %x원\n", temp.Txs[0].Amount)
			fmt.Printf("Timestamp : %x\n", temp.Txs[0].Timestamp)
			fmt.Println("-----------------")
			fmt.Printf("========================================\n\n")
		}

	}


	
	//-----------------------RPC서버 생성 ----------------------//
	rpc.Register(new(RPCWallet))				// Wallet(지갑) 타입의 인스턴스를 생성하여 RPC 서버에 등록
	ln, err := net.Listen("tcp", ":6000")	// TCP 프로토콜에 6000번 포트로 연결을 받음
	if err != nil{
		fmt.Println(err)
		return
	}
	defer ln.Close()						// main 함수가 종료되기 직전에 연결 대기를 닫음

	for {
		conn,err := ln.Accept()				// 클라이언트가 연결되면 TCP 연결을 리턴
		if err != nil {
			continue
		}
		defer conn.Close()					// main 함수가 끝나기 직전에 TCP 연결을 닫음

		go rpc.ServeConn(conn)				// RPC를 처리하는 함수를 고루틴으로 실행
	}
	
}