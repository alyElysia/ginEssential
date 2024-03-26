package controller

import (
	"ginEssential/common"
	"ginEssential/dto"
	"ginEssential/model"
	"ginEssential/response"
	"ginEssential/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

var db = common.InitDB()

func Register(ctx *gin.Context) {

	//获取参数
	name := ctx.PostForm("name")
	tel := ctx.PostForm("tel")
	pwd := ctx.PostForm("pwd")
	//数据验证
	if len(tel) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
	}
	if len(pwd) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")

	}

	//如果没有传递名称，测随机生成
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	//判断手机号是否存在
	if isTelExist(db, tel) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "该手机号已注册")

	}

	//创建用户

	//为密码加密
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密失败")
	}
	newUser := model.User{
		Name: name,
		Tel:  tel,
		Pwd:  string(hashPwd),
	}
	db.Create(&newUser)
	//返回结果
	response.Success(ctx, nil, "注册成功")
}

func Login(ctx *gin.Context) {
	//获取数据
	name := ctx.PostForm("name")
	tel := ctx.PostForm("tel")
	pwd := ctx.PostForm("pwd")

	//判断输入的用户名或手机号是否正确
	//未输入用户名或手机号
	if len(name) == 0 || len(tel) == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户名或手机号不能为空！")
	}
	//输入错误
	var checkUser model.User
	db.Raw("select * from users where name = ? or tel = ?;", name, tel).Scan(&checkUser)
	if checkUser.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户名或手机号输入错误！")
	}

	//判断密码是否正确
	//格式不正确
	if len(pwd) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码长度必须大于6！")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(checkUser.Pwd), []byte(pwd)); err != nil {
		response.Response(ctx, http.StatusNotFound, 404, nil, "密码错误!")
	}

	//发放token
	token, err := common.ReleaseClaims(checkUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常！"})
		return
	}

	//返回结果
	response.Success(ctx, gin.H{"code": 200, "token": token}, "登录成功！")

}

func isTelExist(db *gorm.DB, tel string) bool {
	var user model.User
	db.Where("tel = ?", tel).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

// Info 获取用户信息
func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.Success(ctx, gin.H{"code": 200, "data": gin.H{"user": dto.GetDto(user.(model.User))}}, "用户信息获取成功")
}
