package main

import (
	"fmt"
)

func main() {
	const (
		a = iota    // 0	 iota 0
		b           // 1	 iota 1
		c           // 2	 iota 2
		d = "Xaiat" // Xaiat iota 3
		e           // Xaiat iota 4
		f = 100     // 100	 iota 5
		g           // 100	 iota 6
		h = iota    // 7	 iota 7
		i           // 8	 iota 8
	)

	const (
		j = iota // 0	 iota 0
		k        // 1	 iota 1
	)
	fmt.Println("a = ", a, "b = ", b, "c = ", c, "d = ", d, "e = ", e, "f = ", f, "g = ", g, "h = ", h, "i = ", i, "j = ", j, "k = ", k)
}
