package firewallconfigs

type GeeTestConfig struct {
	IsOn       bool   `yaml:"isOn" json:"isOn"`
	CaptchaId  string `yaml:"captchaId" json:"captchaId"`
	CaptchaKey string `yaml:"captchaKey" json:"captchaKey"`
}

type HTTPFirewallCaptchaAction struct {
	IsPrior bool `yaml:"isPrior" json:"isPrior"`

	CaptchaType CaptchaType `yaml:"captchaType" json:"captchaType"` // 类型

	Life              int32 `yaml:"life" json:"life"`                           // 有效期
	MaxFails          int   `yaml:"maxFails" json:"maxFails"`                   // 最大失败次数
	FailBlockTimeout  int   `yaml:"failBlockTimeout" json:"failBlockTimeout"`   // 失败拦截时间
	FailBlockScopeAll bool  `yaml:"failBlockScopeAll" json:"failBlockScopeAll"` // 是否全局有效

	// 验证码相关配置

	CountLetters int8 `yaml:"countLetters" json:"countLetters"` // 字符数量

	UIIsOn          bool   `yaml:"uiIsOn" json:"uiIsOn"`                   // 是否使用自定义UI
	UITitle         string `yaml:"uiTitle" json:"uiTitle"`                 // 消息标题
	UIPrompt        string `yaml:"uiPrompt" json:"uiPrompt"`               // 消息提示
	UIButtonTitle   string `yaml:"uiButtonTitle" json:"uiButtonTitle"`     // 按钮标题
	UIShowRequestId bool   `yaml:"uiShowRequestId" json:"uiShowRequestId"` // 是否显示请求ID
	UICss           string `yaml:"uiCss" json:"uiCss"`                     // CSS样式
	UIFooter        string `yaml:"uiFooter" json:"uiFooter"`               // 页脚
	UIBody          string `yaml:"uiBody" json:"uiBody"`                   // 内容轮廓

	CookieId string `yaml:"cookieId" json:"cookieId"` // TODO

	Lang string `yaml:"lang" json:"lang"` // 语言，zh-CN, en-US ... TODO 需要实现，目前是根据浏览器Accept-Language动态获取

	// 极验相关配置
	// MUST be struct
	GeeTestConfig GeeTestConfig `yaml:"geeTestConfig" json:"geeTestConfig"`
}

func NewHTTPFirewallCaptchaAction() *HTTPFirewallCaptchaAction {
	return &HTTPFirewallCaptchaAction{
		Life:              600,
		MaxFails:          100,
		FailBlockTimeout:  3600,
		FailBlockScopeAll: true,
		UIShowRequestId:   true,
	}
}
