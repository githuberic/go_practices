package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"strings"
)

var db *gorm.DB
var err error

type Person struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	FirstName string `json:"firstname" gorm:"type:varchar(30);not null"`
	LastName  string `json:"lastname" gorm:"type:varchar(30);not null"`
	Mobile    string `json:"mobile" gorm:"type:varchar(20);not null"`
	City      string `json:"city" gorm:"type:varchar(50);not null"`
}

func NewConn() *gorm.DB {
	const (
		conn = "mysql://root:root_mysql@tcp(127.0.0.1:3307)/data_center?autocommit=true&charset=utf8mb4"
	)
	DBEngine := strings.Replace(conn, "mysql://", "", -1)
	db, err := gorm.Open("mysql", DBEngine)
	if err != nil {
		panic("DB connection fail:" + err.Error())
	}
	return db
}

/*
func init() {
	const (
		conn = "mysql://root:root_mysql@tcp(127.0.0.1:3307)/data_center?autocommit=true&charset=utf8"
	)
	DBEngine := strings.Replace(conn, "mysql://", "", -1)
	db, err := gorm.Open("mysql", DBEngine)
	if err != nil {
		panic("failed to connect database")
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// 启用Logger，显示详细日志
	db.LogMode(true)
}*/

func Dbinit() *gorm.DB {
	db := NewConn()

	//SetMaxOpenConns用于设置最大打开的连接数
	//SetMaxIdleConns用于设置闲置的连接数
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// 启用Logger，显示详细日志
	db.LogMode(true)

	// 自动迁移模式
	//db.AutoMigrate(&Model.UserModel{},
	//	&Model.UserDetailModel{},
	//	&Model.UserAuthsModel{},
	//)
	return db
}

/*
func DeletePerson(c *gin.Context) {
	id := c.Params.ByName("id")
	var person Person
	db = Dbinit()
	d := db.Where("id = ?", id).Delete(&person)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func UpdatePerson(c *gin.Context) {
	var person Person
	id := c.Params.ByName("id")

	db = Dbinit()
	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&person)

	db.Save(&person)
	c.JSON(200, person)
}
*/

func CreatePerson(c *gin.Context) {
	catchException(c)

	var person Person
	err := c.BindJSON(&person)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"code": http.StatusBadGateway,
			"msg":  "参数错误",
			"data": err,
		})
		return
	}

	db := Dbinit()
	create := db.Create(&person)
	if create.RowsAffected > 0 {
		c.JSON(http.StatusBadGateway, gin.H{
			"code": http.StatusOK,
			"msg":  "添加成功",
			"data": true,
		})
		return
	}

	/*
		db = Dbinit()
		db.Create(&person)
		c.JSON(200, person)*/
}

/*
func GetPerson(c *gin.Context) {
	id := c.Params.ByName("id")

	var person Person
	db = Dbinit()
	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, person)
	}
}

func GetPeople(c *gin.Context) {
	var people []Person
	db = Dbinit()
	if err := db.Find(&people).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, people)
	}
}
*/
func catchException(c *gin.Context) {
	// 捕获异常
	defer func() {
		err := recover()
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"code": http.StatusBadGateway,
				"msg":  err,
				"data": err,
			})
		}
	}()
}

func main() {
	db := Dbinit()

	db.AutoMigrate(&Person{})
	if !db.HasTable(&Person{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Person{}).Error; err != nil {
			panic(err)
		}
	}

	r := gin.Default()
	//localhost:8080/people {"id":4,"firstname":"Elvis","lastname":"Presley","city":"beijing"}
	r.POST("/people/create", CreatePerson)

	/*
		//localhost:8080/people
		r.GET("/people/", GetPeople)
		//localhost:8080/people/1
		r.GET("/people/:id", GetPerson)

		//localhost:8080/people {"id":4,"firstname":"Elvis","lastname":"Presley","city":"beijing"}
		r.PUT("/people/:id", UpdatePerson)
		//localhost:8080/people/1
		r.DELETE("/people/:id", DeletePerson)
	*/
	r.Run(":8080")
}
