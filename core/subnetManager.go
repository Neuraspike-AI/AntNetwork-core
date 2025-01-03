package main

import (
	"errors"
	"strings"
)

const (
	base58 string = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
)

func assignSubnet(address string) (int, error) {
	subnetID := string([]rune(address)[0])
	if strings.Contains(base58, subnetID) {
		return strings.Index(base58, subnetID) + 1, nil
	} else {
		return 0, errors.New("invalid address")
	}
}
