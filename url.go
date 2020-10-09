package hstr

import (
	"regexp"
	"strings"
)

// 根据key取URL中参数, 参数不存在返回空字符串
func UrlGetParams(url, k string) (v string) {
	strs := regexp.MustCompile("(\\?|&)" + k + "=([^&]*)(&|$)").FindStringSubmatch(url)
	if len(strs) < 3 {
		return
	}
	v = strs[2]
	return
}

// URL拼接参数
func UrlAddParams(url, k, v string) string {
	if strings.Contains(url, "#") {
		ss := strings.Split(url, "#")
		s2 := ss[1]
		if strings.Contains(s2, "?") {
			s2 += "&"
		} else {
			s2 += "?"
		}
		s2 += k + "=" + v
		return ss[0] + "#" + s2
	}

	if strings.Contains(url, "?") {
		url += "&"
	} else {
		url += "?"
	}
	url += k + "=" + v
	return url
}

// URL设置参数，存在则替换
func UrlSetParams(url, k, v string) string {
	kpr := k + "="
	if strings.Contains(url, kpr) {
		return strings.ReplaceAll(url, kpr+UrlGetParams(url, k), kpr+v)
	}
	return UrlAddParams(url, k, v)
}
