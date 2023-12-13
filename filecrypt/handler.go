package filecrypt

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"

	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"

	"golang.org/x/crypto/pbkdf2"
)

const (
	numberOfBytes      = 32
	numberOfIterations = 4096
)

func Encrypt(srcFile string, password []byte) {
	// open the file then defer to close src file by io readall method
	// file, err := os.Open(srcFile)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer file.Close()

	plaintext, err := ioutil.ReadFile(srcFile)
	if err != nil {
		panic(err.Error())
	}

	// generate empty nonce then randomize it
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	// pbkdf2 new method with iterations 4096, bytes length 32, password, sha-1
	dk := pbkdf2.Key(password, nonce, numberOfIterations, numberOfBytes, sha1.New)

	// pass the derived key dk to the aes
	block, err := aes.NewCipher(dk)
	if err != nil {
		panic(err.Error())
	}

	// block from the aes new method
	// pass block into cipher new gcm method
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	// aesgcm seal method to get the cipher text
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	// append nonce to the generated cipher text
	ciphertext = append(ciphertext, nonce...)

	// os write method for writing ciphertext to a destination file
	dstFile, err := os.Create(srcFile)
	if err != nil {
		panic(err.Error())
	}
	// defer dstFile.Close()

	// if err := dstFile.Write(ciphertext); err != nil {
	// 	panic(err.Error())
	// }
	if _, err := io.Copy(dstFile, bytes.NewReader(ciphertext)); err != nil {
		panic(err.Error())
	}
}

func Decrypt(encryptedFile string, password []byte) {}
