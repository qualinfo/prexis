package prexis

type AuthenticationStruct struct {
	Cached  string `json:"cached"`
	Code    int    `json:"code"`
	Jwt     string `json:"jwt"`
	Key     string `json:"key"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type BalanceStruct struct {
	Code        int `json:"code"`
	CurrentTime int `json:"current_time"`
	Result      struct {
		Credits int    `json:"credits"`
		Message string `json:"message"`
		Rescode int    `json:"rescode"`
	} `json:"result"`
	Status string `json:"status"`
}

type HashStruct struct {
	Code        int `json:"code"`
	CurrentTime int `json:"current_time"`
	Result      struct {
		Credits int    `json:"credits"`
		Message string `json:"message"`
		Rescode int    `json:"rescode"`
		Hash    string `json:"hash"`
	} `json:"result"`
	Status string `json:"status"`
}

type TransactionStruct struct {
	Code        int    `json:"code"`
	CurrentTime int    `json:"current_time"`
	Key         string `json:"key"`
	Result      struct {
		Hash        string `json:"hash"`
		Message     string `json:"message"`
		Rescode     int    `json:"rescode"`
		Transaction struct {
			BlockHash         string `json:"blockHash"`
			BlockNumber       int    `json:"blockNumber"`
			BlockchainURL     string `json:"blockchainUrl"`
			CumulativeGasUsed int    `json:"cumulativeGasUsed"`
			From              string `json:"from"`
			GasUsed           int    `json:"gasUsed"`
			Timestamp         int    `json:"timestamp"`
			TransactionHash   string `json:"transactionHash"`
			TransactionIndex  int    `json:"transactionIndex"`
		} `json:"transaction"`
	} `json:"result"`
	Status string `json:"status"`
}
