func calPoints(operations []string) int {
	stack := []int{}
	var totalSum, num int
	for _, val := range operations {
		if val == "+" {
			num = stack[len(stack) - 1] + stack[len(stack) - 2]
			stack = append(stack, num)
		} else if val == "D" {
			num = stack[len(stack) - 1] * 2
			stack = append(stack, num)
		} else if val == "C" {
			num = stack[len(stack) - 1] * -1
			stack = stack[:len(stack) - 1]
		} else {
			num, _ = strconv.Atoi(val)
			stack = append(stack, num)
		}
		totalSum += num
	}
	return totalSum
}
