package main

import "fmt"

func main() {
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)

	set1[1] = true
	set1[2] = true
	set1[3] = true
	set1[4] = true
	set1[5] = true

	set2[4] = true
	set2[5] = true

	for key := range set1 {
		if set2[key] {
			fmt.Println(key)
		}
	}
}
