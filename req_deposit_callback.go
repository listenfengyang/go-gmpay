package go_gmpay

import (
	"encoding/json"
	"errors"

	"github.com/listenfengyang/go-gmpay/utils"
	"github.com/mitchellh/mapstructure"
)

// 充值-成功回调
func (cli *Client) DepositCallback(req GmPayCallbackReq, processor func(GmPayCallbackReq) error) error {
	//验证签名
	var params map[string]string
	mapstructure.Decode(req, &params)

	flag, err := utils.Verify(params, cli.Params.MerchantInfo.SecretKey)
	if err != nil {
		cli.logger.Errorf("gmPay deposit back verify error, req: %s, err: %s", req, err.Error())
		return err
	}
	if !flag {
		//签名校验失败
		reqJson, _ := json.Marshal(req)
		cli.logger.Errorf("gmPay deposit back verify fail, req: %s", string(reqJson))
		return errors.New("sign verify error")
	}

	//开始处理
	return processor(req)
}
