package main

import (
	"fmt"
	"math"
)

type Complex struct {
	a float64
	b float64
}

func Add(com1, com2 Complex) {
	fmt.Printf("%.2f", com1.a+com2.a)
	fmt.Printf("%.2fi\n", com1.b+com2.b)
}

func Sub(com1, com2 Complex) {
	fmt.Printf("%.2f", com1.a-com2.a)
	if com1.b-com2.b < 0 {
		fmt.Printf("%.2f", com1.b-com2.b)
	} else {
		fmt.Printf("+%.2fi\n", com1.b-com2.b)
	}
}

func Mul(com1, com2 Complex) {
	fmt.Printf("%.2f", com1.a*com2.a-com1.b*com2.b)
	if com1.a*com2.b+com2.a*com1.b > 0 {
		fmt.Printf("+")
	}
	fmt.Printf("%.2fi\n", com1.a*com2.b+com2.a*com1.b)
}

func Lenth(com1 Complex) float64 {
	fmt.Printf("%.2f\n", math.Sqrt(com1.a*com1.a+com1.b*com1.b))
	return math.Sqrt(com1.a*com1.a + com1.b*com1.b)
}

func Div(com1, com2 Complex) {
	fmt.Printf("%.2f", (com1.a*com2.a+com1.b*com2.b)/Lenth(com2))
	if (-com1.a*com2.b+com2.a*com1.b)/Lenth(com2) > 0 {
		fmt.Printf("+")
	}
	if -com1.a*com2.b+com2.a*com1.b != 0 {
		fmt.Printf("%.2fi\n", (-com1.a*com2.b+com2.a*com1.b)/Lenth(com2))
	}
}

func main() {
	fmt.Println("请输入虚数：")
	var com1, com2 Complex
	fmt.Printf("输入真值：")
	fmt.Scan(&com1.a)
	fmt.Printf("输入虚值：")
	fmt.Scan(&com1.b)
	fmt.Printf("输入真值：")
	fmt.Scan(&com2.a)
	fmt.Printf("输入虚值：")
	fmt.Scan(&com2.b)

	fmt.Printf("加减乘除")
	Add(com1, com2)
	Mul(com1, com2)
	Sub(com1, com2)
	Div(com1, com2)
	Lenth(com1)
	Lenth(com2)

}
