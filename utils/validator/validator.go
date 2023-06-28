package validator

import (
	"fmt"
	"mcblog/utils/errno"
	"reflect"

	"github.com/go-playground/locales/zh_Hans_CN"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhs "github.com/go-playground/validator/v10/translations/zh"
)

// Validator 传入待检验的对象实例。错误则返回汉化的错误信息，正确则返回空串
// 在添加用户时调用
func Validator(data interface{}) (string, error) {
	validate := validator.New()                 // 实例化验证器
	uni := ut.New(zh_Hans_CN.New())             // 设置成中文翻译器
	trans, _ := uni.GetTranslator("zh_Hans_CN") //获取翻译词典

	// 注册翻译器
	_ = zhs.RegisterDefaultTranslations(validate, trans)

	//当用户名出错时，返回的信息为："Username长度不能超过12个字符" 希望字段也是中文：
	//1. 在 Model 的字段后tag添加label对应中文名
	//2. 使用 RegisterTagNameFunc() 获取并返回该中文
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})

	//验证器验证
	err := validate.Struct(data)
	if err != nil {
		// err 类型断言成功，确实为 ValidationErrors 将值赋给 v
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), errno.New(errno.ERROR, err)
		}
	}
	return "", nil
}

// 验证单一变量
func VarValidator(data interface{}, tag string) (string, error) {
	validate := validator.New()                 // 实例化验证器
	uni := ut.New(zh_Hans_CN.New())             // 设置成中文翻译器
	trans, _ := uni.GetTranslator("zh_Hans_CN") //获取翻译词典

	// 注册翻译器
	_ = zhs.RegisterDefaultTranslations(validate, trans)

	//验证器验证
	err := validate.Var(data, tag)
	if err != nil {
		// err 类型断言成功，确实为 ValidationErrors 将值赋给 v
		for _, v := range err.(validator.ValidationErrors) {
			fmt.Println(v.Translate(trans))
			return v.Translate(trans), errno.New(errno.ERROR, err)
		}
	}
	return "", nil
}
