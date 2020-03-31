package uniform_service_message

import (
	"errors"
	"fmt"
	"wechat/mini/utils"
)
type UniformServiceMessage struct {
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
	WeappTemplateMsg struct {
		TemplateID      string `json:"template_id"`
		Page            string `json:"page"`
		FormID          string `json:"form_id"`
		Data            string `json:"data"`
		EmphasisKeyword string `json:"emphasis_keyword"`
	} `json:"weapp_template_msg"`
	MpTemplateMsg struct {
		TemplateID  string `json:"template_id"`
		Url         string `json:"url"`
		MiniProgram string `json:"miniprogram"`
		Data        string `json:"data"`
	} `json:"mp_template_msg"`
}

//下发小程序和公众号统一的服务消息
var sendURL = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/uniform_send?access_token=%s"

func (this *UniformServiceMessage) Send(model SendModel) error {

	data:=map[string]interface{}{}
	MpTemplateMsg:=map[string]string{}
	_=utils.Transfer(model,&data)
	_=utils.Transfer(model.MpTemplateMsg,&MpTemplateMsg)
	MpTemplateMsg["appid"]=this.appID
	data["mp_template_msg"]=MpTemplateMsg

	uri:=fmt.Sprintf(sendURL,model.AccessToken)
	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	resJson,err:=utils.Post(uri,data,header)

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
		return errors.New("模板id不正确，weapp_template_msg.template_id或者mp_template_msg.template_id")
	case 41028:
		return errors.New("weapp_template_msg.form_id过期或者不正确")
	case 41029:
		return errors.New("weapp_template_msg.form_id已被使用")
	case 41030:
		return errors.New("weapp_template_msg.page不正确")
	case 45009:
		return errors.New("接口调用超过限额")
	case 40003:
		return errors.New("touser不是正确的openid")
	case 40013:
		return errors.New("appid不正确，或者不符合绑定关系要求")
	default:
		return errors.New(fmt.Sprintf("非法错误：code= %d , msg= %s", result.ErrCode, result.ErrMsg))
	}
}
