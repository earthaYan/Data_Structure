package course

func DelEleInCeratinPosition1(arr []int, position int) {
	arrLen := len(arr)
	for j := position; j < arrLen-1; j++ {
		arr[j] = arr[j+1]
	}
	if position < arrLen {
		arr = arr[0 : len(arr)-1]
	}
}


func DelEleInCeratinPosition2(arr []int, position int) []int{
	if(position<len(arr)){
		arr=append(arr[:position],arr[position+1:]...)
	}
	
	return arr
}