package main

import (
	"fmt"
	"net"
)

/*
List all IP address in given range in golang
Get all IP address in given network in goglang
https://go.dev/play/p/O9yL-jcXpJD
*/

func Hosts(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	// remove network address and broadcast address
	return ips[1 : len(ips)-1], nil
}

// http://play.golang.org/p/m8TNTtygK0
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func main() {
	hosts, _ := Hosts("192.168.11.0/24")

	for i, ip := range hosts {

		fmt.Printf("%d: %s\n", i+1, ip)
	}

}
