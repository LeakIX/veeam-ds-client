package v10

import (
	"encoding/xml"
	"time"
)

type CacheRequest struct {
	RequestSpec
	FilePath string `xml:"FilePath,attr"`
}

type UploadRequest struct {
	RequestSpec
	SystemType                 string  `xml:"SystemType,attr"`
	Host                       string  `xml:"Host,attr"`
	User                       string  `xml:"User,attr"`
	Password                   string  `xml:"Password,attr"`
	TaskType                   string  `xml:"TaskType,attr"`
	FixProductType             string  `xml:"FixProductType,attr"`
	FixProductVeresion         string  `xml:"FixProductVeresion,attr"`
	FixIssueNumber             int     `xml:"FixIssueNumber,attr"`
	SshCredentials             string  `xml:"SshCredentials,attr"`
	SshFingerprint             string  `xml:"SshFingerprint,attr"`
	SshTrustAll                bool    `xml:"SshTrustAll,attr"`
	IsWindows                  bool    `xml:"IsWindows,attr"`
	IsFix                      bool    `xml:"IsFix,attr"`
	CheckSignatureBeforeUpload bool    `xml:"CheckSignatureBeforeUpload,attr"`
	DefaultProtocol            int     `xml:"DefaultProtocol,attr"`
	FileRelativePath           string  `xml:"FileRelativePath,attr"`
	FileProxyPath              string  `xml:"FileProxyPath,attr"`
	FileRemotePath             string  `xml:"FileRemotePath,attr"`
	FilePath                   string  `xml:"FilePath,attr"`
	HostIps                    HostIps `xml:"HostIps"`
}
type HostIps struct {
	String []String `xml:"String"`
}
type String struct {
	Value string `xml:"Value,attr"`
}

type RequestSpec struct {
	XMLName       xml.Name `xml:"EpCmMessage"`
	EpCmScopeData int      `xml:"EpCmScopeData,attr"`
}

type AgentIdentifier struct {
	BiosUuid              string `xml:"BiosUuid,attr"`
	CertificateThumbprint string `xml:"CertificateThumbprint,attr"`
}
type LicenseInfo struct {
	Mode                  uint32    `xml:"Mode,attr"`
	ExpirationDate        time.Time `xml:"ExpirationDate,attr"`
	SupportExpirationDate time.Time `xml:"SupportExpirationDate,attr"`
	Support               bool      `xml:"Support,attr"`
	Workstations          uint32    `xml:"Workstations,attr"`
	Servers               uint32    `xml:"Servers,attr"`
	ContactName           string    `xml:"ContactName,attr"`
	ContactMail           string    `xml:"ContactMail,attr"`
	ProductVersion        string    `xml:"ProductVersion,attr"`
}

type SendAgentLicenseInfo struct {
	RequestSpec
	AgentIdentifier AgentIdentifier `xml:"AgentIdentifier"`
	LicenseInfo     LicenseInfo     `xml:"LicenseInfo"`
}

type ConfigurationServiceInitialize struct {
	RequestSpec
	DbConfig                         string `xml:"DbConfig,attr"`
	VbrServerName                    string `xml:"VbrServerName,attr"`
	ForeignInvokerNegotiateProxyPort uint32 `xml:"ForeignInvokerNegotiateProxyPort,attr"`
	ForeignInvokerSslProxyPort       uint32 `xml:"ForeignInvokerSslProxyPort,attr"`
}
