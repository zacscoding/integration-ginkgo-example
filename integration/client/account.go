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

type ErrResponse struct {
	StatusCode int
	Message    string `json:"message"`
}

func (e *ErrResponse) Error() string {
	return fmt.Sprintf("StatusCode:%d, Message:%s", e.StatusCode, e.Message)
}

func (c *AccountClient) GetAccounts() ([]*Account, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/accounts", c.Endpoint), nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("unknown error: " + string(b))
	}
	var accounts []*Account
	err = json.Unmarshal(b, &accounts)
	return accounts, err
}

func (c *AccountClient) SaveAccount(username, email string) (*Account, error) {
	body := map[string]interface{}{
		"username": username,
		"email":    email,
	}
	b, _ := json.Marshal(&body)
	request, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/account", c.Endpoint), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, _ = ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return nil, errors.New("unknown error: " + string(b))
	}
	var account Account
	err = json.Unmarshal(b, &account)
	return &account, err
}

func (c *AccountClient) GetAccount(id string) (*Account, error) {
	panic("")
}
