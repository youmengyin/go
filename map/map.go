package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	map_1()
	fmt.Println("-------")
	map_2()
	fmt.Println("-------")
	map_3()
}

func map_1()  {
	// 声明 Map
	var p1 map[int]string;
	p1 = make(map[int]string)
	p1[1] = "null"
	p1[0] = "0"
	fmt.Println(p1)

	// 合并
	var p2 map[string]string = make(map[string]string)
	p2["name"] = "nangong"
	fmt.Println(p2)

	var p3 map[int]string = map[int]string{1:"gh"}
	p3[2] = "gg"
	fmt.Println(p3)

	// 也可以简写
	p4 := map[string]string{ "name": "kiku", "sex": "femal"}
	fmt.Println(p4)
}
// 生成json数据
func map_2()  {
	res := make(map[string]interface{})
	res["statusCode"] = 200
	res["message"] = "success"
	res["object"] = map[string]interface{}{
		"username": "admin",
		"token": "fjrerekkkkkkkekemw",
		"lastUpdateTime": "2011-20-20 30:12:22",
		"tags": []string{"阴阳师","原神"},
	}

	fmt.Println(res)
	fmt.Println("")
	// 序列化
	jsons, errs := json.Marshal(res)
	if errs != nil {
		fmt.Println("errors: ", errs)
	}
	fmt.Println("")
	fmt.Println("json: ", string(jsons))

	//反序列化
	res2 := make(map[string]interface{})
	errs = json.Unmarshal([]byte(jsons), &res2)
	if errs != nil {
		fmt.Println("errors: ", errs)
	}
	fmt.Println("")
	fmt.Println("map: ", res2)
}

// 编辑和删除
func map_3()  {
	res := make(map[string]interface{})
	res["statusCode"] = 200
	res["message"] = "success"
	res["object"] = map[string]interface{}{
		"username": "admin",
		"token": "fjrerekkkkkkkekemw",
		"lastUpdateTime": "2011-20-20 30:12:22",
		"tags": []string{"阴阳师","原神"},
	}
	fmt.Println(res)
	fmt.Println("")

	res["statusCode"] = 500
	fmt.Println("after update: ", res)
	fmt.Println("")

	delete(res, "message")
	fmt.Println("after delete: ", res)
	fmt.Println("")

}