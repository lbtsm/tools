package group

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"testing"
)

var testWebHook = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=ec479b3f-6173-4a18-8187-161b0473f240"

func getImageMessage() (string, string) {
	// 图片的
	f, _ := ioutil.ReadFile("/Users/zmm/Desktop/other/235000-1584114600db79.jpg")
	picBase64 := base64.StdEncoding.EncodeToString(f)
	picMd5 := fmt.Sprintf("%x", md5.Sum(f))
	return picBase64, picMd5
}

func getNewsMessage() []NewsMsg {
	return []NewsMsg{
		{
			Title:       "这是一条图文消息",
			Description: "洗澡啦，洗澡啦",
			Url:         "https://work.weixin.qq.com/api/doc/90000/90136/91770#%E6%96%87%E6%9C%AC%E7%B1%BB%E5%9E%8B",
			PicUrl:      "https://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
		},
		{
			Title:       "这是二条图文消息",
			Description: "洗澡啦，洗澡啦",
			Url:         "https://work.weixin.qq.com/api/doc/90000/90136/91770#%E6%96%87%E6%9C%AC%E7%B1%BB%E5%9E%8B",
			PicUrl:      "https://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
		},
		{
			Title:       "这是三条图文消息",
			Description: "洗澡啦，洗澡啦",
			Url:         "https://work.weixin.qq.com/api/doc/90000/90136/91770#%E6%96%87%E6%9C%AC%E7%B1%BB%E5%9E%8B",
			PicUrl:      "https://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
		},
		{
			Title:       "这是四条图文消息",
			Description: "洗澡啦，洗澡啦",
			Url:         "https://work.weixin.qq.com/api/doc/90000/90136/91770#%E6%96%87%E6%9C%AC%E7%B1%BB%E5%9E%8B",
			PicUrl:      "https://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
		},
		{
			Title:       "这是五条图文消息",
			Description: "洗澡啦，洗澡啦",
			Url:         "https://work.weixin.qq.com/api/doc/90000/90136/91770#%E6%96%87%E6%9C%AC%E7%B1%BB%E5%9E%8B",
			PicUrl:      "https://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
		},
		{
			Title:       "这是六条图文消息",
			Description: "洗澡啦，洗澡啦",
			Url:         "https://work.weixin.qq.com/api/doc/90000/90136/91770#%E6%96%87%E6%9C%AC%E7%B1%BB%E5%9E%8B",
			PicUrl:      "https://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
		},
		{
			Title:       "这是七条图文消息",
			Description: "洗澡啦，洗澡啦",
			Url:         "https://work.weixin.qq.com/api/doc/90000/90136/91770#%E6%96%87%E6%9C%AC%E7%B1%BB%E5%9E%8B",
			PicUrl:      "https://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
		},
		{
			Title:       "这是八条图文消息",
			Description: "洗澡啦，洗澡啦",
			Url:         "https://work.weixin.qq.com/api/doc/90000/90136/91770#%E6%96%87%E6%9C%AC%E7%B1%BB%E5%9E%8B",
			PicUrl:      "https://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
		},
	}
}

func TestRobot_News(t *testing.T) {
	type fields struct {
		webHook string
	}
	type args struct {
		news []NewsMsg
	}
	var tests = []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "standard",
			fields: fields{
				webHook: testWebHook,
			},
			args: args{
				news: getNewsMessage(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Robot{
				webHook: tt.fields.webHook,
			}
			if err := r.News(tt.args.news); (err != nil) != tt.wantErr {
				t.Errorf("News() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRobot_File(t *testing.T) {
	type fields struct {
		webHook string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "standard",
			fields: fields{
				webHook: testWebHook,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Robot{
				webHook: tt.fields.webHook,
			}
			if err := r.File(); (err != nil) != tt.wantErr {
				t.Errorf("File() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRobot_Image(t *testing.T) {
	type fields struct {
		webHook string
	}
	type args struct {
		base64 string
		md5    string
	}
	imgBase64, imgMd5 := getImageMessage()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "standard",
			fields: fields{
				webHook: testWebHook,
			},
			args: args{
				base64: imgBase64,
				md5:    imgMd5,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Robot{
				webHook: tt.fields.webHook,
			}
			if err := r.Image(tt.args.base64, tt.args.md5); (err != nil) != tt.wantErr {
				t.Errorf("Image() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRobot_Markdown(t *testing.T) {
	type fields struct {
		webHook string
	}
	type args struct {
		content string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "standard",
			fields: fields{
				webHook: testWebHook,
			},
			args: args{
				content: "<font color='info' size='30'>绿色</font> \n " +
					"<font color='comment'>灰色</font> \n" +
					"<font color='warning'>橙红色</font> \n" +
					"[请点击这里,跳转](https://work.weixin.qq.com/api/doc/90000/90136/91770#%E6%96%87%E6%9C%AC%E7%B1%BB%E5%9E%8B) \n" +
					"## 这是一个标题 \n" +
					"**这是一个加粗的markdown消息** \n" +
					"> 引用 \n" +
					"",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Robot{
				webHook: tt.fields.webHook,
			}
			if err := r.Markdown(tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("Markdown() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRobot_Text(t *testing.T) {
	type fields struct {
		webHook string
	}
	type args struct {
		content             string
		mentionedList       []string
		mentionedMobileList []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "standard",
			fields: fields{
				webHook: testWebHook,
			},
			args: args{
				content:             "这是一条文本消息",
				mentionedList:       nil,
				mentionedMobileList: []string{All},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Robot{
				webHook: tt.fields.webHook,
			}
			if err := r.Text(tt.args.content, tt.args.mentionedList, tt.args.mentionedMobileList); (err != nil) != tt.wantErr {
				t.Errorf("Text() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
