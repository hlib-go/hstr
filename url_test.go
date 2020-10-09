package hstr

import "testing"

func TestUrlAddParams(t *testing.T) {
	t.Log(UrlAddParams("http://sss.com/sss", "code", "vv"))
	t.Log(UrlAddParams("http://sss.com/sss?k1=v1", "code", "vv"))
	t.Log(UrlAddParams("http://sss.com/sss#/ss?k1=v1", "code", "vv"))
	t.Log(UrlAddParams("http://sss.com/sss?k1=v1#/ss?k1=v1", "code", "vv"))
}

func TestUrlGetParams(t *testing.T) {
	t.Log(UrlGetParams("http://sss.com/sss?k1=v1&k2=v2", "k1"))
	t.Log(UrlGetParams("http://sss.com/sss?k1=v1&k2=v2", "k2"))
	t.Log(UrlGetParams("http://sss.com/sss?k1=v1&k2=v2#/sss?kk1=vv1&kk2=vv2", "kk1"))
	t.Log(UrlGetParams("http://sss.com/sss?k1=v1&k2=v2#/sss?kk1=vv1&kk2=vv2", "kk2"))
	t.Log(UrlGetParams("http://sss.com/sss?k1=v1&k2=v2#/sss?k1=vv1", "k1"))
	t.Log(UrlGetParams("http://sss.com/sss", "k1"))
}
