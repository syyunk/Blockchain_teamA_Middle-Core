package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"src/httpServer"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/GenerateTx", httpServer.GenerateTx).Methods("POST")

	http.Handle("/", router)
	err := http.ListenAndServe(":10000", nil)

	if err != nil {
		fmt.Println("Failed To ListenAndServe : ", err)
	}
}
