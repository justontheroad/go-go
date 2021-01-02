package main

import (
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

// Booking contains binded and validated data.
type Booking struct {
	CheckIn  time.Time `form:"check_in" json:"check_in" binding:"required,bookabledate" time_format:"2006-01-02" label:"输入时间"`
	CheckOut time.Time `form:"check_out" json:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02" label:"输出时间"`
}

// 翻译器
var trans ut.Translator

// 预定日期验证方法
var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

// CheckIn CheckOut翻译
var BookingTrans = map[string]string{"CheckIn": "输入时间", "CheckOut": "输出时间"}

// 翻译tag name
func TransTagName(libTans, err interface{}) interface{} {
	switch err.(type) {
	case validator.ValidationErrorsTranslations:
		var errs map[string]string
		errs = make(map[string]string, 0)
		for k, v := range err.(validator.ValidationErrorsTranslations) {
			for key, value := range libTans.(map[string]string) {
				v = strings.Replace(v, key, value, -1)
			}
			errs[k] = v
		}
		return errs
	case string:
		var errs string
		for key, value := range libTans.(map[string]string) {
			errs = strings.Replace(errs, key, value, -1)
		}
		return errs
	default:
		return err
	}
}

// booking 验证器注册
func ValidatorBooingRegister() bool {
	flag := false
	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")
	// 2. 绑定自定义验证器，必须在router绑定handler之前处理
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册翻译器
		_ = zh_translations.RegisterDefaultTranslations(v, trans)
		//注册自定义函数
		_ = v.RegisterValidation("bookabledate", bookableDate)

		//注册一个函数，获取struct tag里自定义的label作为字段名
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := fld.Tag.Get("label")
			return name
		})
		//根据提供的标记注册翻译
		v.RegisterTranslation("bookabledate", trans, func(ut ut.Translator) error {
			return ut.Add("bookabledate", "{0}不能早于当前时间或{1}格式错误!", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("bookabledate", fe.Field(), fe.Field())
			return t
		})

		flag = true
	}

	return flag
}
