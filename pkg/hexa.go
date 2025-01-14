package pkg

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
)


func GenerateHexKey() string {
	const charset = "0123456789abcdef"
	const keyLength = 5

	var shortKey []byte
	for i := 0; i < keyLength; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		shortKey = append(shortKey, charset[num.Int64()])
	}
	return string(shortKey)
}


func DecodeHexKey(hexKey string) ([]byte, error) {
	decoded, err := hex.DecodeString(hexKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex key: %v", err)
	}
	return decoded, nil
}