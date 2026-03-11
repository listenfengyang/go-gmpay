package go_gmpay

type GmPayInitParams struct {
	MerchantInfo `yaml:",inline" mapstructure:",squash"`

	DepositUrl        string `json:"depositUrl" mapstructure:"depositUrl" config:"depositUrl"  yaml:"depositUrl"`                             // 入金地址
	WithdrawUrl       string `json:"withdrawUrl" mapstructure:"withdrawUrl" config:"withdrawUrl"  yaml:"withdrawUrl"`                         // 出金地址
	ReturnUrl         string `json:"returnUrl" mapstructure:"returnUrl" config:"returnUrl"  yaml:"returnUrl"`                                 // 回调地址
	DepositNotifyUrl  string `json:"depositNotifyUrl" mapstructure:"depositNotifyUrl" config:"depositNotifyUrl"  yaml:"depositNotifyUrl"`     // 入金回调地址
	WithdrawNotifyUrl string `json:"withdrawNotifyUrl" mapstructure:"withdrawNotifyUrl" config:"withdrawNotifyUrl"  yaml:"withdrawNotifyUrl"` // 出金回调地址
}

type MerchantInfo struct {
	ApiKey    string `json:"apiKey" mapstructure:"apiKey" config:"apiKey"  yaml:"apiKey"`             // apiKey
	SecretKey string `json:"secretKey" mapstructure:"secretKey" config:"secretKey"  yaml:"secretKey"` // secretKey
}

//============================================================

// gmpay入金
type GmPayDepositReq struct {
	ApiKey          string `json:"api_key" mapstructure:"api_key" config:"api_key"  yaml:"api_key"`                                         // apiKey
	RefNo           string `json:"ref_no" mapstructure:"ref_no" config:"ref_no"  yaml:"ref_no"`                                             // 订单号
	Amount          string `json:"amount" mapstructure:"amount" config:"amount"  yaml:"amount"`                                             // 金额
	PaymentMethodId string `json:"payment_method_id" mapstructure:"payment_method_id" config:"payment_method_id"  yaml:"payment_method_id"` // Online Banking = 1 PayNow = 2
	Currency        string `json:"currency" mapstructure:"currency" config:"currency"  yaml:"currency"`                                     // 币种
	ReturnUrl       string `json:"return_url" mapstructure:"return_url" config:"return_url"  yaml:"return_url"`                             // 回调地址
	CallbackUrl     string `json:"callback_url" mapstructure:"callback_url" config:"callback_url"  yaml:"callback_url"`                     // 回调地址
	PlayerId        string `json:"player_id" mapstructure:"player_id" config:"player_id"  yaml:"player_id"`                                 // 玩家id
	Hash            string `json:"hash" mapstructure:"hash" config:"hash"  yaml:"hash"`                                                     // 签名  *Amount must have exactly 2 decimal places.(SHA 256) Hash Format: ref_no + amount + currency
}

type GmPayDepositRsp struct {
	Status  bool        `json:"status" mapstructure:"status"`
	Message string      `json:"message" mapstructure:"message"`
	Data    DepositData `json:"data" mapstructure:"data"`
}

type DepositData struct {
	CreatedAt            string `json:"created_at" mapstructure:"created_at"`                       //建立时间
	TransactionReference string `json:"transaction_reference" mapstructure:"transaction_reference"` // 交易引用号
	RefNo                string `json:"ref_no" mapstructure:"ref_no"`                               // 订单号
	Currency             string `json:"currency" mapstructure:"currency"`                           // 币种
	PaymentMethod        string `json:"payment_method" mapstructure:"payment_method"`               // 支付方式
	Amount               string `json:"amount" mapstructure:"amount"`                               //金额（小数点后取2位）
	PlatformCharge       string `json:"platform_charge" mapstructure:"platform_charge"`             //平台手续费（小数点后取2位）
	FinalAmount          string `json:"final_amount" mapstructure:"final_amount"`                   //最终到账金额（小数点后取2位）
	ReturnUrl            string `json:"return_url" mapstructure:"return_url"`                       // 回调地址
	CallbackUrl          string `json:"callback_url" mapstructure:"callback_url"`                   // 回调地址
	IpAddress            string `json:"ip_address" mapstructure:"ip_address"`                       // 客户IP地址
	PaymentLink          string `json:"payment_link" mapstructure:"payment_link"`                   // 支付链接
	IsSandbox            bool   `json:"is_sandbox" mapstructure:"is_sandbox"`                       // 是否沙箱环境
}

// gmpay出金
type GmPayWithdrawReq struct {
	ApiKey         string `json:"api_key" mapstructure:"api_key" config:"api_key"  yaml:"api_key"`                                     // apiKey
	RefNo          string `json:"ref_no" mapstructure:"ref_no" config:"ref_no"  yaml:"ref_no"`                                         // 订单号
	Amount         string `json:"amount" mapstructure:"amount" config:"amount"  yaml:"amount"`                                         // 金额
	Currency       string `json:"currency" mapstructure:"currency" config:"currency"  yaml:"currency"`                                 // 币种
	CallbackUrl    string `json:"callback_url" mapstructure:"callback_url" config:"callback_url"  yaml:"callback_url"`                 // 回调地址
	BankName       string `json:"bank_name" mapstructure:"bank_name" config:"bank_name"  yaml:"bank_name"`                             // 银行名称
	BankholderName string `json:"bank_holder_name" mapstructure:"bank_holder_name" config:"bank_holder_name"  yaml:"bank_holder_name"` // 银行账户名
	BankAccount    string `json:"bank_account" mapstructure:"bank_account" config:"bank_account"  yaml:"bank_account"`                 // 银行账户号
	Remarks        string `json:"remarks" mapstructure:"remarks" config:"remarks"  yaml:"remarks"`                                     // 备注
	PlayerId       string `json:"player_id" mapstructure:"player_id" config:"player_id"  yaml:"player_id"`                             // 玩家id
	Hash           string `json:"hash" mapstructure:"hash" config:"hash"  yaml:"hash"`                                                 // 签名
}

type GmPayWithdrawRsp struct {
	Status  bool         `json:"status" mapstructure:"status"`
	Message string       `json:"message" mapstructure:"message"`
	Result  WithdrawData `json:"result" mapstructure:"result"`
}

type WithdrawData struct {
	StatusId             int32  `json:"status_id" mapstructure:"status_id"`
	Status               string `json:"status" mapstructure:"status"`
	RefNo                string `json:"ref_no" mapstructure:"ref_no"`
	Currency             string `json:"currency" mapstructure:"currency"`
	Amount               int32  `json:"amount" mapstructure:"amount"`
	PlatformCharge       int32  `json:"platform_charge" mapstructure:"platform_charge"`
	FinalAmount          int32  `json:"final_amount" mapstructure:"final_amount"`
	BankName             string `json:"bank_name" mapstructure:"bank_name"`
	BankholderName       string `json:"bankholder_name" mapstructure:"bankholder_name"`
	BankAccount          string `json:"bank_account" mapstructure:"bank_account"`
	Remarks              string `json:"remarks" mapstructure:"remarks"`
	CallbackUrl          string `json:"callback_url" mapstructure:"callback_url"`
	TransactionReference string `json:"transaction_reference" mapstructure:"transaction_reference"`
}

// 入金和出金回调
type GmPayCallbackReq struct {
	Status               string `json:"status" form:"status" mapstructure:"status"`                                              // Completed、 Failed
	Currency             string `json:"currency" form:"currency" mapstructure:"currency"`                                        //币种
	Amount               string `json:"amount" form:"amount" mapstructure:"amount"`                                              //金额（小数点后取2位）
	PlatformCharge       string `json:"platform_charge" form:"platform_charge" mapstructure:"platform_charge"`                   //平台手续费（小数点后取2位）
	RefNo                string `json:"ref_no" form:"ref_no" mapstructure:"ref_no"`                                              //平台订单号
	TransactionReference string `json:"transaction_reference" form:"transaction_reference" mapstructure:"transaction_reference"` //交易引用号
	UpdatedAt            string `json:"updated_at" form:"updated_at" mapstructure:"updated_at"`                                  //更新时间
	IsSandBox            string `json:"is_sandbox" form:"is_sandbox" mapstructure:"is_sandbox"`                                  //是否沙箱环境
	// Callback hash using SHA256 transaction_reference + amount + currency
	Hash string `json:"hash" form:"hash" mapstructure:"hash"` //签名
}

type GmPayCallbackRsp struct {
	Status               string `json:"status" form:"status" mapstructure:"status"`                                              // Completed、 Failed
	Currency             string `json:"currency" form:"currency" mapstructure:"currency"`                                        //币种
	Amount               string `json:"amount" form:"amount" mapstructure:"amount"`                                              //金额（小数点后取2位）
	PlatformCharge       string `json:"platform_charge" form:"platform_charge" mapstructure:"platform_charge"`                   //平台手续费（小数点后取2位）
	RefNo                string `json:"ref_no" form:"ref_no" mapstructure:"ref_no"`                                              //平台订单号
	TransactionReference string `json:"transaction_reference" form:"transaction_reference" mapstructure:"transaction_reference"` //交易引用号
	UpdatedAt            string `json:"updated_at" form:"updated_at" mapstructure:"updated_at"`                                  //更新时间
	IsSandBox            int32  `json:"is_sandbox" form:"is_sandbox" mapstructure:"is_sandbox"`                                  //是否沙箱环境
	Hash                 string `json:"hash" form:"hash" mapstructure:"hash"`
}
