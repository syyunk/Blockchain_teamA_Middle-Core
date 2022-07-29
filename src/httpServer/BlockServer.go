package httpServer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"src/block"
)

type BlockResponse struct {
	Hash string
}

func GenerateBlock(rw http.ResponseWriter, req *http.Request) {
	var body TxResponse

	resp, err := http.Post("http://localhost:10000/GenerateTx", "application/json", req.Body)

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
