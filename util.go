package main

import (
	"math/rand"
	"time"
)

func getRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

// generate `n` numbers between `start` and `end`
func randomSlice(start int, end int, n int) []int {
	if end < start || (end-start) < n {
		return nil
	}
	nums := make([]int, 0)
	for len(nums) < n {
		num := getRand().Intn((end - start)) + start
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}
