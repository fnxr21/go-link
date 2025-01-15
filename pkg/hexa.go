package pkg

import (
	"crypto/rand"
	"math/big"
)

func GenerateHexKey() string {
	// only hexadecimal characters (0-9, a-f)
	// Define the length of the key (5)
	// with thants gonna generate 1,048,576 unique keys as maximal.
	
	const charset = "0123456789abcdef"
	const keyLength = 5

	var shortKey []byte
	for i := 0; i < keyLength; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		
		shortKey = append(shortKey, charset[num.Int64()])
	}
	return string(shortKey)
}
