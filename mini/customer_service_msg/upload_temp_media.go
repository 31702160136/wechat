package customer_service_msg

import (
	"errors"
	"fmt"
	"github.com/31702160136/wechat/mini/utils"
)

type UploadTempMediaResult struct {
	ErrorResult
	Type      string `json:"type"`
	MediaID   string `json:"media_id"`
	CreatedAt int64  `json:"created_at"`
}

//获取客服消息内的临时素材。即下载临时的多媒体文件。目前小程序仅支持下载图片文件。
var UploadTempMediaURL = "https://api.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s"

func (this *CustomerServiceMsg) UploadTempMedia(accessToken,path string) (*UploadTempMediaResult, error) {

	uri := fmt.Sprintf(UploadTempMediaURL, accessToken, "image")
	files:=map[string]string{}
	files["media"]=path
	resJson, err := utils.PostFormData(uri,files, map[string]string{})
	result := UploadTempMediaResult{}
	err = utils.Transfer(resJson, &result)
	if err != nil {
		return nil, errors.New("数据转换失败：" + err.Error())
	}
	if result.ErrCode==40004 {
		return nil,errors.New("无效媒体文件类型")
	}
	return &result, nil
}
