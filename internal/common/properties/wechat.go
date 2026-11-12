package properties

// WeChatProperties 微信配置
// （对应 Java 版的 @ConfigurationProperties("sky.wechat")，
// 实现微信登录/微信支付功能时使用）。
type WeChatProperties struct {
	Appid              string `yaml:"appid"`                 // 小程序的appid
	Secret             string `yaml:"secret"`                // 小程序的秘钥
	Mchid              string `yaml:"mchid"`                 // 商户号
	MchSerialNo        string `yaml:"mch-serial-no"`         // 商户API证书的证书序列号
	PrivateKeyFilePath string `yaml:"private-key-file-path"` // 商户私钥文件
	APIV3Key           string `yaml:"api-v3-key"`            // APIv3密钥
	WeChatPayCertPath  string `yaml:"wechat-pay-cert-path"`  // 微信支付平台证书文件路径
	NotifyURL          string `yaml:"notify-url"`            // 支付成功的回调地址
	RefundNotifyURL    string `yaml:"refund-notify-url"`     // 退款成功的回调地址
}
