package main

import (
	"fmt"
	"math/rand"
  "math"
	"time"
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
    fmt.Println("(part)", i, j)
		_split_sort(SA, ISA, l, i, 2*h) // 次はSA[l:i)
		// [i:j) は同じものが並んでる 順位確定してるのでISA更新
		for x := i; x < j; x++ {
			ISA[SA[x]] = j - 1
		}
		_split_sort(SA, ISA, j, r, 2*h) // 次はSA[j:r]
	}
}


// 課題2 Larsson-SadakaneアルゴリズムによるSuffixArray構築
func suffix_array_LS(S []int) []int {
	SA, ISA, count := _rad_sort(S)

  fmt.Println("(S)", S, SA,  count)
	for i := 0; i < len(ISA); i++ {
		fmt.Println(ISA[SA[i]], " ")
	}

	fmt.Println("(first rad finished)", SA, ISA, count)

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
	// もし間に不等号が一つでもあれば、次の関係は不等号になる
	for i := 0; i < len(SA); i++ {
		// もしすでに書き込まれていて、それの一個長いものがLtype(=1)だったら
		if SA[i] > 0 && t[SA[i]-1] == 1 {
			// 次の位置に書き込んでインクリメント
			x := SA[i] - 1 // SA[i]の大小関係を元に、SA[i]-1の順位を確定する
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
			x := SA[i] - 1 // 推論した先の要素の添え字
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
	// S/L/SMLの分類をする S=0, T=0
	N := len(S)
	t = make([]int, N, N)
	t[N-1] = 0 // $はstypeとする
	for i := N - 2; 0 <= i; i-- {
		// となりが同じ文字だったら、その隣のtypeを引き継ぐ
		if S[i] == S[i+1] {
			t[i] = t[i+1]
		} else {
			// 文字が違ったら、その大小で決まる
			if S[i] < S[i+1] {
				t[i] = 0
			} else {
				t[i] = 1
			}
		}
	}
	return
}

// SAからISAを取得
func getISA(SA []int) (ISA []int) {
	ISA = make([]int, len(SA), len(SA))
	for i := 0; i < len(SA); i++ {
		ISA[SA[i]] = i
	}
	return
}

// DONE BWTを取得
func getBWT(SA []int, S []int) (BWT []int) {
	BWT = make([]int, len(SA), len(SA))
	for i := 0; i < len(SA); i++ {
		if SA[i] == 0 {
			BWT[i] = S[len(SA)-1]
		} else {
			BWT[i] = S[SA[i]-1]
		}
	}
	return
}

// 問い合わせ配列を探すためのOcc,Cを作る
func getOccAndC(BWT []int, base int) (Occ [][]int, C []int) {
  // Occ[x][y] : 文字xが、BWT中の位置y以下で出現した回数
  // initialize
  Occ = make([][]int, base, base)
  for k:=0; k<base; k++ {
    Occ[k] = make([]int, len(BWT), len(BWT))
  }
  // Occ作成
  Occ[BWT[0]][0] ++
  for i:=1; i<len(BWT); i++ {
    for k:=0; k<base; k++ {
      Occ[k][i] = Occ[k][i-1]
    }
    Occ[BWT[i]][i] ++
  }

  // C作成
  C = make([]int, base, base)
  C[0] = 0
  for k:=1; k<base; k++ {
    C[k] = C[k-1] + Occ[k-1][len(BWT)-1]
  }
  return
}

// BWTを使って文字列探索をする
func searchBWT(BWT []int, Occ [][]int, C []int, query []int) (lb, ub int) {
  // 初期値
  lb = 0
  ub = len(BWT) -1
  // ループ
  for k:=len(query)-1; k>=0; k-- {
    x := query[k]
    if lb-1 == -1 {
      lb = C[x] + 0
    } else {
      lb = C[x] + Occ[x][lb-1]
    }
    if ub == -1 {
      ub = C[x] + 0 -1
    } else {
      ub = C[x] + Occ[x][ub] -1
    }
  }
  return
}

// SAとBWTで文字列探索するやつ
func search(S []int, query []int) (pos []int) {
  fmt.Println("(s)", S, "(q)", query)
  SA := suffix_array_IS(S)
  BWT := getBWT(SA, S)
  Occ, C := getOccAndC(BWT, 5)
  lb, ub := searchBWT(BWT, Occ, C, query)
  size := ub-lb + 1
  pos = make([]int, size, size)
  for i:=0; i<size; i++ {
    pos[i] = SA[lb+i]
  }
  return
}


// 線形時間SA構築アルゴリズム
func suffix_array_IS(S []int) (SA []int) {
	// (0) Ltype Stypeの分類をする
	t := typeLS(S)
	// (1) バケットを作る
	// Sの中の最大値maxを探して、max+1を文字の種類数とする
	max := S[0]
	for i := 1; i < len(S); i++ {
		if max < S[i] {
			max = S[i]
		}
	}
	// バケットbは、バケットの境界の位置を保存した配列
	b := countArray(S, max+1)
	// (2) ソート済み配列を用意
	//定理 Suffix Array上で先頭文字xについて、x[Ltype], x[Stype]の順になる
	SA = LMSsorted(S, t, b)
	// (3-1) 左からinduce
	induceL(S, SA, b, t)
	// (3-2) 右からinduce
	induceR(S, SA, b, t)
	return
}

// Induceするために、LMSについて順序が決定しているバケットを返す
func LMSsorted(S []int, t []int, B []int) (SA2 []int) {
	SA := make([]int, len(S), len(S))
	SA2 = make([]int, len(S), len(S))
	for i := 0; i < len(S); i++ {
		SA[i] = -1
		SA2[i] = -1
	}
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
	// よくわからないけど順番が付いてしまった
	// LMSだけを適当に並べたSAについてinduceしてみる
	// (a)順番は関係なくLMSだけを並べる
	b := make([]int, len(B)-1, len(B)-1)
	for i := 0; i < len(B)-1; i++ {
		b[i] = B[i+1] - 1
	}
	b2 := make([]int, len(B)-1, len(B)-1)
	copy(b2, b)

	// LMSの登録
	for i := 0; i < len(S); i++ {
		if tLMS[i] == 1 {
			// LMSは、SAに登録する
			SA[b[S[i]]] = i
			b[S[i]]--
		}
	}
	// (b)induceする
	// (b-1)実際にinduceする
	induceL(S, SA, B, t)
	induceR(S, SA, B, t)
	// (c)LMSについての順序が決定しているか？していなければ再帰呼び出し
	// 順序けっていしているかどうか
	l := 1
	prev := 0
	reccursion := 0
	for i := 1; i < len(SA); i++ {
		if tLMS[SA[i]] >= 1 { // LMSで
			// prevと順序がついてるかどうかを調べる
			// SA[prev]とSA[i]の比較
			c := 0
			for {
				if S[SA[prev]+c] != S[SA[i]+c] {
					l++
					break
				}
				if c != 0 && tLMS[SA[prev]+c] >= 1 && tLMS[SA[i]+c] >= 1 {
					// 同時に終了したら、同じ文字列
					reccursion++
					break
				}
				if c != 0 && (tLMS[SA[prev]+c] >= 1 || tLMS[SA[i]+c] >= 1) {
					// 片方だけだったら、違う文字列だけどそこで探索終了
					l++
					break
				}
				if prev+c >= len(SA)-1 || i+c >= len(SA)-1 {
					break
				}
				c++
			}
			tLMS[SA[i]] = l
			prev = i
		}
	}
	for i := 0; i < len(S); i++ {
		SA2[i] = -1
	}
	if reccursion > 0 {
		fmt.Println("再帰呼び出しをします")
		//fmt.Println(tLMS)
		// 再帰したい 新しい文字列つくる
		// 文字列の長さ
		length := 0 // LMSの長さ収録文字列の長さ
		for i := 0; i < len(tLMS); i++ {
			// 全長の計算
			if tLMS[i] >= 1 {
				length++
			}
		}
		newS := make([]int, length+1, length+1) //再帰するべき新しい文字列
		Ss := make([]int, length+1, length+1)   //再帰から戻ってきた時に、元の順序を復元するためのメモ
		j := 0
		for i := 0; i < len(tLMS); i++ {
			if tLMS[i] >= 1 {
				newS[j] = tLMS[i]
				Ss[j] = i
				j++
			}
		}
		newS[j] = 0 //末尾に0($)を付加
		newSA := suffix_array_IS(newS)
		//fmt.Println(newS, newSA)
		// 元の順位を復元 newSAを右から(大きいsuffixから)走査して、順番につめる
		for i := len(newSA) - 1; i >= 0; i-- {
			x := Ss[newSA[i]]
			//fmt.Println(x)
			SA2[b2[S[x]]] = x
			b2[S[x]]--
		}
		// SA2に詰めて返却
	} else {
		// 再帰しなくていいから、詰めてそのまま返す
		// 詰める作業
		for i := len(S) - 1; i >= 0; i-- {
			if tLMS[SA[i]] >= 1 { // SA[i]がLMSの時
				x := SA[i]
				SA2[b2[S[x]]] = SA[i]
				b2[S[x]]--
			}
		}
	}
	return
}

// 等号不等号のテーブルを持ったバケットを表示する 使っていない
func dispBucket(SA []int, rl []int) {
	fmt.Print(SA[0])
	for i := 1; i < len(SA); i++ {
		if len(rl) > 0 && rl[i] == 1 {
			// 等号
			fmt.Print("=")
		} else {
			// 不等号
			fmt.Print("<")
		}
		fmt.Print(SA[i])
	}
	fmt.Print("\n")
}

// suffix arrayが正しく構築されたかを確認するための、SA表示ツール
func dispSA(S []int, SA []int) {
  fmt.Println("(**SA**)")
  for i:=0; i<len(SA); i++ {
    fmt.Println("(",i,":",SA[i],")",S[SA[i]:])
  }
}

func main() {
	a, c, g, t, n := 1, 2, 3, 4, 0 // nは終端文字
	//S := []int{a, t, a, a, t, a, c, g, a, t, a, a, t, a, a, n}
	//S := []int{a, t, a, t, c, g, t, a, t, c, g, a, a, t, a, g, c, t, t, t, c, a, t, a, c, g, a, t, a, a, t, a, a, n}
	fmt.Println(a, c, g, t, n)
	//S := []int{t, a, a, t, a, a, t, a, a, t, c, n}
	//S := []int{a, t, a, a, t, a, c, g, a, t, a, a, t, a, a, n}
	//S := []int{2,2,3,1,0}
	//S := []int{a, t, a, a, t, c, a, t, c, a, t, c, g, t, a, a, t, a, a, n}
	//SA := make([]int, len(S), len(S))
	//ISA := make([]int, len(SA), len(SA))

	//N := 100000
	//S := make([]int, N, N)
	rand.Seed(time.Now().UnixNano())
	//for i := 0; i < N-1; i++ {
	//	S[i] = rand.Intn(4) + 1 // 配列を作る
	//}
	//S[N-1] = 0 // 終端文字
	//S := []int{2,1,1,2,1,1,2,1,1,2,1,1,0}

	fmt.Println("(input)")
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
	//SA := suffix_array_LS(S)
	//fmt.Println(SA)
	// suffix array 構築
	// SA := suffix_array_naive(S)
	// SA := suffix_array_IS(S)
	//fmt.Println("(result)", SA)
  //test_SAIS()
  test_BWT()
}

// 課題3-3 ランダム配列に対してo(n)を確認する
func test_SAIS() {
  for k:=1; k<=8; k++ {
    N := int(math.Pow10(k))
    fmt.Println("(testing", N, ")")
    S := make([]int, N, N)
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < N-1; i++ {
      S[i] = rand.Intn(4) + 1 // 配列を作る
    }
    S[N-1] = 0 // 終端文字

    start := time.Now()
    SA := suffix_array_IS(S)
    end := time.Now()
    fmt.Println("(finished)", len(SA), "time", end.Sub(start).Nanoseconds())
  }
}

// 課題4-2 長さnのランダムな配列に対して、長さkの問い合わせ配列をランダムに生成、問い合わせの平均時間を調べる
func test_BWT() {
  for t:=5; t<=7; t++ {
    // 文字列Sの作成
    N := int(math.Pow10(t))
    fmt.Println("(testing", N, ")")
    S := make([]int, N, N)
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < N-1; i++ {
      S[i] = rand.Intn(4) + 1 // 配列を作る
    }
    S[N-1] = 0 // 終端文字
    // suffix arrayの構築
    SA := suffix_array_IS(S)
    BWT := getBWT(SA, S)
    Occ, C := getOccAndC(BWT, 5)

    // k 問い合わせ文字列の長さ
    for k:=10; k<=20; k++ {
      // r0 queryに関する試行回数 同じ長さの違うqueryをr0種類作る
      for r0:=0; r0<10; r0++ {
        // 問い合わせ文字列queryの作成
        query := make([]int, k, k)
        rand.Seed(time.Now().UnixNano())
        for i := 0; i < k; i++ {
          query[i] = rand.Intn(4) + 1 // 配列を作る
        }
        fmt.Print(k, ",", query, ",")

        // 何回か試行する 10回やったその平均で時間を算出
        for r:=0; r<50; r++ {
          start := time.Now()
          _, _ = searchBWT(BWT, Occ, C, query)
          end := time.Now()
          fmt.Print(",", end.Sub(start).Nanoseconds())
        }
        fmt.Print("\n")
      }
    }
  }
}

func test_SAIS2() {
	a, c, g, t, n := 1, 2, 3, 4, 0 // nは終端文字
	S := []int{a, t, a, a, t, a, c, g, a, t, a, a, t, a, a, n}
	start := time.Now()
	SA := suffix_array_IS(S)
  dispSA(S, SA)
	end := time.Now()
	fmt.Println("SA", SA)
	fmt.Println(end.Sub(start).Nanoseconds())
	//BWT := getBWT(SA, S)
  //Occ, C:=getOccAndC(BWT, 5)
  query := []int{a,a,t}
  //lb, ub := searchBWT(BWT, Occ, C, query)
  fmt.Println("(search)", search(S, query))
}
