package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"reflect"
	"strings"
	"time"
)

// RegisterForm 在web开发中一个不可避免的环节就是对请求参数进行校验，
// 通常我们会在代码中定义与请求参数相对应的模型（结构体），
// 借助模型绑定快捷地解析请求中的参数，例如 gin 框架中的Bind和ShouldBind系列方法。
// gin框架使用github.com/go-playground/validator进行参数校验，
// 目前已经支持github.com/go-playground/validator/v10了，
// 我们需要在定义结构体时使用 binding tag标识相关校验规则，可以查看validator文档查看支持的所有 tag。
type RegisterForm struct {
	UserName   string `json:"username" binding:"required,min=3,max=7"`
	Password   string `json:"password" binding:"required,len=8"`
	RePassword string `json:"re_password" binding:"required,len=8,eqfield=Password"`
	Age        uint32 `json:"age" binding:"required,gte=1,lte=150"`
	Sex        uint32 `json:"sex" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	// 需要使用自定义校验方法checkDate做参数校验的字段Date
	// datetime=2006-01-02是内置的用于校验日期类参数是否满足指定格式要求的tag
	// 如果传入的date参数不满足2006-01-02这种格式就会提示如下错误：{"msg":{"date":"date的格式必须是2006-01-02"}}
	Date string `json:"date" binding:"required,datetime=2006-01-02,checkDate"`
}

type LoginForm struct {
	UserName   string `json:"username" binding:"required,min=3,max=7"`
	Password   string `json:"password" binding:"required,len=8"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password,len=8"`
}

// 定义一个全局翻译器 trans
var trans ut.Translator

func main() {
	// locale 实际从请求头的 Accept-Language中获取，这里仅模拟
	locale := "zh"
	if err := InitializeTrans(locale); err != nil {
		// 初始化 翻译器 失败
		fmt.Println(err.Error())
		panic(err)
	}

	r := gin.Default()
	r.POST("/login", LoginHandler)
	r.POST("/register", RegisterHandler)
	r.Run()
}

// RegisterHandler 注册
func RegisterHandler(c *gin.Context) {
	var r RegisterForm
	if err := c.ShouldBindJSON(&r); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)

		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			c.JSON(http.StatusOK, gin.H{
				"code": 40010,
				"msg":  "注册失败，请检查参数",
				"err":  err.Error(),
			})
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		c.JSON(http.StatusOK, gin.H{
			"code": 40004,
			//"err":  errs.Translate(trans),
			"err": removeTopStruct(errs.Translate(trans)),
		})
		return
	}

	// 注册成功
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "注册成功",
		"data": r,
	})
}

// LoginHandler 登录
func LoginHandler(c *gin.Context) {
	var l LoginForm
	if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 40001,
			"msg":  "登录失败，请检查参数",
			"err":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "登陆成功",
		"data": l.UserName,
	})
}

// InitializeTrans 初始化翻译器
//func InitializeTrans(locale string) (err error) {
func InitializeTrans(locale string) (err error) {
	// 修改 gin 框架中的 Validator 引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取 json tag 的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

			if name == "-" {
				return ""
			}

			return name
		})

		// 在校验器注册自定义的校验方法
		if err := v.RegisterValidation("checkDate", customFunc); err != nil {
			return err
		}

		v.RegisterStructValidation(RegisterFormStructLevelValidation, RegisterForm{})

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		// 第一个参数是备用 fallback 语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		uni := ut.New(enT, zhT)

		// locale 通常取决于 http 请求头的 Accept-Language
		var ok bool
		// 也可以使用 uni.Findtranslator(...)传入多个locale进行查找
		trans, ok = uni.GetTranslator(locale)

		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// 注册翻译器
		switch locale {
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}

		if err != nil {
			return err
		}
		// 注意！因为这里会使用到 trans 实例
		// 所以这一步注册要放到 trans 初始化的后面
		if err := v.RegisterTranslation(
			"checkDate",
			trans,
			registerTranslator("checkDate", "{0}必须要晚于当前日期"),
			translate,
		); err != nil {
			return err
		}
		return
	}

	return
}

func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

// RegisterFormStructLevelValidation 自定义 RegisterForm 结构体校验函数
func RegisterFormStructLevelValidation(sl validator.StructLevel) {
	l := sl.Current().Interface().(RegisterForm)

	if l.Password != l.RePassword {
		// 输出错误提示信息，最后一个参数就是传递的 param
		sl.ReportError(l.RePassword, "re_password", "RePassword", "eqfield", "password")
	}
}

// customFunc 自定义字段级别校验方法
func customFunc(fl validator.FieldLevel) bool {
	date, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}
	if date.Before(time.Now()) {
		return false
	}
	return true
}

// 自定义字段级别的校验方法的错误提示信息很“简单粗暴”
// registerTranslator 为自定义字段添加翻译功能
func registerTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

// translate 自定义字段的翻译方法
func translate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	return msg
}
