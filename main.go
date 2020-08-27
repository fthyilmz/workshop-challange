package main

import "fmt"

func main() {
	test :="abaab"
	k := 2
	n := len(test)
	c := make(map[string]int)

	for l := 1; l <= k; l++ {

		for i := 0; i <= (n-l); i++ {

			j:= i + l - 1
			var d string
			for k := i; k <= j; k++ {
				d = fmt.Sprintf("%s%s",d, test[k: (k + 1)])
			}

			if _, ok := c[d]; ok {
				c[d]++;
			} else {
				c[d] = 0;
			}

		}

	}

	for k,v := range c{

		if v <= 0 {
			delete(c, k)
		}

	}

	println(len(c))

}