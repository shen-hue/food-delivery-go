package properties

// AliOssProperties 阿里云OSS配置
// （对应 Java 版的 @ConfigurationProperties("sky.alioss")，
// 实现 CommonController 文件上传功能时使用）。
type AliOssProperties struct {
	Endpoint        string `yaml:"endpoint"`          // 域名
	AccessKeyID     string `yaml:"access-key-id"`     // AccessKey ID
	AccessKeySecret string `yaml:"access-key-secret"` // AccessKey Secret
	BucketName      string `yaml:"bucket-name"`       // 存储空间名
}
