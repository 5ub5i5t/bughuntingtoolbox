package custom

import (
	"bytes"
	"fmt"
	"os"
	"unicode/utf8"

	"5ub5i5t/bughuntingtoolbox/model"

	"github.com/kardianos/mitmproxy/proxy"

	log "github.com/sirupsen/logrus"
)

func createKeyValuePairs(m map[string][]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
	}
	return b.String()
}

// LogAddon log connection and flow
type SaveFlowAddon struct {
	proxy.BaseAddon
}

func (addon *SaveFlowAddon) Request(f *proxy.Flow) {
	//start := time.Now()
	go func() {
		<-f.Done()

		var flow model.CustomFlow

		flow.Type = "Request"

		flow.FlowId = f.Id.String()

		flow.RequestMethod = f.Request.Method
		flow.RequestProto = f.Request.Proto
		flow.RequestHeaders = fmt.Sprintf("%v", f.Request.Header)
		flow.RequestURL = f.Request.URL.String()

		//RequestBody         []byte `gorm:"type:text"`
		//flow.RequestRaw = fmt.Sprintf("%v", f.Request.Raw)
		//flow.RequestRaw = string(f.Request.content)
		flow.RequestUrlScheme = f.Request.URL.Scheme
		flow.RequestUrlOpaque = f.Request.URL.Opaque
		//flow.RequestUrlUser =
		//flow.RequestUrlUserPassword = f.Request.URL.User.Password()
		flow.RequestUrlHost = f.Request.URL.Host
		flow.RequestUrlPath = f.Request.URL.Path
		flow.RequestUrlRawPath = f.Request.URL.RawPath
		flow.RequestUrlOmitHost = f.Request.URL.OmitHost
		flow.RequestUrlForceQuery = f.Request.URL.ForceQuery
		flow.RequestUrlRawQuery = f.Request.URL.RawQuery
		flow.RequestUrlFragment = f.Request.URL.Fragment
		flow.RequestUrlRawFragment = f.Request.URL.RawFragment

		flow.RequestContentLength = f.Request.Raw().ContentLength
		flow.RequestUserAgent = f.Request.Raw().UserAgent()

		flow.Save()

		//savedEntry, err := flow.Save()

		//if err != nil {
		//	log.Fatalln(err)
		//}

		//if savedEntry != nil {
		//fmt.Printf("Record saved.")
		//}
	}()
}

func (addon *SaveFlowAddon) Response(f *proxy.Flow) {
	//start := time.Now()
	go func() {
		<-f.Done()
		var StatusCode int
		if f.Response != nil {
			StatusCode = f.Response.StatusCode
		}
		var ResponseContentLen int
		if f.Response != nil && f.Response.Body != nil {
			ResponseContentLen = len(f.Response.Body)
		}
		//log.Infof("Response: %v %v %v %v %v - %v ms\n", f.ConnContext.ClientConn.Conn.RemoteAddr(), f.Request.Method, f.Request.URL.String(), StatusCode, contentLen, time.Since(start).Milliseconds())
		//fmt.Println("")

		var flow model.CustomFlow

		flow.Type = "Response"

		flow.FlowId = f.Id.String()

		flow.RequestMethod = f.Request.Method
		flow.RequestProto = f.Request.Proto
		flow.RequestHeaders = fmt.Sprintf("%v", f.Request.Header)
		flow.RequestURL = f.Request.URL.String()

		//RequestBody         []byte `gorm:"type:text"`
		//flow.RequestRaw = fmt.Sprintf("%v", f.Request.Raw)
		//flow.RequestRaw = string(f.Request.content)
		flow.RequestUrlScheme = f.Request.URL.Scheme
		flow.RequestUrlOpaque = f.Request.URL.Opaque
		//flow.RequestUrlUser =
		//flow.RequestUrlUserPassword = f.Request.URL.User.Password()
		flow.RequestUrlHost = f.Request.URL.Host
		flow.RequestUrlPath = f.Request.URL.Path
		flow.RequestUrlRawPath = f.Request.URL.RawPath
		flow.RequestUrlOmitHost = f.Request.URL.OmitHost
		flow.RequestUrlForceQuery = f.Request.URL.ForceQuery
		flow.RequestUrlRawQuery = f.Request.URL.RawQuery
		flow.RequestUrlFragment = f.Request.URL.Fragment
		flow.RequestUrlRawFragment = f.Request.URL.RawFragment

		//flow.RequestContentLength = f.Request.Raw().ContentLength
		flow.RequestUserAgent = f.Request.Raw().UserAgent()

		flow.ResponseStatusCode = StatusCode
		flow.ResponseContentLength = ResponseContentLen

		//errStr = strings.ToValidUTF8(errStr, "")
		//flow.ResponseBody = strings.ToValidUTF8(string(f.Response.Body), "")
		//flow.ResponseBody = fmt.Sprintf("%v", string(f.Response.Body))

		if !utf8.Valid(f.Response.Body) {
			fmt.Println("Invalid UTF-8 byte sequence")
			flow.ResponseBody = ""
		} else {
			flow.ResponseBody = string(f.Response.Body)
		}

		flow.ResponseHeaders = fmt.Sprintf("%v", f.Response.Header)

		savedEntry, err := flow.Save()

		if err != nil {
			log.Fatalln(err)
		}

		if savedEntry != nil {
			//fmt.Printf("Record saved.")
			writeToFile(flow, "testfile.txt")
		}

	}()
}

func writeToFile(flow model.CustomFlow, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write struct data to file
	_, err = fmt.Fprintf(file, "Host: %s\nURL: %d\n", flow.RequestUrlHost, flow.RequestURL)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
