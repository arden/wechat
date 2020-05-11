package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// AESCBCEncrypt AES CBC encrypt with PKCS#7 padding
func AESCBCEncrypt(plainText, key []byte, iv ...byte) ([]byte, error) {
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	plainText = PKCS7Padding(plainText, len(key))

	cipherText := make([]byte, len(plainText))

	if len(iv) == 0 {
		iv = key[:block.BlockSize()]
	}

	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(cipherText, plainText)

	return cipherText, nil
}

// AESCBCDecrypt AES CBC decrypt with PKCS#7 unpadding
func AESCBCDecrypt(cipherText, key []byte, iv ...byte) ([]byte, error) {
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	plainText := make([]byte, len(cipherText))

	if len(iv) == 0 {
		iv = key[:block.BlockSize()]
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	blockMode.CryptBlocks(plainText, cipherText)

	return PKCS7UnPadding(plainText, len(key)), nil
}

// PKCS7Padding PKCS#7 padding
func PKCS7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize

	if padding == 0 {
		padding = blockSize
	}

	padText := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(cipherText, padText...)
}

// PKCS7UnPadding PKCS#7 unpadding
func PKCS7UnPadding(plainText []byte, blockSize int) []byte {
	l := len(plainText)
	unpadding := int(plainText[l-1])

	if unpadding < 0 || unpadding > blockSize {
		unpadding = 0
	}

	return plainText[:(l - unpadding)]
}

// RSAEncrypt rsa encrypt with public key
func RSAEncrypt(data, publicKey []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)

	if block == nil {
		return nil, errors.New("invalid rsa public key")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)

	if err != nil {
		return nil, err
	}

	key, ok := pubKey.(*rsa.PublicKey)

	if !ok {
		return nil, errors.New("invalid rsa public key")
	}

	return rsa.EncryptPKCS1v15(rand.Reader, key, data)
}

// RSADecrypt rsa decrypt with private key
func RSADecrypt(cipherText, privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)

	if block == nil {
		return nil, errors.New("invalid rsa private key")
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	if err != nil {
		return nil, err
	}

	return rsa.DecryptPKCS1v15(rand.Reader, key, cipherText)
}
