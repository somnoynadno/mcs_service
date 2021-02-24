package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"unicode/utf8"
)

func AES256(plaintext string, key string, iv string, blockSize int) (string, error) {
	if utf8.ValidString(plaintext) == false {
		return "", errors.New("plaintext encoding is not valid")
	}

	bKey := []byte(key)
	bIV := []byte(iv)
	bPlaintext := PKCS5Padding([]byte(plaintext), blockSize, len(plaintext))
	block, _ := aes.NewCipher(bKey)
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)
	return hex.EncodeToString(ciphertext), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	pad := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, pad...)
}

func PadKey(key string, l int, padding string) string {
	padNum := l - len(key)
	for i := 0; i < padNum; i++ {
		key += padding
	}
	return key
}