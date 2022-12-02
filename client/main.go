package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type AccountsClient struct {
	url    string
	client *http.Client
}

func NewAccountsClient(baseUrl string) *AccountsClient {
	return &AccountsClient{baseUrl + "/v1/organisation/accounts", &http.Client{}}
}

func Get(accountId string) (*AccountsResponse, error) {
	if accountId == "" {
		return nil, errors.New("Invalid account id")
	}

	requestURL := fmt.Sprintf("http://localhost:%d/v1/organisation/accounts", 8080)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)
	defer res.Body.Close()

	accountsResponse := AccountsResponse{}
	fmt.Println("Response: ", res.Body)
	error := readBody(res, &accountsResponse)
	fmt.Println(accountsResponse)
	return &accountsResponse, error
}

func post(accountData *AccountData) (*AccountResponse, error) {

	requestURL := fmt.Sprintf("http://localhost:%d/v1/organisation/accounts", 8080)
	res, err := http.Get(requestURL)

	req := AccountRequest{
		Data: *accountData,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	createRequest := AccountResponse{}
	createReq, err := http.Post(requestURL, "application/json", bytes.NewBuffer(requestBody))

	defer res.Body.Close()

	fmt.Println("Response: ", createReq.Body)
	error := json.NewDecoder(createReq.Body).Decode(&createRequest)
	fmt.Println(createRequest)
	return &createRequest, error

}

func readBody(res *http.Response, v interface{}) error {
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, v)
}
