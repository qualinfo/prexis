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
	Auth        = "/auth"
	Credits     = "/credits"
	Hash        = "/registerhash"
	Document    = "/registerdoc"
	Transaction = "/transaction"
)

func Authenticate() *AuthenticationStruct {

	config.Init()
	conf := config.GetConfig()

	var response AuthenticationStruct

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

func CreditBalance(token string) *BalanceStruct {

	conf := config.GetConfig()

	req, err := http.NewRequest("GET", conf.Prexis.ApiURL+Credits, nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	headerValue := "Bearer " + token

	req.Header.Set("Authorization", headerValue)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}

	var response BalanceStruct

	json.Unmarshal(body, &response)

	return &response

}

func RegisterHash(token, hash string) *HashStruct {

	conf := config.GetConfig()

	req, err := http.NewRequest("POST", conf.Prexis.ApiURL+Hash, nil)
	if err != nil {
		log.Fatalln(err)
	}

	q := req.URL.Query()

	q.Add("hash", hash)
	q.Add("email", conf.Prexis.Email)

	req.URL.RawQuery = q.Encode()

	headerValue := "Bearer " + token

	req.Header.Set("Authorization", headerValue)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}

	var response HashStruct

	json.Unmarshal(body, &response)

	return &response

}

func GetTransaction(token, hash string) *TransactionStruct {

	conf := config.GetConfig()

	req, err := http.NewRequest("POST", conf.Prexis.ApiURL+Transaction, nil)
	if err != nil {
		log.Fatalln(err)
	}

	q := req.URL.Query()

	q.Add("hash", hash)

	req.URL.RawQuery = q.Encode()

	headerValue := "Bearer " + token

	req.Header.Set("Authorization", headerValue)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}

	var response TransactionStruct

	json.Unmarshal(body, &response)

	return &response

}
