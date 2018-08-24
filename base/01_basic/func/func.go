package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// + - * / 运算
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf(
			"unsupported operation: %s", op)
	}
}

// 带余除法，注意：go的除法(/)默认不带余数
func div(a, b int) (q, r int) {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+
		"(%d, %d)\n", opName, a, b)

	return op(a, b)
}

// 可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

// 交换参数
func swap(a, b int) (int, int) {
	return b, a
}

func swapPointer(a, b *int) {
	*b, *a = *a, *b
}

func main() {
	fmt.Println("Error handling")
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	q, r := div(13, 3)
	fmt.Printf("13 div 3 is %d mod %d\n", q, r)

	// 这里有用到匿名函数
	fmt.Println("pow(3, 4) is:", apply(
		func(a int, b int) int {
			return int(math.Pow(
				float64(a), float64(b)))
		}, 3, 4))


	fmt.Println("1+2+...+5 =", sum(1, 2, 3, 4, 5))

	// 指针的用法
	// a, b := 3, 4
	// swapPointer(&a, &b)
	// fmt.Println(a, b)

	a, b := 3, 4
	a, b = swap(a, b)
	fmt.Println("a, b after swap is:", a, b)
}
