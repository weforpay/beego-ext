package controllers

import (
	"strings"
)

const (
	API_OK  ApiResult = true
	API_ERR ApiResult = false
)
const (
	ERR_INCORRECT_PARAMETERS        = "ERR_INCORRECT_PARAMETERS"
	ERR_CREATE_CAPTCHA_ERR          = "ERR_CREATE_CAPTCHA_ERR"
	ERR_CHECK_CAPTCHA_ERR           = "ERR_CHECK_CAPTCHA_ERR"
	ERR_SEND_SMS_FAILED             = "ERR_SEND_SMS_FAILED"
	ERR_FAIL_TO_CHECK_VERIFY_CODE   = "ERR_FAIL_TO_CHECK_VERIFY_CODE"
	ERR_USER_OR_PWD_WRONG           = "ERR_USER_OR_PWD_WRONG"
	ERR_USER_JSON_CODING_WRONG      = "ERR_USER_JSON_CODING_WRONG" //user json编码错误
	ERR_EMAIL_NOT_EXIST             = "ERR_EMAIL_NOT_EXIST"        //您填写的邮箱不存在！
	ERR_WRONG_LINK                  = "ERR_WRONG_LINK"             //非法连接
	ERR_FAIL_TO_QUERY               = "ERR_FAIL_TO_QUERY"
	ERR_LOGIN_FIRST                 = "ERR_LOGIN_FIRST"                 //请先登录
	ERR_FAIL_TO_CHANGE_PWD          = "ERR_FAIL_TO_CHANGE_PWD"          //修改密码失败
	ERR_UPDATE_USER_ERROR           = "ERR_UPDATE_USER_ERROR"           //更新用户数据失败
	ERR_FAIL_TO_UPDATE_USER_IMAGE   = "ERR_FAIL_TO_UPDATE_USER_IMAGE"   //更新头像文件失败
	ERR_FAIL_TO_DECODE_IMG_BASE64   = "ERR_FAIL_TO_DECODE_IMG_BASE64"   //解码头像数据失败
	ERR_FAIL_TO_WRITE_USER_IMG_FILE = "ERR_FAIL_TO_WRITE_USER_IMG_FILE" //写头像文件失败
	ERR_FAIL_TO_DECODE_JSONFEEBACK  = "ERR_FAIL_TO_DECODE_JSONFEEBACK"  //
	ERR_FAIL_TO_ADD_FEEBACK_TO_DB   = "ERR_FAIL_TO_ADD_FEEBACK_TO_DB"   //写反馈到数据库失败
	ERR_FAIL_TO_ENCODE_JSON         = "ERR_FAIL_TO_ENCODE_JSON"         //json 编码失败
	ERR_FAIL_TO_DECODE_JSON         = "ERR_FAIL_TO_DECODE_JSON"         //json 解码失败
	ERR_FAIL_TO_ADD_WELFARE_CODE    = "ERR_FAIL_TO_ADD_WELFARE_CODE"    //数据库错误
	ERR_DB_OP_WRONG                 = "ERR_DB_OP_WRONG"                 //数据库错误
	ERR_FEEBACK_ID_NOT_EXIST        = "ERR_FEEBACK_ID_NOT_EXIST"        //feebackId不存在
	ERR_FAIL_TO_ADD_REPLAY_TO_DB    = "ERR_FAIL_TO_ADD_REPLAY_TO_DB"    //
	ERR_CAPTCHA_CODE_NOT_TIMEOUT    = "ERR_CAPTCHA_CODE_NOT_TIMEOUT"    //验证码还没有过期
	ERR_REGISTRATION_MOBILE         = "ERR_REGISTRATION_MOBILE"         //手机号已经被注册
	ERR_GET_DEVICES                 = "ERR_GET_DEVICES"                 //获取用户下的设备出错
	ERR_NO_USER                     = "ERR_NO_USER"                     //无此用户
	ERR_UNBIND_DEVICE               = "ERR_UNBIND_DEVICE"               //取消绑定出错
	ERR_NO_DEVICE                   = "ERR_NO_DEVICE"                   //无此设备
	ERR_SAVE_IMG_FILE               = "ERR_SAVE_IMG_FILE"               //保存图像文件失败
	ERR_NOT_SET_GPS_SVR_ADDRESS     = "ERR_NOT_SET_GPS_SVR_ADDRESS"     //数据里没有设置GPS服务器地址
	ERR_MOBILE_EXIST                = "ERR_MOBILE_EXIST"                //手机号已经存在
	ERR_BAIDU_LBS                   = "ERR_BAIDU_LBS"
	ERR_GET_ALIPAY_ORDERINFO        = "ERR_GET_ALIPAY_ORDERINFO"    //生成支付宝支付参数错误
	ERR_DEVICE_SERVICE_TYPE         = "ERR_DEVICE_SERVICE_TYPE"     //服务类型和充值类型有冲突，比如已经是终身的又充值包年
	ERR_DEVICE_EXPIREDATE_EMPTY     = "ERR_DEVICE_EXPIREDATE_EMPTY" //设备到期日期出错，只有终身的到期日期为空
	ERR_DEVICE_PAYING               = "ERR_DEVICE_PAYING"           //设备付款中
	ERR_DEVICE_AREADY_BIND          = "ERR_DEVICE_AREADY_BIND"      //设备已经绑定
	ERR_OLD_PASSWORD                = "ERR_OLD_PASSWORD"            //旧密码错误
	ERR_DEVICE_EXIST                = "ERR_DEVICE_EXIST"
	ERR_BIND_DEVICE                 = "ERR_BIND_DEVICE" //绑定设备失败
	ERR_SIGN                        = "ERR_SIGN"
	ERR_SHARE_SELF                  = "ERR_SHARE_SELF"         //不能自己给自己分享
	ERR_NO_SHARE_USER               = "ERR_NO_SHARE_USER"      //没有此分享用户
	ERR_DEVICE_OFFLINE              = "ERR_DEVICE_OFFLINE"     //设备不在线
	ERR_DEVICE_NO_RESPONSE          = "ERR_DEVICE_NO_RESPONSE" //设备没有应答
	ERR_UNKNOWN                     = "ERR_UNKNOWN"            //未知错误
	ERR_VALUE_TO_LANG               = "ERR_VALUE_TO_LANG"      //修改的值太长，比如密码太长
)

type ErrorController struct {
	ClientApiBase
}

func (this *ErrorController) Prepare() {
	this.ClientApiBase.Prepare()
	this.Out.SetResult(API_ERR)
	_, actionName := this.GetControllerAndAction()
	this.Out.SetErrMsg(strings.TrimPrefix(actionName, "Error"))

}

func (this *ErrorController) ErrorERR_INCORRECT_PARAMETERS()        {}
func (this *ErrorController) ErrorERR_USER_OR_PWD_WRONG()           {}
func (this *ErrorController) ErrorERR_LOGIN_FIRST()                 {}
func (this *ErrorController) ErrorERR_CREATE_CAPTCHA_ERR()          {}
func (this *ErrorController) ErrorERR_CHECK_CAPTCHA_ERR()           {}
func (this *ErrorController) ErrorERR_SEND_SMS_FAILED()             {}
func (this *ErrorController) ErrorERR_FAIL_TO_CHECK_VERIFY_CODE()   {}
func (this *ErrorController) ErrorERR_USER_JSON_CODING_WRONG()      {}
func (this *ErrorController) ErrorERR_EMAIL_NOT_EXIST()             {}
func (this *ErrorController) ErrorERR_WRONG_LINK()                  {}
func (this *ErrorController) ErrorERR_FAIL_TO_QUERY()               {}
func (this *ErrorController) ErrorERR_FAIL_TO_CHANGE_PWD()          {}
func (this *ErrorController) ErrorERR_UPDATE_USER_ERROR()           {}
func (this *ErrorController) ErrorERR_FAIL_TO_UPDATE_USER_IMAGE()   {}
func (this *ErrorController) ErrorERR_FAIL_TO_DECODE_IMG_BASE64()   {}
func (this *ErrorController) ErrorERR_FAIL_TO_WRITE_USER_IMG_FILE() {}
func (this *ErrorController) ErrorERR_FAIL_TO_DECODE_JSONFEEBACK()  {}
func (this *ErrorController) ErrorERR_OLD_PASSWORD()                {}
func (this *ErrorController) ErrorERR_DEVICE_PAYING()               {}
func (this *ErrorController) ErrorERR_DEVICE_SERVICE_TYPE()         {}
func (this *ErrorController) ErrorERR_BAIDU_LBS()                   {}
func (this *ErrorController) ErrorERR_MOBILE_EXIST()                {}
func (this *ErrorController) ErrorERR_NOT_SET_GPS_SVR_ADDRESS()     {}
func (this *ErrorController) ErrorERR_SAVE_IMG_FILE()               {}
func (this *ErrorController) ErrorERR_NO_DEVICE()                   {}
func (this *ErrorController) ErrorERR_UNBIND_DEVICE()               {}
func (this *ErrorController) ErrorERR_NO_USER()                     {}
func (this *ErrorController) ErrorERR_GET_DEVICES()                 {}
func (this *ErrorController) ErrorERR_REGISTRATION_MOBILE()         {}
func (this *ErrorController) ErrorERR_CAPTCHA_CODE_NOT_TIMEOUT()    {}
func (this *ErrorController) ErrorERR_FAIL_TO_ADD_REPLAY_TO_DB()    {}
func (this *ErrorController) ErrorERR_FEEBACK_ID_NOT_EXIST()        {}
func (this *ErrorController) ErrorERR_DB_OP_WRONG()                 {}
func (this *ErrorController) ErrorERR_FAIL_TO_ADD_WELFARE_CODE()    {}
func (this *ErrorController) ErrorERR_DEVICE_EXIST()                {}
func (this *ErrorController) ErrorERR_BIND_DEVICE()                 {}
func (this *ErrorController) ErrorERR_SIGN()                        {}
func (this *ErrorController) ErrorERR_SHARE_SELF()                  {}
func (this *ErrorController) ErrorERR_FAIL_TO_ENCODE_JSON()         {}
func (this *ErrorController) ErrorERR_FAIL_TO_DECODE_JSON()         {}
func (this *ErrorController) ErrorERR_NO_SHARE_USER()               {}
func (this *ErrorController) ErrorERR_DEVICE_OFFLINE()              {}
func (this *ErrorController) ErrorERR_DEVICE_NO_RESPONSE()          {}
func (this *ErrorController) ErrorERR_UNKNOWN()                     {}
func (this *ErrorController) ErrorERR_VALUE_TO_LANG()               {}
