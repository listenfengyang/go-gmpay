package go_gmpay

import (
	"encoding/json"
	"errors"

	"github.com/listenfengyang/go-nepay/utils"
	"github.com/mitchellh/mapstructure"
)

// 出金-成功回调
func (cli *Client) WithdrawCallback(req NePayCallbackReq, processor func(req NePayCallbackReq) error) error {
	//验证签名
	var params map[string]interface{}
	mapstructure.Decode(req.Data, &params)

	// Verify signature
	flag := utils.VerifyCallback(params, cli.Params.AccessKey)
	if !flag {
		//签名校验失败
		reqJson, _ := json.Marshal(req)
		cli.logger.Errorf("nepay withdraw back verify fail, req: %s", string(reqJson))
		return errors.New("sign verify error")
	}

	//开始处理
	return processor(req)
}
