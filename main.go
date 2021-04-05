package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"html/template"
	"log"
	"mbui/models"
	"mbui/sync"
	"net/http"
	"strconv"
)

func main()  {

	go sync.Start()

	r := gin.Default()

	r.SetHTMLTemplate(template.Must(template.ParseGlob("views/*")))

	v1 := r.Group("/v1")
	{
		v1.GET("/records",List)
		v1.GET("/",Index)
	}

	err := r.Run(":" + strconv.Itoa(viper.GetInt("web.port")))
	if err != nil {
		log.Fatal(err)
	}
}

func List(c *gin.Context)  {
	recordDb := models.Db.Model(&models.Record{})

	//参数
	page := c.DefaultQuery("page","1")
	pageSize := c.DefaultQuery("pageSize","10")
	intPage,_ := strconv.Atoi(page)
	intPageSize,_ := strconv.Atoi(pageSize)
	database := c.Query("database")
	table := c.Query("table")
	qtype := c.Query("type")
	sql := c.Query("sql")
	if database != "" {
		recordDb = recordDb.Where("`database` = ?",database)
	}
	if table != "" {
		recordDb = recordDb.Where("`table` = ?",table)
	}
	if qtype != "" {
		recordDb = recordDb.Where("`type` = ?",qtype)
	}
	if sql != "" {
		recordDb = recordDb.Where("`query`","like","%"+sql+"%")
	}

	var count int32
	recordDb.Count(&count)

	var list []models.Record
	recordDb.Order("id DESC").Offset((intPage - 1)*intPageSize).Limit(intPageSize).Find(&list)

	log.Println(list)

	c.JSON(http.StatusOK,gin.H{
		"data": list,
		"total": count,
	})
}

func Index(c *gin.Context)  {
	c.HTML(http.StatusOK,"index.html",nil)
}