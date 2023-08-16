package stocks

import (
	"io"
	"log"
	"net"
	"net/http"
)

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

type BaseResponse struct {
	Status    bool
	Message   string
	ErrorCode string
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

func getLocalIP() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Panicln("Can't fetch interfaces")
	}
	for _, addr := range interfaces {
		if len(addr.HardwareAddr) != 0 {
			addrs, err := addr.Addrs()
			if err != nil {
				log.Panicln("Can't fetch interface address")
			}

			return addrs[0].String()
		}
	}
	log.Panicln("Failed fetching local IP")
	return ""
}

func getMACAddr() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Panicln("Can't fetch interfaces")
	}
	for _, addr := range interfaces {
		if len(addr.HardwareAddr) != 0 {
			return addr.HardwareAddr.String()
		}
	}
	log.Panicln("Failed fetching MAC address")
	return ""
}

func getPublicIP() string {
	res, err := http.Get("https://api.ipify.org")
	if err != nil {
		log.Panicln("Can't fetch public IP address")
	}

	publicIP, err := io.ReadAll(res.Body)
	if err != nil {
		log.Panicln("Failed to read public IP address")
	}
	return string(publicIP)
}
