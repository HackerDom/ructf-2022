package meta_unpacker

import (
	"bytes"
	"compress/zlib"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/mergermarket/go-pkcs7"
	"github.com/usernamedt/doctor-service/pkg/logging"
)

type Meta struct {
	Token    string `json:"token"`
	Question string `json:"question"`
	UserId   string `json:"userid"`
}

// Cipher key must be 32 chars long because block size is 16 bytes
const CIPHER_KEY = "01234567890123456789012345678901"

func Unpack(id string) (*Meta, error) {
	key := []byte(CIPHER_KEY)
	cipherText, err := hex.DecodeString(id)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(cipherText) < aes.BlockSize {
		return nil, fmt.Errorf("cipherText too short")
	}

	iv := []byte("0123456789012345")
	if len(cipherText)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("cipherText is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	cipherText, err = pkcs7.Unpad(cipherText, aes.BlockSize)
	if err != nil {
		return nil, err
	}
	logging.Infof("decrypted:%s\n raw:%v\n", string(cipherText), cipherText)

	ctReader := bytes.NewReader(cipherText)
	rawMeta, err := zlib.NewReader(ctReader)
	if err != nil {
		return nil, err
	}
	defer rawMeta.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(rawMeta)
	if err != nil {
		return nil, err
	}

	var meta Meta
	err = json.Unmarshal(buf.Bytes(), &meta)
	//err = json.Unmarshal(cipherText, &meta)
	if err != nil {
		return nil, err
	}

	logging.Info("unpacked:%v\n", meta)

	return &meta, nil
}
