# Veeam Distribution Service Golang Client

This library provides a `net.Conn` abstraction to a Veeam compressed connection trough NNS.

```golang
ntlmsspClient, err := ntlmssp.NewClient(ntlmssp.SetCompatibilityLevel(1), ntlmssp.SetUserInfo("", ""))
if err != nil {
    return nil, err
}
nnsConn, err := nns.DialNTLMSSP(addr, ntlmsspClient, 5*time.Second)
if err != nil {
    return nil, err
}
conn := veeam.WrapConnection(nnsConn)
```

An example to download files can be found in [cmd/veeamdl](cmd/veeamdl)