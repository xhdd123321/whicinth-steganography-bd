package utils

import (
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
)

const (
	XForwardedFor = "X-Forwarded-For"
	XRealIP       = "X-Real-IP"
)

// RemoteIp 返回远程客户端的 IP，如 192.168.1.1
func RemoteIp(c *app.RequestContext) string {
	remoteAddr := c.RemoteAddr().String()
	if ip := string(c.GetHeader(XRealIP)); ip != "" {
		remoteAddr = ip
	} else if ip = string(c.GetHeader(XForwardedFor)); ip != "" {
		remoteAddr = ip
	}

	if remoteAddr == "" {
		remoteAddr = "127.0.0.1"
	}
	remoteAddr = strings.Split(remoteAddr, ":")[0]
	return remoteAddr
}
