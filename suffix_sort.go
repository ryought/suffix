package main

import (
	"fmt"
)

// DONE
// _sort do a sort
func _rad_sort(S []int) (SA []int, ISA []int, count [5]int) {
	N := len(S)
	SA = make([]int, N, N)
	// 各文字の出現回数を記録する表
	count = [5]int{0, 0, 0, 0, 0}
	// 各桁の各文字の出現回数
	for i := 0; i < N; i++ {
		count[S[i]]++
	}
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

	// ISA作る 順位の最大値を取ってくる
	ISA = make([]int, N, N)
	for i := 0; i < N; i++ {
		ISA[SA[i]] = i
	}
	return
}

func doubling(SA []int, ISA []int, left int, right int, h int) {
	if left >= right {
		fmt.Println("hoge")
	}

}

// _partition parts SA[l:r) into [l:i) [i:j) [j:r) by elem's size
func _partition(SA []int, ISA []int, l int, r int, h int) (i, j int) {
	pivot := SA[r-1] // 右をピボットに取る
	mi, mj := l, r-1 // pivotと同じものゾーンの左 最初は [0:0)と[r-1:r)がpivotゾーン
	i, j = l, r-1    // pivotより小ゾーン[l:i), [j:r) の範囲にする
	for {
		for ; i < j && SA[i] <= pivot; i++ { // iゾーン
			if SA[i] == pivot {
				// 端っこに寄せる
				SA[mi], SA[i] = SA[i], SA[mi]
				mi++
			}
		}
		for ; i < j && pivot <= SA[j-1]; j-- { // jゾーンを増やす 違うとこで止める
			if SA[j-1] == pivot {
				// 端っこに寄せる
				SA[mj-1], SA[j-1] = SA[j-1], SA[mj-1]
				mj--
			}
		}
		if j <= i {
			break
		}
		// 入れ替えて、そうすると条件を満たすはずなのでインクリメントしてループに戻る
		SA[i], SA[j-1] = SA[j-1], SA[i]
		i++
		j--
	}
	// pivotと同じ奴を元に戻す
	for l < mi {
		SA[i-1], SA[mi-1] = SA[mi-1], SA[i-1]
		mi--
		i--
	}
	for mj < r {
		SA[j], SA[mj] = SA[mj], SA[j]
		mj++
		j++
	}
	return
}

// SAをISA  SA[l:r)をソートする関数
func _split_sort(SA []int, ISA []int, l int, r int, h int) {
	if l < r-1 {
		i, j := _partition(SA, ISA, l, r, h)
		_split_sort(SA, ISA, l, i, h) // 次はSA[l:i)
		// [i:j) は同じものが並んでる
		_split_sort(SA, ISA, j, r, h) // 次はSA[j:r]
	}
}

func main() {
	//a, c, g, t, n := 1, 2, 3, 4, 0 // nは終端文字
	//S := make([]int, N, N)
	//S := []int{a, t, a, a, t, a, c, g, a, t, a, a, t, a, a, n}
	S := []int{10, 500, 20, 10, 50, 40, 50, 20, 40, 30, 100, 200, 50}
	SA := make([]int, len(S), len(S))
	ISA := make([]int, len(SA), len(SA))

	// rand.Seed(time.Now().UnixNano())
	// for i := 0; i < N; i++ {
	// 	A[i] = rand.Intn(4) // 配列を作る
	// }

	fmt.Println("(input)", S)
	//insertionSort(A, 0, N-1)
	//SA := suffixArrayIS(S, N)
	//SA, ISA, count := _rad_sort(S)
	// for k := 0; k < 4; k++ {
	// 	doubling(SA, ISA)
	// }
	//fmt.Println("(SA1)", SA, "(count)", count, "(ISA)", ISA)

	_split_sort(S, ISA, 0, len(S), 0)
	fmt.Println("(sorted)", S)
}
