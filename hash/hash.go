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

// NewFromDID creates hash struct from DID string
func NewFromDID(hash string, logger *log.Logger) (Hash, error) {
	var h Hash

	hashDec, err := hashToDec(hash, logger)
	if err != nil {
		return h, err
	}

	h.hash = hash
	h.dec = hashDec

	return h, nil
}

// NewHash creates a new OBADA hash
func NewHash(value []byte, logger *log.Logger) (Hash, error) {
	var hash Hash
	var debugStr string

	h := sha256.New()

	if _, err := h.Write(value); err != nil {
		return hash, fmt.Errorf("cannot wite bytes %v to hasher: %w", value, err)
	}

	hashStr := hex.EncodeToString(h.Sum(nil))

	if logger != nil {
		logger.Printf("SHA256(%q) -> %q", value, hashStr)
	}

	hashDec, err := hashToDec(hashStr, logger)

	if err != nil {
		logger.Println(debugStr)
		return hash, err
	}

	hash.hash = hashStr
	hash.dec = hashDec

	return hash, nil
}

// hashToDec convert hash which is hex string into decimal
func hashToDec(hash string, logger *log.Logger) (uint64, error) {
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

	if logger != nil {
		logger.Printf("Get8CharsFromHash(%q) -> %q -> Hex2Dec(%q) -> %d", hash, partialHash, partialHash, decimal)
	}

	return decimal, nil
}

// GetHash returns hash string
func (h Hash) GetHash() string {
	return h.hash
}

// GetDec returns decimal value
func (h Hash) GetDec() uint64 {
	return h.dec
}

// SumHashes returns sum of given hashes
func SumHashes(logger *log.Logger, hashes ...Hash) uint64 {
	sum := uint64(0)

	for _, hash := range hashes {
		sum += hash.GetDec()
	}

	if logger != nil {
		var dec []uint64

		for _, hash := range hashes {
			dec = append(dec, hash.GetDec())
		}

		logger.Printf(
			"\n<|%s|> => SumHashes(%v) -> (%v) -> %d",
			"Computing sum of hashes",
			hashes,
			dec,
			sum,
		)
	}

	return sum
}
