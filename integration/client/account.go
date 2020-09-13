package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type AccountClient struct {
	Endpoint string
}

type Account struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

func (c *AccountClient) GetAccounts() ([]*Account, int, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/accounts", c.Endpoint), nil)
	if err != nil {
		return nil, 0, err
	}

	var accounts []*Account
	code, err := doRequest(request, &accounts)
	return accounts, code, err
}

func (c *AccountClient) SaveAccount(username, email string) (*Account, int, error) {
	body := map[string]interface{}{
		"username": username,
		"email":    email,
	}
	b, _ := json.Marshal(&body)
	request, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/account", c.Endpoint), bytes.NewBuffer(b))
	if err != nil {
		return nil, 0, err
	}

	var account Account
	code, err := doRequest(request, &account)
	return &account, code, err
}

func (c *AccountClient) GetAccount(id string) (*Account, int, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/account/%s", c.Endpoint, id), nil)
	if err != nil {
		return nil, 0, err
	}

	var account Account
	code, err := doRequest(request, &account)
	return &account, code, err
}

func (c *AccountClient) UpdateAccount(id string, username string) (*StatusResponse, int, error) {
	body := map[string]interface{}{
		"username": username,
	}
	b, _ := json.Marshal(&body)
	request, err := http.NewRequest("PUT", fmt.Sprintf("%s/v1/account/%s", c.Endpoint, id), bytes.NewBuffer(b))
	if err != nil {
		return nil, 0, err
	}

	var response StatusResponse
	code, err := doRequest(request, &response)
	return &response, code, err
}

func (c *AccountClient) DeleteAccount(id string) (*StatusResponse, int, error) {
	request, err := http.NewRequest("DELETE", fmt.Sprintf("%s/v1/account/%s", c.Endpoint, id), nil)
	if err != nil {
		return nil, 0, err
	}

	var response StatusResponse
	code, err := doRequest(request, &response)
	return &response, code, err
}

func doRequest(request *http.Request, v interface{}) (int, error) {
	request.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return resp.StatusCode, errors.New("unknown error: " + string(b))
	}
	return resp.StatusCode, json.Unmarshal(b, v)
}
