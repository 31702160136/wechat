package subscribe_msg

import (
	"errors"
	"fmt"
	"wechat/mini/utils"
)
type Subscribe struct {
	appID  string
	secret string
}

type ErrorResult struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
type SendModel struct {
	AccessToken      string `json:"access_token"`
	ToUser           string `json:"touser"`
	TemplateID       string `json:"template_id"`
	Page             string `json:"page"`
	MiniProgramState string `json:"miniprogram_state"`
	Lang             string `json:"lang"`
}

//下发小程序和公众号统一的服务消息
var sendURL = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=%s"

func (this *Subscribe) Send(model SendModel,data interface{}) error {

	body := map[string]interface{}{}
	_ = utils.Transfer(model, &body)
	body["data"] = data

	uri := fmt.Sprintf(sendURL, model.AccessToken)
	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	resJson, err := utils.Post(uri, body, header)

	if err != nil {
		return errors.New("发送失败：" + err.Error())
	}
	result := ErrorResult{}
	err = utils.Transfer(resJson, &result)
	if err != nil {
		return errors.New("数据转换失败：" + err.Error())
	}
	switch result.ErrCode {
	case 0:
		return nil
	case 40037:
		return errors.New("订阅模板id为空不正确")
	case 43101:
		return errors.New("用户拒绝接受消息，如果用户之前曾经订阅过，则表示用户取消了订阅关系")
	case 47003:
		return errors.New("模板参数不准确，可能为空或者不满足规则，errmsg会提示具体是哪个字段出错")
	case 41030:
		return errors.New("page路径不正确，需要保证在现网版本小程序中存在，与app.json保持一致")
	case 40003:
		return errors.New("touser字段openid为空或者不正确")
	default:
		return errors.New(fmt.Sprintf("非法错误：code= %d , msg= %s", result.ErrCode, result.ErrMsg))
	}
}
