package BlockAPI

type ManagementRequest struct {
	From      string `json:"From"`
	To        string `json:"To"`
	Purpose   string `json:"Purpose"`
	Amount    int64  `json:"Amount"`
	TxidByhex string `json:"TxidByhex"`
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
	Txid string `json:"Txid"`
}

type RefBlockArgs struct {
	TxidByhex string `json:"TxidByHex"`
}

type BlockResponse struct {
	Hash string
	Txid string `json:"Txid"`
}

type GetBlockResponse struct {
	Hash      string // BlkID O
	Data      string // copyrights
	Txid      string // 확인용
	Timestamp []byte // local time
}
