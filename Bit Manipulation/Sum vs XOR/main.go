package main

import (
	"fmt"
	"math/bits"
)

// 给定一个正整数n,有一个x(0<=x<=n)，使得n + x = n^x,求这样的x的个数
// 思路： 只要确保在n的基础上加上x后不会引发进位，所以只要求得n二进制中非前缀0的个数
// 在这些位数上，各有0，1两种选择，所以总共个数是pow(2,个数)

func main() {
	var n uint
	fmt.Scanf("%d\n", &n)
	fmt.Println(1 << uint(bits.Len(n)-bits.OnesCount(n)))
}
