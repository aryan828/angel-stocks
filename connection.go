package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

const (
	loginurl = "https://apiconnect.angelbroking.com/rest/auth/angelbroking/user/v1/loginByPassword"
)

func (c *Client) Connect() (data ConnectionResponse, err error) {
	requestData, err := json.Marshal(c.credentials)
	if err != nil {
		return
	}

	request, err := http.NewRequest(http.MethodPost, loginurl, bytes.NewBuffer(requestData))
	if err != nil {
		return
	}

	c.setHeaders(request)

	resp, err := c.client.Do(request)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return
	}

	c.tokens.jwt = data.Data.JwtToken
	c.tokens.refresh = data.Data.RefreshToken
	c.tokens.feed = data.Data.FeedToken

	return
}
