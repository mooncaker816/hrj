package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scanf("%s\n", &s)
	b := []byte(s)
	var oldnew []string
	for _, v := range b {
		oldnew = append(oldnew, string(v)+string(v), "")
	}
	r := strings.NewReplacer(oldnew...)
	out := reducePair(s, r)
	if out == " " || out == "" {
		out = "Empty String"
	}
	fmt.Println(out)
}

func reducePair(str string, r *strings.Replacer) string {
	s := r.Replace(string(str))
	if len(s) == len(str) {
		return s
	}
	return reducePair(s, r)
}
