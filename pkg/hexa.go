package pkg

import (
	"crypto/rand"
	"math/big"
)


func GenerateHexKey() string {
	const charset = "0123456789abcdef"
	//  generate 1,048,576 unique keys as maximal.
	const keyLength = 5

	var shortKey []byte
	for i := 0; i < keyLength; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		shortKey = append(shortKey, charset[num.Int64()])
	}
	return string(shortKey)
}


