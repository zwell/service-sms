package internal

import (
	"errors"
	"github.com/zwell/service-sms/internal/factory"
	"github.com/zwell/service-sms/internal/yunpian"
	"github.com/zwell/service-sms/internal/yunxin"
)

func GetSmsService(typeService string) (factory.Factory, error) {
    switch typeService {
        case "yunpian":
            return yunpian.Yunpian{}, nil
		case "yunxin":
			return yunxin.Yunxin{}, nil
        default:
        	return nil, errors.New("短信供应商不存在")
    }
}