package main

import(
	"fmt"
	"encoding/json"
	"bytes"
	"net/http"
	
)
// 사용자에게 반환할 블록
type JsonBlockResponse struct {
	Hash     	[]byte 	// BlkID O
	Txid     	[]byte 	// O
	bits    	int64 	// targetBytes of Pow
	Height  	int64
	Timestamp 	[]byte 	// local time
	Data 		[]byte 	// copyrights
	//Txs 		[]*Tx
}


func main(){
	blk := JsonBlockResponse{}
	//var u = User {"Gopher", 7}
	b, _ := json.Marshal(blk)

	buff := bytes.NewBuffer(b)
	fmt.Println(b)
	fmt.Println(string(b))
	fmt.Println(buff)
	
	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(res http.ResponseWriter, req *http.Request){
		res.Write([]byte(string(b)))
	})
	//req, _ := http.NewRequest("POST", "/movie", buff)
	//response := executeRequest(req)
	http.ListenAndServe(":80", mux)
}