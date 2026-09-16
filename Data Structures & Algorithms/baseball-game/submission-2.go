func calPoints(operations []string) int {
    stack := []int{}
    res := 0
    for _, op := range operations {
        if op == "+" {
            n := len(stack)
            newTop := stack[n-1] + stack[n-2]
            stack = append(stack, newTop)
            res += newTop
        } else if op == "D" {
            val := 2 * stack[len(stack)-1]
            stack = append(stack, val)
            res += val
        } else if op == "C" {
            res -= stack[len(stack)-1]
            stack = stack[:len(stack)-1]
        } else {
            num, _ := strconv.Atoi(op)
            stack = append(stack, num)
            res += num
        }
    }
    return res
}