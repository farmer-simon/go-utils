package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

type rsaEncrypt struct {
	public []byte
	private []byte
}

// New
func New(public, private []byte) *rsaEncrypt {
	return &rsaEncrypt{
		public:  public,
		private: private,
	}
}

// Encrypt
func (e *rsaEncrypt) Encrypt(origData []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(e.public)
	if block == nil {
		return nil, fmt.Errorf("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func (e *rsaEncrypt) Decrypt(cipherText []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(e.private)
	if block == nil {
		return nil, fmt.Errorf("private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv.(*rsa.PrivateKey), cipherText)
}
