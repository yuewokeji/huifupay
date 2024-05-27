package sign

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

func NewRSASigner(publicKey, privateKey string) *RSASigner {
	r := &RSASigner{
		pubKey: publicKey,
		priKey: privateKey,
	}
	return r
}

type RSASigner struct {
	pubKey string
	priKey string

	rsaPubKey *rsa.PublicKey
	rsaPriKey *rsa.PrivateKey
}

// Encrypt 签名
func (s *RSASigner) Encrypt(plainText []byte) (cipher []byte, err error) {
	var key *rsa.PrivateKey
	key, err = s.getRSAPrivateKey()
	if err != nil {
		return
	}

	h := sha256.New()
	h.Write(plainText)
	hash := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, hash[:])
	if err != nil {
		return
	}

	cipher = make([]byte, base64.StdEncoding.EncodedLen(len(signature)))
	base64.StdEncoding.Encode(cipher, signature)
	return
}

// Verify 验证签名
func (s *RSASigner) Verify(plainText []byte, signature []byte) error {
	key, err := s.getRSAPublicKey()
	if err != nil {
		return err
	}

	dst := make([]byte, base64.StdEncoding.EncodedLen(len(signature)))
	n, err := base64.StdEncoding.Decode(dst, signature)
	if err != nil {
		return err
	}

	hash := sha256.New()
	hash.Write(plainText)
	bytes := hash.Sum(nil)

	err = rsa.VerifyPKCS1v15(key, crypto.SHA256, bytes, dst[:n])
	if err != nil {
		if errors.Is(rsa.ErrVerification, err) {
			return ErrSignVerification
		}
	}

	return err
}

func (s *RSASigner) getRSAPublicKey() (key *rsa.PublicKey, err error) {
	if s.rsaPubKey == nil {
		key, err = ParseRSAPublicKey(s.pubKey)
		if err != nil {
			return
		}
		s.rsaPubKey = key
	}
	return s.rsaPubKey, nil
}

func (s *RSASigner) getRSAPrivateKey() (key *rsa.PrivateKey, err error) {
	if s.rsaPriKey == nil {
		key, err = ParseRSAPrivateKey(s.priKey)
		if err != nil {
			return
		}
		s.rsaPriKey = key
	}
	return s.rsaPriKey, nil
}

const (
	rsaPublicKeyBegin = "-----BEGIN PUBLIC KEY-----"
	rsaPublicKeyEnd   = "-----END PUBLIC KEY-----"

	rsaPrivateKeyPemBegin = "-----BEGIN RSA PRIVATE KEY-----"
	rsaPrivateKeyPemEnd   = "-----END RSA PRIVATE KEY-----"
)

// ParseRSAPublicKey 读取公钥文件
func ParseRSAPublicKey(pubKey string) (*rsa.PublicKey, error) {
	if !strings.HasSuffix(pubKey, rsaPublicKeyBegin) {
		pubKey = fmt.Sprintf("%s\n%s\n%s", rsaPublicKeyBegin, pubKey, rsaPublicKeyEnd)
	}
	block, _ := pem.Decode([]byte(pubKey))
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return key.(*rsa.PublicKey), nil
}

// ParseRSAPrivateKey 读取私钥文件
func ParseRSAPrivateKey(priKey string) (*rsa.PrivateKey, error) {
	if !strings.HasSuffix(priKey, rsaPrivateKeyPemBegin) {
		priKey = fmt.Sprintf("%s\n%s\n%s", rsaPrivateKeyPemBegin, priKey, rsaPrivateKeyPemEnd)
	}
	block, _ := pem.Decode([]byte(priKey))
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return key.(*rsa.PrivateKey), nil
}
