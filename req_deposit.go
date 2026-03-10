package go_gmpay

import (
	"crypto/tls"
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/listenfengyang/go-gmpay/utils"
	"github.com/mitchellh/mapstructure"
)

// 下单
func (cli *Client) Deposit(req GmPayDepositReq) (*GmPayDepositRsp, error) {

	rawURL := cli.Params.DepositUrl

	var params map[string]string
	mapstructure.Decode(req, &params)

	params["amount"] = fmt.Sprintf("%2d", req.Amount) // 必须保留2位小数
	params["callback_url"] = cli.Params.CallbackUrl
	params["return_url"] = cli.Params.ReturnUrl

	// Generate signature
	signStr, _ := utils.Sign(params)
	params["sign"] = signStr
	fmt.Println(params)
	var result GmPayDepositRsp

	resp2, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetFormData(params).
		SetBody(params).
		SetHeaders(getHeaders()).
		SetDebug(cli.debugMode).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp2))
	cli.logger.Infof("PSPResty#gmPay#deposit->%s", string(restLog))

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
