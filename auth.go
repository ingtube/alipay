package alipay

func (this *AliPay) UserAuthorization(param AlipayUserAuthorization) (results *AlipayUserAuthorizationResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

