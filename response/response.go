package response

type UnionOpenGoodsJingfenQueryResponse struct {
	Code       int32          `json:"code"`
	Message    string         `json:"message"`
	RequestId  string         `json:"requestId"`
	TotalCount int64          `json:"totalCount"`
	Data       []*JFGoodsResp `json:"data"`
}

// 数据明细
type JFGoodsResp struct {
	CategoryInfo          *CategoryInfo   `json:"categoryInfo"`          // 类目信息
	Comments              int64           `json:"comments"`              // 评论数
	CommissionInfo        *CommissionInfo `json:"commissionInfo"`        // 佣金信息
	CouponInfo            *CouponInfo     `json:"couponInfo"`            // 优惠券信息，返回内容为空说明该SKU无可用优惠券
	GoodCommentsShare     float64         `json:"goodCommentsShare"`     // 商品好评率
	ImageInfo             *ImageInfo      `json:"imageInfo"`             // 图片信息
	InOrderCount30Days    int64           `json:"inOrderCount30Days"`    // 30天内引单数量
	MaterialUrl           string          `json:"materialUrl"`           // 商品落地页
	PriceInfo             *PriceInfo      `json:"priceInfo"`             // 价格信息
	ShopInfo              *ShopInfo       `json:"shopInfo"`              // 店铺信息
	SkuId                 int64           `json:"skuId"`                 // 商品ID
	SkuName               string          `json:"skuName"`               // 商品名称
	IsHot                 uint8           `json:"isHot"`                 // 是否爆款，1：是，0：否
	Spuid                 float64         `json:"spuid"`                 // spuid，其值为同款商品的主skuid
	BrandCode             string          `json:"brandCode"`             // 品牌code
	BrandName             string          `json:"brandName"`             // 品牌名
	Owner                 string          `json:"owner"`                 // g=自营，p=pop
	PinGouInfo            *PinGouInfo     `json:"pinGouInfo"`            // 拼购信息
	ResourceInfo          *ResourceInfo   `json:"resourceInfo"`          // 资源信息
	InOrderCount30DaysSku int64           `json:"inOrderCount30DaysSku"` // 30天引单数量(sku维度)
	SeckillInfo           *SeckillInfo    `json:"seckillInfo"`           // 秒杀信息(可选）
	JxFlags               []int32         `json:"jxFlags"`               // 京喜商品类型，1京喜、2京喜工厂直供、3京喜优选（包含3时可在京东APP购买）(可选）
	VideoInfo             *VideoInfo      `json:"videoInfo"`             // 视频信息(可选）
	DocumentInfo          *DocumentInfo   `json:"documentInfo"`          // 段子信息(可选）
}

// 类目信息
type CategoryInfo struct {
	Cid1     int64  `json:"cid1"`     // 一级类目ID
	Cid1Name string `json:"cid1Name"` // 一级类目名称
	Cid2     int64  `json:"cid2"`     // 二级类目ID
	Cid2Name string `json:"cid2Name"` // 二级类目名称
	Cid3     int64  `json:"cid3"`     // 三级类目ID
	Cid4Name string `json:"cid3Name"` // 三级类目名称
}

// 佣金信息
type CommissionInfo struct {
	Commission          float64 `json:"commission"`          // 佣金
	CommissionShare     float64 `json:"commissionShare"`     // 佣金比例
	CouponCommission    float64 `json:"couponCommission"`    // 券后佣金(非必选)
	PlusCommissionShare float64 `json:"plusCommissionShare"` // plus佣金比例(即将上线)(非必选)
}

// 优惠券信息
type CouponInfo struct {
	CouponList []*Coupon `json:"couponList"` // 优惠券合集
}

// 优惠券
type Coupon struct {
	BindType     int32   `json:"bindType"`     // 券种类 (优惠券种类：0 - 全品类，1 - 限品类（自营商品），2 - 限店铺，3 - 店铺限商品券)
	Discount     float64 `json:"discount"`     // 券面额
	Link         string  `json:"link"`         // 券链接
	PlatformType int32   `json:"platformType"` // 券使用平台 (平台类型：0 - 全平台券，1 - 限平台券)
	Quota        float64 `json:"quota"`        // 券消费限额
	GetStartTime int64   `json:"getStartTime"` // 领取开始时间(时间戳，毫秒)
	GetEndTime   int64   `json:"getEndTime"`   // 券领取结束时间(时间戳，毫秒)
	UseStartTime int64   `json:"useStartTime"` // 券有效使用开始时间(时间戳，毫秒)
	UseEndTime   int64   `json:"useEndTime"`   // 券有效使用结束时间(时间戳，毫秒)
	IsBest       uint8   `json:"isBest"`       // 最优优惠券，1：是；0：否
	HotValue     uint8   `json:"hotValue"`     // 券热度，值越大热度越高，区间:[0,10]
}

// 图片信息
type ImageInfo struct {
	ImageList []*UrlInfo `json:"imageList"` // 图片合集
}

// 图片合集
type UrlInfo struct {
	Url string `json:"url"` // 图片链接地址，第一个图片链接为主图链接
}

// 价格信息
type PriceInfo struct {
	Price             float64 `json:"price"`             // 无线价格
	LowestPrice       float64 `json:"lowestPrice"`       // 最低价格(可选)
	LowestPriceType   uint8   `json:"lowestPriceType"`   // 最低价格类型，1：无线价格；2：拼购价格； 3：秒杀价格
	LowestCouponPrice float64 `json:"lowestCouponPrice"` // 最低价后的优惠券价(当商品无最优券时，不返回该字段)
}

// 店铺信息
type ShopInfo struct {
	ShopName  string  `json:"shopName"`  // 店铺名称
	ShopId    int64   `json:"shopId"`    // 店铺ID
	ShopLevel float64 `json:"shopLevel"` // 店铺评分
}

// 拼购信息
type PinGouInfo struct {
	PinGouPrice     float64 `json:"pingouPrice"`     // 拼购价格
	PinGouTmCount   int64   `json:"pingouTmCount"`   // 拼购成团所需人数
	PinGouUrl       string  `json:"pingouUrl"`       // 拼购落地页URL
	PinGouStartTime int64   `json:"pingouStartTime"` // 拼购开始时间(时间戳，毫秒)
	PinGouEndTime   int64   `json:"pingouEndTime"`   // 拼购结束时间(时间戳，毫秒)
}

// 资源信息
type ResourceInfo struct {
	EliteId   int32  `json:"eliteId"`   // 频道ID
	EliteName string `json:"eliteName"` // 频道名称
}

// 秒杀信息
type SeckillInfo struct {
	SeckillOriPrice  float64 `json:"seckillOriPrice"`  // 秒杀原价
	SeckillPrice     float64 `json:"seckillPrice"`     // 秒杀价
	SeckillStartTime int64   `json:"seckillStartTime"` // 秒杀开始时间(时间戳，毫秒)
	SeckillEndTime   int64   `json:"seckillEndTime"`   // 秒杀结束时间(时间戳，毫秒)
}

// 视频信息
type VideoInfo struct {
	VideoList []*Video `json:"videoList"` // 视频集合
}

// 视频明细
type Video struct {
	Width     int32  `json:"width"`     // 宽
	High      int32  `json:"high"`      // 高
	ImageUrl  string `json:"imageUrl"`  // 视频图片地址
	VideoType uint8  `json:"videoType"` // 1:主图，2：商详
	PlayUrl   string `json:"playUrl"`   // 播放地址
	PlayType  string `json:"playType"`  // low：标清，high：高清
	Duration  int32  `json:"duration"`  // // 时长(单位:s)
}

// 段子信息
type DocumentInfo struct {
	Document string `json:"document"` // 描述文案
	Discount string `json:"discount"` // 优惠力度文案
}
