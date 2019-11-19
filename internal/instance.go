package internal

import (
	"zwell.github/mic-server/sms/internal/factory"
	"zwell.github/mic-server/sms/internal/yunpian"
	"zwell.github/mic-server/sms/internal/yunxin"
)

func GetSmsService(typeService string) factory.Factory {
    switch typeService {
        case "yunpian":
            return yunpian.Yunpian{}
		case "yunxin":
			return yunxin.Yunxin{}
        default:
        	panic("无效运算符号")
        	return nil
    }
}