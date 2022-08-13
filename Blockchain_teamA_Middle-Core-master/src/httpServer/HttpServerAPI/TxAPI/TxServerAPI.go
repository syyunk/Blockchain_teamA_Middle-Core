package TxAPI

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"src/transaction"
)

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

func GetTx(wr http.ResponseWriter, req *http.Request) {
	var body RefTxArgs

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	fmt.Println(req.Body)
	_ = decoder.Decode(&body)

	// if err != nil {
	// 	panic(err)
	// }

	convertTxid, _ := hex.DecodeString(body.TxidByhex)

	Reftx := transaction.GetTransaction(convertTxid)

	fmt.Println(Reftx)

	resp := &RefTxResponse{Txid: hex.EncodeToString(Reftx.Txid), From: string(Reftx.From), To: string(Reftx.To), Amount: Reftx.Amount}
	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode(resp)

}
