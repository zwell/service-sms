package internal

import (
	"errors"
	"zwell.github/mic-server/sms/internal/factory"
	"zwell.github/mic-server/sms/internal/yunpian"
	"zwell.github/mic-server/sms/internal/yunxin"
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