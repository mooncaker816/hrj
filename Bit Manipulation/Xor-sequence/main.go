package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scanf("%d\n", &t)
	for i := 0; i < t; i++ {
		var l, r int
		fmt.Scanf("%d", &l)
		fmt.Scanf("%d", &r)
		fmt.Println(G(l-1) ^ G(r))
	}
}

// 1 ^ 2 ^ ...^ n
func xorupto(n int) int {
	// 1- Find the remainder of n by moduling it with 4.
	// 2- If rem = 0, then xor will be same as n.
	// 3- If rem = 1, then xor will be 1.
	// 4- If rem = 2, then xor will be n+1.
	// 5- If rem = 3 ,then xor will be 0.
	// n % 4
	switch n & 3 {
	case 0:
		return n
	case 1:
		return 1
	case 2:
		return n + 1
	case 3:
		return 0
	default:
		return 0
	}
}

// A[i] = 1 ^ 2 ^ 3 ... ^ i
// A[i]^A[i+1]^A[i+2]...^A[j-1]^A[j]=
// A[1]^A[2]^...^A[i-1]^
// A[i]^A[i+1]^A[i+2]...^A[j-1]^A[j]^
// A[1]^A[2]^...^A[i-1]=
// (A[1]^...^A[j])^(A[1]...^A[i-1]) = G[j] ^ G[i-1]
// G[n] = A[1]^...^A[n]
// G[0] = 0
// G[1] = 1
// G[2] = 2
// G[3] = 2
// G[4] = 6
// G[5] = 7
// G[6] = 0
// G[7] = 0
func G(n int) int {
	switch n & 7 {
	case 0, 1:
		return n
	case 2, 3:
		return 2
	case 4, 5:
		return n + 2
	case 6, 7:
		return 0
	default:
		return 0
	}
}
