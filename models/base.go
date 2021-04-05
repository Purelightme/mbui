package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"gopkg.in/redis.v4"
	"log"
	"strconv"
)

var Db *gorm.DB
var Redis *redis.Client

func init()  {
	var err error
	viper.AddConfigPath("config")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	viper.WatchConfig()
	var constr string
	host := viper.GetString("persist.host")
	port := viper.GetInt("persist.port")
	user := viper.GetString("persist.user")
	password := viper.GetString("persist.password")
	database := viper.GetString("persist.database")
	constr = fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	Db, err = gorm.Open("mysql", constr)
	if err != nil {
		log.Fatal(err)
	}

	Db.LogMode(true)

	Db.AutoMigrate(&Record{})

	redisHost := viper.GetString("redis.host")
	redisPort := viper.GetInt("redis.port")
	redisPassword := viper.GetString("redis.password")
	redisDatabase := viper.GetInt("redis.database")
	Redis = redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + strconv.Itoa(redisPort),
		Password: redisPassword,
		DB:       redisDatabase,
	})
}