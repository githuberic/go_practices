package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go_practices/go_gorm/db_conn"
	"net/http"
)

type Person struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	FirstName string `json:"firstname" gorm:"type:varchar(30);not null"`
	LastName  string `json:"lastname" gorm:"type:varchar(30);not null"`
	Mobile    string `json:"mobile" gorm:"type:varchar(20);not null"`
	City      string `json:"city" gorm:"type:varchar(50);not null"`
}

func init() {
	db_conn.MysqlInit()

	db_conn.DB.AutoMigrate(&Person{})
	if !db_conn.DB.HasTable(&Person{}) {
		if err := db_conn.DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Person{}).Error; err != nil {
			panic(err)
		}
	}
}

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

	create := db_conn.DB.Create(&person)
	if create.RowsAffected > 0 {
		c.JSON(http.StatusBadGateway, gin.H{
			"code": http.StatusOK,
			"msg":  "添加成功",
			"data": true,
		})
		return
	}
}

func GetPerson(c *gin.Context) {
	catchException(c)

	id := c.Params.ByName("id")

	var person Person
	if err := db_conn.DB.Where("id=?", id).First(&person).Error; err != nil {
		//c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "success",
			"data": person,
		})
	}
}

func GetMultiplePerson(c *gin.Context) {
	catchException(c)

	lastname := c.Params.ByName("lastname")

	persons := make([]Person, 10)
	tx := db_conn.DB.Where("last_name=?", lastname).Find(&persons).Limit(10)
	if tx.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadGateway,
			"msg":  "error",
			"data": tx.Error,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "success",
			"data": persons,
		})
	}
	return
}

func UpdatePerson(c *gin.Context) {
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
	fmt.Println(person)

	// 判定是否存在
	personExist := Person{}
	if err := db_conn.DB.First(&personExist, "id=?", person.ID).Error; err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"code": http.StatusOK,
			"msg":  "更新失败",
			"data": err,
		})
		return
	}

	updates := db_conn.DB.Model(&Person{}).Where("id=?", person.ID).Updates(&person)
	if updates.RowsAffected > 0 {
		c.JSON(http.StatusBadGateway, gin.H{
			"code": http.StatusOK,
			"msg":  "更新成功",
			"data": "ok",
		})
		return
	} else {
		c.JSON(http.StatusBadGateway, gin.H{
			"code": http.StatusOK,
			"msg":  "更新失败",
			"data": updates.Error,
		})
		return
	}
}

func DeletePerson(c *gin.Context) {
	catchException(c)

	id := c.Params.ByName("id")
	tx := db_conn.DB.Where("id=?", id).Delete(&Person{})
	if tx.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "删除成功",
			"data": "ok",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "删除失败",
			"data": tx.Error,
		})
	}
}

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
	r := gin.Default()
	r.POST("/person/create", CreatePerson)
	r.GET("/person/:id", GetPerson)
	r.GET("/person/multiple/:lastname", GetMultiplePerson)
	r.POST("/person/update", UpdatePerson)
	r.DELETE("/person/delete/:id", DeletePerson)
	r.Run(":8080")
}
