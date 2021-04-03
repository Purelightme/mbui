package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"html/template"
	"log"
	"mbui/models"
	"net/http"
	"strconv"
)

//maxwell --user='root' --password='root' --host='127.0.0.1' --port=33060 --producer=redis --redis_type=lpush --output_ddl=true
//maxwell --user='root' --password='root' --host='127.0.0.1' --port=33060 --producer=stdout

var Db *gorm.DB

func init()  {
	var err error
	var constr string
	constr = fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "root", "127.0.0.1", 33060, "mbui")
	Db, err = gorm.Open("mysql", constr)
	if err != nil {
		log.Fatal(err)
	}

	Db.LogMode(true)

	Db.AutoMigrate(&models.Record{})
}

func main()  {
	r := gin.Default()

	r.SetHTMLTemplate(template.Must(template.ParseGlob("views/*")))

	v1 := r.Group("/v1")
	{
		v1.GET("/records",List)
		v1.GET("/",Index)
	}

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func List(c *gin.Context)  {
	recordDb := Db.Model(&models.Record{})

	//参数
	page := c.DefaultQuery("page","1")
	pageSize := c.DefaultQuery("pageSize","10")
	intPage,_ := strconv.Atoi(page)
	intPageSize,_ := strconv.Atoi(pageSize)
	database := c.Query("database")
	table := c.Query("table")
	qtype := c.Query("type")
	data := c.Query("data")
	old := c.Query("old")
	if database != "" {
		recordDb = recordDb.Where("database = ?",database)
	}
	if table != "" {
		recordDb = recordDb.Where("table = ?",table)
	}
	if qtype != "" {
		recordDb = recordDb.Where("type = ?",qtype)
	}
	if data != "" {
		recordDb = recordDb.Where("data","like","%"+data+"%")
	}
	if old != "" {
		recordDb = recordDb.Where("old","like","%"+old+"%")
	}

	var count int32
	recordDb.Count(&count)

	list := []models.Record{}
	recordDb.Order("id DESC").Offset((intPage - 1)*intPageSize).Limit(intPageSize).Find(&list)

	c.JSON(http.StatusOK,gin.H{
		"data": list,
		"total": count,
	})
}

func Index(c *gin.Context)  {
	c.HTML(http.StatusOK,"index.html",nil)
}