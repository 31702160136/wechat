package datacube

import (
	"errors"
	"fmt"
	"wechat/mini/utils"
)

type visit struct {
	Province  *[]NameValueV3 `json:"province"`
	City      *[]NameValueV3 `json:"city"`
	Genders   *[]NameValueV3 `json:"genders"`
	Platforms *[]NameValueV3 `json:"platforms"`
	Devices   *[]NameValueV2 `json:"devices"`
	Ages      *[]NameValueV3 `json:"ages"`
}
type UserPortraitResult struct {
	RefDate    string `json:"refDate"`
	VisitUvNew *visit  `json:"visitUvNew"`
	VisitUv    *visit  `json:"visitUv"`
}

// 获取小程序新增或活跃用户的画像分布数据。
// 时间范围支持昨天、最近7天、最近30天。
// 其中，新增用户数为时间范围内首次访问小程序的去重用户数，活跃用户数为时间范围内访问过小程序的去重用户数。
var UserPortraitURL = "https://api.weixin.qq.com/datacube/getweanalysisappiduserportrait?access_token=%s"

func (this *Datacube) GetUserPortrait(accessToken, beginDate, endDate string) (*UserPortraitResult, error) {
	uri := fmt.Sprintf(UserPortraitURL, accessToken)

	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	data := make(map[string]interface{})
	data["begin_date"] = beginDate
	data["end_date"] = endDate
	resJson, err := utils.Post(uri, data, header)

	result := UserPortraitResult{}
	err = utils.Transfer(resJson, &result)
	if err != nil {
		return nil, errors.New("数据转换失败：" + err.Error())
	}
	return &result, nil
}
