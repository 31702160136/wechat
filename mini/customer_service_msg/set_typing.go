package customer_service_msg

import (
	"errors"
	"fmt"
	"zytool/common/sdk/wechat/wxtools"
)

type Typing struct {
	AccessToken string `json:"access_token"`
	ToUser      string `json:"touser"`
	Command     string `json:"command"`//可选值：1.Typing 对用户下发"正在输入"状态、2.CancelTyping 取消对用户的"正在输入"状态
}
//下发客服当前输入状态给用户。详见 客服消息输入状态
var setTyping = "https://api.weixin.qq.com/cgi-bin/message/custom/typing?access_token=%s"
func (this *CustomerServiceMsg) SetTyping(model Typing) error {
	uri := fmt.Sprintf(setTyping, model.AccessToken)
	header := make(map[string]string)
	header["Content-Type"] = "application/json"

	resJson, err := wxtools.Post(uri, model, header)
	if err != nil {
		return errors.New("发送失败：" + err.Error())
	}

	result := ErrorResult{}
	err = wxtools.Transfer(resJson, &result)
	if err != nil {
		return errors.New("数据转换失败：" + err.Error())
	}
	switch result.ErrCode {
	case 0:
		return nil
	case -1:
		return errors.New("合法错误：系统繁忙，此时请开发者稍候再试")
	case 40001:
		return errors.New("获取 access_token 时 AppSecret 错误，" +
			"或者 access_token 无效。请开发者认真比对 AppSecret 的正确性，" +
			"或查看是否正在为恰当的小程序调用接口")
	case 40002:
		return errors.New("不合法的凭证类型")
	case 40003:
		return errors.New("不合法的 OpenID，请开发者确认 OpenID 是否是其他小程序的 OpenID")
	case 45015:
		return errors.New("回复时间超过限制")
	case 45047:
		return errors.New("客服接口下行条数超过上限")
	case 48001:
		return errors.New("API 功能未授权，请确认小程序已获得该接口")
	default:
		return errors.New(fmt.Sprintf("非法错误：code= %d , msg= %s", result.ErrCode, result.ErrMsg))
	}
}