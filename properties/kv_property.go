package properties

import (
	"fmt"
	"github.com/obada-foundation/sdk-go/hash"
	"log"
	"strconv"
)

// KvProperty ...
type KvProperty struct {
	records []Record
	hash    hash.Hash
}

// Record ...
type Record struct {
	key   StringProperty
	value StringProperty
	hash  hash.Hash
}

// NewRecord creates a new key/value record
func NewRecord(key string, value string, log *log.Logger, debug bool) (Record, error) {
	var r Record

	if debug {
		log.Printf("\nNewRecord(%q, %q)", key, value)
	}

	k, err := NewStringProperty(key, log, debug)
	v, err := NewStringProperty(value, log, debug)

	if err != nil {
		return r, err
	}

	kh := k.GetHash()
	vh := k.GetHash()
	kvDec := kh.GetDec() + vh.GetDec()

	if debug {
		log.Printf("(%d + %d) -> %d", kh.GetDec(), vh.GetDec(), kvDec)
	}

	h, err := hash.NewHash(strconv.FormatUint(kvDec, 10), log, debug)

	if err != nil {
		return r, err
	}

	r.key = k
	r.value = v
	r.hash = h

	return r, nil
}

// GetKey returns a record key
func (r *Record) GetKey() StringProperty {
	return r.key
}

// GetValue returns a record value
func (r *Record) GetValue() StringProperty {
	return r.value
}

// GetHash returns a record hash
func (r *Record) GetHash() hash.Hash {
	return r.hash
}

// NewMapProperty creates map property
func NewMapProperty(kv map[string]string, log *log.Logger, debug bool) (KvProperty, error) {
	var mp KvProperty
	var kvDec uint64

	if debug {
		log.Printf("\nNewMapProperty(%v)", kv)
	}

	for key, value := range kv {
		r, err := NewRecord(key, value, log, debug)

		if err != nil {
			return mp, err
		}

		rh := r.GetHash()
		kvDec += rh.GetDec()

		mp.records = append(mp.records, r)
	}

	h, err := hash.NewHash(strconv.FormatUint(kvDec, 10), log, debug)

	if err != nil {
		return mp, fmt.Errorf("cannot hash %q: %w", kvDec, err)
	}

	mp.hash = h

	return mp, nil
}

// GetAll returns slice of records
func (mp *KvProperty) GetAll() []Record {
	return mp.records
}

// GetHash returns hash of all records
func (mp *KvProperty) GetHash() hash.Hash {
	return mp.hash
}
