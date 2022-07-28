package httpServer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TxArgs struct {
	From   []byte
	To     []byte
	Amount int64
}

func NewTx(rw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	b, _ := ioutil.ReadAll(req.Body)

	args := &TxArgs{}

	err := json.Unmarshal(b, &args)

	if err != nil {
		panic(err)
	}

	fmt.Println(args.From)

	// tx := transaction.NewTransaction(args.From, args.To, args.Amount)

	// response, _ := json.Marshal(tx)

	// rw.Write(response)
}
