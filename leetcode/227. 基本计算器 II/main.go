package main

// 基本思路, 将正负号,放入栈用,最后累加
// 遇到前一个是*/就计算,计算好入栈,比如1+2*3-5+3, 当进入符号位-的时候,前一个是*,这个时候就进行计算,放入2*3=6
// 如果前一个是-/+,比如执行到最后一个+号,则将-5放入栈, 另外注意,最后一位,因为没有下一个符号了,也需要进入计算
func calculate(s string) int {
	prevSign := '+'

	stack := make([]int, 0)
	result := 0
	num := 0
	for i, ch := range s {
		isDigit := ch >= '0' && ch <= '9'

		if isDigit {
			num = num*10 + int(ch-'0')
		}

		if (!isDigit && ch != ' ') || i == len(s)-1 { // 最后没有符号了,也需要进入计算
			switch prevSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
			num = 0
			prevSign = ch
		}
	}

	for len(stack) > 0 {
		result += stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}

	return result
}
