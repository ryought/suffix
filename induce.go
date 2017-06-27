package main

import (
	"fmt"
)

// 線形時間SA構築アルゴリズム
func suffixArrayIS(S []int, N int) (SA []int) {
	// step0: Ltype(1) Stype(0)を格納する o(n)
	t := make([]int, N, N)
	t[N-1] = 0
	for i := N - 2; 0 <= i; i-- {
		if S[i] == S[i+1] {
			t[i] = t[i+1]
		} else {
			if S[i] < S[i+1] {
				t[i] = 0
			} else {
				t[i] = 1
			}
		}
	}
	// LMS(-1)を見つける o(n)
	for i := 1; i < N; i++ {
		if t[i-1] == 1 && t[i] == 0 { // 左がLなStypeはLMS
			t[i] = -1
		}
	}
	fmt.Println(t)

	// step1: Lだけの順序をつける
	base := 4 // 文字の種類
	// バケットを作る
	SA = make([]int, N, N)
	// 出現回数
	occ := make([]int, base, base)
	for i := 0; i < N; i++ {
		if S[i] != -1 {
			occ[S[i]]++
		}
	}
	fmt.Println("occ", occ)
	// 先頭の位置を示す
	h := make([]int, base+1, base+1)
	for i := 0; i < N; i++ {

	}
	fmt.Println("h", h)

	return
}

func main() {
	a, c, g, t, n := 1, 2, 3, 4, 0 // nは終端文字
	N := 16
	//S := make([]int, N, N)
	S := []int{a, t, a, a, t, a, c, g, a, t, a, a, t, a, a, n}

	// rand.Seed(time.Now().UnixNano())
	// for i := 0; i < N; i++ {
	// 	A[i] = rand.Intn(4) // 配列を作る
	// }

	fmt.Println(S)
	//insertionSort(A, 0, N-1)
	//SA := suffixArrayIS(S, N)
	SA := _sort(S)
	fmt.Println(SA)
}
