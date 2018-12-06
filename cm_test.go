package sktch

import (
	"testing"
)

func TestCountMinSketch(t *testing.T) {
	W := 1000
	d := 100
	cm := NewCountMinSketch(W, d)
	if len(cm.table) != W*d {
		t.Error("Expected length of 20, got ", len(cm.table))
	}
	cm.add([]byte("Markus Brown"), 2)
	if cm.count([]byte("Markus Brown")) != 2 {
		t.Errorf("CM should return the exact result if this only one event")
	}
	if cm.count([]byte("Ulysses")) != 0 {
		t.Errorf("Very unlikely")
	}
}
