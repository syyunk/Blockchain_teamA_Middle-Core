package main

import (
	"fmt"
	"net/http"
	"src/httpServer/HttpServerAPI/BlockAPI"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/MainBlockServer", BlockAPI.BlockManagement).Methods("POST")

	http.Handle("/", router)
	err := http.ListenAndServe(":10000", nil)

	if err != nil {
		fmt.Println("Failed To ListenAndServe : ", err)
	}
}
