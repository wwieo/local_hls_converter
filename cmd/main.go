package main

import (
	"main/service"
)

// @title		HLS_PROVIDER
// @version	1.0
// @basePath	/hls_provider
// @schemes	http
func main() {
	service.HlsProvider().Run()
}
