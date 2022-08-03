package BlockAPI

type ManagementRequest struct {
	From    string `json:"From"`
	To      string `json:"To"`
	Purpose string `json:"Purpose"`
	Amount  int64  `json:"Amount"`
}

type RequestMsg struct {
	Timestamp int64  `json:"timestamp"`
	ClientID  string `json:"clientID"`
	Operation string `json:"operation"`
}

type MakeBlockArgs struct {
	From   string `json:"From"`
	To     string `json:"To"`
	Amount int64  `json:"Amount"`
}

type MakeBlockResponse struct {
	Hash string
}

type RefBlockArgs struct {
	From string `json:"From"`
}

type BlockResponse struct {
	Hash string
}
