package customer_service_msg

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/31702160136/wechat/mini/utils"
)


type TempMediaResult struct {
	ContentType string       `json:"content_type"`
	Buffer      bytes.Buffer `json:"Buffer"`
	ErrCode     int          `json:"errCode"`
	ErrMsg      string       `json:"errMsg"`
}
//获取客服消息内的临时素材。即下载临时的多媒体文件。目前小程序仅支持下载图片文件。
var TempMediaURL = "https://api.weixin.qq.com/cgi-bin/media/get?access_token=%s&media_id=%s"
func (this *CustomerServiceMsg) GetTempMedia(accessToken,mediaID string) (*TempMediaResult,error) {
	uri:=fmt.Sprintf(TempMediaURL,accessToken,mediaID)
	resJson, err := utils.Get(uri, map[string]string{})
	result := TempMediaResult{}
	err = utils.Transfer(resJson, &result)
	if err != nil {
		return nil, errors.New("数据转换失败：" + err.Error())
	}
	return &result,nil
}
