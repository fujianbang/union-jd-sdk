package response

type Result struct {
	Code       int32       `json:"code"`
	Message    string      `json:"message"`
	RequestId  string      `json:"requestId"`
	TotalCount int64       `json:"totalCount"`
	Data       interface{} `json:"data"`
}
