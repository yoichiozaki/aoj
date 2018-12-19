package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Scan()
	buf := strings.Split(scanner.Text(), " ")
	N, _ := strconv.Atoi(buf[0])
	D, _ := strconv.Atoi(buf[1])
	I, J, K, R := PF(D)

	// そもそもサイコロの目は1, 2, 3, 4, 5, 6のみだから、n回振った時の出た目の積を素因数分解すると2, 3, 5しか出てこない
	// dp[n][i][j][k]: サイコロをn回振った時の出た目の積が2^i * 3^j * 5^kとなる確率
	// ただしi >= I, j >= J, k >= Kなるi, j, kについては dp[n][I][J][K]にまとめる
	// n = 0, 1, 2, ..., N-1, N
	// D = 2^I * 3^J * 5^Kであったとして
	// i = 0, 1, 2, ..., I-1, I
	// j = 0, 1, 2, ..., J-1, J
	// k = 0, 1, 2, ..., K-1, K
	dp := make([][][][]float64, N+2)
	for n := range dp {
		dp[n] = make([][][]float64, I+2)
		for i := range dp[n] {
			dp[n][i] = make([][]float64, J+2)
			for j := range dp[n][i] {
				dp[n][i][j] = make([]float64, K+2)
			}
		}
	}

	if R != 1 && R != 0 {
		fmt.Printf("%.9f", 0.0)
		return
	}

	// 初期値
	dp[0][0][0][0] = 1.0

	for nn := 0; nn <= N; nn++ {
		for ii := 0; ii <= I; ii++ {
			for jj := 0; jj <= J; jj++ {
				for kk := 0; kk <= K; kk++ {
					dp[nn+1][ii][jj][kk] += dp[nn][ii][jj][kk] / 6.0
					dp[nn+1][min(I, ii+1)][jj][kk] += dp[nn][ii][jj][kk] / 6.0
					dp[nn+1][ii][min(J, jj+1)][kk] += dp[nn][ii][jj][kk] / 6.0
					dp[nn+1][min(I, ii+2)][jj][kk] += dp[nn][ii][jj][kk] / 6.0
					dp[nn+1][ii][jj][min(K, kk+1)] += dp[nn][ii][jj][kk] / 6.0
					dp[nn+1][min(I, ii+1)][min(J, jj+1)][kk] += dp[nn][ii][jj][kk] / 6.0
				}
			}
		}
	}

	// 求めたいのはサイコロをn回振った時の出た目の積がI個以上の2とJ個以上の3とK個以上の5の積である確率
	fmt.Printf("%.9f", dp[N][I][J][K])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func PF(D int) (int, int, int, int) {
	store := D
	I, J, K := 0, 0, 0
	for D%2 == 0 {
		I++
		D /= 2
	}
	D = store
	for D%3 == 0 {
		J++
		D /= 3
	}
	D = store
	for D%5 == 0 {
		K++
		D /= 5
	}
	R := D / (intPow(2, I) * intPow(3, J) * intPow(5, K))
	return I, J, K, R
}

func intPow(x, y int) int {
	ret := 1
	for y > 0 {
		ret *= x
		y--
	}
	return ret
}
