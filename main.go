package main

import (
	"fmt"
	"leet_code_go/course"
)

func main() {

}
func init() {
	name := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	res := course.Binary_Search(name, 0, len(name)-1,4)
	fmt.Println(res)
}
