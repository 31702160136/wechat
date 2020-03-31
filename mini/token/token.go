package token

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/31702160136/wechat/mini/utils"
	"time"
)

type Token struct {
	appID  string
	secret string
}

var wx = &Token{}

func Instance(appID, secret string) *Token {
	if wx.appID == "" || wx.secret == "" {
		return &Token{appID: appID, secret: secret}
	}
	return wx
}

type AccessTokenModel struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

//获取accessToken
var accessTokenUrl = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
var accessToken = &AccessTokenModel{}

func (this *Token) GetAccessToken() (string, error) {
	url := fmt.Sprintf(accessTokenUrl, this.appID, this.secret)
	//token有效时间>当前时间则token未过期
	if (accessToken.ExpiresIn) > time.Now().Unix() {
		return accessToken.AccessToken, nil
	}
	bt, err := utils.Get(url, map[string]string{})
	if err != nil {
		return "", err
	}
	if err = json.Unmarshal([]byte(bt), accessToken); err != nil {
		return "", err
	}
	switch accessToken.ErrCode {
	case 0:
		//token有效时间+当前时间-600秒
		accessToken.ExpiresIn += time.Now().Unix() - 600
		return accessToken.AccessToken, nil
	case -1:
		return "", errors.New("合法错误：系统繁忙，此时请开发者稍候再试")
	case 40001:
		return "", errors.New("AppSecret 错误或者 AppSecret 不属于这个小程序，请开发者确认 AppSecret 的正确性")
	case 40002:
		return "", errors.New("请确保 grant_type 字段值为 client_credential")
	case 40013:
		return "", errors.New("不合法的 AppID，请开发者检查 AppID 的正确性，避免异常字符，注意大小写")
	default:
		return "", errors.New(fmt.Sprintf("非法错误：code= %d , msg= %s", accessToken.ErrCode, accessToken.ErrMsg))
	}
}
