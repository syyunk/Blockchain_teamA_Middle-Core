package main

import(
	"fmt"
	"net"
	"net/rpc"
)

type Calc int	//PRC 서버에 등록하기 위해 임의의 타입으로 정의

// 매개변수
type Args struct {
	A, B int
}

// 리턴값
type Reply struct {
	C int
}
// 덧셈
func (C *Calc) Sum(args Args, reply *Reply) error {
	reply.C = args.A + args.B		// 두 값을 더하여 리턴값 구조체에 넣어줌
	return nil
}
// 뺄셈
func (C *Calc) Minus(args Args, reply *Reply) error {
	reply.C = args.A - args.B
	return nil
}
// 곱셈
func (C *Calc) Mul(args Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}
// 나눗셈
func (C *Calc) Div(args Args, reply *Reply) error {
	reply.C = args.A / args.B
	return nil
}

func main(){
	rpc.Register(new(Calc))	// Calc 타입의 인스턴스를 생성하여 RPC 서버에 등록
	ln, err := net.Listen("tcp", ":6000")	// TCP 프로토콜에 6000번 포트로 연결을 받음
	if err != nil{
		fmt.Println(err)
		return
	}
	defer ln.Close()			// main 함수가 종료되기 직전에 연결 대기를 닫음

	for {
		conn,err := ln.Accept()	// 클라이언트가 연결되면 TCP 연결을 리턴
		if err != nil {
			continue
		}
		defer conn.Close()		// main 함수가 끝나기 직전에 TCP 연결을 닫음

		go rpc.ServeConn(conn)	// RPC를 처리하는 함수를 고루틴으로 실행
	}
}