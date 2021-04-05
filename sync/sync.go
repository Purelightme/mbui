package sync

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"mbui/models"
	"time"
)

func Start()  {
	key := viper.GetString("redis.key")
	for {
		item,err := models.Redis.RPop(key).Result()
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

		models.Db.Create(&record)

		time.Sleep(time.Second)
	}
}