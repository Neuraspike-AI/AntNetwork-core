package main

import (
	"crypto/sha256"
	"fmt"
)

func ToSha256(input string) string {
	h := sha256.New()

	h.Write([]byte(input))

	bs := h.Sum(nil)

	return fmt.Sprintf("%x\n", bs)
}
