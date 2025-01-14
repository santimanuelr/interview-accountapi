package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
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

	return req, nil
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

	defer res.Body.Close()

	return errors.New(errorBody.ErrorMessage)
}

func (c AccountsClient) Fetch(id string) (*AccountData, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/%s", id), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req)

	err = checkStatus(res, err)
	if err != nil {
		return nil, err
	}

	var account AccountResponse
	error := readBody(res, &account)
	log.Println("Response: ", account)
	return &account.Data, error
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

	error := readBody(res, &createRequest)
	log.Println("Response: ", createRequest)
	return &createRequest, error
}

func (c AccountsClient) Delete(accountData *AccountData) error {
	version := strconv.FormatInt(*accountData.Version, 10)
	req, err := c.newRequest("DELETE", "/"+accountData.ID+"?version="+version, nil)
	if err != nil {
		return err
	}

	res, err := c.client.Do(req)

	return checkStatus(res, err)
}

func readBody(res *http.Response, v interface{}) error {
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, v)
}
