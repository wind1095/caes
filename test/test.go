package main

import (
	"fmt"
	"github.com/wind1095/caes"
)

func main() {
	key := "CATCHALL20221228" //16바이트
	s := `icatch2016`

	err := caes.Chipper_key(key)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	ciphertext := caes.Encrypt(block, []byte(s))
	fmt.Printf("%x\n", ciphertext)

	plaintext := caes.Decrypt(block, ciphertext)
	fmt.Println(string(plaintext))
}
