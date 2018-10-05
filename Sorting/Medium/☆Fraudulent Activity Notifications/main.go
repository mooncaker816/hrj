package main

import (
	"fmt"
)

func main() {
	// pprof.StartCPUProfile(os.Stdout)
	// defer pprof.StopCPUProfile()
	var n, d int
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &d)
	// fmt.Scanln()
	inputs := make([]int, n)
	for i := 0; i < int(n); i++ {
		var num int
		fmt.Scanf("%d", &num)
		inputs[i] = num
	}
	// fmt.Println(inputs)
	// start := time.Now()
	// fmt.Println(activityNotifications(inputs, d))
	// activityNotifications(inputs, d)
	fmt.Println(activityNotifications2(inputs, d))
	// fmt.Println()
	// fmt.Println(time.Since(start))
}
