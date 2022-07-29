package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"src/block"
	"src/httpServer"
)

func main() {
	router := mux.NewRouter()
	block.NewBlockchain(block.GenesisBlock())
	fmt.Println(len(block.Blockchain))

	router.HandleFunc("/GenerateBlock", httpServer.GenerateBlock).Methods("POST")

	http.Handle("/", router)
	err := http.ListenAndServe(":9000", nil)

	if err != nil {
		fmt.Println("Failed To ListenAndServe : ", err)
	}
}
