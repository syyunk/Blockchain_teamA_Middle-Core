package main

import (
	"fmt"
	"net/http"
	"src/httpServer/HttpServerAPI/TxAPI"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/GenerateTx", TxAPI.GenerateTx).Methods("POST")
	router.HandleFunc("/GetTx", TxAPI.GetTx).Methods("POST")

	http.Handle("/", router)
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Println("Failed To ListenAndServe : ", err)
	}
}
