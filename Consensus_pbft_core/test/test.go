package main

//import (
//	"bytes"
//	"encoding/json"
//	"fmt"
//	"net/http"
//	"time"
//)
//
//var counter int = 0
//
//type RequestMsg struct {
//	Timestamp int64  `json:"timestamp"`
//	ClientID  string `json:"clientID"`
//	Operation string `json:"operation"`
//}
//
//func main() {
//	for i := 0; i < 1050; i++ {
//		requestMsg := RequestMsg{
//			Timestamp: 102938172,
//			ClientID:  "kabigon",
//			Operation: "GetMyName",
//		}
//
//		pbytes, _ := json.Marshal(requestMsg)
//		data := bytes.NewBuffer(pbytes)
//
//		http.Post("http://localhost:5000/req", "application/json", data)
//
//		counter += 1
//		fmt.Println(counter)
//
//		time.Sleep(time.Millisecond * 50)
//	}
//}
