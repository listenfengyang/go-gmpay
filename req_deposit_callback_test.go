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
		MerchantInfo: MerchantInfo{
			ApiKey:    API_KEY,
			SecretKey: SECRET_KEY,
		},
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
		Status:               "failed",
		Currency:             "SGD",
		Amount:               "100.00",
		Hash:                 "14797451ca4461ab0bd67016b7206ed12b4ca827d7270df0bb4fc6a7fa7ce8d1",
		RefNo:                "20230824152007",
		PlatformCharge:       "2.00",
		TransactionReference: "9149DBF33A7DBB330C8097D15F083FF4",
		UpdatedAt:            "2026-03-10 16:38:15",
		IsSandBox:            "1",
	}
}
