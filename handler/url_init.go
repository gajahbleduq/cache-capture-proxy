package handler

import (
	"net/url"
)

var targetURL *url.URL

func init() {
	targetURL, _ = url.Parse("https://ordera.id/")
}
