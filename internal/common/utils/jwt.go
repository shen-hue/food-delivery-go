// Package utils 对应 Java 版 sky-common 的 com.sky.utils 包。
package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// CreateJWT 生成jwt令牌，使用HS256算法（对应 Java 版 JwtUtil.createJWT）。
//
// 参数：
//   - secretKey: jwt秘钥
//   - ttlMillis: jwt过期时间(毫秒)
//   - claims: 设置的信息，如 {"empId": 1}
func CreateJWT(secretKey string, ttlMillis int64, claims jwt.MapClaims) (string, error) {
	now := time.Now()
	claims["exp"] = now.Add(time.Duration(ttlMillis) * time.Millisecond).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// ParseJWT 解析jwt令牌（对应 Java 版 JwtUtil.parseJWT）。
// 令牌无效、签名不符或已过期时返回错误。
//
// 注意：秘钥一定要保留好在服务端，不能暴露出去，
// 否则sign就可以被伪造。如果对接多个客户端建议改造成多个。
func ParseJWT(secretKey, tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		// 校验签名算法，防止算法混淆攻击
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
