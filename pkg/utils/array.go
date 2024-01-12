package utils

func InArrayStr(src string, arr []string) bool {
	for _, v := range arr {
		if src == v {
			return true
		}
	}
	return false
}

func InArrayInt(d int, arr []int) bool {
	for _, v := range arr {
		if d == v {
			return true
		}
	}
	return false
}
