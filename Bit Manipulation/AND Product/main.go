package main

import (
	"fmt"
	"math/bits"
)

//保留相同的高位，后续置零
func main() {
	var t int
	fmt.Scanf("%d\n", &t)
	for i := 0; i < t; i++ {
		var a, b, res uint32
		fmt.Scanf("%d", &a)
		fmt.Scanf("%d", &b)
		// //fmt.Printf("%b,%b\n", a, b)
		// la, lb := bits.Len32(a), bits.Len32(b)
		// for la == lb {
		// 	a &^= 1 << uint(la-1)
		// 	b &^= 1 << uint(lb-1)
		// 	//fmt.Printf("%b,%b\n", a, b)
		// 	res |= 1 << uint(la-1)
		// 	la, lb = bits.Len32(a), bits.Len32(b)
		// }
		res = ^(1<<uint(bits.Len32(a^b)) - 1) & a
		fmt.Println(res)
	}
}

//^(1<<uint(bits.Len32(a^b))-1)&a
// 1. a^b 将a和b左边相同的高位置零
// 2. 取比a^b大的最小的2的指数幂-1，即将a除去相同高位的部分都置1
// 3. 取反恢复原先相同高位部分，并将前缀0置1
// 4. 与a做与运算，得到相同高位不变后续置零的数
