package go_gmpay

import (
	"fmt"
	"testing"
)

type VLog struct {
}

func (l VLog) Debugf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Infof(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Warnf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Errorf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func TestCallback(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &GmPayInitParams{
		MerchantInfo: MerchantInfo{MERCHANT_ID, ACCESS_KEY},
	})

	err := cli.DepositCallback(GenCallbackRequestDemo(), func(GmPayCallbackReq) error { return nil })
	if err != nil {
		cli.logger.Errorf("Error:%s", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", err)
}

func GenCallbackRequestDemo() GmPayCallbackReq {
	return GmPayCallbackReq{
		Data: CallbackData{
			OrderNumber:       "202602060808160273",
			SystemOrderNumber: "GX20260206140816639893",
			UserName:          "CPT01",
			Amount:            "2205.00",
			Status:            5,
			Sign:              "3edd86a4daf46ccac3b4b28296f282f7",
		},
		HttpStatusCode: 200,
		ErrorCode:      0,
		Message:        "u5f02u6b65u56deu8c03",
	}
}
