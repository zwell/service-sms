package model

import (
	"database/sql"
	"errors"
	"zwell.github/mic-server/sms/internal/database"
)

type TemplateSupplierModel struct {

}

// 验证码类型。1短信验证码，2通知，3营销短信
const TYPE_CODE = 1

type TemplateSupplier struct {
	Id        int `db:"id"`
	SupplierId      int `db:"supplier_id"`
	TemplateId      int `db:"template_id"`
	TemplateCode      string `db:"template_code"`
	TemplateParams      string `db:"template_params"`
	TemplateContent      string `db:"template_content"`
	TemplateType      int
}

func (*TemplateSupplierModel) GetOne (templateId int, templateType int) (*TemplateSupplier, error) {

	var querySql = "Select id, supplier_id, template_id, template_code, template_params, template_content " +
		" from template_supplier " +
		" where id = ?"
	templateSupplier := TemplateSupplier{}
	err := database.GetDB().Get(&templateSupplier, querySql, templateId)
	if err == sql.ErrNoRows {
		return nil, errors.New("短信模板不存在")
	}
	if err != nil {
		return nil, err
	}
	templateSupplier.TemplateType = templateType

	return &templateSupplier, nil
}