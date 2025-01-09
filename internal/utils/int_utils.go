package utils

import (
	"crypto/rand"
	"math/big"
)

func GetRandomInt(min, max int64) int64 {
	n, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))
	return n.Int64() + min
}
