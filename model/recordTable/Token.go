package recordTable

import (
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"log"
	"record_go_api/global"
	"strings"
	"time"
)

type Token struct {
	UserId      int64
	Token       string "json:token"
	CreateTime  int64
	InvalidTime int64
}

func (Token) TableName() string {
	return "token"
}

/**
创建一个token关联上指定用户
@param userId 指定用户的id
@return 生成的token
*/
func (t Token) CreateTokenModel(userId int64) (*Token, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	token := strings.ReplaceAll(uuid.String(), "-", "")
	timestamp := time.Now().Unix()
	// 12天
	var addDay = fmt.Sprint(24*12, "h")
	day, _ := time.ParseDuration(addDay)
	invalidTime := time.Now().Add(day).Unix()
	var model = Token{
		UserId: userId, Token: token, CreateTime: timestamp, InvalidTime: invalidTime,
	}
	result, err := global.GVA_DB.Insert(&model)
	fmt.Println(result)
	return &model, nil
}

/**
从字符串解析token
*/
func (t Token) GetToken(authentication string) (model *Token, err error) {
	if authentication == "" {
		return nil, errors.New("authentication  is null")
	}
	_, error := global.GVA_DB.Where("token = ? ", authentication).Get(&t)
	if err != nil {
		return nil, error
	}
	return &t, nil
}

/**
检查token是否有效
*/
func (t Token) CheckToken(authentication string) (check bool, err error) {

	if authentication == "" {
		return false, errors.New("authentication  is null")
	}
	_, error := global.GVA_DB.Where("token = ? ", authentication).Limit(1).Desc("id").Get(&t)
	fmt.Println(t.InvalidTime)
	if err != nil {
		return false, error
	}

	if t.InvalidTime < time.Now().Unix() {
		return false, errors.New("authentication 401 ")
	}

	return true, nil
}

/**
  清除 token
*/
func (t Token) DeleteToken(userId int64) (check bool, err error) {
	_, err = global.GVA_DB.Where("user_id = ?", userId).Delete(&t)
	if err != nil {
		return false, err
	}
	return true, nil
}
