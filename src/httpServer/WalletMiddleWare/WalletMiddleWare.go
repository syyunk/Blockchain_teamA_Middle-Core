package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"src/httpServer/HttpServerAPI/WalletAPI"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/MakeWallet", WalletAPI.MakeWallet).Methods("POST")
	router.HandleFunc("/GetWalletInfo", WalletAPI.GetWalletInfo).Methods("POST")

	http.Handle("/", router)
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Println("Failed To ListenAndServe : ", err)
	}
}
