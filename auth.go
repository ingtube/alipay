package alipay

func (this *AliPay) UserAuthorization(param AliPayFundTransToAccountTransfer) (results *AliPayFundTransToAccountTransferResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}


