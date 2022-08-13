package main

import (
	"fmt"
	"net/http"
	"src/httpServer/HttpServerAPI/WalletAPI"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/MakeWallet", WalletAPI.MakeWallet).Methods("POST")
	router.HandleFunc("/GetWalletInfo", WalletAPI.GetWalletInfo).Methods("POST")

	http.Handle("/", router)
	err := http.ListenAndServe(":3030", nil)

	if err != nil {
		fmt.Println("Failed To ListenAndServe : ", err)
	}
}
