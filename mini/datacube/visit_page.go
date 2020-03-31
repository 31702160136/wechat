package datacube

import (
	"errors"
	"fmt"
	"wechat/mini/utils"
)

type VisitPageResult struct {
	PagePath       string `json:"page_path"`
	PageVisitPv    int    `json:"page_visit_pv"`
	PageVisitUv    int    `json:"page_visit_uv"`
	PageStaytimePv int    `json:"page_staytime_pv"`
	EntrypagePv    int    `json:"entrypage_pv"`
	ExitpagePv     int    `json:"exitpage_pv"`
	PageSharePv    int    `json:"page_share_pv"`
	PageShareUv    int    `json:"page_share_uv"`
}

//访问页面。目前只提供按 page_visit_pv 排序的 top200。
var VisitPageURL = "https://api.weixin.qq.com/datacube/getweanalysisappidvisitpage?access_token%s"

func (this *Datacube) GetVisitPage(accessToken, beginDate, endDate string) (*VisitPageResult, error) {
	uri := fmt.Sprintf(VisitPageURL, accessToken)

	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	data := make(map[string]interface{})
	data["begin_date"] = beginDate
	data["end_date"] = endDate
	resJson, err := utils.Post(uri, data, header)

	result := VisitPageResult{}
	err = utils.Transfer(resJson, &result)
	if err != nil {
		return nil, errors.New("数据转换失败：" + err.Error())
	}
	return &result, nil
}
