package main

import "fmt"

func main() {
	var x1, v1, x2, v2 int
	fmt.Scanf("%d %d %d %d\n", &x1, &v1, &x2, &v2)
	if v1 > v2 && (x2-x1)%(v1-v2) == 0 {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}
