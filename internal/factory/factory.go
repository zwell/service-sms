package factory

import "github.com/zwell/service-sms/internal/model"

type Response struct {
	Code    int32
	Message string
}

type Factory interface {
	Send(*model.TemplateSupplier, string, map[string]interface{}) (response *Response, smsContent string, err error)
}
