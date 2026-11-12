package utils

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// httpClient 对应 Java 版 HttpClientUtil 中设置的超时时间（5秒）。
var httpClient = &http.Client{Timeout: 5 * time.Second}

// DoGet 发送GET请求（对应 HttpClientUtil.doGet）。
// 参数paramMap中的键值对会拼接为查询字符串。
func DoGet(rawURL string, paramMap map[string]string) (string, error) {
	if len(paramMap) > 0 {
		params := url.Values{}
		for k, v := range paramMap {
			params.Set(k, v)
		}
		rawURL = rawURL + "?" + params.Encode()
	}

	resp, err := httpClient.Get(rawURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// DoPost 发送POST请求（表单提交，对应 HttpClientUtil.doPost）。
func DoPost(rawURL string, paramMap map[string]string) (string, error) {
	form := url.Values{}
	for k, v := range paramMap {
		form.Set(k, v)
	}

	resp, err := httpClient.Post(rawURL,
		"application/x-www-form-urlencoded",
		strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// DoPost4JSON 发送POST请求（JSON提交，对应 HttpClientUtil.doPost4Json）。
func DoPost4JSON(rawURL string, jsonBody string) (string, error) {
	resp, err := httpClient.Post(rawURL,
		"application/json",
		strings.NewReader(jsonBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
