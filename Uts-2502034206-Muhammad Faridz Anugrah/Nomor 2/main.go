package main

import "fmt"

func compare(array1 [3]string, array2 [3]string) {
	for i := 0; i < len(array1); i++ {
		if array1[i] != array2[i] {
			fmt.Printf("index ke %d berbeda\n", i)
		}
	}
}

func main() {
	array1 := [3]string{"a", "c", "d"}
	array2 := [3]string{"a", "d", "c"}

	compare(array1, array2)
}