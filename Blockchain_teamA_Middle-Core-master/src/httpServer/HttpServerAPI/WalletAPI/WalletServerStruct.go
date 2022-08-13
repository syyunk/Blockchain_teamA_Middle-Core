package WalletAPI

type WalletArgs struct {
	Alias   string
	Address string
}

type WalletReply struct {
	Alias   string
	Address string
	PrvKey  []byte
	PubKey  []byte
}
