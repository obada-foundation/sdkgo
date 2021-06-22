package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

// Hash struct represent string value as a hash or as a decimal
type Hash struct {
	hash string
	dec  uint64
}

// NewHash creates a new OBADA hash
func NewHash(value string, log *log.Logger, debug bool) (Hash, error) {
	var hash Hash
	var debugStr string

	h := sha256.New()

	if _, err := h.Write([]byte(value)); err != nil {
		return hash, fmt.Errorf("cannot wite bytes %v to hasher: %w", []byte(value), err)
	}

	hashStr := hex.EncodeToString(h.Sum(nil))

	if debug {
		log.Printf("SHA256(%q) -> %q", value, hashStr)
	}

	hashDec, err := hashToDec(hashStr, log, debug)

	if err != nil {
		log.Println(debugStr)
		return hash, err
	}

	hash.hash = hashStr
	hash.dec = hashDec

	return hash, nil
}

// hashToDec convert hash which is hex string into decimal
func hashToDec(hash string, log *log.Logger, debug bool) (uint64, error) {
	match, err := regexp.MatchString(`^[0-9a-fA-F]+$`, hash)
	partialHash := hash

	if err != nil {
		return 0, fmt.Errorf("cannot check if given string %q is valid hex: %w", hash, err)
	}

	if !match {
		return 0, fmt.Errorf("given string string %q is not valid hex", hash)
	}

	if len(hash) > 8 {
		partialHash = hash[:8]
	}

	decimal, err := strconv.ParseUint(partialHash, 16, 32)

	if err != nil {
		return 0, err
	}

	if debug {
		log.Printf("Get8CharsFromHash(%q) -> %q -> Hex2Dec(%q) -> %d", hash, partialHash, partialHash, decimal)
	}

	return decimal, nil
}

// GetHash returns hash
func (h Hash) GetHash() string {
	return h.hash
}

// GetDec returns decimal value
func (h Hash) GetDec() uint64 {
	return h.dec
}
