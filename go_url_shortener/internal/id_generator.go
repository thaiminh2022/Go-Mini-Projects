package internal

import "math/rand"

const encode = "abcdefghijklmnopqrstuvwhyzABCDEFGHIJKLMNOPQRSTUVWHYZ0123456789"

func generateId(length int) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = encode[rand.Intn(length)]
	}
	return string(b)

}

func GetRandomID() string {
	return generateId(10)
}
