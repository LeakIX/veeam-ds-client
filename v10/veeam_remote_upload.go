package v10

import (
	"encoding/xml"
	"github.com/LeakIX/nns"
	"github.com/LeakIX/ntlmssp"
	vdsclient "github.com/LeakIX/veeam-ds-client"
	"net"
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

func SendRequest(addr string, input interface{}, output interface{}) error {
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
	return xmlDecoder.Decode(output)
}
