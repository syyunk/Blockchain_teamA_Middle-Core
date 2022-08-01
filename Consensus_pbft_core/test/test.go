package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var counter int = 0

type RequestMsg struct {
	Timestamp int64  `json:"timestamp"`
	ClientID  string `json:"clientID"`
	Operation string `json:"operation"`
}

func main() {
	// 시작 시간
    startTime := time.Now()
	for i := 0; i < 130; i++ {
		requestMsg := RequestMsg{
			Timestamp: 102938172,
			ClientID:  "kabigon",
			Operation: "GetMyName",
		}

		pbytes, _ := json.Marshal(requestMsg)
		data := bytes.NewBuffer(pbytes)

		http.Post("http://localhost:5000/req", "application/json", data)

		counter += 1
		fmt.Println(counter)

		time.Sleep(time.Millisecond * 75)

		//i=130, sleep = 80 12.31초 106완료
		//i=130. sleep = 75 11.06초 107완료 (ResolvingTimeDuration = time.Millisecond / 100)
		//i=130, sleep=70 10.1316 102완료

		//500 -> 51초 96Complete
		//300 -> 31초 94Complete
		//200 -> 21초 90Complete
		//160 -> 17초 86Complete
		//110 -> 12.64 83Complete
		//100 -> 11.10 80Complete
		//80 -> 9.54 77Complete
		//consensusPBFT.exe P1
	}
	 // 경과 시간
	 elapsedTime := time.Since(startTime)
     
	 fmt.Printf("실행시간: %s\n", elapsedTime)
}
