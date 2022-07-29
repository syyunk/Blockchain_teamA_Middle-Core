package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"src/httpServer"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/MakeWallet", httpServer.MakeWallet).Methods("POST")
	router.HandleFunc("/GetWalletInfo", httpServer.GetWalletInfo).Methods("POST")

	http.Handle("/", router)
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Println("Failed To ListenAndServe : ", err)
	}
}
