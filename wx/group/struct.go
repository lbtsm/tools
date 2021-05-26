package group

type RobotMsgBack struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type NewsMsg struct {
	Title       string `json:"title"`       // 标题，不超过128个字节，超过会自动截断
	Description string `json:"description"` // 描述，不超过512个字节，超过会自动截断
	Url         string `json:"url"`         // 点击后跳转的链接。
	PicUrl      string `json:"pic_url"`     // 图文消息的图片链接，支持JPG、PNG格式，较好的效果为大图 1068*455，小图150*150。
}
