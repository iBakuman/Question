package main

import (
	"math/rand"
	"testing"
)

func Test1(t *testing.T) {
	arrLen := 100
	arr := make([]int, arrLen)
	for i := 0; i < arrLen; i++ {
		arr[i] = rand.Intn(arrLen)
	}
	resolve(arr)
}
