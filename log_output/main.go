package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length := 30

	b := make([]byte, length)

	for i := range b {
		b[i] = charset[rand.N(len(charset))]
	}

	// Convert the byte slice to a string and print it
	timestamp := time.Now().Format(time.RFC3339)
	randomString := timestamp + "-" + string(b)
	fmt.Println(randomString)

	for {
		time.Sleep(5 * time.Second)
		timestamp = time.Now().Format(time.RFC3339)
		randomString := timestamp + "-" + string(b)
		fmt.Println(randomString)
	}
}
