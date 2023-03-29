package main

/*
This code generates list of subnets in a given Network using golang
https://go.dev/play/p/qltb77htF5u
Many thanks to Neo Anderson for solution

*/

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"net"
	"strconv"
)

func GenSubnetsInNetwork(netCIDR string, subnetMaskSize int) ([]string, error) {
	ip, ipNet, err := net.ParseCIDR(netCIDR)
	if err != nil {
		return nil, err
	}
	if !ip.Equal(ipNet.IP) {
		return nil, errors.New("netCIDR is not a valid network address")
	}
	netMaskSize, _ := ipNet.Mask.Size()
	if netMaskSize > int(subnetMaskSize) {
		return nil, errors.New("subnetMaskSize must be greater or equal than netMaskSize")
	}

	totalSubnetsInNetwork := math.Pow(2, float64(subnetMaskSize)-float64(netMaskSize))
	totalHostsInSubnet := math.Pow(2, 32-float64(subnetMaskSize))
	subnetIntAddresses := make([]uint32, int(totalSubnetsInNetwork))
	// first subnet address is same as the network address
	subnetIntAddresses[0] = ip2int(ip.To4())
	for i := 1; i < int(totalSubnetsInNetwork); i++ {
		subnetIntAddresses[i] = subnetIntAddresses[i-1] + uint32(totalHostsInSubnet)
	}

	subnetCIDRs := make([]string, 0)
	for _, sia := range subnetIntAddresses {
		subnetCIDRs = append(
			subnetCIDRs,
			int2ip(sia).String()+"/"+strconv.Itoa(int(subnetMaskSize)),
		)
	}
	return subnetCIDRs, nil
}

func ip2int(ip net.IP) uint32 {
	if len(ip) == 16 {
		panic("cannot convert IPv6 into uint32")
	}
	return binary.BigEndian.Uint32(ip)
}
func int2ip(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}
func main() {
	Runner()
}

func Runner() {
	a, _ := GenSubnetsInNetwork("192.168.0.0/16", 22)
	fmt.Println(len(a))
	for _, j := range a {
		fmt.Println(j)
	}
}
