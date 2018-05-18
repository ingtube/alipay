package alipay

type AlipayUserAuthorization struct {
	AppAuthToken string `json:"-"` // 可选
	GrantType    string `json:"grant_type"`// 值为authorization_code时，代表用code换取；值为refresh_token时，代表用refresh_token换取
}

func (this AlipayUserAuthorization) APIName() string {
	return "alipay.system.oauth.token"
}

func (this AlipayUserAuthorization) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	m["grant_type"] = this.GrantType
	return m
}

func (this AlipayUserAuthorization) ExtJSONParamName() string {
	return ""
}

func (this AlipayUserAuthorization) ExtJSONParamValue() string {
	return marshal(this)
}

type AlipayUserAuthorizationResponse struct {
	Body struct {
		Code         string `json:"code"`
		Msg          string `json:"msg"`
		SubCode      string `json:"sub_code"`
		SubMsg       string `json:"sub_msg"`
		UserId       string `json:"user_id"`
		AccessToken  string `json:"access_token"`
		ExpiresIn    string `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		ReExpiresIn  string `json:"re_expires_in"` // 商户转账唯一订单号：发起转账来源方定义的转账单据号。请求时对应的参数，原样返回
	} `json:"alipay_system_oauth_token_response"`
	Sign string `json:"sign"`
}

func (this *AlipayUserAuthorizationResponse) IsSuccess() bool {
	if this.Body.Code == K_SUCCESS_CODE {
		return true
	}
	return false
}