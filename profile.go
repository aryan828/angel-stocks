package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ProfileResponse struct {
	BaseResponse
	Data struct {
		ClientCode    string
		Name          string
		Email         string
		Mobileno      string
		Exchanges     []string
		Products      []string
		LastLoginTime string
		BrokerID      string
	}
}

func (c *Client) GetProfile() (profile ProfileResponse, err error) {
	request, err := http.NewRequest(http.MethodGet, "https://apiconnect.angelbroking.com/rest/secure/angelbroking/user/v1/getProfile", &bytes.Buffer{})
	if err != nil {
		return
	}

	c.setHeaders(request)

	response, err := c.client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &profile)
	if err != nil {
		return
	}

	return
}

func (pr ProfileResponse) String() string {
	return fmt.Sprintf("User Data:\n- Angel Client Code: %s\n- Name: %s\n- Email: %s\n- Mobile No: %s\n- Exchanges: %s\n- Products: %s", pr.Data.ClientCode, pr.Data.Name, pr.Data.Email, pr.Data.Mobileno, pr.Data.Exchanges, pr.Data.Products)
}
