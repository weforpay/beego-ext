package controllers

import (
	"encoding/json"
	"path"
	"runtime"
	"strconv"

	"github.com/weforpay/beego-ext/mylog"

	"github.com/astaxie/beego"
)

type ApiResult bool

type Output interface {
	SetResult(ApiResult)
	SetErrMsg(string)
	GetResult() ApiResult
	GetErrMsg() string
}

//输入参数自己在接口函数内部定义
//所有给客户端的输出参数都要从此结构继承
type OutputBase struct {
	Result ApiResult `json:"result"`
	ErrMsg string    `json:"errMsg"`
}

func (this *OutputBase) SetResult(r ApiResult) {
	this.Result = r
}

func (this *OutputBase) SetErrMsg(msg string) {
	this.ErrMsg = string(msg)
}

func (this *OutputBase) GetResult() ApiResult {
	return this.Result
}

func (this *OutputBase) GetErrMsg() string {
	return this.ErrMsg
}

//所有给客户端的接口的Controller都从此结构继承
type ClientApiBase struct {
	beego.Controller
	Out Output
}

func (this *ClientApiBase) Prepare() {
	var out OutputBase
	this.Out = &out
	var Origin = this.Ctx.Input.Header("Origin")
	if len(Origin) == 0 {
		Origin = "*"
	}
	this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	this.Ctx.Output.Header("Access-Control-Allow-Origin", Origin)

	//this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	this.Ctx.Output.Header("Content-Type", "application/json;charset=utf-8")
}

func (this *ClientApiBase) Finish() {
	if this.Out == nil {
		this.Out = &OutputBase{}
	}
	if len(this.Out.GetErrMsg()) == 0 {
		this.Out.SetResult(API_OK)
	} else {
		mylog.Error("api err:", string(this.Ctx.Input.RequestBody), this.Out)
	}
	this.Data["json"] = this.Out
	this.ServeJSON()
}

func (this *ClientApiBase) UnmarshalBody(v interface{}) (err error) {
	err = json.Unmarshal(this.Ctx.Input.RequestBody, v)
	return
}

type SessionApiBase struct {
	ClientApiBase
	UserId int64
}

type SessionOutputBase struct {
	OutputBase
}

func (this *SessionApiBase) Prepare() {
	this.ClientApiBase.Prepare()
	if this.GetSession("Id") == nil {
		this.Error(ERR_LOGIN_FIRST)
		return
	}

	this.UserId = this.GetSession("Id").(int64)

	if this.UserId <= 0 {
		this.Out.SetResult(API_ERR)
		this.Out.SetErrMsg(ERR_LOGIN_FIRST)
		mylog.Error("login first")
		return
	}
}

func (this *ClientApiBase) CheckError(code string, err error) {
	if err != nil {

		_, file, line, ok := runtime.Caller(1)
		if !ok {
			file = "???"
			line = 0
		}
		_, filename := path.Split(file)

		mylog.Error("CheckError[" + filename + ":" + strconv.FormatInt(int64(line), 10) + "] " + err.Error())
		panic(code)
	}
}
func (this *ClientApiBase) Error(code string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	_, filename := path.Split(file)

	mylog.Error("CheckError[" + filename + ":" + strconv.FormatInt(int64(line), 10) + "] ")
	panic(code)
}
