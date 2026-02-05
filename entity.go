package go_nepay

type NePayInitParams struct {
	MerchantInfo `yaml:",inline" mapstructure:",squash"`

	ChannelCode       string `json:"channelCode" mapstructure:"channelCode" config:"channelCode"  yaml:"channelCode"` // 渠道码 固定的
	NotifyUrl         string `json:"notifyUrl" mapstructure:"notifyUrl" config:"notifyUrl"  yaml:"notifyUrl"`         // 通知地址
	ReturnUrl         string `json:"returnUrl" mapstructure:"returnUrl" config:"returnUrl"  yaml:"returnUrl"`         // 回调地址
	DepositUrl        string `json:"depositUrl" mapstructure:"depositUrl" config:"depositUrl"  yaml:"depositUrl"`
	WithdrawUrl       string `json:"withdrawUrl" mapstructure:"withdrawUrl" config:"withdrawUrl"  yaml:"withdrawUrl"`
	WithdrawNotifyUrl string `json:"withdrawNotifyUrl" mapstructure:"withdrawNotifyUrl" config:"withdrawNotifyUrl"  yaml:"withdrawNotifyUrl"`
}

type MerchantInfo struct {
	UserName  string `json:"username" mapstructure:"username" config:"username"  yaml:"username"`     // 商户号
	AccessKey string `json:"accessKey" mapstructure:"accessKey" config:"accessKey"  yaml:"accessKey"` // accessKey
}

//============================================================

// nepay入金
type NePayDepositReq struct {
	ChannelCode string `json:"channelCode" form:"channelCode" mapstructure:"channel_code"` // 渠道码 固定的
	UserName    string `json:"username" form:"username" mapstructure:"username"`           //商户orderNo
	Amount      string `json:"amount" form:"amount" mapstructure:"amount"`                 //订单金额
	OrderNumber string `json:"orderNumber" form:"orderNumber" mapstructure:"order_number"` //唯一订单号
	RealName    string `json:"realName" form:"realName" mapstructure:"real_name"`          //客户真实姓名，非必填
	ClientIp    string `json:"clientIp" form:"clientIp" mapstructure:"client_ip"`          //客户IP地址

	//这个不需要业务侧使用,而是sdk帮计算和补充
	//Sign        string `json:"sign" mapstructure:"sign"`               //签名
}

type NePayDepositRsp struct {
	HttpStatusCode int         `json:"httpStatusCode" mapstructure:"httpStatusCode"`
	Message        string      `json:"message" mapstructure:"message"`
	ErrorCode      string      `json:"errorCode" mapstructure:"errorCode"`
	Data           DepositData `json:"data" mapstructure:"data"`
}

type DepositData struct {
	UserName           string `json:"user_name" mapstructure:"user_name"`                 //商户号              //客户真实姓名，非必填
	Amount             string `json:"amount" mapstructure:"amount"`                       //金额（小数点后取2位）
	OrderNumber        string `json:"orderNumber" mapstructure:"order_number"`            //商户订单号            //唯一订单号
	SystemOrderNumber  string `json:"systemOrderNumber" mapstructure:"systemOrderNumber"` //平台订单号
	NotifyUrl          string `json:"notifyUrl" mapstructure:"notifyUrl"`                 //异步通知地址
	ReturnUrl          string `json:"returnUrl" mapstructure:"returnUrl"`                 //跳转URL
	CreatedAt          string `json:"createdAt" mapstructure:"createdAt"`                 //建立时间
	ConfirmedAt        string `json:"confirmedAt" mapstructure:"confirmedAt"`             //成功时间
	Status             int32  `json:"status" mapstructure:"status"`                       //订单状态：1、2、3、7、11 处理中 | 4、5 成功 | 6、8 失败
	Sign               string `json:"sign" mapstructure:"sign"`                           //签名
	CasherUrl          string `json:"casherUrl" mapstructure:"casherUrl"`                 //以下特定通道才有
	QrcodeUrl          string `json:"qrcodeUrl" mapstructure:"qrcodeUrl"`
	SchemeUrl          string `json:"schemeUrl" mapstructure:"schemeUrl"`
	ReceiverAccount    string `json:"receiverAccount" mapstructure:"receiverAccount"`
	ReceiverBankName   string `json:"receiverBankName" mapstructure:"receiverBankName"`
	ReceiverBankBranch string `json:"receiverBankBranch" mapstructure:"receiverBankBranch"`
	ReceiverName       string `json:"receiverName" mapstructure:"receiverName"`
	Note               string `json:"note" mapstructure:"note"`
}

// nepay出金
type NePayWithdrawReq struct {
	Amount             string `json:"amount" mapstructure:"amount"`                               //金额（小数点后取2位）
	OrderNumber        string `json:"order_number" mapstructure:"order_number"`                   //商户订单号            //唯一订单号
	BankCardHolderName string `json:"bank_card_holder_name" mapstructure:"bank_card_holder_name"` //银行账户名
	BankCardNumber     string `json:"bank_card_number" mapstructure:"bank_card_number"`           //银行账户号
	BankName           string `json:"bank_name" mapstructure:"bank_name"`                         //银行名称
	// BankProvince       string `json:"bank_province" mapstructure:"bank_province"`                 //银行省份，非必填
	// BankCity           string `json:"bank_city" mapstructure:"bank_city"`                         //银行城市，非必填
	// UserName           string `json:"username" mapstructure:"username"`                           //商户号
	// Sign string `json:"sign" mapstructure:"sign"` //签名NotifyUrl          string `json:"notify_url" mapstructure:"notify_url"`                       //异步通知地址
	// NotifyUrl          string `json:"notify_url" mapstructure:"notify_url"`                       //异步通知地址
}

type NePayWithdrawRsp struct {
	HttpStatusCode int          `json:"httpStatusCode" mapstructure:"httpStatusCode"`
	Message        string       `json:"message" mapstructure:"message"`
	ErrorCode      string       `json:"errorCode" mapstructure:"errorCode"`
	Data           WithdrawData `json:"data" mapstructure:"data"`
}

type WithdrawData struct {
	UserName           string `json:"username" mapstructure:"username"`                           //商户号
	Amount             string `json:"amount" mapstructure:"amount"`                               //金额（小数点后取2位）
	OrderNumber        string `json:"order_number" mapstructure:"order_number"`                   //商户订单号            //唯一订单号
	SystemOrderNumber  string `json:"system_order_number" mapstructure:"system_order_number"`     //平台订单号
	NotifyUrl          string `json:"notify_url" mapstructure:"notify_url"`                       //异步通知地址
	CreatedAt          string `json:"created_at" mapstructure:"created_at"`                       //建立时间
	ConfirmedAt        string `json:"confirmed_at" mapstructure:"confirmed_at"`                   //成功时间
	BankCardHolderName string `json:"bank_card_holder_name" mapstructure:"bank_card_holder_name"` //银行账户名
	BankCardNumber     string `json:"bank_card_number" mapstructure:"bank_card_number"`           //银行账户号
	BankName           string `json:"bank_name" mapstructure:"bank_name"`                         //银行名称
	BankProvince       string `json:"bank_province" mapstructure:"bank_province"`                 //银行省份，非必填
	BankCity           string `json:"bank_city" mapstructure:"bank_city"`                         //银行城市，非必填
	Sign               string `json:"sign" mapstructure:"sign"`                                   //签名
}

// 入金&出金回调
type NePayCallbackReq struct {
	HttpStatusCode int          `json:"httpStatusCode" form:"httpStatusCode" mapstructure:"httpStatusCode"`
	Message        string       `json:"message" form:"message" mapstructure:"message"`
	ErrorCode      int64        `json:"errorCode" form:"errorCode" mapstructure:"errorCode"`
	Data           CallbackData `json:"data" form:"data" mapstructure:"data"`
}

type CallbackData struct {
	UserName          string `json:"username" form:"username" mapstructure:"username"`                                  //商户号
	Amount            string `json:"amount" form:"amount" mapstructure:"amount"`                                        //金额（小数点后取2位）
	OrderNumber       string `json:"order_number" form:"order_number" mapstructure:"order_number"`                      //商户订单号            //唯一订单号
	SystemOrderNumber string `json:"system_order_number" form:"system_order_number" mapstructure:"system_order_number"` //平台订单号
	Status            int32  `json:"status" form:"status" mapstructure:"status"`                                        //订单状态：1、2、3、7、11 处理中 | 4、5 成功 | 6、8 失败
	Sign              string `json:"sign" form:"sign" mapstructure:"sign"`                                              //签名
}
