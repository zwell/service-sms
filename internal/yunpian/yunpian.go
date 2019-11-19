package yunpian

import (
	"github.com/micro/go-micro/util/log"
	ypclnt "github.com/yunpian/yunpian-go-sdk/sdk"
	"strings"
	"github.com/zwell/service-sms/internal/config"
	"github.com/zwell/service-sms/internal/factory"
	"github.com/zwell/service-sms/internal/model"
)

// 云片。主要用来发送一些大厂发送不了的短信模板
type Yunpian struct {
}

func (Yunpian) Send(templateModel *model.TemplateSupplier, phone string, params map[string]interface{}) (response *factory.Response, smsContent string, err error) {

	// 变量替换
	templateParams := strings.Split(templateModel.TemplateParams, ",")
	smsContent = templateModel.TemplateContent
	for k, v := range params {
		for _, p := range templateParams {
			if k == p {
				temp := strings.Replace(smsContent, "#"+k+"#", v.(string), 1)
				smsContent = temp
				break
			}
		}
	}

	conf := config.GetConf()

	// 发送
	clnt := ypclnt.New(conf.YunPian.ApiKey)
	param := ypclnt.NewParam(2)
	param[ypclnt.MOBILE] = phone
	param[ypclnt.TEXT] = smsContent
	r := clnt.Sms().SingleSend(param)

	log.Info("云片短信发送", param, r)

	var code int32
	if r.Code == 0 {
		code = 200
	} else {
		code = 500
	}

	var message string
	if code == 200 {
		message = "success"
	} else {
		message = "fail"
	}

	return &factory.Response{Code: code, Message: message}, smsContent, nil
}
