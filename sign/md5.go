package sign

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
)

type Append uint8

const (
	AppendLeft = iota
	AppendRight
)

func NewMd5Signer(accessKey []byte, append Append) *MD5Signer {
	return &MD5Signer{
		accessKey: accessKey,
		append:    append,
	}
}

type MD5Signer struct {
	accessKey []byte
	append    Append
}

func (m *MD5Signer) Encrypt(plainText []byte) ([]byte, error) {
	data := plainText
	if m.append == AppendLeft {
		data = append(m.accessKey, plainText...)
	} else {
		data = append(plainText, m.accessKey...)
	}

	hash := md5.Sum(data)
	cipher := make([]byte, hex.EncodedLen(len(hash)))

	n := hex.Encode(cipher, hash[:])
	return cipher[:n], nil
}

func (m *MD5Signer) Verify(plainText []byte, signature []byte) error {
	cipherText, err := m.Encrypt(plainText)
	if err != nil {
		return err
	}
	if !bytes.Equal(cipherText, signature) {
		return ErrSignVerification
	}
	return nil
}
