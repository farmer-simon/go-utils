package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

type aesEncrypt struct {
	key []byte
	iv  []byte
}

func New(key []byte, iv []byte) *aesEncrypt {
	return &aesEncrypt{
		key: key,
		iv:  iv,
	}
}

// padding 填充数据
func padding(src []byte, blockSize int) []byte {
	padNum := blockSize - len(src)%blockSize
	pad := bytes.Repeat([]byte{byte(padNum)}, padNum)
	return append(src, pad...)
}

// unPadding 去掉填充数据
func unPadding(src []byte) []byte {
	n := len(src)
	unPadNum := int(src[n-1])
	return src[:n-unPadNum]
}

// Encrypt 加密
func (encrypt aesEncrypt) Encrypt(src []byte) ([]byte, error) {
	block, err := aes.NewCipher(encrypt.key)
	if err != nil {
		return nil, err
	}
	src = padding(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, encrypt.iv)
	blockMode.CryptBlocks(src, src)
	return src, nil
}

// Decrypt 解密
func (encrypt aesEncrypt) Decrypt(src []byte) ([]byte, error) {
	block, err := aes.NewCipher(encrypt.key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, encrypt.iv)
	blockMode.CryptBlocks(src, src)
	src = unPadding(src)
	return src, nil
}
