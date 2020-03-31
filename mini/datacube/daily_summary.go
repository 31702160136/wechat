package datacube

import (
	"errors"
	"fmt"
	"wechat/mini/utils"
)

type DailySummaryResult struct {
	List *[]struct {
		RefDate    string `json:"ref_date"`
		VisitTotal int    `json:"visit_total"`
		SharePv    int    `json:"share_pv"`
		ShareUv    int    `json:"share_uv"`
	} `json:"list"`
}

//获取用户访问小程序数据概况
var DailySummaryURL = "https://api.weixin.qq.com/datacube/getweanalysisappiddailysummarytrend?access_token=%s"

func (this *Datacube) GetDailySummary(accessToken, beginDate, endDate string) (*DailySummaryResult, error) {
	uri := fmt.Sprintf(DailySummaryURL, accessToken)

	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	data := make(map[string]interface{})
	data["begin_date"] = beginDate
	data["end_date"] = endDate
	resJson, err := utils.Post(uri, data, header)

	result := DailySummaryResult{}
	err = utils.Transfer(resJson, &result)
	if err != nil {
		return nil, errors.New("数据转换失败：" + err.Error())
	}
	return &result, nil
}
