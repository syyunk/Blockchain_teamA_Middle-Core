package main

import "github.com/gorilla/mux"

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/MakeWallet")
}
