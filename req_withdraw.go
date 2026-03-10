package go_gmpay

import (
	"crypto/tls"
	"fmt"
	"strconv"

	jsoniter "github.com/json-iterator/go"
	"github.com/listenfengyang/go-gmpay/utils"
	"github.com/mitchellh/mapstructure"
)

// 提现
func (cli *Client) WithdrawReq(req GmPayWithdrawReq) (*GmPayWithdrawRsp, error) {

	rawURL := cli.Params.WithdrawUrl
	// 2. Convert struct to map for signing
	var params map[string]string
	mapstructure.Decode(req, &params)
	amount, _ := strconv.ParseFloat(req.Amount, 64)
	params["amount"] = strconv.FormatFloat(amount, 'f', 2, 64) // 必须保留2位小数
	params["callback_url"] = cli.Params.WithdrawNotifyUrl

	// Generate signature
	signStr, _ := utils.Sign(params, cli.Params.MerchantInfo.SecretKey)
	params["hash"] = signStr
	var result GmPayWithdrawRsp
	fmt.Println(params)

	resp2, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(params).
		SetHeaders(getHeaders()).
		SetDebug(cli.debugMode).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp2))
	cli.logger.Infof("PSPResty#gmpay#withdraw->%s", string(restLog))

	if err != nil {
		return nil, err
	}

	if resp2.StatusCode() != 200 {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("status code: %d", resp2.StatusCode())
	}

	if resp2.Error() != nil {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("%v, body:%s", resp2.Error(), resp2.Body())
	}

	return &result, nil
}
