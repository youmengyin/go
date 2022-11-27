package main

import (
	"fmt"
	"time"
)

func main() {
	defer_1()
}

// defer 执行顺序按定义顺序逆序输出
func defer_1() {
	go_2()
	time.Sleep(1 * time.Second)
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func go_1() {
	panic("error")
}
func go_2() {
	defer (func() {
		if errs := recover(); errs != nil {
			fmt.Println("panic:" + fmt.Sprintf("%s", errs))
		}
	})()
	panic("error")
	// go go_1()
}
