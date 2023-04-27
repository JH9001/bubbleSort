package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	pl("BUBBLE SORT!")

	arr := []int{9, 4, 2, 7, 1}

	pl(arr)

	for index := 0; index < len(arr); index++ {
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
	pl(arr)
}
