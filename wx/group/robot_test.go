package group

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func getImageMessage() (string, string) {
	// 图片的
	f, _ := ioutil.ReadFile("/Users/zmm/Desktop/other/235000-1584114600db79.jpg")
	picBase64 := base64.StdEncoding.EncodeToString(f)
	picMd5 := fmt.Sprintf("%x", md5.Sum(f))
	return picBase64, picMd5
}

func getNewsMessage() string {
	news := []NewsMsg{
		{
			Title:       "这是一条图文消息",
			Description: "洗澡啦，洗澡啦",
			Url:         "https://work.weixin.qq.com/api/doc/90000/90136/91770#%E6%96%87%E6%9C%AC%E7%B1%BB%E5%9E%8B",
			PicUrl:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
		},
		{
			Title:       "这是二条图文消息",
			Description: "洗澡啦，洗澡啦",
			Url:         "https://work.weixin.qq.com/api/doc/90000/90136/91770#%E6%96%87%E6%9C%AC%E7%B1%BB%E5%9E%8B",
			PicUrl:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
		},
		{
			Title:       "这是三条图文消息",
			Description: "洗澡啦，洗澡啦",
			Url:         "https://work.weixin.qq.com/api/doc/90000/90136/91770#%E6%96%87%E6%9C%AC%E7%B1%BB%E5%9E%8B",
			PicUrl:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
		},
		{
			Title:       "这是四条图文消息",
			Description: "洗澡啦，洗澡啦",
			Url:         "https://work.weixin.qq.com/api/doc/90000/90136/91770#%E6%96%87%E6%9C%AC%E7%B1%BB%E5%9E%8B",
			PicUrl:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
		},
		{
			Title:       "这是五条图文消息",
			Description: "洗澡啦，洗澡啦",
			Url:         "https://work.weixin.qq.com/api/doc/90000/90136/91770#%E6%96%87%E6%9C%AC%E7%B1%BB%E5%9E%8B",
			PicUrl:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
		},
		{
			Title:       "这是六条图文消息",
			Description: "洗澡啦，洗澡啦",
			Url:         "https://work.weixin.qq.com/api/doc/90000/90136/91770#%E6%96%87%E6%9C%AC%E7%B1%BB%E5%9E%8B",
			PicUrl:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
		},
		{
			Title:       "这是七条图文消息",
			Description: "洗澡啦，洗澡啦",
			Url:         "https://work.weixin.qq.com/api/doc/90000/90136/91770#%E6%96%87%E6%9C%AC%E7%B1%BB%E5%9E%8B",
			PicUrl:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
		},
		{
			Title:       "这是八条图文消息",
			Description: "洗澡啦，洗澡啦",
			Url:         "https://work.weixin.qq.com/api/doc/90000/90136/91770#%E6%96%87%E6%9C%AC%E7%B1%BB%E5%9E%8B",
			PicUrl:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
		},
		// 超过九条则消息发送失败
		// {
		// 	Title:       "这是九条图文消息",
		// 	Description: "洗澡啦，洗澡啦",
		// 	Url:         "https://work.weixin.qq.com/api/doc/90000/90136/91770#%E6%96%87%E6%9C%AC%E7%B1%BB%E5%9E%8B",
		// 	PicUrl:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
		// },
	}

	data, _ := json.Marshal(news)
	return string(data)
}
