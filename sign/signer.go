package sign

import "errors"

type Signer interface {
	Encrypt(plainText []byte) (cipherText []byte, err error)
	Verify(plainText []byte, signature []byte) error
}

var (
	ErrSignVerification = errors.New("signature verification error")
)
