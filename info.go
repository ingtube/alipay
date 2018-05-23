package alipay

func (this *AliPay) UserInfo(param AlipayUserInfo) (results *AlipayUserInfoResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}
