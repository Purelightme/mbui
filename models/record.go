package models

import (
	"database/sql"
	"encoding/json"
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
type Record struct {
	gorm.Model
	Database string `json:"database"`
	Table string `json:"table"`
	Type string `json:"type"`
	Ts int `json:"ts"`
	Xid int `json:"xid"`
	Commit bool `json:"commit"`
	Data Data `json:"data" gorm:"type:json"`
	Old MyNullString `json:"old" gorm:"type:json default:null"`
}

type Data string

func (t *Data) MarshalJSON() ([]byte, error) {
	return []byte(*t), nil
}

func (t *Data) UnmarshalJSON(data []byte) error {
	*t = Data(data)
	return nil
}

type MyNullString struct {
	sql.NullString
}

func (ns MyNullString) MarshalJSON()(b []byte, err error) {
	if ns.String == "" && !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}