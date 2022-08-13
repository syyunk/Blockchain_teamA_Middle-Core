package ConsensusAPI

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var replyMsgs = make([]*ReplyMsg, 0)

type ReplyMsg struct {
	ViewID    int64  `json:"viewID"`
	Timestamp int64  `json:"timestamp"`
	ClientID  string `json:"clientID"`
	NodeID    string `json:"nodeID"`
	Result    string `json:"result"`
}

func ReplyFromConsensus(rw http.ResponseWriter, req *http.Request) {
	var msg ReplyMsg

	err := json.NewDecoder(req.Body).Decode(&msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	replyMsgs = append(replyMsgs, &msg)

	if len(replyMsgs) == 4 {
		fmt.Println("합의 완료!!!!=====================")
		for i := 0; i < 4; i++ {
			fmt.Println(replyMsgs[i].NodeID)
		}
		fmt.Println("=================================")

		replyMsgs = make([]*ReplyMsg, 0)

		http.Get("http://localhost:9000/SetConcensusCompleteFlag")
		if err != nil {
			fmt.Println(err)
		}
	}
}
