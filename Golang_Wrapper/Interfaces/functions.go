package prexis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	config "../Configuration"
)

const (
	Auth             = "/auth"
	Credits          = "/credits"
	RegisterHash     = "/registerhash"
	RegisterDocument = "/registerdoc"
	GetTransaction   = "/transaction"
)

func Authenticate() *Authentication {
	config.Init()
	conf := config.GetConfig()

	var response Authentication

	req, err := http.NewRequest("POST", conf.Prexis.ApiURL+Auth, nil)
	if err != nil {
		log.Fatalln(err)
	}

	q := req.URL.Query()

	q.Add("key", conf.Prexis.Key)
	q.Add("password", conf.Prexis.Password)

	req.URL.RawQuery = q.Encode()

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending the request: ", err.Error())
	}
	defer resp.Body.Close()

	content, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(content, &response)

	return &response

}
