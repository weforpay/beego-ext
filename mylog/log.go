package mylog

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var mylog = logs.NewLogger(1024)

func Start() (err error) {
	beego.SetLogger("file", `{"filename":"log.txt","maxsize":102400000}`)
	mylog.SetLogger("console", "")
	mylog.SetLogger("file", `{"filename":"log.txt","maxsize":102400000}`)
	mylog.SetLogFuncCallDepth(3)
	mylog.EnableFuncCallDepth(true)
	//	beego.SetLevel(beego.LevelInformational)
	return
}
func SetLevel(l int) {
	mylog.SetLevel(l)
}
func Debugf(f string, p ...interface{}) {
	mylog.Debug(f, p...)

}
func Infof(f string, p ...interface{}) {
	mylog.Info(f, p...)
}
func Warnf(f string, p ...interface{}) {
	mylog.Warn(f, p...)
}
func Errorf(f string, p ...interface{}) {
	mylog.Error(f, p...)
}

func Debug(p ...interface{}) {
	mylog.Debug("%s", fmt.Sprint(p...))
}
func Info(p ...interface{}) {
	mylog.Info("%s", fmt.Sprint(p...))
}
func Warn(p ...interface{}) {
	mylog.Warn("%s", fmt.Sprint(p...))
}
func Error(p ...interface{}) {
	mylog.Error("%s", fmt.Sprint(p...))
}
