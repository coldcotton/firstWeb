package logic

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/coldcotton/firstWeb/app/model"
	"github.com/coldcotton/firstWeb/app/tools"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name         string `json:"name" form:"name"` // json,form的值必须和前端一样
	Password     string `json:"password" form:"password"`
	CaptchaId    string `json:"captcha_id" form:"captcha_id"`
	CaptchaValue string `json:"captcha_value" form:"captcha_value"`
}

func GetLogin(context *gin.Context) { // 获得登录页面
	context.HTML(http.StatusOK, "login.tmpl", nil)
}

func DoLogin(context *gin.Context) { // 用户登录
	var user User

	if err := context.ShouldBind(&user); err != nil {
		context.JSON(http.StatusOK, tools.ECode{
			Message: err.Error(),
		})
		return
	}

	if !tools.CaptchaVerify(tools.CaptchaData{
		CaptchaId: user.CaptchaId,
		Data:      user.CaptchaValue,
	}) {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10010,
			Message: "验证码错误",
		})
		return
	}

	ret := model.GetUser(user.Name)
	if ret.Id < 1 || ret.Password != encryptV1(user.Password) {
		context.JSON(http.StatusOK, tools.UserErr)

		return
	}

	// 设置cookie
	// context.SetCookie("name", user.Name, 3600, "/", "", false, false) // 第6个是安全性
	// context.SetCookie("id", fmt.Sprint(ret.Id), 3600, "/", "", false, false)

	_ = model.SetSession(context, user.Name, ret.Id) // 用户名、用户id

	context.JSON(http.StatusOK, tools.ECode{
		Message: "登录成功",
	})

	return
}

func Logout(context *gin.Context) { // 用户退出登录
	// context.SetCookie("name", "", 3600, "/", "", false, false)
	// context.SetCookie("id", "", 3600, "/", "", false, false)

	_ = model.FlushSession(context)
	context.Redirect(http.StatusFound, "/login")
}

// model和logic两层结构体，增加安全性，解耦
type CUser struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
	Password2 string `json:"password_2"`
}

func CreateUser(context *gin.Context) { // 用户注册
	var user CUser
	if err := context.ShouldBind(&user); err != nil {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10001,
			Message: err.Error(),
		})
		return
	}

	// 校验
	if user.Name == "" || user.Password == "" || user.Password2 == "" {
		context.JSON(http.StatusOK, tools.ParamErr)
		return
	}
	if user.Password != user.Password2 {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10003,
			Message: "两次输入的密码不同",
		})
		return
	}

	nameLen := len(user.Name)
	password := len(user.Password)
	if nameLen > 16 || nameLen < 8 || password > 16 || password < 8 {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10005,
			Message: "账号和密码的长度应在8-16位",
		})
		return
	}

	regex := regexp.MustCompile("^[0-9]+$")
	if regex.MatchString(user.Password) {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10006,
			Message: "密码不能为纯数字",
		})
		return
	}

	if oldUser := model.GetUser(user.Name); oldUser.Id > 0 {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10004,
			Message: "用户名已存在！！",
		})
		return
	}

	newUser := model.User{
		Name:        user.Name,
		Password:    encryptV1(user.Password),
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
		Uuid:        tools.GetUUID(),
	}
	if err := model.CreateUser(&newUser); err != nil {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10007,
			Message: "用户创建失败",
		})
		return
	}

	context.JSON(http.StatusOK, tools.OK)

}

// 加密v1
func encrypt(pwd string) string {
	hash := md5.New()
	hash.Write([]byte(pwd))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	fmt.Printf("加密后的密码：%s\n", hashString)

	return hashString
}

// 加密v2，加盐
func encryptV1(pwd string) string {
	newPwd := pwd + "加盐"
	hash := md5.New()
	hash.Write([]byte(newPwd))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	fmt.Printf("加密后的密码：%s\n", hashString)

	return hashString
}

// 加密v3
func encryptV2(pwd string) string {
	//基于Blowfish 实现加密。简单快速，但有安全风险
	//golang.org/x/crypto/ 中有大量的加密算法
	newPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("密码加密失败：", err)
		return ""
	}
	newPwdStr := string(newPwd)
	fmt.Printf("加密后的密码：%s\n", newPwdStr)
	return newPwdStr
}
