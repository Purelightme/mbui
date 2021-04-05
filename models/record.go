package models

import (
	"database/sql/driver"
	"github.com/jinzhu/gorm"
)

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
	Query    string `json:"query" gorm:"type:text"`
}

type Data string

//func (t *Data) MarshalJSON() ([]byte, error) {
//	return []byte(*t), nil
//}

func (t *Data) UnmarshalJSON(data []byte) error {
	*t = Data(data)
	return nil
}

func (t Data) Value() (driver.Value, error) {
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