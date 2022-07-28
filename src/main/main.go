package main

import (
	"fmt"
	"net/http"
	"src/httpServer"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/MakeWallet", httpServer.MakeWallet).Methods("POST")
	router.HandleFunc("/GetWalletInfo", httpServer.GetWalletInfo).Methods("POST")
	router.HandleFunc("/NewTx", httpServer.NewTx).Methods("POST")

	http.Handle("/", router)
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Println("Failed To ListenAndServe : ", err)
	}
}
