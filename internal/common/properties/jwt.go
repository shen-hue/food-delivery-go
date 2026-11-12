package properties

// JwtProperties JWT配置（对应 Java 版的 @ConfigurationProperties("sky.jwt")）。
type JwtProperties struct {
	AdminSecretKey string `yaml:"admin-secret-key"` // 管理端jwt秘钥
	AdminTTL       int64  `yaml:"admin-ttl"`        // 管理端jwt过期时间（毫秒）
	AdminTokenName string `yaml:"admin-token-name"` // 管理端jwt令牌请求头名称

	// 用户端（微信用户）配置 —— 实现用户登录功能时使用
	UserSecretKey string `yaml:"user-secret-key"` // 用户端jwt秘钥
	UserTTL       int64  `yaml:"user-ttl"`        // 用户端jwt过期时间（毫秒）
	UserTokenName string `yaml:"user-token-name"` // 用户端jwt令牌请求头名称
}
