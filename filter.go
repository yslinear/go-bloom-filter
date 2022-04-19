package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"hash"
)

type filter struct {
	bitfield [10]bool
}

var hasher = sha1.New()

func createHash(h hash.Hash, input string) int {
	bits := h.Sum([]byte(input))
	buffer := bytes.NewBuffer(bits)
	result, _ := binary.ReadVarint(buffer)
	return int(result)
}

func (f *filter) hashPosition(s string) int {
	hs := createHash(hasher, s)
	if hs < 0 {
		hs = -hs
	}
	return hs % len(f.bitfield)
}

func (f *filter) Set(s string) {
	pos := f.hashPosition(s)
	f.bitfield[pos] = true
}
func (f *filter) Get(s string) bool {
	return f.bitfield[f.hashPosition(s)]
}
