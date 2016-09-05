package validate

import (
	"github.com/gin-gonic/gin"
	"time"
	"gopkg.in/go-playground/validator.v9"
	"strconv"
)

type Form struct {
	Str  string `form:"str"`
	Num  string `form:"num"`
	Date string `form:"date"`
}

var validate *validator.Validate

func ValidateForm(c *gin.Context) (*Form, error) {
	validate = validator.New()
	// 独自に定義したバリデーションを登録する
	validate.RegisterStructValidation(FormStructLevelValidation, Form{})

	obj := &Form{}
	c.Bind(obj)
	return obj, validate.Struct(obj)
}

// 独自に追加するバリデーションを定義する
func FormStructLevelValidation(sl validator.StructLevel) {
	form := sl.Current().Interface().(Form)

	// 数値チェック
	if form.Num != "" {
		i, ok := strconv.Atoi(form.Num)
		if ok != nil {
			sl.ReportError(form.Num, "Num", "num", "not numeric type", "")
		} else {
			if i < 1 || 10 < i {
				sl.ReportError(form.Num, "Num", "num", "not 1~10", "")
			}
		}
	}

	// 日付チェック
	if form.Date != "" {
		_, ok := time.Parse("2006/01/02", form.Date)
		if ok != nil {
			sl.ReportError(form.Date, "Date", "date", "not date type", "")
		}
	}

	// 相関チェック（今回は相関必須チェック）
	if form.Str == "" && form.Num == "" && form.Date == "" {
		sl.ReportError("input Fields", "", "", "Requiring at least one field", "")
	}
}