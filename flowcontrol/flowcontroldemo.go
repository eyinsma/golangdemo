package flowcontrol

import (
	"fmt"
	"time"
	"os"
)

func init() {
	fmt.Println("package fc init")
}

func DemoIf() {
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}


func DemoGoto() {
	i := 0
	Here:
	fmt.Println(i)
	i++
	if i < 2 {
		goto Here
	}
}

func DemoFor() {
	i := 1
	for i < 3 {
		fmt.Println(i)
		i++
	}

	for {
		fmt.Println("loop")
		break
	}
}

func DemoSwitch() {
	v := "a"
	switch v {
	case "b":
		fmt.Println("b")
	case "a":
		fmt.Println("a")
	}

	switch v {
	case "b":
		fmt.Println("b")
	case "a":
		fmt.Println("a")
		fallthrough
	case "c":
		fmt.Println("c")
	default:
		fmt.Println("default")
	}

	switch  {
	case false:
		fmt.Println("false")
	case true:
		fmt.Println("true")
	default:
		fmt.Println("default")
	}


	switch  {
	case false:
		fmt.Println("false")
	default:
		fmt.Println("default")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("before 12")
	default:
		fmt.Println("after 12")
	}
}

func DemoVarPara(arg ... int) {
	for _, n := range arg {
		fmt.Println("number:", n)
	}
}

func DemoDefer() {
	f, err := os.Open("not")
	defer f.Close()

	fmt.Println("error:", err)
	fmt.Println("file:", f)

	for i := 0; i < 5; i++ {
		defer fmt.Println("v:", i)
	}
}

type testInt func(int) bool

func isOdd(integer int) bool {
	if integer % 2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer % 2 == 0 {
		return true
	}

	return false
}

func filter(ints []int, f testInt) []int {
	var result []int
	for _, value := range ints {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func DemoFuncPara() {
	slice := []int{1, 2, 3, 4}
	odd := filter(slice, isOdd)
	fmt.Println("slice:", odd)
}