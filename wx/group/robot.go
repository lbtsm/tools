package group

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Robot struct {
	webHook string
}

func NewRobot(webHook string) *Robot {
	return &Robot{
		webHook: webHook,
	}
}

func (r *Robot) Text(content string, mentionedList, mentionedMobileList []string) error {
	text := fmt.Sprintf(textContentFormat, content, mentionedList, mentionedMobileList)
	return r.request(r.webHook, strings.NewReader(text))
}

func (r *Robot) Markdown(content string) error {
	text := fmt.Sprintf(markdownContentFormat, content)
	return r.request(r.webHook, strings.NewReader(text))
}

func (r *Robot) Image(base64, md5 string) error {
	text := fmt.Sprintf(imageContentFormat, base64, md5)
	return r.request(r.webHook, strings.NewReader(text))
}

func (r *Robot) News(news []NewsMsg) error {
	data, err := json.Marshal(news)
	if err != nil {
		return nil
	}

	text := fmt.Sprintf(newsContentFormat, string(data))
	return r.request(r.webHook, strings.NewReader(text))
}

func (r *Robot) File() error {

	return nil
}

func (r *Robot) request(url string, body io.Reader) error {
	if url == "" {
		return WebHookIsEmpty
	}

	resp, err := http.Post(url, applicationJson, body)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	rmb := &RobotMsgBack{}
	err = json.Unmarshal(data, rmb)
	if err != nil {
		return err
	}
	if rmb.ErrCode != 0 && rmb.ErrMsg != "ok" {
		return errors.New(fmt.Sprintf("enterPrise wx robot back code is %v, msg is %v", rmb.ErrCode, rmb.ErrMsg))
	}

	return nil
}

// todo 图片base64的
// todo 获取文件的md5
