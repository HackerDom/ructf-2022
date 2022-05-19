package crypto

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

const BlockSize = 16
const Rounds = 13

var S = [256]byte{120, 112, 116, 124, 232, 224, 228, 236, 121, 113, 117, 125, 233, 225, 229, 237, 57, 49, 53, 61, 169, 161, 165, 173, 56, 48, 52, 60, 168, 160, 164, 172, 216, 208, 212, 220, 72, 64, 68, 76, 217, 209, 213, 221, 73, 65, 69, 77, 153, 145, 149, 157, 9, 1, 5, 13, 152, 144, 148, 156, 8, 0, 4, 12, 88, 80, 84, 92, 200, 192, 196, 204, 89, 81, 85, 93, 201, 193, 197, 205, 25, 17, 21, 29, 137, 129, 133, 141, 24, 16, 20, 28, 136, 128, 132, 140, 248, 240, 244, 252, 104, 96, 100, 108, 249, 241, 245, 253, 105, 97, 101, 109, 185, 177, 181, 189, 41, 33, 37, 45, 184, 176, 180, 188, 40, 32, 36, 44, 122, 114, 118, 126, 234, 226, 230, 238, 123, 115, 119, 127, 235, 227, 231, 239, 59, 51, 55, 63, 171, 163, 167, 175, 58, 50, 54, 62, 170, 162, 166, 174, 218, 210, 214, 222, 74, 66, 70, 78, 219, 211, 215, 223, 75, 67, 71, 79, 155, 147, 151, 159, 11, 3, 7, 15, 154, 146, 150, 158, 10, 2, 6, 14, 90, 82, 86, 94, 202, 194, 198, 206, 91, 83, 87, 95, 203, 195, 199, 207, 27, 19, 23, 31, 139, 131, 135, 143, 26, 18, 22, 30, 138, 130, 134, 142, 250, 242, 246, 254, 106, 98, 102, 110, 251, 243, 247, 255, 107, 99, 103, 111, 187, 179, 183, 191, 43, 35, 39, 47, 186, 178, 182, 190, 42, 34, 38, 46}
var SInv = [256]byte{61, 53, 189, 181, 62, 54, 190, 182, 60, 52, 188, 180, 63, 55, 191, 183, 89, 81, 217, 209, 90, 82, 218, 210, 88, 80, 216, 208, 91, 83, 219, 211, 125, 117, 253, 245, 126, 118, 254, 246, 124, 116, 252, 244, 127, 119, 255, 247, 25, 17, 153, 145, 26, 18, 154, 146, 24, 16, 152, 144, 27, 19, 155, 147, 37, 45, 165, 173, 38, 46, 166, 174, 36, 44, 164, 172, 39, 47, 167, 175, 65, 73, 193, 201, 66, 74, 194, 202, 64, 72, 192, 200, 67, 75, 195, 203, 101, 109, 229, 237, 102, 110, 230, 238, 100, 108, 228, 236, 103, 111, 231, 239, 1, 9, 129, 137, 2, 10, 130, 138, 0, 8, 128, 136, 3, 11, 131, 139, 93, 85, 221, 213, 94, 86, 222, 214, 92, 84, 220, 212, 95, 87, 223, 215, 57, 49, 185, 177, 58, 50, 186, 178, 56, 48, 184, 176, 59, 51, 187, 179, 29, 21, 157, 149, 30, 22, 158, 150, 28, 20, 156, 148, 31, 23, 159, 151, 121, 113, 249, 241, 122, 114, 250, 242, 120, 112, 248, 240, 123, 115, 251, 243, 69, 77, 197, 205, 70, 78, 198, 206, 68, 76, 196, 204, 71, 79, 199, 207, 33, 41, 161, 169, 34, 42, 162, 170, 32, 40, 160, 168, 35, 43, 163, 171, 5, 13, 133, 141, 6, 14, 134, 142, 4, 12, 132, 140, 7, 15, 135, 143, 97, 105, 225, 233, 98, 106, 226, 234, 96, 104, 224, 232, 99, 107, 227, 235}
var P = [16]byte{12, 14, 13, 15, 0, 2, 1, 3, 4, 6, 5, 7, 8, 10, 9, 11}
var PInv = [16]byte{4, 6, 5, 7, 8, 10, 9, 11, 12, 14, 13, 15, 0, 2, 1, 3}

func pad(data []byte) []byte {
	paddingLength := BlockSize - len(data)%BlockSize
	paddingBytes := bytes.Repeat([]byte{byte(paddingLength)}, paddingLength)
	return append(data, paddingBytes...)
}

func unpad(data []byte) []byte {
	return data[:len(data)-int(data[len(data)-1])]
}

func substitute(data [BlockSize]byte, s [256]byte) [BlockSize]byte {
	var result [BlockSize]byte
	for i := 0; i < len(data); i++ {
		result[i] = s[data[i]]
	}
	return result
}

func permutate(data [BlockSize]byte, p [BlockSize]byte) [BlockSize]byte {
	var result [BlockSize]byte
	for i := 0; i < len(data); i++ {
		result[i] = data[p[i]]
	}
	return result
}

func expandKey(key [BlockSize]byte) [BlockSize]byte {
	keyToBlocks := func(key [BlockSize]byte) [4][4]byte {
		var blocks [4][4]byte
		for i := range blocks {
			copy(blocks[i][:], key[i*4:(i+1)*4])
		}
		return blocks
	}

	shiftBlocks := func(blocks [4][4]byte) [4][4]byte {
		var newBlocks [4][4]byte
		for i := 0; i < 3; i++ {
			copy(newBlocks[i][:], blocks[i+1][:])
		}
		for i, x := range blocks[0] {
			newBlocks[3][i] = S[x]
		}
		return newBlocks
	}

	blocksToKey := func(blocks [4][4]byte) [BlockSize]byte {
		var key [BlockSize]byte
		for i, x := range blocks {
			copy(key[i*4:(i+1)*4], x[:])
		}
		return key
	}

	return blocksToKey(shiftBlocks(keyToBlocks(key)))
}

func generateRoundKeys(key [BlockSize]byte) [Rounds][BlockSize]byte {
	var keys [Rounds][BlockSize]byte
	lastKey := key
	for i := 0; i < Rounds; i++ {
		keys[i] = expandKey(lastKey)
		lastKey = keys[i]
	}
	return keys
}

func splitToBlocks(data []byte) ([][BlockSize]byte, error) {
	if len(data)%BlockSize != 0 {
		return nil, fmt.Errorf("[-] Bad data size: %d", len(data))
	}
	numberOfBlocks := len(data) / BlockSize
	dataBlocks := make([][BlockSize]byte, numberOfBlocks)
	for i := 0; i < numberOfBlocks; i++ {
		copy(dataBlocks[i][:], data[i*BlockSize:(i+1)*BlockSize])
	}
	return dataBlocks, nil
}

func joinBlocks(blocks [][BlockSize]byte) []byte {
	data := make([]byte, BlockSize*len(blocks))
	for i, block := range blocks {
		copy(data[i*BlockSize:(i+1)*BlockSize], block[:])
	}
	return data
}

type Cipher struct {
	key [BlockSize]byte
}

type CipherInterface interface {
	Encrypt(data []byte)
	Decrypt(data []byte)
}

func xorBlock(a, b [BlockSize]byte) [BlockSize]byte {
	var result [BlockSize]byte
	for i := 0; i < len(a); i++ {
		result[i] = a[i] ^ b[i]
	}
	return result
}

func encryptRound(block [BlockSize]byte, roundKey [BlockSize]byte) [BlockSize]byte {
	var result [BlockSize]byte
	result = substitute(block, S)
	result = permutate(result, P)
	result = xorBlock(result, roundKey)
	return result
}

func decryptRound(block [BlockSize]byte, roundKey [BlockSize]byte) [BlockSize]byte {
	var result [BlockSize]byte
	result = xorBlock(block, roundKey)
	result = permutate(result, PInv)
	result = substitute(result, SInv)
	return result
}

func (c Cipher) Encrypt(pt []byte) ([]byte, error) {
	ptBlocks, err := splitToBlocks(pad(pt))
	if err != nil {
		return nil, err
	}
	ctBlocks := make([][BlockSize]byte, len(ptBlocks))
	roundKeys := generateRoundKeys(c.key)
	for i, block := range ptBlocks {
		for round := 0; round < len(roundKeys); round++ {
			block = encryptRound(block, roundKeys[round])
		}
		ctBlocks[i] = block
	}
	return joinBlocks(ctBlocks), nil
}

func (c Cipher) Decrypt(ct []byte) ([]byte, error) {
	ctBlocks, err := splitToBlocks(ct)
	if err != nil {
		return nil, err
	}
	ptBlocks := make([][BlockSize]byte, len(ctBlocks))
	roundKeys := generateRoundKeys(c.key)
	for i, block := range ctBlocks {
		for round := len(roundKeys) - 1; round >= 0; round-- {
			block = decryptRound(block, roundKeys[round])
		}
		ptBlocks[i] = block
	}
	return unpad(joinBlocks(ptBlocks)), nil
}

func NewCipher(key []byte) Cipher {
	var goodKey [BlockSize]byte
	copy(goodKey[:], key)
	return Cipher{key: goodKey}
}

func main() {
	pt := []byte("Hello there! Nice to meet you!")
	fmt.Printf("PT: %s\n", hex.EncodeToString(pt))
	var key [BlockSize]byte
	for i := range key {
		key[i] = byte(i)
	}
	c := Cipher{key: key}
	ct, _ := c.Encrypt(pt)
	fmt.Printf("CT: %s\n", hex.EncodeToString(ct))
	maybePt, _ := c.Decrypt(ct)
	fmt.Printf("PT? %s\n", hex.EncodeToString(maybePt))
	fmt.Println(bytes.Compare(pt, maybePt) == 0)
}
