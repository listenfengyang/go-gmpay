package go_nepay

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
	cli := NewClient(vLog, &NePayInitParams{
		MerchantInfo:      MerchantInfo{MERCHANT_ID, ACCESS_KEY},
		DepositUrl:        DEPOSIT_URL,
		WithdrawUrl:       WITHDRAW_URL,
		NotifyUrl:         NOTIFY_URL,
		WithdrawNotifyUrl: WITHDRAW_NOTIFY_URL,
	})

	err := cli.DepositCallback(GenCallbackRequestDemo(), func(NePayCallbackReq) error { return nil })
	if err != nil {
		cli.logger.Errorf("Error:%s", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", err)
}

func GenCallbackRequestDemo() NePayCallbackReq {
	return NePayCallbackReq{
		Data: CallbackData{
			OrderNumber:       "20260205835236672",
			SystemOrderNumber: "GX20260205085851633762",
			UserName:          "CPT02",
			Amount:            "1100.00",
			Status:            5,
			Sign:              "e73afcb8e47df8dcebec0239bd667b51",
		},
		HttpStatusCode: 200,
		ErrorCode:      0,
		Message:        "u5f02u6b65u56deu8c03",
	}
}
