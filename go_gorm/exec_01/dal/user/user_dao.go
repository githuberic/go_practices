package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Id      int    `gorm:"primary_key" json:"id"`
	Name    string `json:"name"`
	Address string `json:"string"`
	Mobile  string `json:"mobile"`
}

func dbConn(User, Password, Host, Db string, Port int) string {
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&autocommit=true&parseTime=True&loc=Local", User, Password, Host, Port, Db)
	return connArgs
}

func main() {
	conn := dbConn("root", "root_mysql", "127.0.0.1", "db_user", 3307)
	//db, err := gorm.Open("mysql", "root:@(localhost:3307)/root_mysql")
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		fmt.Printf("error %v", err)
		//panic("failed to connect database,%v", err)
	}

	db.SingularTable(true) //如果使用gorm来帮忙创建表时，这里填写false的话gorm会给表添加s后缀，填写true则不会
	db.LogMode(true)       //打印sql语句
	//开启连接池
	db.DB().SetMaxIdleConns(100)   //最大空闲连接
	db.DB().SetMaxOpenConns(10000) //最大连接数
	db.DB().SetConnMaxLifetime(30) //最大生存时间(s)

	//关闭数据库连接
	defer db.Close()

	//创建表
	db.AutoMigrate(&User{})

	// 插入
	db.Create(&User{Name: "L1213", Address: "杭州", Mobile: "13588827425"})

	// 读取
	var user User
	db.First(&user, 100)                   // 查询id为1的user
	//db.First(&user, "name = ?", "L1213") // 查询code为l1212的user

	// 更新
	db.Model(&user).Update("Mobile", "13034214214")
}
