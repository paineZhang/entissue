package util

// 三项表达式
func Ternary(expression bool, trueResult interface{}, falseResult interface{}) interface{} {
	if expression {
		return trueResult
	}
	return falseResult
}
