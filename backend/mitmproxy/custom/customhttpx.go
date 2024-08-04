package custom

import (
	"fmt"
	"time"

	"github.com/kardianos/mitmproxy/proxy"

	"github.com/projectdiscovery/goflags"
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
	"github.com/projectdiscovery/httpx/runner"
	log "github.com/sirupsen/logrus"
)

// LogAddon log connection and flow
type CustomHttpxAddon struct {
	proxy.BaseAddon
}

func (addon *CustomHttpxAddon) Response(f *proxy.Flow) {
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

		gologger.DefaultLogger.SetMaxLevel(levels.LevelVerbose) // increase the verbosity (optional)

		options := runner.Options{
			Methods: "GET",
			//StoreResponseDir: "data",
			InputTargetHost: goflags.StringSlice{f.Request.URL.String()},
			OnResult: func(r runner.Result) {
				// handle error
				if r.Err != nil {
					fmt.Printf("[Err] %s: %s\n", r.Input, r.Err)
					return
				}
				//fmt.Println("")
				//fmt.Println("RESULT:")
				//fmt.Printf("%s %s %d\n", r.Input, r.Host, r.StatusCode)
				//fmt.Printf("%+v\n", r)
				//fmt.Println("")

			},
		}

		if err := options.ValidateOptions(); err != nil {
			log.Fatal(err)
		}

		httpxRunner, err := runner.New(&options)
		if err != nil {
			log.Fatal(err)
		}
		defer httpxRunner.Close()

		httpxRunner.RunEnumeration()

	}()
}
