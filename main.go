package main

import (
	"fmt"

	"github.com/AlexanderGuan/GoTools/tools"
)

func main() {
	a, b := 2, 3
	max := tools.If(a > b, a, b)
	fmt.Println(max)
}
