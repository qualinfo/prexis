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
