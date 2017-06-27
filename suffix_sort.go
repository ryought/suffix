package main

import (
	"fmt"
)

// suffix同士の大小比較  s[x,:) > s[y,;)だったら1、そうでなければ-1を返す
func suffix_comp(S []int, x int, y int) int {
	if S[x] > S[y] {
		return 1
	} else if S[x] < S[y] {
		return -1
	} else { // S[x] == S[y]
		if S[x] == 0 { // 絶対xが短い
			return -1
		} else if S[y] == 0 {
			return 1
		} else {
			return suffix_comp(S, x+1, y+1)
		}
	}
}

// suffix array構築 ナイーブな実装
func suffix_array_naive(S []int) (SA []int) {
	SA = make([]int, len(S), len(S))

	for i := 0; i < len(S); i++ {
		SA[i] = i
	}
	fmt.Println(SA)
	for c := 0; c < len(S); c++ {
		d := c
		for d > 0 && suffix_comp(S, d, d-1) > 0 {
			SA[d], SA[d-1] = SA[d-1], SA[d]
			d--
		}
	}
	return
}

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
	M := 4
	X := N - 1
	for i := N - 1; i >= 0; i-- {

		ISA[SA[i]] = X
		if i == count[M] {
			M--
			X = i - 1
		}
	}
	return
}

// _partition parts SA[l:r) into [l:i) [i:j) [j:r) by elem's size
func _partition(SA []int, ISA []int, l int, r int, h int) (i, j int) {
	pivot := SA[(l+r)/2] // 中央値をピボットに取る
	mi, mj := l, r-1     // pivotと同じものゾーンの左 最初は [0:0)と[r-1:r)がpivotゾーン
	i, j = l, r-1        // pivotより小ゾーン[l:i), [j:r) の範囲にする
	for {
		for ; i < j && ISA[SA[i]+h] <= pivot; i++ { // iゾーン
			if ISA[SA[i]+h] == pivot {
				// 端っこに寄せる
				SA[mi], SA[i] = SA[i], SA[mi]
				mi++
			}
		}
		for ; i < j && pivot <= ISA[SA[j-1]+h]; j-- {
			// jゾーンを増やす 違うとこで止める
			if ISA[SA[j-1]+h] == pivot {
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

// _split_sortは SA[l:r)を ISA[SA[i]+h] の大小によってソートする関数
func _split_sort(SA []int, ISA []int, l int, r int, h int) {
	if l < r-1 {
		fmt.Println("(called)", l, r, h, SA, ISA)
		i, j := _partition(SA, ISA, l, r, h)
		_split_sort(SA, ISA, l, i, 2*h) // 次はSA[l:i)
		// [i:j) は同じものが並んでる 順位確定してるのでISA更新
		for x := i; x < j; x++ {
			ISA[SA[x]] = j - 1
		}
		_split_sort(SA, ISA, j, r, 2*h) // 次はSA[j:r]
	}
}

func suffix_array_LS(S []int) []int {
	SA, ISA, count := _rad_sort(S)

	for i := 0; i < len(ISA); i++ {
		fmt.Println(ISA[SA[i]], " ")
	}

	fmt.Println(SA, ISA, count)

	for i := 0; i < 4; i++ {
		_split_sort(SA, ISA, count[i], count[i+1], 1)
	}
	_split_sort(SA, ISA, count[4], len(SA), 1)
	return SA
}

func induceL(S []int, SA []int, B []int, t []int) {
	// Step2 SAを左側から走査して、Ltypeの順位を、一つ左側を根拠に誘導する
	// ^テーブルbを作る
	b := make([]int, len(B)-1, len(B)-1)
	for i := 0; i < len(B)-1; i++ {
		b[i] = B[i]
	}
	// SAを左から(0~)走査
	for i := 0; i < len(SA); i++ {
		// もしすでに書き込まれていて、それの一個長いものがLtype(=1)だったら
		if SA[i] > 0 && t[SA[i]-1] == 1 {
			// 次の位置に書き込んでインクリメント
			x := SA[i] - 1
			SA[b[S[x]]] = x
			b[S[x]]++
		}
	}
}

func induceR(S []int, SA []int, B []int, t []int) {
	// Step3 SAを右側から走査
	// ^テーブルbの準備
	b := make([]int, len(B)-1, len(B)-1)
	for i := 0; i < len(B)-1; i++ {
		b[i] = B[i+1] - 1
	}
	// induceする
	// SAを右から(n-1 ~ )走査
	for i := len(SA) - 1; i > 0; i-- {
		// すでに数字があって、それの一個長いものがStype(=0)だったら、
		if SA[i] > 0 && t[SA[i]-1] == 0 {
			x := SA[i] - 1
			SA[b[S[x]]] = x
			b[S[x]]--
		}
	}
}

// DONE Sに含まれる各文字(base種類)のbacket上での区切り位置を含む表を返す
func countArray(S []int, base int) []int {
	// 各文字の出現回数を記録する表
	b := make([]int, base+1, base+1)
	// 各桁の各文字の出現回数
	for i := 0; i < len(S); i++ {
		b[S[i]]++
	}
	// 累積する
	for k := 1; k < base+1; k++ {
		b[k] = b[k-1] + b[k]
	}
	// ずらす
	for k := base; k > 0; k-- {
		b[k] = b[k-1]
	}
	b[0] = 0
	return b
}

func typeLS(S []int) (t []int) {
	// S/L/SMLの分類 S=0, T=0
	N := len(S)
	t = make([]int, N, N)
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
	fmt.Println(t)
	return
}

func suffix_array_IS(S []int) (SA []int) {
	N := len(S)

	// バケットを作る
	b := countArray(S, 5)
	// LMSprefixのソート
	//SAp := make([]int, N, N)
	t := typeLS(S)

	// LMSsuffixがソート済みと仮定して、induce
	// それぞれの文字の個数を引き出す

	// Ltypeを詰めていく
	//定理 Suffix Array上で先頭文字xについて、x[Ltype], x[Stype]の順になる
	for i := 0; i < N; i++ {
		if SA[i] != -1 && t[i-1] == 1 {
			fmt.Println("found")
			SA[b[i]] = i
			b[i]++
		}
	}
	// Stypeを詰めていく
	for i := N - 1; i > 0; i-- {
		if SA[i] != -1 && t[i-1] == 0 {
			// iの順序通りに、i-1を入れる
		}
	}
	//
	SA = S
	return
}

func LMSsubstring(S []int, t []int, b []int) {

}

func LMSsorted(S []int, t []int, B []int) (SA2 []int) {
	SA := make([]int, len(S), len(S))
	SA2 = make([]int, len(S), len(S))
	// とりあえずLMSだけぶち込んでソートしてみる
	// LMSを分類する tLMSは、iがLMSの時(=iがsで、i-1がlのもの)、1になる
	tLMS := make([]int, len(t), len(t))
	nLMS := 0 //LMSの数
	for i := 1; i < len(S); i++ {
		if t[i] == 0 && t[i-1] == 1 {
			tLMS[i] = 1
			nLMS++
		}
	}
	fmt.Println("(tLMS)", tLMS)
	// よくわからないけど順番が付いてしまった
	// LMSだけを適当に並べたSAについてinduceしてみる
	// (a)順番は関係なくLMSだけを並べる
	b := make([]int, len(B)-1, len(B)-1)
	for i := 0; i < len(B)-1; i++ {
		b[i] = B[i+1] - 1
	}
	b2 := make([]int, len(B)-1, len(B)-1)
	copy(b2, b)
	fmt.Println(b)
	for i := 0; i < len(S); i++ {
		if tLMS[i] == 1 {
			// LMSは、SAに登録する
			SA[b[S[i]]] = i
			b[S[i]]--
		}
	}
	fmt.Println("(b2)", b2)
	// (b)induceする
	induceL(S, SA, B, t)
	induceR(S, SA, B, t)
	// (c)LMSについての順序が決定しているか？していなければ再帰呼び出し
	fmt.Println("(SA)", SA)
	for i := 0; i < len(S); i++ {
		SA2[i] = -1
	}
	for i := len(S) - 1; i >= 0; i-- {
		if tLMS[SA[i]] == 1 { // SA[i]がLMSの時
			// 大きさが大事
			x := SA[i]
			SA2[b2[S[x]]] = SA[i]
			b2[S[x]]--
		}
	}
	fmt.Println("(SA2)", SA2)
	//SA = []int{15, -1, -1, -1, -1, -1, 10, 2, 5, 8, -1, -1, -1, -1, -1, -1}
	//fmt.Println("(ref)", SA)
	return
}

func main() {
	a, c, g, t, n := 1, 2, 3, 4, 0 // nは終端文字
	S := []int{a, t, a, a, t, a, c, g, a, t, a, a, t, a, a, n}
	//S := []int{a, t, a, a, t, c, a, t, c, a, t, c, g, t, a, a, t, a, a, n}
	//SA := make([]int, len(S), len(S))
	//ISA := make([]int, len(SA), len(SA))

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
	//for i := 0; i < len(S); i++ {
	//	fmt.Println(suffix_comp(S, i, i+1))
	//}
	// SA := suffix_array_LS(S)
	//fmt.Println(SA)
	// suffix array 構築
	// SA := suffix_array_naive(S)
	// SA := suffix_array_IS(S)
	//fmt.Println("(result)", SA)
	ty := typeLS(S)
	b := countArray(S, 5)
	fmt.Println(ty, b)
	//SA := LMSsubstring(S, ty, b)
	SA := LMSsorted(S, ty, b)
	fmt.Println(SA)
	induceL(S, SA, b, ty)
	fmt.Println("(induceL)", SA)
	induceR(S, SA, b, ty)
	fmt.Println("(induceR)", SA)
	fmt.Println("(finished)")
}
