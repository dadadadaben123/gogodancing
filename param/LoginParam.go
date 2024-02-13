package param

type LoginParam struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type LoginPwdParams struct {
	Name         string `json:"name"`
	Password     string `json:"password"`
	CaptchaId    string `json:"captcha_id"`
	CaptchaValue string `json:"captcha_value"`
}
