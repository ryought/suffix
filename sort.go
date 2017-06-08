package main

import (
	"fmt"
	"math/rand"
	"time"
)

func insertionSort(A []int, left int, right int) {
	for c := left + 1; c <= right; c++ {
		d := c
		for d > 0 && A[d] < A[d-1] {
			A[d], A[d-1] = A[d-1], A[d]
			d--
		}
	}
}

func digit(A []int, j int, m int) {
	q := 1
}

func radixSort(A []int, N int, m int) {
	tmp := make([]int, N)

	// jは桁数
	for j := 1; j <= maxDigit; j++ {
		tmp[0] = 1
	}
}

func suffix_array_LS(S []int, N int) (SA []int) {
	SA := make([]int, N) // suffix array
	return
}

func main() {
	N := 10
	A := make([]int, N, N)
	B := map[int]string{0: "hoge", 1: "yahoo"}
	fmt.Println(B)

	// 配列生成
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < N; i++ {
		A[i] = rand.Intn(4)
	}
	fmt.Println(A)
	insertionSort(A, 0, N-1)
	fmt.Println(A)
}
