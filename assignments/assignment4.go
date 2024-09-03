package main

import "slices"

func assignment4() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	pl(arr1)
	slices.Reverse(arr1)
	pl(arr1)
	slices.Reverse(arr1)

	evenArr := []int{}
	oddArr := []int{}

	for x := 0; x < len(arr1); x++ {
		if arr1[x]%2 == 0 {
			evenArr = append(evenArr, arr1[x])
		} else {
			oddArr = append(oddArr, arr1[x])
		}
	}
	pl(evenArr)
	pl(oddArr)
}
