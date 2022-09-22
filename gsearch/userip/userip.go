package userip

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

// 为了防止key冲突
type key int

const userIPKey key = 0

func FromRequest(req *http.Request) (net.IP, error) {
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
	}
	return net.ParseIP(ip), nil
}

func NewContext(ctx context.Context, userIP net.IP) context.Context {
	return context.WithValue(ctx, userIPKey, userIP)
}

// 从context中取得
func FromContext(ctx context.Context) (net.IP, bool) {
	// ctx.Value returns nil if ctx has no value for the key;
	// the net.IP type assertion returns ok=false for nil.
	userIP, ok := ctx.Value(userIPKey).(net.IP)
	return userIP, ok
}
