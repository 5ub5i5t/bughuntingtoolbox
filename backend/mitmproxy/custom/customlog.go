package custom

import (
	"fmt"
	"time"

	"github.com/kardianos/mitmproxy/proxy"

	log "github.com/sirupsen/logrus"
)

// LogAddon log connection and flow
type CustomLogAddon struct {
	proxy.BaseAddon
}

func (addon *CustomLogAddon) Request(f *proxy.Flow) {
	start := time.Now()
	go func() {
		<-f.Done()
		var StatusCode int
		if f.Response != nil {
			StatusCode = f.Response.StatusCode
		}
		var contentLen int
		if f.Response != nil && f.Response.Body != nil {
			contentLen = len(f.Response.Body)
		}
		log.Infof("%v %v %v %v %v %v - %v ms\n", f.ConnContext.ClientConn.Conn.RemoteAddr(), f.Request.Method, f.Request.URL.String(), StatusCode, contentLen, time.Since(start).Milliseconds(), f.Request.Body)
		fmt.Println("")
	}()
}

func (addon *CustomLogAddon) Response(f *proxy.Flow) {
	start := time.Now()
	go func() {
		<-f.Done()
		var StatusCode int
		if f.Response != nil {
			StatusCode = f.Response.StatusCode
		}
		var contentLen int
		if f.Response != nil && f.Response.Body != nil {
			contentLen = len(f.Response.Body)
		}
		log.Infof("Response: %v %v %v %v %v - %v ms\n", f.ConnContext.ClientConn.Conn.RemoteAddr(), f.Request.Method, f.Request.URL.String(), StatusCode, contentLen, time.Since(start).Milliseconds())
		fmt.Println("")
	}()
}
