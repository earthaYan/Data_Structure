package main

import (
	"fmt"
	"leet_code_go/course"
)

func main() {

}
func init() {
	names := []int{0, 1, 2}
	ages := []int{2, 3, 7}
	res := course.MergeArr(names, ages)
	fmt.Println(res)
}
