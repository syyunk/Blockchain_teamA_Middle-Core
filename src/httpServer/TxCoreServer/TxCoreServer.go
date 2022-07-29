package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"src/httpServer/HttpServerAPI/TxAPI"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/GenerateTx", TxAPI.GenerateTx).Methods("POST")

	http.Handle("/", router)
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Println("Failed To ListenAndServe : ", err)
	}
}
