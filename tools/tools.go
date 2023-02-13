package tools

import (
	"fmt"
	"io"
	"os"
)

/* 三元表达式 */
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

/* 打开文件 */
func OpenFile(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	if f != nil {
		defer func(f io.Closer) {
			if err = f.Close(); err != nil {
				fmt.Printf("defer close file %v %v\n", fileName, err)
			}
		}(f) // 利用了defer 先求值，再延迟调用 的性质，先在注册时指定要关闭哪个文件，函数结束时就会正确关闭这个文件。这个方法在打开多个文件后很有用。
	}

	return nil
}
