package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"src/block"
	"src/httpServer/HttpServerAPI/BlockAPI"
)

func main() {
	router := mux.NewRouter()
	block.NewBlockchain(block.GenesisBlock())
	fmt.Println(len(block.Blockchain))

	router.HandleFunc("/GenerateBlock", BlockAPI.GenerateBlock).Methods("POST")
	router.HandleFunc("/SetConcensusCompleteFlag", BlockAPI.SetConcensusCompleteFlag).Methods("GET")

	http.Handle("/", router)
	err := http.ListenAndServe(":9000", nil)

	if err != nil {
		fmt.Println("Failed To ListenAndServe : ", err)
	}
}
