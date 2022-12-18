package main

import (
	"fmt"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main()  {
	// dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:123456@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize: 171,
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "qx_",   // table name prefix, table for `User` would be `t_users`
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase: false, // skip the snake_casing of names
			NameReplacer: strings.NewReplacer("CID", "cid"), // use name replacer to change struct/field name before convert it to db name
		},
		DisableForeignKeyConstraintWhenMigrating: true,// 自动创建外键约束
	})
	// User表
	type User struct {
		Username string;
		Password int8;
		CID int8;
	}
	// 自动迁移
	_ = db.AutoMigrate(&User{})
	fmt.Println("db", db)
	if err != nil {
		fmt.Println(err)
	}
}