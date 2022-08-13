package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"src/httpServer/HttpServerAPI/ConsensusAPI"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/CompleteConsensus", ConsensusAPI.ReplyFromConsensus).Methods("POST")

	http.Handle("/", router)
	err := http.ListenAndServe(":7000", nil)

	if err != nil {
		fmt.Println("Failed To ListenAndServe : ", err)
	}
}
