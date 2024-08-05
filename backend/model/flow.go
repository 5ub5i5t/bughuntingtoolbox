package model

import (
	"5ub5i5t/bughuntingtoolbox/database"

	"gorm.io/gorm"
)

type CustomFlow struct {
	gorm.Model
	Type           string `gorm:"type:text"`
	FlowId         string `gorm:"type:text"`
	RequestMethod  string `gorm:"type:text"`
	RequestURL     string `gorm:"type:text"`
	RequestProto   string `gorm:"type:text"`
	RequestHeaders string `gorm:"type:text"`
	//RequestBody         []byte `gorm:"type:text"`
	RequestRaw string `gorm:"type:text"`

	RequestUrlScheme       string `gorm:"type:text"`
	RequestUrlOpaque       string `gorm:"type:text"` // encoded opaque data
	RequestUrlUserInfo     string `gorm:"type:text"` // username and password information
	RequestUrlUser         string `gorm:"type:text"` // username and password information
	RequestUrlUserPassword string `gorm:"type:text"` // username and password information
	RequestUrlHost         string `gorm:"type:text"` // host or host:port (see Hostname and Port methods)
	RequestUrlPath         string `gorm:"type:text"` // path (relative paths may omit leading slash)
	RequestUrlRawPath      string `gorm:"type:text"` // encoded path hint (see EscapedPath method)
	RequestUrlOmitHost     bool   // do not emit empty host (authority)
	RequestUrlForceQuery   bool   // append a query ('?') even if RawQuery is empty
	RequestUrlRawQuery     string `gorm:"type:text"` // encoded query values, without '?'
	RequestUrlFragment     string `gorm:"type:text"` // fragment for references, without '#'
	RequestUrlRawFragment  string `gorm:"type:text"` // encoded fragment hint (see EscapedFragment method)

	RequestContentLength int64
	RequestUserAgent     string `gorm:"type:text"`

	ResponseStatusCode int
	ResponseHeader     map[string][]string `gorm:"type:text"`
	//ResponseBody        []byte `gorm:"type:text"`
	//ResponseDecodedBody []byte `gorm:"type:text"`
	ResponseDecoded       bool
	ResponseDecodedErr    string `gorm:"type:text"`
	ResponseContentLength int
	ResponseBody          string `gorm:"type:text"`
	ResponseHeaders       string `gorm:"type:text"`
}

func (flow *CustomFlow) Save() (*CustomFlow, error) {
	err := database.Database.Create(&flow).Error
	if err != nil {
		return &CustomFlow{}, err
	}
	return flow, nil
}
