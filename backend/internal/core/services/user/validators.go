package user

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"os"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s Service) CheckExistentUser(ctx context.Context, email string) bool {
	_, err := s.userRepository.GetByEmail(ctx, email)
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

func (s Service) PasswordsMatch(ctx context.Context, hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (s Service) PasswordMatchesRepeatPassword(ctx context.Context, password, repeatPassword string) bool {
	return password == repeatPassword
}

func (s Service) EncryptPassword(ctx context.Context, password string) string {
	keyString := os.Getenv("CIPHER_KEY")
	key, _ := hex.DecodeString(keyString)

	// Create a new AES cipher block
	block, _ := aes.NewCipher(key)

	ivString := os.Getenv("CIPHER_IV")
	iv, _ := hex.DecodeString(ivString)

	// Create a new CBC cipher mode using the AES cipher and the
	// initialization vector
	mode := cipher.NewCBCEncrypter(block, iv)

	// Convert the password to a byte slice and pad it so that its length is
	// a multiple of the block size
	input := []byte(password)
	padding := aes.BlockSize - len(input)%aes.BlockSize
	input = append(input, make([]byte, padding)...)

	// Encrypt the input using the CBC cipher mode
	encrypted := make([]byte, len(input))
	mode.CryptBlocks(encrypted, input)

	// Print the encrypted input as a hexadecimal string
	return hex.EncodeToString(encrypted)
}

func (s Service) DecryptPassword(ctx context.Context, cipheredPassword string) (string, error) {
	// Convert the ciphertext from a hexadecimal string to a byte slice
	encrypted, err := hex.DecodeString(cipheredPassword)
	if err != nil {
		return "", err
	}

	keyString := os.Getenv("CIPHER_KEY")
	key, _ := hex.DecodeString(keyString)

	// Create a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ivString := os.Getenv("CIPHER_IV")
	iv, _ := hex.DecodeString(ivString)

	// Create a new CBC cipher mode using the AES cipher and the
	// initialization vector
	mode := cipher.NewCBCDecrypter(block, iv)

	// Decrypt the ciphertext using the CBC cipher mode
	decrypted := make([]byte, len(encrypted))
	mode.CryptBlocks(decrypted, encrypted)

	// Remove the padding from the decrypted password
	padding := decrypted[len(decrypted)-1]
	decrypted = decrypted[:len(decrypted)-int(padding)]

	password := string(decrypted)

	// Now you can use bcrypt to compare the decrypted password with a hashed password
	err = bcrypt.CompareHashAndPassword([]byte(cipheredPassword), []byte(password))
	if err != nil {
		// Passwords do not match
		return "", err
	}

	return password, nil
}
