package msg

import (
	"github.com/webx-top/echo"
)

type Message struct {
	TimeLocal                                   string
	RemoteAddr, XRealIP, XForwardFor, LocalAddr string
	User, Version                               string
	Referer, UserAgent, Path, Method, Scheme    string
	BrowerName, BrowerType                      string
	BytesSent                                   uint64
	StatusCode                                  uint
	UpstreamTime, RequestTime                   float64
}

func (m *Message) SetByMap(data echo.Store) *Message {
	m.TimeLocal = data.String(`TimeLocal`)
	m.RemoteAddr = data.String(`RemoteAddr`)
	m.XRealIP = data.String(`XRealIP`)
	m.XForwardFor = data.String(`XForwardFor`)
	m.LocalAddr = data.String(`LocalAddr`)
	m.User = data.String(`User`)
	m.Version = data.String(`Version`)
	m.Referer = data.String(`Referer`)
	m.UserAgent = data.String(`UserAgent`)
	m.Path = data.String(`Path`)
	m.Method = data.String(`Method`)
	m.Scheme = data.String(`Scheme`)
	m.BrowerName = data.String(`BrowerName`)
	m.BrowerType = data.String(`BrowerType`)
	m.BytesSent = data.Uint64(`BytesSent`)
	m.StatusCode = data.Uint(`StatusCode`)
	m.UpstreamTime = data.Float64(`UpstreamTime`)
	m.RequestTime = data.Float64(`RequestTime`)
	return m
}

func (m *Message) ToMap() echo.Store {
	data := echo.Store{}
	data.Set(`TimeLocal`, m.TimeLocal)
	data.Set(`RemoteAddr`, m.RemoteAddr)
	data.Set(`XRealIP`, m.XRealIP)
	data.Set(`XForwardFor`, m.XForwardFor)
	data.Set(`LocalAddr`, m.LocalAddr)
	data.Set(`User`, m.User)
	data.Set(`Version`, m.Version)
	data.Set(`Referer`, m.Referer)
	data.Set(`UserAgent`, m.UserAgent)
	data.Set(`Path`, m.Path)
	data.Set(`Method`, m.Method)
	data.Set(`Scheme`, m.Scheme)
	data.Set(`BrowerName`, m.BrowerName)
	data.Set(`BrowerType`, m.BrowerType)
	data.Set(`BytesSent`, m.BytesSent)
	data.Set(`StatusCode`, m.StatusCode)
	data.Set(`UpstreamTime`, m.UpstreamTime)
	data.Set(`RequestTime`, m.RequestTime)
	return data
}
