package main

import (
	"fmt"
	"net/http"
	"src/httpServer"
	"src/restAPI"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/MakeWallet", restAPI.MakeWallet).Methods("POST")
	router.HandleFunc("/GetWalletInfo", restAPI.GetWalletInfo).Methods("POST")
	router.HandleFunc("/BlockManagement", restAPI.BlockManagement).Methods("POST")

	router.HandleFunc("/GenerateTx", httpServer.GenerateTx).Methods("POST")
	router.HandleFunc("/GenerateBlock", httpServer.GenerateBlock).Methods("POST")

	http.Handle("/", router)
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Println("Failed To ListenAndServe : ", err)
	}
}
