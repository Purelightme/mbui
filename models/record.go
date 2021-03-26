package models

import (
	"database/sql/driver"
	"github.com/jinzhu/gorm"
)

//{
//    "database": "bookstore",
//    "table": "books",
//    "type": "update",
//    "ts": 1616145995,
//    "xid": 1608,
//    "commit": true,
//    "data": {
//        "id": 4,
//        "book": "PHP Practise",
//        "price": 10
//    },
//    "old": {
//        "book": "PHP AND Mysql3",
//        "price": 20
//    }
//}

//{
//    "type": "table-alter",
//    "database": "mbui",
//    "table": "qq",
//    "old": {
//        "database": "mbui",
//        "charset": "utf8mb4",
//        "table": "qq",
//        "columns": [
//            {
//                "type": "int",
//                "name": "id11",
//                "signed": false
//            }
//        ],
//        "primary-key": [
//            "id11"
//        ]
//    },
//    "def": {
//        "database": "mbui",
//        "charset": "utf8mb4",
//        "table": "aaa",
//        "columns": [
//            {
//                "type": "int",
//                "name": "id11",
//                "signed": false
//            }
//        ],
//        "primary-key": [
//            "id11"
//        ]
//    },
//    "ts": 1616728512000,
//    "sql": "RENAME TABLE `mbui`.`qq` TO `mbui`.`aaa`"
//}
type Record struct {
	gorm.Model
	Database string `json:"database"`
	Table    string `json:"table"`
	Type     string `json:"type"`
	Ts       int64  `json:"ts" gorm:"type:bigint(15)"`
	Xid      int    `json:"xid"`
	Commit   bool   `json:"commit"`
	Data     Data   `json:"data" gorm:"type:json"`
	Old      Data   `json:"old" gorm:"type:json"`
	Def      Data   `json:"def" gorm:"type:json"`
	Sql      string `json:"sql" gorm:"type:text"`
}

type Data string

func (t *Data) MarshalJSON() ([]byte, error) {
	return []byte(*t), nil
}

func (t *Data) UnmarshalJSON(data []byte) error {
	*t = Data(data)
	return nil
}

func (t Data) Value() (driver.Value, error) {
	//add this
	if string(t) == "" {
		return string("[]"), nil
	}
	return string(t), nil
}

func (t *Data) Scan(src interface{}) error {
	s, ok := src.([]byte)
	if !ok {
		return nil
	}
	*t = Data(s)
	return nil
}