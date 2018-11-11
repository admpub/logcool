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
	m.TimeLocal = data.Get(`TimeLocal`).String()
	m.RemoteAddr = data.Get(`RemoteAddr`).String()
	m.XRealIP = data.Get(`XRealIP`).String()
	m.XForwardFor = data.Get(`XForwardFor`).String()
	m.LocalAddr = data.Get(`LocalAddr`).String()
	m.User = data.Get(`User`).String()
	m.Version = data.Get(`Version`).String()
	m.Referer = data.Get(`Referer`).String()
	m.UserAgent = data.Get(`UserAgent`).String()
	m.Path = data.Get(`Path`).String()
	m.Method = data.Get(`Method`).String()
	m.Scheme = data.Get(`Scheme`).String()
	m.BrowerName = data.Get(`BrowerName`).String()
	m.BrowerType = data.Get(`BrowerType`).String()
	m.BytesSent = data.Get(`BytesSent`).Uint64()
	m.StatusCode = data.Get(`StatusCode`).Uint()
	m.UpstreamTime = data.Get(`UpstreamTime`).Float64()
	m.RequestTime = data.Get(`RequestTime`).Float64()
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
