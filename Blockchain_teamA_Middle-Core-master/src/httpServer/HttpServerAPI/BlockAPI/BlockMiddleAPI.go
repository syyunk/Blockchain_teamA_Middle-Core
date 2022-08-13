package BlockAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"src/httpServer/HttpServerAPI/TxAPI"
)

func BlockManagement(rw http.ResponseWriter, req *http.Request) {
	var managementRequest ManagementRequest

	decoder := json.NewDecoder(req.Body)
	decoder.Decode(&managementRequest)

	if managementRequest.Purpose == "블록 조회" {
		//txid를 담고 있을 것
		managementArgs := new(RefBlockArgs)

		managementArgs.TxidByhex = managementRequest.TxidByhex
		fmt.Println("managementRequest :::  ", managementRequest.TxidByhex)
		pbytes, _ := json.Marshal(managementArgs)
		sendData := bytes.NewBuffer(pbytes)
		fmt.Println("sendData :: ", sendData)

		resp, _ := http.Post("http://localhost:9000/GetBlock", "application/json", sendData)

		decoder := json.NewDecoder(resp.Body)
		decoder.DisallowUnknownFields()

		var managementResponse GetBlockResponse

		decoder.Decode(&managementResponse)

		pbytes2, _ := json.Marshal(managementResponse)
		fmt.Println("getBlock 응답값 : ", managementResponse)
		rw.Write(pbytes2)
	}

	if managementRequest.Purpose == "블록 생성" {
		managementArgs := new(MakeBlockArgs)

		managementArgs.From = managementRequest.From
		managementArgs.To = managementRequest.To
		managementArgs.Amount = managementRequest.Amount

		pbytes, _ := json.Marshal(managementArgs)
		sendData := bytes.NewBuffer(pbytes)

		resp, _ := http.Post("http://localhost:9000/GenerateBlock", "application/json", sendData)

		decoder := json.NewDecoder(resp.Body)
		decoder.DisallowUnknownFields()

		var managementResponse MakeBlockResponse

		decoder.Decode(&managementResponse)

		pbytes2, _ := json.Marshal(managementResponse)

		rw.Write(pbytes2)
	}

	if managementRequest.Purpose == "tx 조회" {
		managementArgs := new(TxAPI.RefTxArgs)

		managementArgs.TxidByhex = managementRequest.TxidByhex

		pbytes, _ := json.Marshal(managementArgs)
		sendData := bytes.NewBuffer(pbytes)

		resp, _ := http.Post("http://localhost:9000/RefTxFromBlk", "application/json", sendData)

		decoder := json.NewDecoder(resp.Body)
		decoder.DisallowUnknownFields()

		var managementResponse TxAPI.RefTxResponse

		decoder.Decode(&managementResponse)

		pbytes2, _ := json.Marshal(managementResponse)

		rw.Write(pbytes2)
	}

}
