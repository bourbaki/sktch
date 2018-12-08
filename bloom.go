package sktch

import (
	"hash"
)

// Set is set implemetation
type Set interface {
	Add(element []byte)
	Contains(element []byte) bool
}

// BloomFilter is an implementation of bloom filter
type BloomFilter struct {
	array []byte
	hash  hash.Hash32
	k     uint
	m     uint
}

// NewBloomFilter creates new Bloom Filter with m bits and k hash functions
func NewBloomFilter(m uint, k uint, hash hash.Hash32) *BloomFilter {
	return &BloomFilter{
		make([]byte, m/8+1),
		hash,
		k,
		m,
	}
}

// Contains implements a Set interface for BloomFilter
func (bf *BloomFilter) Contains(element []byte) bool {
	mask := bf.getBitMaskForElement(element)
	isInSet := true
	for i, block := range bf.array {
		blockMask := (block & mask[i]) == mask[i]
		isInSet = isInSet && blockMask
	}
	return isInSet
}

func (bf *BloomFilter) getBitMaskForElement(element []byte) []byte {
	mask := make([]byte, len(bf.array))
	_, err := bf.hash.Write(element)
	if err != nil {
		panic(err)
	}
	for i := 0; i < int(bf.k); i++ {
		bf.hash.Write([]byte(string(i)))
		idx := bf.hash.Sum32() % uint32(bf.m)
		byteIdx := idx / 8
		byteOffset := idx % 8
		mask[byteIdx] |= (byte(1) << byteOffset)
	}
	bf.hash.Reset()
	return mask
}

// Add method implements add operations for bloom filter
func (bf *BloomFilter) Add(element []byte) {
	mask := bf.getBitMaskForElement(element)
	for i := range bf.array {
		bf.array[i] |= mask[i]
	}
}
