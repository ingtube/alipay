package alipay

// https://docs.open.alipay.com/api_9/alipay.system.oauth.token/
type AlipayUserAuthorization struct {
	GrantType    string `json:"grant_type"`// 值为authorization_code时，代表用code换取；值为refresh_token时，代表用refresh_token换取
	Code         string `json:"code"`  // 可选。授权码，用户对应用授权后得到。
	RefreshToken string `json:"refresh_token"` // 可选。刷刷新令牌，上次换取访问令牌时得到
}

func (this AlipayUserAuthorization) APIName() string {
	return "alipay.system.oauth.token"
}

func (this AlipayUserAuthorization) Params() map[string]string {
	var m = make(map[string]string)
	m["grant_type"] = this.GrantType
	m["code"] = this.Code
	return m
}

func (this AlipayUserAuthorization) ExtJSONParamName() string {
	return ""
}

func (this AlipayUserAuthorization) ExtJSONParamValue() string {
	return ""
}

type AlipayUserAuthorizationResponse struct {
	Body struct {
		UserId       string `json:"user_id"`
		AccessToken  string `json:"access_token"`
		ExpiresIn    int64 `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		ReExpiresIn  int64 `json:"re_expires_in"`
	} `json:"alipay_system_oauth_token_response"`
	BodyError struct {
		Code         string `json:"code"`
		Msg          string `json:"msg"`
		SubCode      string `json:"sub_code"`
		SubMsg       string `json:"sub_msg"`
	} `json:"error_response"`
	Sign string `json:"sign"`
}

func (this *AlipayUserAuthorizationResponse) IsSuccess() bool {
	if this.Body.AccessToken != "" {
		return true
	}
	return false
}