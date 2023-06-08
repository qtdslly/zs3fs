package internal

import (
	"github.com/zehuamama/balancer/balancer"
	"net"
	"strings"
)

func Next(lb balancer.Balancer, client string) string {
	targetHost, err := lb.Balance(client)
	if err != nil {
		return ""
	}

	lb.Inc(targetHost)
	defer lb.Done(targetHost)
	return targetHost
}

func GetIps() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	var ips []string
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}

	var result string
	for _, ip := range ips {
		result = result + "," + ip
	}

	return strings.Trim(string(result), ","), nil
}
