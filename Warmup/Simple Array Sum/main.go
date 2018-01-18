package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	r.ReadString('\n')
	in2, _ := r.ReadString('\n')
	s := strings.Fields(string(in2))
	sum := 0
	exp := ""
	for i, v := range s {
		vi, _ := strconv.Atoi(v)
		sum += vi
		if i > 0 {
			exp += "+"
		}
		exp += v
	}
	fmt.Printf("%s = %d\n", exp, sum)
}

/* [Min]
Given an array of  integers, can you find the sum of its elements?

Input Format

The first line contains an integer, , denoting the size of the array.
The second line contains  space-separated integers representing the array's elements.

Output Format

Print the sum of the array's elements as a single integer.

Sample Input

6
1 2 3 4 10 11
Sample Output

31
Explanation

We print the sum of the array's elements, which is: .
*/
