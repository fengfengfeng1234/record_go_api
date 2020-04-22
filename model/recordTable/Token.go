package recordTable

import "github.com/jinzhu/gorm"

type Token struct {
	gorm.Model
	UserId      int64
	Token       string "json:token"
	CreateTime  int64
	InvalidTime int64
}

func (Token) TableName() string {
	return "token"
}
