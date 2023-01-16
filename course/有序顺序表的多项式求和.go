package course

type Element struct {
	Ceof int //系数
	Exp  int //指数
}

func Calculate(arr1 []Element, arr2 []Element) []Element {
	up, down := 0, 0
	var merged []Element
	for up < len(arr1) && down < len(arr2) {
		if arr1[up].Exp < arr2[down].Exp {
			merged = append(merged, arr1[up])
			up++
		} else if arr1[up].Exp > arr2[down].Exp {
			merged = append(merged, arr2[down])
			down++
		} else if arr1[up].Exp == arr2[down].Exp {
			if arr1[up].Ceof+arr2[down].Ceof != 0 {
				merged = append(merged, Element{Ceof: arr1[up].Ceof + arr2[down].Ceof, Exp: arr1[up].Exp})
			}
			up++
			down++
		}
	}
	if up < len(arr1) {
		merged = append(merged, arr1[up:]...)
	} else if down < len(arr2) {
		merged = append(merged, arr2[down:]...)
	}
	return merged

}
