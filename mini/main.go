package mini

import (
	"wechat/mini/login"
	"wechat/mini/token"
)

type WeChat struct {
	appID  string
	secret string
}
func New(appID string, secret string) *WeChat {
	return &WeChat{
		appID,
		secret,
	}
}

//登录模块
func (this *WeChat) LoginMode() *login.Login {
	return login.New(this.appID,this.secret)
}

//令牌模块
func (this *WeChat) Token() *token.Token {
	return token.Instance(this.appID,this.secret)
}

type WeChatInfo struct {
	AppID  string
	Secret string
}
//获取信息
func (this *WeChat) GetInfo() *WeChatInfo {
	return &WeChatInfo{
		AppID:this.appID,
		Secret:this.secret,
	}
}