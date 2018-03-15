package gumi

func StringBackSpace(str string, count int) string {
	temp := []rune(str)
	templen := len(temp)
	if count > templen {
		count = templen
	}
	return string(temp[:templen-count])
}
func StringControlBackSpace(str string, count int) (res string) {
	temp := []rune(str)
	to := len(temp) - 1
	if to < 0 {
		return ""
	}
	//
	for i := 0 ; i < count || to != 0; i ++{
		if temp[to] == ' ' {
			for i := to; i >= 0; i-- {
				to = i
				if temp[i] != ' ' {
					break
				}
			}
		}
		for i := to; i >= 0; i-- {
			to = i
			if temp[i] == ' ' {
				to += 1
				break
			}
		}
	}
	return string(temp[:to])
}
