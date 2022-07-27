package main


import(
	"encoding/json"
	"net/http"
)
//사용자 구조체
type person struct {
	
}

// 사용자에게 반환할 블록
type JsonBlockResponse struct {
	Hash     	[]byte 	// BlkID O
	Txid     	[]byte 	// O
	bits    	int64 	// targetBytes of Pow
	Height  	int64
	Timestamp 	[]byte 	// local time
	Data 		[]byte 	// copyrights
	Txs 		[]*Tx
}

// 상품 구매하기
// func (p *person) purchase(userID, walletID){

// }

// 블록 1개 조회
func getBlock(w http.ResponseWriter, req *http.Request /*구매자가 넘긴값*/){
	bs := &blocks
	b := &block
	blkID := b.getBlockID("txID값")
	blk := bs.getBlock(blkID)
	JsonBlockResponse := bs.getBlock(blkID)

	bb, _ := json.Marshal(JsonBlockResponse)
	fmt.Println(bb)
	fmt.println(string(bb))
	
	var response = JsonResponse{
		BlockID : blk.Hash, 
		TxiD : blk.Txid, 
		Data : blk.Data,
		Timestamp : blk.Timestamp
	}
	json.NewEncoder(w).Encode(response)

}

// Tx 조회
func (p *person) getTx(userID){
	
}
// 잔고 조회
func (p *person) getBalance(userID, walletID){
	
}
// 충전
func (p *person) charge(userID, walletID){
	
}
// 지갑 생성
func (app App) mkWallet(userID){
	
}
