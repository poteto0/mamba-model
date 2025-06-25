package utils

import (
	"crypto/rand"
	"math"
	"math/big"
)

func Logistic(x float64, offset float64, l float64, k float64) float64 {
	return l / (1 + math.Exp(-k*(x-offset)))
}

func BradleyTerry(sum float64, val float64, ot float64) float64 {
	return val / (ot + sum)
}

func RandFloat64() float64 {
	num, _ := rand.Int(rand.Reader, big.NewInt(1000000000))
	return float64(num.Int64()) / 1000000000.0
}
