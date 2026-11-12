package utils

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"

	"sky-take-out/internal/common/properties"
)

// AliOssUtil 阿里云OSS工具类（对应 Java 版的 com.sky.utils.AliOssUtil）。
// 实现管理端/公共文件上传接口（CommonController）时使用。
type AliOssUtil struct {
	props *properties.AliOssProperties
}

// NewAliOssUtil 构造函数（对应 Java 版的 @Autowired 注入 AliOssProperties）。
func NewAliOssUtil(props *properties.AliOssProperties) *AliOssUtil {
	return &AliOssUtil{props: props}
}

// Upload 文件上传（对应 Java 版 AliOssUtil.upload）。
//
// 参数：
//   - data: 上传文件的字节数组
//   - objectName: 上传到OSS的文件名（含路径），如 "dish/2026/11/10/xxx.jpg"
//
// 返回上传后文件的访问URL：https://{bucket}.{endpoint}/{objectName}
func (u *AliOssUtil) Upload(ctx context.Context, data []byte, objectName string) (string, error) {
	// Go版使用官方SDK v2（github.com/aliyun/alibabacloud-oss-go-sdk-v2），
	// 需要region；从endpoint推导，如 oss-cn-hangzhou.aliyuncs.com → cn-hangzhou
	cfg := oss.LoadDefaultConfig().
		WithEndpoint(u.props.Endpoint).
		WithRegion(regionFromEndpoint(u.props.Endpoint)).
		WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				u.props.AccessKeyID, u.props.AccessKeySecret))

	client := oss.NewClient(cfg)

	_, err := client.PutObject(ctx, &oss.PutObjectRequest{
		Bucket: oss.Ptr(u.props.BucketName),
		Key:    oss.Ptr(objectName),
		Body:   bytes.NewReader(data),
	})
	if err != nil {
		return "", fmt.Errorf("上传文件失败: %w", err)
	}

	return fmt.Sprintf("https://%s.%s/%s",
		u.props.BucketName, u.props.Endpoint, objectName), nil
}

// regionFromEndpoint 从endpoint推导region。
func regionFromEndpoint(endpoint string) string {
	host := strings.Split(endpoint, ".")[0] // 如 oss-cn-hangzhou
	return strings.TrimPrefix(host, "oss-") // 如 cn-hangzhou
}

// BuildObjectName 按当前日期生成OSS对象名，避免文件重名覆盖
// （对应 Java 版 AliOssUtil.upload 中按日期构造文件名的逻辑）。
func BuildObjectName(now time.Time, ext string) string {
	return fmt.Sprintf("%d/%02d/%02d/%d%s",
		now.Year(), int(now.Month()), now.Day(), now.UnixNano(), ext)
}
