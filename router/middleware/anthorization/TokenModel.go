package anthorization

type TokenModel struct {
	userId      uint32 `json:"userId"`
	token       string `json:"token"`
	createTime  int64  `json:"createTime"`
	invalidTime int64  `json:"invalidTime"`
}

func (TokenModel) TableName() string {
	return "token"
}
