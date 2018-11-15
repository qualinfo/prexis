package prexis

type Authentication struct {
	Cached  string `json:"cached"`
	Code    int    `json:"code"`
	Jwt     string `json:"jwt"`
	Key     string `json:"key"`
	Message string `json:"message"`
	Status  string `json:"status"`
}
