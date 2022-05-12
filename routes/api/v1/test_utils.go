package v1

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func CreateRandInt(maxNumber int) int {
	return rand.Intn(maxNumber)
}

func CreateRandFloat(maxNumber int) float64 {
	return float64(CreateRandInt(maxNumber)) + rand.Float64()
}

func CreateRandID(length int) string {
	var b = make([]string, length)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		b[i] = strconv.Itoa(CreateRandInt(100))
	}

	return strings.Join(b, "")
}

func CreateHash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

var (
	SYMBOLS = []string{"QQQ", "AAPL", "TSLA", "GOOG", "FB", "AMZN", "TWTR", "BA", "MS", "HD", "LOW", "WFC", "BAC", "JPM", "C", "SPY", "AAL", "GMC", "AMC"}
)

func GetRandomSymbol() string {
	index := CreateRandInt(len(SYMBOLS))

	return SYMBOLS[index]
}
