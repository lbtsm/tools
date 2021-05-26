package group

import "errors"

const All = "@all"

var applicationJson = "application/json"

var (
	textContentFormat     = `{"content":"%s", "mentioned_list":%+v, "mentioned_mobile_list": %v}`
	markdownContentFormat = `{"content":"%s"}`
	imageContentFormat    = `{"base64":"%s", "md5":"%s"}`
	newsContentFormat     = `{"articles":%s}` // 图文消息，一个图文消息支持1到8条图文，超过九条则消息发送失败
	fileContentFormat     = `{"media_id":"%s"}`
)

var (
	formatOfText     = `{"msgtype": "text","text": %s}`
	formatOfMarkdown = `{"msgtype": "markdown","markdown": %s}`
	formatOfImage    = `{"msgtype": "image","image": %s}`
	formatOfNews     = `{"msgtype": "news","news": %s}`
	formatOfFile     = `{"msgtype": "file","file": %s}`
)

var (
	WebHookIsEmpty = errors.New("webHook is empty")
)
