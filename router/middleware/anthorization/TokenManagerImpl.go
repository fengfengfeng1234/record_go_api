package anthorization

import (
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"log"
	"record_go_api/global"
	anthorization "record_go_api/model/recordTable"
	"strings"
	"time"
)

/**
创建一个token关联上指定用户
@param userId 指定用户的id
@return 生成的token
*/
func CreateTokenModel(userId uint32) TokenModel {

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
	var model = TokenModel{
		userId, token, timestamp, invalidTime,
	}
	return model
}

/**
从字符串解析token
*/
func GetToken(authentication string) (model *anthorization.Token, err error) {
	if authentication == "" {
		return nil, errors.New("authentication  is null")
	}
	var info anthorization.Token
	err = global.GVA_DB.Unscoped().Last(&info, "token = ? ", authentication).Error
	if err != nil {
		return nil, errors.New("authentication  is null")
	}
	return &info, nil
}

/**
检查token是否有效
*/
func CheckToken(authentication string) (check bool, err error) {

	if authentication == "" {
		return false, errors.New("authentication  is null")
	}

	var info anthorization.Token
	err = global.GVA_DB.Unscoped().Last(&info, "token = ? ", authentication).Error
	if err != nil {
		return false, errors.New("not found result")
	}

	if info.InvalidTime < time.Now().Unix() {
		return false, errors.New("authentication 401 ")
	}

	return true, nil
}

func GetApiById(authentication string) (err error, api anthorization.Token) {
	err = global.GVA_DB.Where("token = ?", authentication).First(&api).Error
	return
}

/**
  清除 token
*/
func DeleteToken(userId uint32) bool {
	return true
}
