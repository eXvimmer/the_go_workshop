package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)

func encrypt(data []byte, key string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return []byte{}, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, err
	}
	var dst []byte
	nonce := make([]byte, gcm.NonceSize())
	dst = nonce
	if _, err := rand.Read(nonce); err != nil {
		return []byte{}, err
	}
	return gcm.Seal(dst, nonce, data, []byte("test")), nil
}

func decrypt(data []byte, key string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return []byte{}, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, err
	}
	ciphertext := data[gcm.NonceSize():]
	nonce := data[:gcm.NonceSize()]
	resp, err := gcm.Open(nil, nonce, ciphertext, []byte("test"))
	if err != nil {
		return resp, fmt.Errorf("error decrypting data: %v", err)
	}
	return resp, nil
}
func main() {
	const key = "mysecurepassword"
	encrypted, err := encrypt([]byte("Hello World!"), key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Encrypted Text: ", string(encrypted))
	decrypted, err := decrypt(encrypted, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Decrypted Text: ", string(decrypted))
}
