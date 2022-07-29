package TxAPI

type TxRequest struct {
	From   string `json:"From"`
	To     string `json:"To"`
	Amount int64  `json:"Amount"`
}

type TxResponse struct {
	Txid   []byte `json:"Txid"`
	From   string `json:"From"`
	To     string `json:"To"`
	Amount int64  `json:"Amount"`
}
