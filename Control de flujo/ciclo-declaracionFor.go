package main

import "fmt"

func main() {

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	j := 0

	for {
		if j > 9 {
			break
		}
		fmt.Println(j)
		j++

	}

}
