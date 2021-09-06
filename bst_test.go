package main

import (
	"sort"
	"testing"
)

func TestBst_Insert(t *testing.T) {

}

// func TestBst_Delete(t *testing.T) {}
// func TestBst_Search(t *testing.T) {}

// func TestBst_Min(t *testing.T) {}
// func TestBst_Max(t *testing.T) {}

func TestBst_InOrderWalk(t *testing.T) {

	arr := randomSlice(0, 20, 30)
	bst := newBst()
	for _, v := range arr {
		bst.Insert(v)
	}

	got := bst.InOrderWalk(bst.GetRoot())
	sort.Ints(arr)

	for i := 0; i < len(arr); i++ {
		if got[i] != arr[i] {
			t.Errorf("Node.key = %d, want %d", got[i], arr[i])
		}
	}
}

// func TestBst_PreOrderWalk(t *testing.T) {}
// func TestBst_PostOrderWalk(t *testing.T) {}
