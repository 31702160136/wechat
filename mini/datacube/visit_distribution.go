package datacube

import (
	"errors"
	"fmt"
	"wechat/mini/utils"
)

type VisitDistributionResult struct {
	RefDate string `json:"ref_date"`
	List    *[]struct {
		Index    string        `json:"index"`
		ItemList *[]KeyValueV2 `json:"item_list"`
	} `json:"list"`
}

//获取用户小程序访问分布数据
var VisitDistribution = "https://api.weixin.qq.com/datacube/getweanalysisappidvisitdistribution?access_token=%s"

func (this *Datacube) GetVisitDistribution(accessToken, beginDate, endDate string) (*VisitDistributionResult, error) {
	uri := fmt.Sprintf(VisitDistribution, accessToken)

	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	data := make(map[string]interface{})
	data["begin_date"] = beginDate
	data["end_date"] = endDate
	resJson, err := utils.Post(uri, data, header)

	result := VisitDistributionResult{}
	err = utils.Transfer(resJson, &result)
	if err != nil {
		return nil, errors.New("数据转换失败：" + err.Error())
	}
	return &result, nil
}
