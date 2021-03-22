package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/redis.v4"
	"log"
	"mbui/models"
	"time"
)

var Db *gorm.DB
var Redis *redis.Client

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

	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func main()  {
	for {
		item,err := Redis.RPop("maxwell").Result()
		if err != nil {
			log.Println(err)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println(item)

		//入库
		record := models.Record{}
		err = json.Unmarshal([]byte(item),&record)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(record.Old)

		Db.Create(&record)

		time.Sleep(time.Second)
	}
}
