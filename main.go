package main

import (
	"fmt"
	"leet_code_go/LeetCode_GoLang"
)

func main() {

}
func init() {
	names := []string{"Mary", "John", "Emma"}
	heights := []int{155, 185, 150}
	res := LeetCode_GoLang.SortPeople(names, heights)
	fmt.Print(res)
}
