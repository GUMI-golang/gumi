package eventline

func textComplete(completed string, editing rune) string {
	if editing == 0{
		return completed
	}
	return completed + string(editing)
}
