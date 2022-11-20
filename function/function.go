package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

func main()  {
	function_1()
	fmt.Println("--------")
	
	fmt.Printf("MD5(%s): %s\n", "123456", function_2("123456"))
	fmt.Println("--------")
	function_3()
	fmt.Println("--------")

	params := map[string]interface{} {
		"name" : "Tom",
		"pwd"  : "123456",
		"age"  : 30,
	}
	fmt.Printf("sign : %s\n", function_4(params))
}
type Result struct {
	StatusCode int "json:StatusCode"
	Object string "json:Object"
}

func function_1()  {
	var res Result
	res.StatusCode    = 200
	res.Object = "success"
	toJson(&res)
	
	setData(&res)
	toJson(&res)	
}
func setData (res *Result) {
	res.StatusCode    = 500
	res.Object = "fail"
}

func toJson (res *Result) {
	jsons, errs := json.Marshal(res)
	if errs != nil {
			fmt.Println("json marshal error:", errs)
	}
	fmt.Println("json data :", string(jsons))
}

func function_2(str string) string {
	sercet := md5.New()
	// fmt.Println(sercet)
	sercet.Write([]byte(str))
	
	return hex.EncodeToString(sercet.Sum(nil))
}
func function_3()  {
	// 当前时间戳
	t := time.Now().Unix()
	// 当前时间字符串
	formatTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(t, formatTime)
}
func function_4(params map[string]interface{})  string{
	var key []string
	var str = ""
	for k, _ := range params {
		key = append(key, k)
	}
	sort.Strings(key)
	fmt.Println("key ==>", key)
	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v", key[i], params[key[i]])
		} else {
			str = str + fmt.Sprintf("&xl_%v=%v", key[i], params[key[i]])
		}
	}
	fmt.Println("str: ", str)
	// 自定义密钥
	var secret = "123456789"

	// 自定义签名算法
	sign := function_2(function_2(str) + function_2(secret))
	return sign
}
