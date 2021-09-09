package protocol

// 创建账号
type AccountCreateRequest struct {
	BaseRequest
	DevId           string
	Version         int64
	IMEI            int64
	AccessIP        string
	MarketId        string
	UserType        string
	AdvertisementId string
	OSType          string
	OSVersion       string
}
type AccountCreateResponse struct {
	BaseResponse
}

// 验证账号
type AccountAuthRequest struct {
	BaseRequest
	// 客户端的版本信息
	Version  int64
	DevId    string
	IMEI     int64
	AccessIP string
	// gps = Google Play Store, aas = Apple App Store, one = One Store
	MarketId string
	// 用户类型(GM、tester、AI、BOT等 普通用户为NULL)
	UserType string
	// platform_device_id_type
	AdvertisementId string
	// iOS,Android
	OSType    string
	OSVersion string
	// 设备固有id (unity的systeminfo deviceuniqueidentifier)。
	DeviceUniqueId string
}
type AccountAuthResponse struct {
	BaseResponse
}

//
type AccountGetTutorialRequest struct {
	BaseRequest
}
type AccountGetTutorialResponse struct {
	BaseResponse
	TutorialIds []int64
}

//
type AccountSetTutorialRequest struct {
	BaseRequest
	TutorialIds []int64
}
type AccountSetTutorialResponse struct {
	BaseResponse
}

//
type AccountPassCheckRequest struct {
	BaseRequest
	DevId string
}
type AccountPassCheckResponse struct {
	BaseResponse
}

// ???
type AccountLinkRewardRequest struct {
	BaseRequest
}
type AccountLinkRewardResponse struct {
	BaseResponse
}

// 检查悠星账号
type AccountCheckYostarRequest struct {
	BaseRequest
	UID         int64
	YostarToken string
	// 可忽略yostar认证的结果。
	PassCheckYostarServer bool
	EnterTicket           string
}
type AccountCheckYostarResponse struct {
	BaseResponse
	ResultState   int
	ResultMessage string
	Birth         string
}
