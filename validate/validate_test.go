package validate

import (
	"testing"
	"gopkg.in/go-playground/validator.v9"
)

func TestFormStructLevelValidation(t *testing.T) {
	validate = validator.New()
	validate.RegisterStructValidation(FormStructLevelValidation, Form{})
	m := map[*Form]bool{}
	// 数値型チェック
	m[&Form{Str:"a", Num:"1", Date:"2006/01/02"}] = true
	m[&Form{Str:"a", Num:"a", Date:"2006/01/02"}] = false
	// 数値下限チェック
	m[&Form{Str:"a", Num:"1", Date:"2006/01/02"}] = true
	m[&Form{Str:"a", Num:"0", Date:"2006/01/02"}] = false
	// 数値上限チェック
	m[&Form{Str:"a", Num:"10", Date:"2006/01/02"}] = true
	m[&Form{Str:"a", Num:"11", Date:"2006/01/02"}] = false
	// 日付チェック
	m[&Form{Str:"a", Num:"1", Date:"2006/12/31"}] = true
	m[&Form{Str:"a", Num:"1", Date:"2008/02/29"}] = true
	m[&Form{Str:"a", Num:"1", Date:"2006/12/32"}] = false
	m[&Form{Str:"a", Num:"1", Date:"2007/02/29"}] = false
	// 相関必須チェック
	m[&Form{Str:"a", Num:"", Date:""}] = true
	m[&Form{Str:"", Num:"1", Date:""}] = true
	m[&Form{Str:"", Num:"", Date:"2006/12/31"}] = true
	m[&Form{Str:"", Num:"", Date:""}] = false

	for key, value := range m {
		err := validate.Struct(key)
		if value && err != nil {
			t.Fatalf("Error Occured: %v %v %v", key.Str, key.Num, key.Date)
		}
		if !value && err == nil {
			t.Fatalf("Error Occured: %v %v %v", key.Str, key.Num, key.Date)
		}
	}

}
