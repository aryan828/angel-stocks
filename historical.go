package stocks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	NSE = "NSE"
	NFO = "NFO"
)

const (
	ONE_MINUTE     = "ONE_MINUTE"
	THREE_MINUTE   = "THREE_MINUTE"
	FIVE_MINUTE    = "FIVE_MINUTE"
	TEN_MINUTE     = "TEN_MINUTE"
	FIFTEEN_MINUTE = "FIFTEEN_MINUTE"
	THIRTY_MINUTE  = "THIRTY_MINUTE"
	ONE_HOUR       = "ONE_HOUR"
	ONE_DAY        = "ONE_DAY"
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

func (c *Client) History(exchange, symbol, interval, from, to string) (data HistoricalDataResponse, err error) {
	payload := map[string]string{
		"exchange":    exchange,
		"symboltoken": symbol,
		"interval":    interval,
		"fromdate":    from,
		"todate":      to,
	}
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

func (r Record) String() string {
	return fmt.Sprintf("{\n\ttime: %s\n\topen: %f\n\thigh: %f\n\tlow: %f\n\tclose: %f\n\tvolume: %d\n}", r.Timestamp, r.Open, r.High, r.Low, r.Close, r.Volume)
}
