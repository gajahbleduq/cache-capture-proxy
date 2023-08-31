package handler

import (
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

func ProxyHandler(c *gin.Context) {
	// Menampilkan cache jika tersedia
	if cachedData := GetCachedData(c.Request.URL.String()); cachedData != nil {
		c.Header("X-Cache", "HIT")
		c.Writer.Write(cachedData)
		return
	}

	// Mengubah header host agar sesuai dengan URL target
	c.Request.Host = targetURL.Host

	// Mengambil data dari server tujuan
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.ServeHTTP(c.Writer, c.Request)

	// Mengambil respons yang diterima dan menyimpannya didalam cache
	cachedResponse := CaptureResponse(c.Writer)
	if cachedResponse != nil {
		CacheResponse(c.Request.URL.String(), cachedResponse)
	}

}
