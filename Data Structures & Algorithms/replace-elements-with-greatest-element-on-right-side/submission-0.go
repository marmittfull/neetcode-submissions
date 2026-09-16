func replaceElements(arr []int) []int {
	for firstIdx := range arr {
		greatestElement := -1
		if firstIdx <= len(arr)-1 {
			for _, val := range arr[firstIdx+1:] {
				if val < greatestElement {
					continue
				}
				greatestElement = val
			}
			arr[firstIdx] = greatestElement
		}
	}
	return arr
}