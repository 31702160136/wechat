package customer_service_msg
type CustomerServiceMsg struct {
	appID  string
	secret string
}

type ErrorResult struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}