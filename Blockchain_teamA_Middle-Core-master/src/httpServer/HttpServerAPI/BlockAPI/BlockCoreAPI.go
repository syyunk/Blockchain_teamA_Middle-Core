package BlockAPI

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"src/block"
	"src/httpServer/HttpServerAPI/TxAPI"
)

var isConsensusComplete bool = false

func GenerateBlock(rw http.ResponseWriter, req *http.Request) {
	// Timestamp는 거래가 발생한 시간을 기입할것!
	// ClientID는 지갑주소나 임의의 값을 지정
	requestMsg := RequestMsg{
		Timestamp: 102938172,
		ClientID:  "kabigon",
		Operation: "Trade",
	}
	pbytes, _ := json.Marshal(requestMsg)
	data := bytes.NewBuffer(pbytes)

	http.Post("http://localhost:5000/req", "application/json", data)

	// time.Sleep(50 * time.Millisecond) //블록 무작위로 생성안될수 수 있음
	for isConsensusComplete != true {
		continue
	}
	if isConsensusComplete == true {
		var body TxAPI.TxResponse

		var inputdata = ""

		isConsensusComplete = false

		resp, err := http.Post("http://localhost:8000/GenerateTx", "application/json", req.Body)

		if err != nil {
			panic(err)
		}

		decoder := json.NewDecoder(resp.Body)
		decoder.DisallowUnknownFields()

		err = decoder.Decode(&body)

		//에러 체크
		if err != nil {
			panic(err)
		}

		if body.To == "companywalletid" { //수정 필요***
			inputdata = "충전 내역"
		} else {
			inputdata = "거래 내역"
		}

		b := block.NewBlock(
			block.GetCurrentBlockId(),
			body.Txid,
			int64(len(block.Blockchain)-1),
			[]byte(inputdata),
		)

		block.AddBlock(b)
		response := &BlockResponse{Hash: fmt.Sprintf("%x", b.Hash), Txid: fmt.Sprintf("%x", b.Txid)}
		// response := &BlockResponse{Hash: fmt.Sprintf("%x", b.Hash)}
		rw.Header().Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(response)
	}
}

// 합의가 완료되었을 때 발동하는 라우팅 함수
// isConsensusComplete 값을 false에서 true 바꿔 합의가 끝났음을 처리해준다.
func SetConcensusCompleteFlag(rw http.ResponseWriter, req *http.Request) {
	isConsensusComplete = true
}

//블록 1개 조회
func GetBlock(rw http.ResponseWriter, req *http.Request) {
	//byte 값 다를 경우 예외처리
	fmt.Println("##################333")
	var body RefBlockArgs
	fmt.Println("111")
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&body)

	fmt.Printf("222")
	if err != nil {
		panic(err)
	}

	Txid, _ := hex.DecodeString(body.TxidByhex)
	fmt.Println("333")
	fmt.Println("sendData로 보내주는 Txid :: ", body.TxidByhex)
	b := block.GetBlockOne(Txid)

	managementArgs := new(GetBlockResponse)

	managementArgs.Hash = hex.EncodeToString(b.Hash)
	managementArgs.Data = string(b.Data)
	managementArgs.Timestamp = b.Timestamp
	managementArgs.Txid = hex.EncodeToString(b.Txid)

	fmt.Println("GetBlockResponse :::  ", managementArgs.Hash)

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(managementArgs)
}

func RefTxFromBlk(rw http.ResponseWriter, req *http.Request) {

	var body TxAPI.RefTxResponse

	fmt.Println("tx 조회시 txidbyhex ::: ", req.Body)
	resp, err := http.Post("http://localhost:8000/GetTx", "application/json", req.Body)

	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(resp.Body)
	decoder.DisallowUnknownFields()

	err = decoder.Decode(&body)

	//에러 체크
	if err != nil {
		panic(err)
	}

	response := &TxAPI.RefTxResponse{From: body.From, To: body.To, Amount: body.Amount, Txid: body.Txid}
	// response := &BlockResponse{Hash: fmt.Sprintf("%x", b.Hash)}
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(response)

}
