func isValid(s string) bool {
	var openers []rune
	for _, val := range s {
		if val == '{' || val == '[' || val == '(' {
			openers = append(openers, val)
		} else if len(openers) > 0{
			resultString := string([]rune{openers[len(openers) - 1], val})	
			if resultString == "{}" || resultString == "[]" || resultString == "()" {
				openers = openers[:len(openers) - 1]
			} else {
				return false
			}
		} else {
			return false
		}
	}
	if len(openers) > 0 {
		return false
	}
	return true
}
