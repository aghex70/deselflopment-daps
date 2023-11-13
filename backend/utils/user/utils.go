package utils

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"github.com/aghex70/daps/internal/pkg"
	domain2 "github.com/aghex70/daps/internal/ports/domain"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type MyCustomClaims struct {
	UserID uint `json:"user_id"`
	Admin  bool `json:"admin"`
	jwt.RegisteredClaims
}

func EncryptPassword(ctx context.Context, password string) string {
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

func DecryptPassword(ctx context.Context, encryptedPassword string) (string, error) {
	// Convert the ciphertext from a hexadecimal string to a byte slice
	encrypted, err := hex.DecodeString(encryptedPassword)
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
	err = bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password))
	if err != nil {
		// Passwords do not match
		return "", err
	}

	return password, nil
}

func PasswordsMatch(ctx context.Context, hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func GenerateJWT(ctx context.Context, u domain2.User) (string, int, error) {
	claims := NewCustomClaims(ctx, u)
	mySigningKey := pkg.HmacSampleSecret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", 0, err
	}

	return ss, int(u.ID), nil
}

func PasswordMatchesRepeatPassword(ctx context.Context, password, repeatPassword string) bool {
	return password == repeatPassword
}

func NewCustomClaims(ctx context.Context, u domain2.User) MyCustomClaims {
	return MyCustomClaims{
		UserID: u.ID,
		Admin:  u.Admin,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   u.Email,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(96 * time.Hour)),
		},
	}
}
