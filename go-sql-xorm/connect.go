package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var (
	userName      string = "root"
	password      string = "123456"
	passwordProd  string = "Mcw-1340739521"
	ipAddress     string = "127.0.0.1"
	ipAddressProd string = "101.132.70.183"
	port          int    = 3306
	db_name       string = "go_blog"
	charset       string = "utf8mb4"
)

func main() {
	// root:123456@tcp(127.0.0.1:3306)/go_blog?charset=utf8mb4
	// MysqlDataSources := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddress, port, db_name, charset)
	MysqlProdDataSources := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, passwordProd, ipAddressProd, port, db_name, charset)
	// engine, err := xorm.NewEngine("mysql", MysqlDataSources)
	engine, err := xorm.NewEngine("mysql", MysqlProdDataSources)
	if err != nil {
		fmt.Println("--------数据库连接失败！")
		panic(err)
	}

	type User struct {
		Id       int64
		Username string
		Avatar   string
		Password string    `xorm:"varchar(200)"`
		Created  time.Time `xorm:"created"`
		Updated  time.Time `xorm:"updated"`
	}

	tableStructErr := engine.Sync(new(User))
	if tableStructErr != nil {
		fmt.Println("表结构同步失败")
		panic(tableStructErr)
	}

	// insert
	// affected, err := engine.Table("user").Insert([]map[string]interface{}{
	// 	{
	// 		"Username": "admin",
	// 		"Password": "123456",
	// 	},
	// 	{
	// 		"Username": "luo",
	// 		"Password": "123456",
	// 	},
	// })
	// fmt.Println("sql结果", affected)
	// results, err := engine.QueryInterface("select * from user")
	results, err := engine.Where("id = 2").Query()
	fmt.Println("sql结果", results)
}
