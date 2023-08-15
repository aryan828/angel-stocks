package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

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
