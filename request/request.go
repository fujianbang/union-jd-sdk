package request

type Config struct {
	Version     string `url:"v"`
	Method      string `url:"method"`
	AccessToken string `url:"access_token"`
	AppKey      string `url:"app_key"`
	SignMethod  string `url:"sign_method"`
	Format      string `url:"format"`
	Timestamp   string `url:"timestamp"`
	Sign        string `url:"sign"`
	ParamJson   string `url:"param_json"`
}

// UnionOpenGoodsJingfenQuery 京粉精选商品查询接口
type UnionOpenGoodsJingfenQueryRequest struct {
	EliteId   string `json:"eliteId,omitempty"`   // 频道ID
	PageIndex string `json:"pageIndex,omitempty"` // 页码 默认1
	PageSize  string `json:"PageSize,omitempty"`  // 每页数量，默认20，上限50
	SortName  string `json:"sortName,omitempty"`  // 排序字段
	Sort      string `json:"sort,omitempty"`      // asc,desc升降序,默认降序
	Pid       string `json:"pid,omitempty"`       // 联盟id_应用id_推广位id，三段式
	Fields    string `json:"fields,omitempty"`    // 支持出参数据筛选，逗号','分隔，目前可用：videoInfo,documentInfo
}
