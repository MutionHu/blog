package controllers

import (
	"errors"

	uuid "github.com/satori/go.uuid"

	"github.com/astaxie/beego"
	"github.com/blog/liteblog/models"
	"github.com/blog/liteblog/syserrors"
)

const SESSION_USER_KEY = "SESSION_USER_KEY"

type NestPrepare interface {
	NestPrepare()
}
type BaseController struct {
	beego.Controller
	IsLogin bool
	User    models.User
	Dao     *models.DB
}

func (ctx *BaseController) Prepare() {
	//log.Println("BaseControll")
	ctx.Data["Path"] = ctx.Ctx.Request.RequestURI
	ctx.Dao = models.NewDB()

	ctx.IsLogin = false
	if u, ok := ctx.GetSession(SESSION_USER_KEY).(models.User); ok {
		ctx.User = u
		ctx.Data["User"] = u
		ctx.IsLogin = true
	}
	ctx.Data["IsLogin"] = ctx.IsLogin
	if app, ok := ctx.AppController.(NestPrepare); ok {
		app.NestPrepare()
	}
}

func (ctx *BaseController) Abort500(err error) {
	ctx.Data["error"] = err
	ctx.Abort("500")
}

func (ctx *BaseController) MustLogin() {
	if !ctx.IsLogin {
		ctx.Abort500(syserrors.NoUserError{})
	}
}

func (ctx *BaseController) GetMustString(key, msg string) string {
	email := ctx.GetString(key, "")
	if len(email) == 0 {
		ctx.Abort500(errors.New(msg))
	}
	return email
}

type H map[string]interface{}

type ResultJsonValue struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Action string      `json:"action,omitempty"`
	Count  int         `json:"count,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func (ctx *BaseController) JSONOK(msg string, actions ...string) {
	var action string
	if len(actions) > 0 {
		action = actions[0]
	}
	ctx.Data["json"] = &ResultJsonValue{
		Code:   0,
		Msg:    msg,
		Action: action,
	}
	ctx.ServeJSON()
}

func (ctx *BaseController) JSONOKH(msg string, maps H) {
	if maps == nil {
		maps = H{}
	}
	maps["code"] = 0
	maps["msg"] = msg
	ctx.Data["json"] = maps
	ctx.ServeJSON()
}

func (ctx *BaseController) JSONOKData(count int, data interface{}) {
	ctx.Data["json"] = &ResultJsonValue{
		Code:  0,
		Count: count,
		Msg:   "成功！",
		Data:  data,
	}
	ctx.ServeJSON()
}

func (this *BaseController) UUID() string {
	u, err := uuid.NewV4()
	if err != nil {
		this.Abort500(syserrors.NewError("系统错误", err))
	}
	return u.String()
}
