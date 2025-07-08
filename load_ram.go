package perf

import "fmt"

type loader struct {
	name string
}

func LoadRam() {
	ar := make([]*loader, 100)
	count := len(ar)
	for {
		l := &loader{name: "SomeName"}
		ar = append(ar, l)
		if count%10_000 == 0 {
			fmt.Printf("Created %d values", count)
		}
	}
}
