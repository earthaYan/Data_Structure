package main

import (
	"fmt"
	"leet_code_go/course"
)

func main() {

}
func init() {
	names := []int{0, 1, 2}
	res := course.BinarySearch(names, 10, 0, len(names)-1)
	fmt.Println(res)
}
