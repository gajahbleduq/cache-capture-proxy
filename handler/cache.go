package handler

import (
	"net/http"
	"sync"
)

var (
	cache      = make(map[string][]byte)
	cacheMutex sync.RWMutex
)

func GetCachedData(url string) []byte {
	cacheMutex.RLock()
	defer cacheMutex.RUnlock()
	return cache[url]
}

func CacheResponse(url string, data []byte) {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	cache[url] = data
}

func CaptureResponse(w http.ResponseWriter) []byte {
	// Menangkap semua output dari ResponseWriter
	capturedWriter := &capturedResponseWriter{ResponseWriter: w}
	return capturedWriter.capturedData
}

type capturedResponseWriter struct {
	http.ResponseWriter
	capturedData []byte
}

func (c *capturedResponseWriter) Write(data []byte) (int, error) {
	c.capturedData = append(c.capturedData, data...)
	return c.ResponseWriter.Write(data)
}
