package illumiocore

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"hash/crc32"
	"io"
)

// convenience method to convert int literals into int pointer
func intPointer(number int) *int { return &number }

// encrypt using AES-GCM algorithm
func aesGcmEncrypt(key, plaintext string) (string, string, error) {
	k, err := hex.DecodeString(key)
	if err != nil {
		return "", "", errors.New("could not decode AES GCM key")
	}
	text := []byte(plaintext)

	block, err := aes.NewCipher(k)
	if err != nil {
		return "", "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", "", err
	}
	nonce := make([]byte, aesgcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", "", errors.New("could not generate nonce")
	}

	// nonce is added as prefix so encrypted string stores the nonce along with ciphertext
	ciphertext := aesgcm.Seal(nil, nonce, text, nil)
	return fmt.Sprintf("%x", ciphertext), fmt.Sprintf("%x", nonce), nil
}

func hashcode(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}
