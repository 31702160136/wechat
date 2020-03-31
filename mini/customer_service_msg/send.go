package customer_service_msg

import (
	"errors"
	"fmt"
	"github.com/31702160136/wechat/mini/utils"
	"reflect"
)

//发送文本消息
//发送文本消息时，支持添加可跳转小程序的文字连接
//文本内容...<a href="http://www.qq.com" data-miniprogram-appid="appid" data-miniprogram-path="pages/index/index">点击跳小程序</a>
type SendTextModel struct {
	AccessToken string `json:"access_token"`
	Touser      string `json:"touser"`
	Msgtype     string `json:"msgtype"`
	Text        struct {
		Content string `json:"content"`
	} `json:"text"`
}
//发送图片消息
type SendImageModel struct {
	AccessToken string `json:"access_token"`
	Touser      string `json:"touser"`
	Msgtype     string `json:"msgtype"`
	Image       struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
}
//发送图文链接
type SendLinkModel struct {
	AccessToken string `json:"access_token"`
	Touser      string `json:"touser"`
	Msgtype     string `json:"msgtype"`
	Link        struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Url         string `json:"url"`
		ThumbUrl    string `json:"thumb_url"`
	} `json:"link"`
}
//发送小程序卡片
type SendMiniProgramPageModel struct {
	AccessToken     string `json:"access_token"`
	Touser          string `json:"touser"`
	Msgtype         string `json:"msgtype"`
	MiniProgramPage struct {
		Title        string `json:"title"`
		PagePath     string `json:"pagepath"`
		ThumbMediaID string `json:"thumb_media_id"`
	} `json:"miniprogrampage"`
}

//发送客服消息给用户。详细规则见 发送客服消息
var sendURL = "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s"

func (this *CustomerServiceMsg) Send(model interface{}) error {
	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	uri := ""
	resJson := ""
	var err error
	tp := getType(model)
	switch tp {
	case "SendTextModel":
		obj := model.(SendTextModel)
		obj.Msgtype="text"
		uri = fmt.Sprintf(sendURL, obj.AccessToken)
		resJson, err = utils.Post(uri, obj, header)
	case "SendImageModel":
		obj := model.(SendImageModel)
		obj.Msgtype="image"
		uri = fmt.Sprintf(sendURL, obj.AccessToken)
		resJson, err = utils.Post(uri, obj, header)
	case "SendLinkModel":
		obj := model.(SendLinkModel)
		obj.Msgtype="link"
		uri = fmt.Sprintf(sendURL, obj.AccessToken)
		resJson, err = utils.Post(uri, obj, header)
	case "SendMiniProgramPageModel":
		obj := model.(SendMiniProgramPageModel)
		obj.Msgtype="miniprogrampage"
		uri = fmt.Sprintf(sendURL, obj.AccessToken)
	}

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
//得到参数类型
func getType(model interface{}) string {
	tp := reflect.TypeOf(model)
	return tp.Name()
}
