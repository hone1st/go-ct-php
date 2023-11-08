package other

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func EncryptData(data string) string {
	s := []byte("WpBiYvT9e4OI8Z0n7848w0bfdr")
	key := make([]byte, 32)
	copy(key, s)
	iv := []byte("EZhWr98dmc83Jnd4")
	cipherText, _ := Encrypt(data, key, iv)
	return base64.StdEncoding.EncodeToString([]byte(cipherText))
}

func DecryptData(data string) string {
	data1, _ := base64.StdEncoding.DecodeString(data)
	s := []byte("WpBiYvT9e4OI8Z0n7848w0bfdr")
	key := make([]byte, 32)
	copy(key, s)
	iv := []byte("EZhWr98dmc83Jnd4")
	cipherText, _ := Decrypt(string(data1), key, iv)
	return cipherText
}

func Encrypt(plainText string, key []byte, iv []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	paddedText := pkcs7Pad([]byte(plainText))
	cipherText := make([]byte, len(paddedText))

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText, paddedText)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func Decrypt(cipherText string, key []byte, iv []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipherTextBytes, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	plainText := make([]byte, len(cipherTextBytes))

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plainText, cipherTextBytes)

	unpaddedText := pkcs7Unpad(plainText)
	return string(unpaddedText), nil
}

func pkcs7Pad(data []byte) []byte {
	padding := aes.BlockSize - (len(data) % aes.BlockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func pkcs7Unpad(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
