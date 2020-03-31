package datacube

import (
	"errors"
	"fmt"
	"github.com/31702160136/wechat/mini/utils"
)

type AccessTrendResult struct {
	List *[]struct {
		RefDate         string  `json:"ref_date"`
		SessionCnt      int     `json:"session_cnt"`
		VisitPv         int     `json:"visit_pv"`
		VisitUv         int     `json:"visit_uv"`
		VisitUvNew      int     `json:"visit_uv_new"`
		StayTimeUv      float64 `json:"stay_time_uv"`
		StayTimeSession float64 `json:"stay_time_session"`
		VisitDepth      float64 `json:"visit_depth"`
	} `json:"list"`
}

//获取用户访问小程序数据日趋势
var DailyVisitTrendURL = "https://api.weixin.qq.com/datacube/getweanalysisappiddailyvisittrend?access_token=%s"

func (this *Datacube) GetDailyVisitTrend(accessToken, beginDate, endDate string) (*AccessTrendResult, error) {
	uri := fmt.Sprintf(DailyVisitTrendURL, accessToken)

	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	data := make(map[string]interface{})
	data["begin_date"] = beginDate
	data["end_date"] = endDate
	resJson, err := utils.Post(uri, data, header)

	result := AccessTrendResult{}
	err = utils.Transfer(resJson, &result)
	if err != nil {
		return nil, errors.New("数据转换失败：" + err.Error())
	}
	return &result, nil
}

//获取用户访问小程序数据月趋势
var MonthlyVisitTrend = "https://api.weixin.qq.com/datacube/getweanalysisappidmonthlyvisittrend?access_token=%s"

func (this *Datacube) GetMonthlyVisitTrend(accessToken, beginDate, endDate string) (*AccessTrendResult, error) {
	uri := fmt.Sprintf(MonthlyVisitTrend, accessToken)

	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	data := make(map[string]interface{})
	data["begin_date"] = beginDate
	data["end_date"] = endDate
	resJson, err := utils.Post(uri, data, header)

	result := AccessTrendResult{}
	err = utils.Transfer(resJson, &result)
	if err != nil {
		return nil, errors.New("数据转换失败：" + err.Error())
	}
	return &result, nil
}

//获取用户访问小程序数据周趋势
var WeeklyVisitTrend = "https://api.weixin.qq.com/datacube/getweanalysisappidweeklyvisittrend?access_token=%s"

func (this *Datacube) GetWeeklyVisitTrend(accessToken, beginDate, endDate string) (*AccessTrendResult, error) {
	uri := fmt.Sprintf(WeeklyVisitTrend, accessToken)

	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	data := make(map[string]interface{})
	data["begin_date"] = beginDate
	data["end_date"] = endDate
	resJson, err := utils.Post(uri, data, header)

	result := AccessTrendResult{}
	err = utils.Transfer(resJson, &result)
	if err != nil {
		return nil, errors.New("数据转换失败：" + err.Error())
	}
	return &result, nil
}
