package perf

import "fmt"

type loader struct {
	name string
}

func LoadRam() {
	ar := make([]*loader, 100)
	count := 0
	for {
		l := &loader{name: "SomeName"}
		ar = append(ar, l)
		count++
		if count%10000 == 0 {
			fmt.Printf("Created %d values", count)
		}
	}
}
