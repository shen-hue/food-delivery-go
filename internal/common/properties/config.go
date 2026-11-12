// Package properties 对应 Java 版 sky-common 的 com.sky.properties 包，
// 并承担 Spring Boot 的配置文件加载职责（application.yaml + profile覆盖）。
package properties

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// DatasourceProperties 数据源配置（对应 application-dev.yml 中的 sky.datasource）。
type DatasourceProperties struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// Config 应用配置根结构，对应 application.yml 的结构。
type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`

	Sky struct {
		Datasource DatasourceProperties `yaml:"datasource"`
		Jwt        JwtProperties        `yaml:"jwt"`
		AliOSS     AliOssProperties     `yaml:"alioss"`
		Wechat     WeChatProperties     `yaml:"wechat"`
	} `yaml:"sky"`
}

// Load 加载配置：先读取 config/application.yaml，
// 再用 config/application-{profile}.yaml 覆盖同名配置
// （相当于 Spring Boot 的 profile 机制）。
//
// profile 通过环境变量 SKY_PROFILE 指定，默认 dev。
func Load(profile string) (*Config, error) {
	cfg := &Config{}

	if err := loadFile("config/application.yaml", cfg); err != nil {
		return nil, err
	}
	if profile != "" {
		if err := loadFile(fmt.Sprintf("config/application-%s.yaml", profile), cfg); err != nil {
			return nil, err
		}
	}

	if cfg.Server.Port == 0 {
		cfg.Server.Port = 8080
	}
	return cfg, nil
}

// ProfileFromEnv 从环境变量读取当前profile。
func ProfileFromEnv() string {
	if p := os.Getenv("SKY_PROFILE"); p != "" {
		return p
	}
	return "dev"
}

// DSN 生成 PostgreSQL 连接串（pgx 使用）。
func (c *Config) DSN() string {
	ds := &c.Sky.Datasource
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		ds.Username, ds.Password, ds.Host, ds.Port, ds.Database)
}

func loadFile(path string, cfg *Config) error {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // 配置文件不存在时跳过（如没有profile文件）
		}
		return fmt.Errorf("读取配置文件 %s 失败: %w", path, err)
	}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return fmt.Errorf("解析配置文件 %s 失败: %w", path, err)
	}
	return nil
}
