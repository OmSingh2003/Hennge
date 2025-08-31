package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base32"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"hash"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	email := "singhom2003.os@gmail.com"
	sharedSecret := email + "HENNGECHALLENGE004"
	// Base32 encode the secret in uppercase without padding
	base32Secret := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString([]byte(sharedSecret))
	base32Secret = strings.ToUpper(base32Secret)

	// Generate 10-digit TOTP code
	totp := generateTOTP(base32Secret, 10, 30, sha512.New)
	fmt.Printf("Generated TOTP: %s\n", totp)

	payload := map[string]string{
		"github_url":        "https://gist.github.com/OmSingh2003/db0f9b6d14c628b5dbcc484d5011c20c",
		"contact_email":     "singhom2003.os@gmail.com",
		"solution_language": "golang",
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	req, err := http.NewRequest("POST", "https://api.challenge.hennge.com/challenges/backend-recursion/004", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		os.Exit(1)
	}

	req.Header.Set("Content-Type", "application/json")

	// Basic Auth header: base64 encode "email:totp"
	auth := base64.StdEncoding.EncodeToString([]byte(email + ":" + totp))
	req.Header.Set("Authorization", "Basic "+auth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)
}

func generateTOTP(secret string, digits int, interval uint64, hashFunc func() hash.Hash) string {
	key, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(secret)
	if err != nil {
		panic(err)
	}

	counter := uint64(time.Now().Unix() / int64(interval))
	var bCounter [8]byte
	binary.BigEndian.PutUint64(bCounter[:], counter)

	hmacHash := hmac.New(hashFunc, key)
	hmacHash.Write(bCounter[:])
	hash := hmacHash.Sum(nil)

	offset := hash[len(hash)-1] & 0x0F
	binCode := (int(hash[offset])&0x7F)<<24 |
		(int(hash[offset+1])&0xFF)<<16 |
		(int(hash[offset+2])&0xFF)<<8 |
		(int(hash[offset+3]) & 0xFF)

	modulo := 1
	for i := 0; i < digits; i++ {
		modulo *= 10
	}

	otp := binCode % modulo
	return fmt.Sprintf("%0*d", digits, otp)
}
