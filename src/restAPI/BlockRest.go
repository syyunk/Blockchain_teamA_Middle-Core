package restAPI

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type ManagementRequest struct {
	From    string `json:"From"`
	To      string `json:"To"`
	Purpose string `json:"Purpose"`
	Amount  int64  `json:"Amount"`
}

type MakeBlockArgs struct {
	From   string `json:"From"`
	To     string `json:"To"`
	Amount int64  `json:"Amount"`
}

type MakeBlockResponse struct {
	Hash string
}

type RefBlockArgs struct {
	From string `json:"From"`
}

func BlockManagement(rw http.ResponseWriter, req *http.Request) {
	var managementRequest ManagementRequest

	decoder := json.NewDecoder(req.Body)
	decoder.Decode(&managementRequest)

	if managementRequest.Purpose == "내역 조회" {
		managementArgs := new(RefBlockArgs)

		managementArgs.From = managementRequest.From
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
}
