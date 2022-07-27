package main

import (
    "fmt"
)

func main(){

    wl := NewWallet()

    fmt.Println(wl.address)
    // wl := &wallet

    // adress := wl.GetAddress()
    // fmt.Println(adress)
}