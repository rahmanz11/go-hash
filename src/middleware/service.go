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
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/sha3"
)

type Algo string

const (
	Sha512     Algo = "sha512"
	Sha256     Algo = "sha256"
	Md5        Algo = "md5"
	Xxhash     Algo = "xxhash"
	Whirlpool  Algo = "whirlpool"
	Sha3256    Algo = "sha3256"
	Sha3384    Algo = "sha3384"
	Sha3512    Algo = "sha3512"
	Blake2b384 Algo = "blake2b384"
	Blake2b512 Algo = "blake2b512"
)

func CreateHash(data []string, algo Algo) string {
	str := strings.Join(data[:], "")
	hash := ""
	switch algo {
	case Sha512:
		hash = getSha512(str)
	case Sha256:
		hash = getSha256(str)
	case Md5:
		hash = getMd5(str)
	case Xxhash:
		hash = getXxhash(str)
	case Whirlpool:
		hash = getWhirlpool(str)
	case Sha3256:
		hash = getSha3256(str)
	case Sha3384:
		hash = getSha3384(str)
	case Sha3512:
		hash = getSha3512(str)
	case Blake2b384:
		hash = getBlake2b384(str)
	case Blake2b512:
		hash = getBlake2b512(str)
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

func getSha3256(text string) string {
	h := sha3.New512()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}

func getSha3384(text string) string {
	h := sha3.New384()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}

func getSha3512(text string) string {
	h := sha3.New512()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}

func getBlake2b384(text string) string {
	h := blake2b.Sum384([]byte(text))
	return hex.EncodeToString(h[:])
}

func getBlake2b512(text string) string {
	h := blake2b.Sum512([]byte(text))
	return hex.EncodeToString(h[:])
}
