package course

import "fmt"

func InsertEleInSpecificPosition1(arr []int, position int, ele int) []int {
	arrLen := len(arr)
	var newArr = make([]int, arrLen+1)
	copy(newArr, arr)
	for j := arrLen - 1; j >= position; j-- {
		newArr[j+1] = arr[j]
	}
	newArr[position] = ele
	fmt.Println(newArr)
	return newArr
}

func InsertEleInSpecificPosition2(arr []int, position int, ele int) []int {
	var newArr []int
	newArr = append(newArr, arr[:position]...)
	newArr = append(newArr, ele)
	newArr = append(newArr, arr[position:]...)
	return newArr
}
