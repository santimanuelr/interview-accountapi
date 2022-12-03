package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type AccountsClient struct {
	url    string
	client *http.Client
}

func NewAccountsClient(baseUrl string) *AccountsClient {
	return &AccountsClient{baseUrl + "/v1/organisation/accounts", &http.Client{}}
}

func (c AccountsClient) newRequest(method string, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, c.url+path, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/vnd.api+json")
	req.Header.Add("Date", time.Now().GoString())

	return req, nil
}

func (c AccountsClient) Fetch(id string) (*AccountData, error) {
	req, err := c.newRequest("GET", "/"+id, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req)

	err = checkStatus(res, err)
	if err != nil {
		return nil, err
	}

	var account AccountResponse
	fmt.Println("Response: ", res.Body)
	error := readBody(res, &account)
	fmt.Println(account)
	return &account.Data, error
}

func checkStatus(res *http.Response, responseError error) error {
	if responseError != nil {
		return responseError
	}
	if res.StatusCode < 400 {
		return nil
	}

	var errorBody ErrorData
	err := readBody(res, &errorBody)
	if err != nil {
		return err
	}

	return errors.New(errorBody.ErrorMessage)
}

func (c AccountsClient) Create(accountData *AccountData) (*AccountResponse, error) {

	accountRequest := AccountRequest{
		Data: *accountData,
	}

	requestBody, err := json.Marshal(accountRequest)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	req, err := c.newRequest("POST", "", bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req)

	err = checkStatus(res, err)
	if err != nil {
		return nil, err
	}

	createRequest := AccountResponse{}

	defer res.Body.Close()

	fmt.Println("Response: ", res.Body)
	error := json.NewDecoder(res.Body).Decode(&createRequest)
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
