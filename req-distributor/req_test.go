package main

import (
	"testing"
)

func TestSelectBucket(t *testing.T) {
	var bInfo [3]*bucketInfo
	for i := 0; i < 3; i++ {
		bInfo[i] = &bucketInfo{}
		bInfo[i].init()
	}

	idx := selectBucket(bInfo)
	if idx != 0 {
		t.Errorf("selectBucket\"\" failed, expected %v, got %v", 0, idx)
	}

	idx = selectBucket(bInfo)
	if idx != 1 {
		t.Errorf("selectBucket\"\" failed, expected %v, got %v", 1, idx)
	}

	idx = selectBucket(bInfo)
	if idx != 2 {
		t.Errorf("selectBucket\"\" failed, expected %v, got %v", 2, idx)
	}

	idx = selectBucket(bInfo)
	if idx != 0 {
		t.Errorf("selectBucket\"\" failed, expected %v, got %v", 0, idx)
	}

	idx = selectBucket(bInfo)
	if idx != 1 {
		t.Errorf("selectBucket\"\" failed, expected %v, got %v", 1, idx)
	}
}
