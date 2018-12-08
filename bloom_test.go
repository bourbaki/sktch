package sktch

import (
	"hash/fnv"
	"testing"
	"testing/quick"
)

func TestBloomConstructor(t *testing.T) {
	nHashFunctions := uint(2)
	bloomWithSize := func(a uint16) bool {
		bloom := NewBloomFilter(uint(a), nHashFunctions, fnv.New32())
		expectedSize := (a / 8) + 1
		return len(bloom.array) == int(expectedSize)
	}
	if err := quick.Check(bloomWithSize, nil); err != nil {
		t.Error(err)
	}
}

func TestBloomAddThenContains(t *testing.T) {
	bloomAddElement := func(x []byte) bool {
		bloom := NewBloomFilter(uint(16), 2, fnv.New32())
		if bloom.Contains(x) {
			return false
		}
		bloom.Add(x)
		return bloom.Contains(x)
	}
	if err := quick.Check(bloomAddElement, nil); err != nil {
		t.Error(err)
	}
}
