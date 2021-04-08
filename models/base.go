package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gopkg.in/redis.v4"
	"log"
	"strconv"
)

var Db *gorm.DB
var Redis *redis.Client

func init()  {
	parseFlag()
	var err error
	viper.AddConfigPath("config")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	viper.WatchConfig()
	var constr string
	host := viper.GetString("mysql.host")
	port := viper.GetInt("mysql.port")
	user := viper.GetString("mysql.user")
	password := viper.GetString("mysql.password")
	database := viper.GetString("mysql.database")
	constr = fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	Db, err = gorm.Open("mysql", constr)
	if err != nil {
		log.Fatal(err)
	}

	//Db.LogMode(true)

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

func parseFlag()  {
	pflag.String("mysql.host","127.0.0.1","mysql host")
	pflag.Int("mysql.port",3306,"mysql port")
	pflag.String("mysql.user","root","user for mysql")
	pflag.String("mysql.password","root","password from mysql")
	pflag.String("mysql.database","mbui","mysql database")
	pflag.String("redis.host","127.0.0.1","redis host")
	pflag.Int("redis.port",6379,"redis port")
	pflag.String("redis.password","","redis auth")
	pflag.Int("redis.database",0,"redis producer database")
	pflag.String("redis.key","maxwell","redis producer list key name")
	pflag.Int("web.port",8080,"web service port")
	pflag.Bool("debug",false,"run gin in debug mode?")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}