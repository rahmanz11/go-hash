package middleware

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/rahmanz11/go-hash/src/payloads"
)

func CreateHash(arr *payloads.Data, algo string) string {
	fmt.Printf("%s", arr)
	concatenated := strings.Join(arr.Value[:], "")
	fmt.Println(concatenated)
	hash := ""
	switch algo {
	case "sha-256":
		hash = generateSha256(concatenated)
	case "md5":
		fmt.Println("Linux.")
	default:
		hash = generateSha256(concatenated)
	}

	return hash
}

func generateSha256(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	hh := h.Sum(nil)
	fmt.Printf("%x\n", hh)
	return base64.URLEncoding.EncodeToString(hh)
}
