package main

import (
	"fmt"
	//	"math/rand"
	//	"time"
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

func _sort(S []int) (SA []int) {
	N := len(S)
	SA = make([]int, N, N)
	// 各文字の出現回数を記録する表
	count := [5]int{0, 0, 0, 0, 0}
	// 各桁の各文字の出現回数
	for i := 0; i < N; i++ {
		count[S[i]]++
	}
	fmt.Println(count)
	// 累積する
	for k := 1; k < 5; k++ {
		count[k] = count[k-1] + count[k]
	}
	//この時点でcountは、各文字の終了位置を示している。
	//それに従って埋めていく
	for i := N - 1; i >= 0; i-- { // 逆側から
		count[S[i]]-- // デクリメントしてから使う
		SA[count[S[i]]] = i
	}
	fmt.Println(SA)
	fmt.Println(count)
	return
}

// doublingして伸ばしていく
func doubling(SA []int, ISA []int) {
	h := 1

}

func suffixArrayLS(A []int, N int) (SA []int) {
	SA = make([]int, N, N)   // suffix array
	ISA := make([]int, N, N) // inversed SA
	//SA1 := _sort()

	for h := 1; h < 100; h *= 2 {
		// 各桁数 radix sort してから、各区間ごとくっつけて次に行く
		fmt.Println("h =", h)
		// radix sort

		//
	}
	return
}

func inverse(SA []int) (ISA []int) {
	for i := 0; i < len(SA); i++ {
		break
	}
	return
}

/**
** @m ソート対象の桁
**/
func _sort_and_ref(SA []int, ISA []int, S []int, left int, right int, h int) {
	N := len(S)
	// 再帰の終了条件
	if left > right {
		break
	}
	// 各文字の出現回数を記録する表
	count := [5]int{0, 0, 0, 0, 0}
	// 各桁の各文字の出現回数
	for i := 0; i < N; i++ {
		if h == 1 {
			// 初回なので、参考にするSAがない
			count[S[i]]++
		} else {
			// 二回目以降は、SAの次の奴をソートする
			count[S[SA[i]+h-1]]++ //  Sのhけた目の数字のところをインックリメント
		}
	}
	// 累積する
	for k := 1; k < 5; k++ {
		count[k] = count[k-1] + count[k]
	}
	//この時点でcountは、各文字の終了位置を示している。
	//それに従って埋めていく
	tmp := make([]int, N, N)
	for i := N - 1; i >= 0; i-- { // 逆側から
		count[S[i]]-- // デクリメントしてから使う
		tmp[count[S[i]]] = S[i]
	}

	// 切り替わりの部分で再帰呼び出しする
	for k := 0; k < 4; k++ {
		// 再帰呼び出し 次は2hでやる
		_sort_and_ref(SA, ISA, S, count[k], count[k+1]-1, h*2)
	}
}

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
