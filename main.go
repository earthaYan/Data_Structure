package main

import (
	"fmt"
	"leet_code_go/course"
)

func main() {

}
func init() {
	names := []course.Element{
		{
			Ceof: 1,
			Exp:  1,
		},
		{
			Ceof: 2,
			Exp:  2,
		},
		{
			Ceof: 2,
			Exp:  3,
		},
	}
	ages := []course.Element{
		{
			Ceof: 1,
			Exp:  2,
		},
		{
			Ceof: 1,
			Exp:  3,
		},
	}
	res := course.Calculate(names, ages)
	fmt.Println(res)
}
