package util

import (
	"errors"
	"net"
)

var IPAddrNotFound = errors.New("ip address is not found in interface")

func GetLocalIpAddr() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		Log().Error("Getting net.InterfaceAddrs() is failed, %s.", err)
		return "", err
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", IPAddrNotFound
}
