func isAnagram(s string, t string) bool {
	caracteresS := make(map[rune]int)
	caracteresT := make(map[rune]int)

	for _, val := range s {
		caracteresS[val] += 1
	}
	for _, val := range t {
		caracteresT[val] += 1
	}
	if len(caracteresS) != len(caracteresT) {
		return false
	}
	for key, val := range caracteresS {
		if val != caracteresT[key] {
			return false
		}
	}
	return true
}