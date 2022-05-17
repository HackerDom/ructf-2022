package main

import (
	"encoding/binary"
	"fmt"
)

const BlockLength = 16

func performXor(data, result, key []byte) {
	for i := 0; i < len(data); i++ {
		result[i] = data[i] ^ key[i%len(key)]
	}
}

func Encrypt(data, key []byte) []byte {
	resultLen := (len(data) / BlockLength) + BlockLength
	result := make([]byte, resultLen)

	performXor(data, result, key)

	ln := make([]byte, 4)
	binary.LittleEndian.PutUint32(ln, uint32(len(data)))

	return append(result, ln...)
}

func Decrypt(data, key []byte) []byte {
	ln := binary.LittleEndian.Uint32(data[len(data)-4:])

	result := make([]byte, len(data)-4)

	performXor(data[:len(data)-4], result, key)

	return result[:ln]
}

func main() {
	data := []byte("Hello world!")
	key := []byte("key")
	mid := Encrypt(data, key)
	res := Decrypt(mid, key)
	fmt.Println(data)
	fmt.Println(mid)
	fmt.Println(res)
}
