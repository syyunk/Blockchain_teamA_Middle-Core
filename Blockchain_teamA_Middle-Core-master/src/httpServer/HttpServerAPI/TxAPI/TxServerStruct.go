package TxAPI

type TxRequest struct {
	From    string `json:"From"`
	To      string `json:"To"`
	Amount  int64  `json:"Amount"`
	Purpose string `json:"Purpose`
}

type TxResponse struct {
	Txid   []byte `json:"Txid"`
	From   string `json:"From"`
	To     string `json:"To"`
	Amount int64  `json:"Amount"`
}

type RefTxArgs struct {
	TxidByhex string `json:"TxidByhex"`
}

type RefTxResponse struct {
	From   string `json:"From"`
	To     string `json:"To"`
	Amount int64  `json:"Amount"`
	Txid   string `json:"Txid"`
}
