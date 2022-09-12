package middleware

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/cespare/xxhash"
	"github.com/rahmanz11/go-hash/src/payloads"
)

func CreateHash(arr *payloads.Data, algo string) string {
	fmt.Printf("%s", arr)
	concatenated := strings.Join(arr.Value[:], "")
	fmt.Println(concatenated)
	hash := ""
	switch algo {
	case "sha256":
		hash = getSha256(concatenated)
	case "md5":
		hash = getMd5(concatenated)
	case "xxhash":
		hash = getXxhash(concatenated)
	default:
		hash = getSha512(concatenated)
	}

	return hash
}

func getSha256(text string) string {
	h := sha256.New()
	h.Write([]byte(text))
	hh := h.Sum(nil)
	return base64.URLEncoding.EncodeToString(hh)
}

func getMd5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func getXxhash(text string) string {
	hash := xxhash.Sum64String(text)
	return strconv.FormatUint(hash, 10)
}

func getSha512(text string) string {
	h := sha512.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}
