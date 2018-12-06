package sktch

import (
	"encoding/binary"
	"hash"
	"hash/fnv"
	"math"
)

const maxInt = math.MaxInt64

// CountMinSketch is implementation of Count Min Sketch
type CountMinSketch struct {
	table        []int
	W            int
	D            int
	hashFunction hash.Hash32
}

func (cm *CountMinSketch) getIndexFor(i, j int) int {
	return cm.W*i + j
}

func (cm *CountMinSketch) incrementByIndex(i, j int) int {
	cm.table[cm.getIndexFor(i, j)] += 1
	return 0
}

func (cm *CountMinSketch) add(event []byte, count int) error {
	_, err := cm.hashFunction.Write(event)

	if err != nil {
		return err
	}

	for i := 0; i < cm.D; i++ {
		err := binary.Write(cm.hashFunction, binary.LittleEndian, uint16(i))
		if err != nil {
			panic(err)
		}
		result := cm.hashFunction.Sum32()
		idx := int(result) % cm.W
		cm.table[cm.W*i+idx] += count
	}

	cm.hashFunction.Reset()

	return nil
}

func (cm *CountMinSketch) count(event []byte) int {
	_, err := cm.hashFunction.Write(event)
	if err != nil {
		panic(err)
	}
	min := maxInt
	for i := 0; i < cm.D; i++ {
		err := binary.Write(cm.hashFunction, binary.LittleEndian, uint16(i))
		if err != nil {
			panic(err)
		}
		result := cm.hashFunction.Sum32()
		idx := int(result) % cm.W
		if cm.table[cm.W*i+idx] < min {
			min = cm.table[cm.W*i+idx]
		}
	}
	return min
}

// NewCountMinSketch creates NewCountMinSketch
func NewCountMinSketch(w, d int) *CountMinSketch {
	return &CountMinSketch{
		table:        make([]int, w*d),
		W:            w,
		D:            d,
		hashFunction: fnv.New32(),
	}
}
