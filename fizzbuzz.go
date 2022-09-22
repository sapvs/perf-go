package perf

import "fmt"

func FizzbuzzV1(num int) []string {
	var ret []string
	for i := 1; i <= num; i++ {
		if i%15 == 0 {
			ret = append(ret, "fizzbuzz")
		} else if i%5 == 0 {
			ret = append(ret, "buzz")
		} else if i%3 == 0 {
			ret = append(ret, "fizz")
		} else {
			ret = append(ret, fmt.Sprint(i))
		}
	}

	return ret
}

func FizzbuzzV2(num int) []string {
	var ret []string
	for i := 1; i <= num; i++ {
		app := ""
		if i%3 == 0 {
			app += "fizz"
		}
		if i%5 == 0 {
			app += "buzz"
		}
		if app == "" {
			ret = append(ret, fmt.Sprint(i))
		} else {
			ret = append(ret, app)
		}
	}
	return ret
}

func FizzbuzzV3(num int) []string {
	var ret []string
	num3, num5 := 0, 0
	for i := 1; i <= num; i++ {
		app := ""
		num3++
		num5++
		if num3 == 3 {
			app += "fizz"
			num3 = 0
		}
		if num5 == 5 {
			app += "buzz"
			num5 = 0
		}
		if app == "" {
			ret = append(ret, fmt.Sprint(i))
		} else {
			ret = append(ret, app)
		}
	}
	return ret
}
