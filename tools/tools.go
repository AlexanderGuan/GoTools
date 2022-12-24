package tools

/* 三元表达式 */
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
