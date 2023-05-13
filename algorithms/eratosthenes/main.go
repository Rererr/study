package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 対象の自然数の上限値
	limit := 20000000

	wg := sync.WaitGroup{}
	// 単純に全件確認
	wg.Add(1)
	go func(limit int) {
		defer wg.Done()
		start := time.Now()
		_ = simple(limit)
		end := time.Now()
		fmt.Printf("simple: %f秒\n", (end.Sub(start)).Seconds())
	}(limit)

	// エラトステネスの篩
	// 所要時間の計測
	wg.Add(1)
	go func(limit int) {
		defer wg.Done()
		start := time.Now()
		_ = eratosthenes(limit)
		end := time.Now()
		fmt.Printf("eratosthenes: %f秒\n", (end.Sub(start)).Seconds())
	}(limit)

	wg.Wait()
}

// simple
//
// 与えられたnまでの自然数が素数かどうかを判定し、map[int]boolで返却する
// int = 自然数, bool = 素数ならtrue,素数でなければfalse
// 素数判定にはisPrimeを用いて、特にアルゴリズムを利用せずに判定する
func simple(n int) map[int]bool {
	isPrimeArray := map[int]bool{}
	for i := 1; i <= n; i++ {
		isPrimeArray[i] = isPrime(i)
	}
	return isPrimeArray
}

// eratosthenes
//
// 与えられたnまでの自然数が素数かどうかを判定し、map[int]boolで返却する
// int = 自然数, bool = 素数ならtrue,素数でなければfalse
// eratosthenesの篩を用いる
func eratosthenes(n int) map[int]bool {
	// 素数かどうかを判定した結果を納める配列
	isPrimeArray := make(map[int]bool, n)

	for i := 1; i <= n; i++ {
		isPrimeArray[i] = true
	}

	// 1は素数ではない
	// 今回のエラトステネスの篩が成立するのは2以上の自然数のみになることや、素因数分解で1が素数だと無限に1が登場できてしまうなど
	// 数字を扱う際に1が素数であると不都合が多いため
	isPrimeArray[1] = false

	// 2から順に素数かどうかを判定する
	for i := 2; i <= n; i++ {
		// isPrimeArray[i] = false つまり既に素数ではないとされた数値の場合はスキップ
		if !isPrimeArray[i] {
			continue
		}
		// iの倍数は(iで割り切れるため)全て素数ではないので、ラベルをfalseに変更する
		for q := i * 2; q <= n; q += i {
			isPrimeArray[i] = false
		}
	}

	return isPrimeArray
}

// isPrime
//
// 与えられた数値nが素数がどうかを判定する関数
func isPrime(n int) bool {
	// 素数は自然数、かつ1を含まないという性質から2より小さい数値が与えられた場合はfalse(素数ではない)を返す
	if n < 2 {
		return false
	}
	// nが素数ではない場合、√nまで試し割りすれば判定できる
	for i := 2; i*i <= n; i++ {
		// 割り切れる場合には素数ではない
		if n%i == 0 {
			return false
		}
	}
	// √nまでの全ての自然数で割り切れなければ素数となる
	return true
}
