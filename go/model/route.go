package model

import (
	"hash/adler32"
	"hash/crc32"
)

const (
	RouteTypeHash  = "hash"
	RouteTypeRange = "range"
)

// We will route to let a SQL exexuted in the exact group
type Route struct {
	DB     string `json:"db"`
	Table  string `json:"table"`
	Column string `json:"column"`
	Type   string `json:"type"`

	// For hash type
	// crc32, adler32, or you can use your own hash function if registered
	// default is crc32
	Hash string `json:"hash"`

	// For range type
	// key format 1-10000, left close and right open
	// value is sub table id
	// e.g, if the ranges is {"0-100" : 0, "100-200" : 1}
	// if a key is 50, we may know that the data will be saved in table_0
	Ranges map[string]int `json:"ranges"`
}

type HashFunc func([]byte) uint32

var hashFuncs map[string]HashFunc

func RegisterHashFunc(name string, fn HashFunc) {
	hashFuncs[name] = fn
}

func Hash(name string, data []byte) uint32 {
	fn, ok := hashFuncs[name]
	if !ok {
		return crc32.ChecksumIEEE(data)
	}

	return fn(data)
}

func init() {
	hashFuncs = make(map[string]HashFunc)

	hashFuncs["crc32"] = crc32.ChecksumIEEE
	hashFuncs["adler32"] = adler32.Checksum
}
