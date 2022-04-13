package v10

import "encoding/xml"

type CacheReply struct {
	Reply
	IsFileInCache bool `xml:"IsFileInCache,attr"`
}

type UploadReply struct {
	Reply
	RemotePath string `xml:"RemotePath,attr"`
}

type Reply struct {
	XMLName              xml.Name       `xml:"EpCmResponse"`
	Exception            string         `xml:"Exception,attr"`
	PlainException       PlainException `xml:"PlainException"`
	PersistentConnection bool           `xml:"PersistentConnection,attr"`
}

type PlainException struct {
	XMLName  xml.Name   `xml:"PlainException"`
	KeyValue []KeyValue `xml:"KeyValue,allowempty"`
}

type KeyValue struct {
	XMLName xml.Name `xml:"KeyValue"`
	Key     string   `xml:"Key,attr"`
	Value   string   `xml:"Value,attr"`
}

func (vr *Reply) HasExceptions() bool {
	return len(vr.Exception) > 0
}
