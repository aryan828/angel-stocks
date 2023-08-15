package main

import (
	"io"
	"log"
	"net"
	"net/http"
)

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
