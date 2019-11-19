package service

import (
	"database/sql"
	"errors"
	"time"
	"zwell.github/mic-server/sms/internal"
	"zwell.github/mic-server/sms/internal/database"
	"zwell.github/mic-server/sms/internal/factory"
	"zwell.github/mic-server/sms/internal/model"
)

type SmsService struct {

}

// 调用指定的供应商接口，发送短信
func (*SmsService) Send (templateResult *TemplateResult, phone string, params map[string]interface{}) (*factory.Response, error) {

	// 获取发送对象
	smsSupplier := internal.GetSmsService(templateResult.SupplierCode)

	var templateSupplierModel = model.TemplateSupplierModel{}
	templateSupplier, err := templateSupplierModel.GetOne(templateResult.TemplateId, templateResult.TemplateType)
	if err != nil {
		return nil, errors.New("短信模板不存在")
	}

	// 发送短信
	var response *factory.Response
	var smsContent string
	response, smsContent, err = smsSupplier.Send(templateSupplier, phone, params)
	if err != nil {
		return nil, err
	}

	// 记录发送日志
	var status int
	if response.Code == 200 {
		status = 1
	} else {
		status = 0
	}
	var insertSql = "insert into send_log (task_id, template_supplier_id, phone, content, status, error_msg, created_at) value (0, ?, ?, ?, ?, ?, ?)"
	database.GetDB().MustExec(insertSql, templateSupplier.Id, phone, smsContent, status, response.Message, time.Now().Unix())

	return response, nil
}

type TemplateResult struct {
	SupplierCode string `db:"supplier_code"`
	TemplateId int `db:"id"`
	TemplateCode string `db:"code"`
	TemplateType int `db:"type"`
}

// 获取发送短信的供应商
// 根据价格，优先级筛选
func (*SmsService) GetSupplier(template string) (*TemplateResult, error) {

	var querySql = "Select s.code as supplier_code, ts.id, t.code, t.type from template t " +
		"left join template_supplier ts On ts.template_id = t.id " +
		"left join supplier s On s.id = ts.supplier_id and s.status = 1 " +
		" where t.code = '" + template + "' Order by ts.price asc, ts.priority desc limit 1"
	templateResult := TemplateResult{}
	err := database.GetDB().Get(&templateResult, querySql)
	if err == sql.ErrNoRows {
		return nil, errors.New("供应商不存在")
	}
	if err != nil {
		return nil, err
	}

	return &templateResult, nil
}

