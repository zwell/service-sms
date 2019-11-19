package yunxin

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/micro/go-micro/util/log"
	"io/ioutil"
	"math/big"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
	"zwell.github/mic-server/sms/internal/config"
	"zwell.github/mic-server/sms/internal/factory"
	"zwell.github/mic-server/sms/internal/model"
)

// 云信。网易云旗下服务
type Yunxin struct {
}

func (Yunxin) Send(templateModel *model.TemplateSupplier, phone string, params map[string]interface{}) (response *factory.Response, smsContent string, err error) {

	var code string
	for k, v := range params {
		if k == "code" {
			code = v.(string)
			break
		}
	}

	// 发送
	var res *factory.Response
	if templateModel.TemplateType == model.TYPE_CODE {
		res, err = sendCode(templateModel.TemplateCode, phone, code)
		if err != nil {
			return nil, "", err
		}
	}

	// 发送验证码内容
	smsContent = templateModel.TemplateContent
	for _, v := range params {
		temp := strings.Replace(smsContent, "%s", v.(string), 1)
		smsContent = temp
	}

	return res, smsContent, nil
}

// 发送验证码
func sendCode(templateId string, phone string, code string) (response *factory.Response, err error) {

	var url string = "https://api.netease.im/sms/sendcode.action"

	return post(url, "templateid="+templateId+"&mobile="+phone+"&authCode="+code)
}

func post(url string, post string) (response *factory.Response, err error) {

	conf := config.GetConf()

	// 随机数
	rand.Seed(time.Now().Unix())
	t := rand.Int()
	nonce := strconv.Itoa(t)

	t1 := time.Now().Unix()
	curTime := strconv.FormatInt(t1, 10)

	// 签名
	h := sha1.New()
	h.Write([]byte(conf.YunXin.AppSecret + nonce + curTime))
	hashBytes := h.Sum(nil)
	hexSha1 := hex.EncodeToString(hashBytes)
	// Integer base16 conversion
	intBase16, ok := new(big.Int).SetString(hexSha1, 16)
	if !ok {
		return nil, errors.New("签名失败")
	}
	checkSum := hex.EncodeToString(intBase16.Bytes())

	// 发送请求
	body := strings.NewReader(post)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Appkey", conf.YunXin.AppKey)
	req.Header.Set("Curtime", curTime)
	req.Header.Set("Checksum", checkSum)
	req.Header.Set("Nonce", nonce)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.Info("云信短信发送", url, post, string(result))

	// 返回结果解析
	type resStruct struct {
		Code int32  `json:"code"`
		Msg  string `json:"msg"`
		Obj  string `json:"obj"`
	}
	var r resStruct
	err = json.Unmarshal(result, &r)
	if err != nil {
		return nil, err
	}

	if r.Code != 200 {
		r.Code = 500
	}

	var message string
	if r.Code == 200 {
		message = "success"
	} else {
		message = "fail"
	}

	return &factory.Response{Code: r.Code, Message: message}, nil
}
