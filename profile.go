package PassGenerator

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

// must be 32 characters
var key = "sfjkloiutynzsmklmaqiebcgfplvbsdu"

var ErrMalformedEncryption = errors.New("malformed encryption")

// Enc will hold the encrypted password
// Platform stores the service for which the password is generated
// password will hold the actual generated password, it is in small letters so that it is not stored
type profile struct {
	Enc, Platform, password string
}

func (p *profile) encrypt() error {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	enc := gcm.Seal(nonce, nonce, []byte(p.password), nil)
	p.Enc = hex.EncodeToString(enc)

	return nil
}

func (p *profile) decrypt() error {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonceSize := gcm.NonceSize()
	if len(p.Enc) < nonceSize {
		return ErrMalformedEncryption
	}

	enc, err := hex.DecodeString(p.Enc)
	if err != nil {
		return err
	}

	password, err := gcm.Open(nil, enc[:nonceSize], enc[nonceSize:], nil)
	if err != nil {
		return err
	}

	p.password = string(password)

	return nil
}
