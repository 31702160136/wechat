package datacube

import (
	"errors"
	"fmt"
	"wechat/mini/utils"
)

type RetainResult struct {
	RefDate    string        `json:"ref_date"`
	VisitUvNew *[]KeyValueV2 `json:"visit_uv_new"`
	VisitUv    *[]KeyValueV2 `json:"visit_uv"`
}

//获取用户访问小程序日留存
var DailyRetainURL = "https://api.weixin.qq.com/datacube/getweanalysisappiddailyretaininfo?access_token=%s"

func (this *Datacube) GetDailyRetain(accessToken, beginDate, endDate string) (*RetainResult, error) {
	uri := fmt.Sprintf(DailyRetainURL, accessToken)

	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	data := make(map[string]interface{})
	data["begin_date"] = beginDate
	data["end_date"] = endDate
	resJson, err := utils.Post(uri, data, header)

	result := RetainResult{}
	err = utils.Transfer(resJson, &result)
	if err != nil {
		return nil, errors.New("数据转换失败：" + err.Error())
	}
	return &result, nil
}

//获取用户访问小程序月留存
var getMonthlyRetainURL = "https://api.weixin.qq.com/datacube/getweanalysisappidmonthlyretaininfo?access_token=%s"

func (this *Datacube) GetMonthlyRetain(accessToken, beginDate, endDate string) (*RetainResult, error) {
	uri := fmt.Sprintf(getMonthlyRetainURL, accessToken)

	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	data := make(map[string]interface{})
	data["begin_date"] = beginDate
	data["end_date"] = endDate
	resJson, err := utils.Post(uri, data, header)

	result := RetainResult{}
	err = utils.Transfer(resJson, &result)
	if err != nil {
		return nil, errors.New("数据转换失败：" + err.Error())
	}
	return &result, nil
}

//获取用户访问小程序周留存
var getWeeklyRetainURL = "https://api.weixin.qq.com/datacube/getweanalysisappidweeklyretaininfo?access_token=%s"

func (this *Datacube) GetWeeklyRetain(accessToken, beginDate, endDate string) (*RetainResult, error) {
	uri := fmt.Sprintf(getWeeklyRetainURL, accessToken)

	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	data := make(map[string]interface{})
	data["begin_date"] = beginDate
	data["end_date"] = endDate
	resJson, err := utils.Post(uri, data, header)

	result := RetainResult{}
	err = utils.Transfer(resJson, &result)
	if err != nil {
		return nil, errors.New("数据转换失败：" + err.Error())
	}
	return &result, nil
}
