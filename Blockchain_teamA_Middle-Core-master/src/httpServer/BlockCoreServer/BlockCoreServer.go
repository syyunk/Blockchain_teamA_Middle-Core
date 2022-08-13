package main

import (
	"fmt"
	"net/http"
	"src/block"
	"src/httpServer/HttpServerAPI/BlockAPI"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	block.NewBlockchain(block.GenesisBlock())
	fmt.Println(len(block.Blockchain))

	router.HandleFunc("/GenerateBlock", BlockAPI.GenerateBlock).Methods("POST")
	router.HandleFunc("/GetBlock", BlockAPI.GetBlock).Methods("POST")
	router.HandleFunc("/RefTxFromBlk", BlockAPI.RefTxFromBlk).Methods("POST")
	router.HandleFunc("/SetConcensusCompleteFlag", BlockAPI.SetConcensusCompleteFlag).Methods("GET")

	http.Handle("/", router)
	err := http.ListenAndServe(":9000", nil)

	if err != nil {
		fmt.Println("Failed To ListenAndServe : ", err)
	}
}
