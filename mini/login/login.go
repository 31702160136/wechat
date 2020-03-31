package login

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/31702160136/wechat/mini/utils"
)

type Login struct {
	appID  string
	secret string
}
var wx=&Login{}
func New(appID, secret string) *Login {
	if wx.appID == "" || wx.secret == "" {
		return &Login{appID:appID,secret:secret}
	}
	return wx
}


//登录
type LoginResult struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

var loginURL = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

func (this *Login) Login(jsCode string) (*LoginResult, error) {
	url := fmt.Sprintf(loginURL, this.appID, this.secret, jsCode)
	resultBT, err := utils.Get(url, map[string]string{})
	if err != nil {
		return nil, err
	}
	result := LoginResult{}
	err = json.Unmarshal([]byte(resultBT), &result)
	if err != nil {
		return nil, err
	}
	switch result.ErrCode {
	case 0:
		return &result, nil
	case -1:
		return nil, errors.New("合法错误：系统繁忙，此时请开发者稍候再试")
	case 40029:
		return nil, errors.New("合法错误：code 无效")
	case 45011:
		return nil, errors.New("合法错误：频率限制，每个用户每分钟100次")
	default:
		return nil, errors.New(fmt.Sprintf("非法错误：code= %d , msg= %s", result.ErrCode, result.ErrMsg))
	}
}
