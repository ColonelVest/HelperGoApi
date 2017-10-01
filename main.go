package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main()  {
	router := gin.Default()
	v1 := router.Group("/api/v1/tasks")
	{
		v1.GET("/", FetchAllTasks)
		//v1.GET("/:id", FetchSingeTask)
	}

	router.Run()
	fmt.Println("123")
}

func FetchAllTasks(c *gin.Context)  {
	var dbTasks []Task

	db := getDbConnection()
	db.Find(&dbTasks)

	if len(dbTasks) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No todo found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "data" : dbTasks})
}

func getDbConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:44834631@/helper_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("Error connecting to DB: %s", err)
		panic("failed to connect DB!!!")
	}
	return db
}

type Task struct {
	gorm.Model
	Title string
	Description string
}