package pkg

// import (
// 	"math/rand"
// 	"time"
// )

// // generate random string
// func GenerateShortKey() string {
// 	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
// 	const keyLength = 6

// 	rand.Seed(time.Now().UnixNano())
// 	shortKey := make([]byte, keyLength)
// 	for i := range shortKey {
// 		shortKey[i] = charset[rand.Intn(len(charset))]
// 	}
// 	return string(shortKey)
// }
