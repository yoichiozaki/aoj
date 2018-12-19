package main

import "fmt"

const (
	// MOD ...
	MOD = 1000000007
)

type combination struct {
	dp  [][]int64
	mod int64
}

func newCombination(n int, mod int64) *combination {
	dp := make([][]int64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int64, n+1)
	}
	// dp[n][k] = dp[n-1][k-1] + dp[n-1][k]
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= n; j++ {
			if j > 0 {
				dp[i][j] = (dp[i-1][j-1] + dp[i-1][j]) % mod
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return &combination{
		dp:  dp,
		mod: mod,
	}
}

func (c *combination) get(n int, r int) int64 {
	if n < r || n < 0 || r < 0 {
		return 0
	}
	return c.dp[n][r]
}

const m0 = 0x5555555555555555 // 01010101 ...
const m1 = 0x3333333333333333 // 00110011 ...
const m2 = 0x0f0f0f0f0f0f0f0f // 00001111 ...

func onesCount64(x uint64) int {
	const m = 1<<64 - 1
	x = x>>1&(m0&m) + x&(m0&m)
	x = x>>2&(m1&m) + x&(m1&m)
	x = (x>>4 + x) & (m2 & m)
	x += x >> 8
	x += x >> 16
	x += x >> 32
	return int(x) & (1<<7 - 1)
}

func main() {
	var R, C int
	var X, Y int
	var D, L int
	fmt.Scanf("%d %d", &R, &C)
	fmt.Scanf("%d %d", &X, &Y)
	fmt.Scanf("%d %d", &D, &L)

	c := newCombination(900, MOD)

	var res int64
	res = 0

	for i := 0; i < (1 << 2); i++ {
		for j := 0; j < (1 << 2); j++ {
			x := X - onesCount64(uint64(i))
			y := Y - onesCount64(uint64(j))
			if x < 0 || y < 0 {
				continue
			}
			if (onesCount64(uint64(i))+onesCount64(uint64(j)))%2 == 0 {
				res = (res + ((c.get(x*y, D) * c.get(x*y-D, L)) % MOD)) % MOD
			} else {
				res = res - ((c.get(x*y, D) * c.get(x*y-D, L)) % MOD)
				for res < 0 {
					res += MOD
				}
				res = res % MOD
			}
		}
	}
	res = ((res * int64((R - X + 1)) % MOD) * int64((C - Y + 1))) % MOD

	fmt.Println(res)
}
