package perf

import "fmt"

func ReverseString(s string) string {
	for _, v := range s {
		fmt.Printf("v: %v  %T %s %d\n", v, v, string(v), len(string(v)))
	}
	return ""
}
