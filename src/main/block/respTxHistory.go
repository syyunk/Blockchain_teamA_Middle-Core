package main

import(
	"fmt"
	"encoding/json"
	"bytes"
	"net/http"
	
)
// 판매자의 1개의 구매내역 정보
type SellHistory struct {
	BlkID     	[]byte 	// 블록ID 
	TxID     	[]byte 	// 트랜잭션즈
	BWallet	    []byte 	// 구매자지갑
	SWallet  	[]byte	// 판매자지갑
	Price 		int64	// 가격
	Timestamp 	[]byte 	// 거래일
}

type SsH struct {
	SsH []*sellHistory
 }


 // 구매 내역 조회
func getSellHistory(){
	txs := &Txs{}
	bs := &blocks{}
	SsH := &SsH{}

	// 거래내역 구조체 생성
	sellHistory := SellHistory{}
	
	//블록아이디로 txid값 가져오기
	txid := bs.GetTxID([]byte{}/*블록아이디*/);

	for i := 0; i < len(*txs); i++ {
		tx := txs.GetTxs([]byte{})

		sellHistory.TxID 		= txid
		sellHistory.BWallet 	= tx.From
		sellHistory.SWallet 	= tx.To
		sellHistory.Price 		= tx.Amount
		sellHistory.Timestamp 	= tx.Timestamp
		sellHistory.BlkID 		= []byte{}//사용자가보냄

		SsH.SsH= append(SsH.SsH, sellHistory)
	 }
		

	// 데이터를 받은 구조체를 json 형식으로 변환
	TxH, _ := json.Marshal(SsH)

	buff := bytes.NewBuffer(TxH)
	fmt.Println(TxH)
	fmt.Println(string(TxH))
	fmt.Println(buff)
	
	mux := http.NewServeMux()
	mux.HandleFunc("/getsellHistory", func(res http.ResponseWriter, req *http.Request){
		res.Write([]byte(string(TxH)))
	})

	//req, _ := http.NewRequest("POST", "/movie", buff)
	//response := executeRequest(req)
	http.ListenAndServe(":80", mux)
}
