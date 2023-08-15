package main

import "net/http"

type Client struct {
	apiKey      string
	client      http.Client
	credentials map[string]string
	tokens      struct {
		jwt     string
		refresh string
		feed    string
	}
}

func (c *Client) setHeaders(req *http.Request) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-ClientLocalIP", getLocalIP())
	req.Header.Add("X-ClientPublicIP", getPublicIP())
	req.Header.Add("X-MACAddress", getMACAddr())
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-PrivateKey", c.apiKey)
	req.Header.Add("X-UserType", "USER")
	req.Header.Add("X-SourceID", "WEB")
	req.Header.Add("Authorization", "Bearer "+c.tokens.jwt)
}
