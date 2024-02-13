package tool

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

type CaptchaResult struct {
	CaptchaId  string `json:"captcha_id"`
	Base64Bold string `json:"base_64_bold"`
	VerifyCode string `json:"verify_code"`
}

func Captcha(context *gin.Context) {
	parameter := base64Captcha.ConfigCharacter{
		Height:     30,
		Width:      60,
		CaptchaLen: 4,
	}

	//生成
	captchaId, captchaInstance := base64Captcha.GenerateCaptcha("", parameter)
	//装换成base64
	base64Bold := base64Captcha.CaptchaWriteToBase64Encoding(captchaInstance)

	captchaResult := CaptchaResult{captchaId, base64Bold, ""}
	//浏览器返回
	Success(context, captchaResult)
}

func VerifyCaptcha(id string, value string) bool {
	result := base64Captcha.VerifyCaptcha(id, value)
	return result
}
