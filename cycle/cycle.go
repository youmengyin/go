package main

import "fmt"

func main()  {
	cycle_1()
	fmt.Println("--------")
	cycle_2()
	fmt.Println("--------")
	cycle_3()
	fmt.Println("--------")
	cycle_4()
	fmt.Println("--------")
	cycle_5()
}

// 循环array
func cycle_1()  {
	person := [4]string {"nangong", "kiku", "kyokyo"}
	fmt.Printf("len=%d, cap=%d, array=%v\n", len(person), cap(person), person)
	fmt.Println("")

	for i := 0; i < len(person); i++ {
		fmt.Printf("person[%d]: %s\n", i, person[i])
	}
	fmt.Println("")

	for k,v := range person {
		fmt.Printf("person[%d]: %s\n", k, v)
	}
	// 只取一位时默认为k,可以提供占位符_替换不需要的变量
}
func cycle_2()  {
	person := []string {"nangong", "kiku", "kyokyo"}
	fmt.Printf("len=%d, cap=%d, array=%v\n", len(person), cap(person), person)
	fmt.Println("")

	for i := 0; i < len(person); i++ {
		fmt.Printf("person[%d]: %s\n", i, person[i])
	}
	fmt.Println("")
	
	for k,v := range person {
		fmt.Printf("person[%d]: %s\n", k, v)
	}
	// 只取一位时默认为k,可以提供占位符_替换不需要的变量
}

// 循环map
func cycle_3()  {
	// person := []string {"nangong", "kiku", "kyokyo"}
	person := map[string]string{ "name": "nangong", "info": "..."}
	fmt.Printf("len=%d, map=%v\n", len(person), person)
	fmt.Println("")
	
	for k,v := range person {
		fmt.Printf("person[%s]: %s\n", k, v)
	}
	// 只取一位时默认为k,可以提供占位符_替换不需要的变量
}
// goto 跳到指定执行位置
func cycle_4()  {
	fmt.Println("begin")

	for i := 1; i <= 10; i++ {
		if i == 6 {
			goto END
		}
		fmt.Println("i =", i)
	}

	END :
		fmt.Println("end")
}

// switch fallthrough
// fallthrough执行完不退出，执行下一个case

func cycle_5()  {
	i := 3
	fmt.Printf("当 i = %d 时：\n", i)

	switch i {
		case 1:
			fmt.Println("输出 i =", 1)
		case 2:
			fmt.Println("输出 i =", 2)
		case 3:
			fmt.Println("输出 i =", 3)
			fallthrough
		case 4,5,6:
			fmt.Println("输出 i =", "4 or 5 or 6")
		default:
			fmt.Println("输出 i =", "xxx")
	}	
}