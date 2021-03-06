# ソフト論 課題2個目

`go run suffix_sort.go`で動きます。

## 課題2 Larsson-Sadakane
`suffix_sort.go`の142行目、`suffix_array_LS()`で実装している。
初回の1桁目のradixsortを`_rad_sort()`で、その後の2桁目以降のDoublingによるソートを`_split_sort()`で実装している。


## 課題3 Induced sorting
`suffix_sort.go`の538行目、`test_SAIS()`を実行すると、10^1~10^8個の文字からなる文字列(文字種は1,2,3,4、終端文字は0)をランダムに生成した上で、Suffix Arrayを作成。そのSAの作成にかかった時間をns単位で表示する。

実際のInduced sortingの実装は、`suffix_array_IS()`(L324)にある。

その結果は、次のようになった。
10	2882
100	46412
1000	158972
10000	2110784
100000	17082284
1000000	407467893
10000000	8114630998
100000000	1.6008E+11
これを両軸対数目盛りでグラフにした(添付のsais_result.png)上で、線形近似すると$y=148.26 x^{1.09}$となった。肩が1に近く、o(n)であるといって良さそうだが、0.09の誤差の生じた理由は?


## 課題4 BWT実装
`suffix_sort.go`の557行目、`test_BWT()`を実行すると、ランダムに生成した文字列に対してSAとBWT、オカレンステーブルを構築したあと、ランダムに生成した検索文字列query(長さk=10~20)に対して検索処理を行い、それにかかった時間を表示する。
BWTの実装は、getBWT()でSAからBWTを構築、getOccAndC()でオカレンステーブルを作成、最後にsearchBWTで作成したテーブルを用いて部分文字列を検索する。
実際にSの中に含まれるqueryの開始位置を知りたい場合は、search()(L308)にあるように、SA上のインデックスがlb~ubであるような要素を取り出せば良い。

動かした結果は、添付の`BWT_time.xlsx`にある。N=100000のときについてグラフを描画すると、外れ値はあるものの線形になっていることがわかった。

