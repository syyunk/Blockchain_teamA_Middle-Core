package httpServer

import (
	"encoding/json"
	"net/http"
	"src/transaction"
)

type TxRequest struct {
	From   string `json:"From"`
	To     string `json:"To"`
	Amount int64  `json:"Amount"`
}

type TxResponse struct {
	Txid   []byte `json:"Txid"`
	From   string `json:"From"`
	To     string `json:"To"`
	Amount int64  `json:"Amount"`
}

func GenerateTx(wr http.ResponseWriter, req *http.Request) {
	var body TxRequest

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&body)

	if err != nil {
		panic(err)
	}

	tx := transaction.NewTransaction([]byte(body.From), []byte(body.To), body.Amount)

	transaction.AddTx(tx)

	response := &TxResponse{Txid: tx.Txid, From: body.From, To: body.To, Amount: body.Amount}

	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode(response)
}
