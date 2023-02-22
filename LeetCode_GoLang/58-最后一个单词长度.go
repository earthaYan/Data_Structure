package LeetCode_GoLang

import (
	"strings"
)

/*
*
获取最后一个飞控单词的长度
**
*/
func LengthOfLastWord(s string) int {
	strArr := strings.Split(s, " ")
	var targetArr []string
	for _, val := range strArr {
		if val != "" {
			targetArr = append(targetArr, val)
		}
	}
	return len(targetArr[len(targetArr)-1])
}
