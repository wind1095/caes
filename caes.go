package caes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func Chipper_key(key string) error {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		//fmt.Println(err)
		return err
	}
	return nil
}
func Encrypt(b cipher.Block, plaintext []byte) []byte {
	if mod := len(plaintext) % aes.BlockSize; mod != 0 {
		padding := make([]byte, aes.BlockSize-mod)
		plaintext = append(plaintext, padding...)
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Println(err)
		return nil
	}
	mode := cipher.NewCBCEncrypter(b, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext
}

func Decrypt(b cipher.Block, ciphertext []byte) []byte {
	if len(ciphertext)%aes.BlockSize != 0 {
		fmt.Println("암호화된 데이터의 길이는 블록 크기의 배수가 되어야 합니다.")
		return nil
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	plaintext := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(b, iv)

	mode.CryptBlocks(plaintext, ciphertext)

	return plaintext
}
