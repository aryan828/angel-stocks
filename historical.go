package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Record struct {
	Timestamp time.Time
	Open      float64
	High      float64
	Low       float64
	Close     float64
	Volume    int64
}

type HistoricalDataResponse struct {
	BaseResponse
	Data []Record
}

func (c *Client) History(payload map[string]string) (data HistoricalDataResponse, err error) {
	requestData, err := json.Marshal(payload)
	if err != nil {
		return
	}

	request, err := http.NewRequest(http.MethodPost, "https://apiconnect.angelbroking.com/rest/secure/angelbroking/historical/v1/getCandleData", bytes.NewBuffer(requestData))
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

	return
}

func (r *Record) UnmarshalJSON(data []byte) error {
	var v []interface{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("Error while decoding %v\n", err)
		return err
	}

	r.Timestamp, err = time.Parse(time.RFC3339, v[0].(string))
	if err != nil {
		return err
	}

	r.Open = v[1].(float64)
	r.High = v[2].(float64)
	r.Low = v[3].(float64)
	r.Close = v[4].(float64)
	r.Volume = int64(v[5].(float64))

	return nil
}
