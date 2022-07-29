package BlockAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"src/block"
	"src/httpServer/HttpServerAPI/TxAPI"
	"time"
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

	time.Sleep(50 * time.Millisecond)

	if isConsensusComplete == true {
		var body TxAPI.TxResponse

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

		b := block.NewBlock(
			block.GetCurrentBlockId(),
			body.Txid,
			int64(len(block.Blockchain)-1),
			[]byte(""),
		)

		block.AddBlock(b)

		response := &BlockResponse{Hash: fmt.Sprintf("%x", b.Hash)}
		rw.Header().Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(response)
	}
}

// 합의가 완료되었을 때 발동하는 라우팅 함수
// isConsensusComplete 값을 false에서 true 바꿔 합의가 끝났음을 처리해준다.
func SetConcensusCompleteFlag(rw http.ResponseWriter, req *http.Request) {
	isConsensusComplete = true
}
