package middleware

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"strconv"
	"strings"

	"github.com/cespare/xxhash"
	"github.com/jzelinskie/whirlpool"
)

func CreateHash(data []string, algo string) string {
	str := strings.Join(data[:], "")
	hash := ""
	switch algo {
	case "sha512":
		hash = getSha512(str)
	case "sha256":
		hash = getSha256(str)
	case "md5":
		hash = getMd5(str)
	case "xxhash":
		hash = getXxhash(str)
	case "whirlpool":
		hash = getWhirlpool(str)
	case "sha3256":
		hash = getWhirlpool(str)
	default:
		hash = getSha512(str)
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

func getWhirlpool(text string) string {
	w := whirlpool.New()
	w.Write([]byte(text))
	return hex.EncodeToString(w.Sum(nil))
}
