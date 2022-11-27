package main

import (
	"fmt"
	"time"
	"xorm.io/xorm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	userName string = "root"
	password string = "123456"
	ipAddress string = "127.0.0.1"
	port int = 3306
	db_name string = "go_blog"
	charset string = "utf8mb4"
)
func main()  {
	// root:123456@tcp(127.0.0.1:3306)/go_blog?charset=utf8mb4
	MysqlDataSources := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddress, port, db_name, charset)
	engine, err := xorm.NewEngine("mysql", MysqlDataSources)
	if err != nil {
		fmt.Println("--------数据库连接失败！")
		panic(err)
	}

	type User struct {
    Id int64
    Username string
    Password string `xorm:"varchar(200)"`
    Created time.Time `xorm:"created"`
    Updated time.Time `xorm:"updated"`
	}

	tableStructErr := engine.Sync(new(User))
	if tableStructErr != nil {
		fmt.Println("表结构同步失败")
		panic(tableStructErr)
	}

}
	