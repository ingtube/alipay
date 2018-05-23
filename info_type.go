package alipay

//https://docs.open.alipay.com/api_2/alipay.user.info.share
type AlipayUserInfo struct {
	AuthToken string `json:"auth_token"`
}

func (this AlipayUserInfo) APIName() string {
	return "alipay.user.info.share"
}

func (this AlipayUserInfo) Params() map[string]string {
	var m = make(map[string]string)
	m["auth_token"] = this.AuthToken
	return m
}

func (this AlipayUserInfo) ExtJSONParamName() string {
	return ""
}

func (this AlipayUserInfo) ExtJSONParamValue() string {
	return ""
}

type AlipayUserInfoResponse struct {
	Body struct {
		Code         string `json:"code"`
		Msg          string `json:"msg"`
		UserId         string `json:"user_id"`
		Avatar          string `json:"avatar"`
		Province      string `json:"province"`
		City       string `json:"city"`
		NickName       string `json:"nick_name"`
		IsStudentCertified  string `json:"is_student_certified"`
		UserType    string `json:"user_type"`
		UserStatus string `json:"user_status"`
		IsCertified  string `json:"is_certified"`
		Gender        string `json:"gender"`
	} `json:"alipay_user_info_share_response"`
	BodyError struct {
		Code         string `json:"code"`
		Msg          string `json:"msg"`
		SubCode      string `json:"sub_code"`
		SubMsg       string `json:"sub_msg"`
	} `json:"error_response"`
	Sign string `json:"sign"`
}

func (this *AlipayUserInfoResponse) IsSuccess() bool {
	if this.Body.Code == K_SUCCESS_CODE {
		return true
	}
	return false
}