package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

type Encryptor struct {
	key []byte
}

func NewEncryptor(key string) Encryptor {
	return Encryptor{key: []byte(key)}
}

func (e *Encryptor) Encrypt(plaintext string) (string, error) {
	block, err := aes.NewCipher(e.key)
	if err != nil {
		return "", err
	}

	bPlaintext := []byte(plaintext)

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	cipherText := make([]byte, len(bPlaintext))
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText, bPlaintext)

	result := append(iv, cipherText...)
	return base64.StdEncoding.EncodeToString(result), nil
}

func (e *Encryptor) Decrypt(encryptText string) (string, error) {
	cipherText, err :=  base64.StdEncoding.DecodeString(encryptText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(e.key)
	if err != nil{
		return "", err 
	}

	if len(cipherText) < aes.BlockSize {
		return "", fmt.Errorf("cipher text is too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)
	return string(cipherText), nil
}