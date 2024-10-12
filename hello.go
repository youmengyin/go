package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"quickstart/cmd"
	"time"
)

// func main() {
// 	fmt.Println("Hello World!")
// }

func parseParams() {
	var config string
	flag.StringVar(&config, "c", "", "config file hhh")

	//定义命令行参数方式1
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")

	//解析命令行参数
	flag.Parse()

	fmt.Printf("config file: %s\n", config)
	// 从系统环境变量读取值
	configEnv := os.Getenv("NUMBER_OF_PROCESSORS")

	// 当前目录的上层目录 返回绝对路径
	path, _ := filepath.Abs("..")
	fmt.Println(name, age, married, delay, configEnv, path)

	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}

func main() {
	// parseParams()
	cmd.Execute()
}
