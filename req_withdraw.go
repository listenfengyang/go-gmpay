package go_nepay

import (
	"crypto/tls"
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/listenfengyang/go-nepay/utils"
	"github.com/mitchellh/mapstructure"
)

func (cli *Client) WithdrawReq(req NePayWithdrawReq) (*NePayWithdrawRsp, error) {

	rawURL := cli.Params.WithdrawUrl
	// 2. Convert struct to map for signing
	var params map[string]string
	mapstructure.Decode(req, &params)
	params["notify_url"] = cli.Params.WithdrawNotifyUrl
	params["username"] = cli.Params.MerchantInfo.UserName

	// Generate signature
	signStr, _ := utils.Sign(params, cli.Params.AccessKey)
	params["sign"] = signStr
	var result NePayWithdrawRsp
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
	cli.logger.Infof("PSPResty#nepay#withdraw->%s", string(restLog))

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
