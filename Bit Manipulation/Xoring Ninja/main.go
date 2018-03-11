package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scanf("%d\n", &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scanf("%d\n", &n)
		ord := uint(0)
		for j := 0; j < n; j++ {
			var m uint
			fmt.Scanf("%d", &m)
			ord |= m
		}

		var res int
		pos := 0
		for ord > 0 {
			if ord&1 == 1 {
				res = (res + calculate(pos, n)) % 1000000007
			}
			pos++
			ord >>= 1
		}
		fmt.Println(res)
	}
}

func calculate(pos, n int) int {
	x := 1 << uint(pos)
	for i := 0; i < n-1; i++ {
		x = 2 * x % 1000000007
	}
	return x
}

/* [Min]
参考：
Understand like this. Consider ith bit of all inputs. Count the number of 1's and the number of 0's.
So suppose at the ith bit we have K 1's and N-K 0's.
Then out of all 2^n subsets, those subsets which have odd number of 1's will give 1 as XOR output at the ith bit and the other subsets will give zero.
Now you can prove that independent of K, exactly 2^(n-1) subsets will have odd number of ones. (provided K>0)

Proof : KC1*2^(N-K) + KC3*2^(N-K) + ... = 2^(K-1) * 2^(N-K) = 2^(N-1)



Consider one bit position at a time. How many of the terms have bit i set? The terms that have bit i set are exactly those that correspond to a subset that contains an odd number of inputs that have bit i set.

If there is any input that has bit i set, then exactly half of the 2^N possible subsets will be of this form, and so they will contribute 2^(N−1+i) to the final sum.

On the other hand, if no input has bit i set, then of course no terms will have that bit set either.

Summing these contributions of 2^(N−1+i) per bit position is easy enough -- the final sum will simply be 2^(N−1) times the bitwise OR of all the inputs.
*/
