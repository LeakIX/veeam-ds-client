package scp_server

import (
	"bufio"
	"github.com/gliderlabs/ssh"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func StartScpServer() {
	ssh.Handle(func(s ssh.Session) {
		log.Println("Received remoted connection")
		if len(s.Command()) < 1 || s.Command()[0] != "scp" || s.Subsystem() != "" {
			s.Close()
			return
		}
		s.Write([]byte{0x00})
		bufreader := bufio.NewReader(s)
		cmdLine, err := bufreader.ReadString('\n')
		if err != nil {
			return
		}
		cmdParts := strings.Split(cmdLine, " ")
		if len(cmdParts) != 3 {
			log.Println("didn't receive a valid copy command")
			return
		}
		cmd, filename := cmdParts[0], strings.TrimRight(cmdParts[2], "\n")
		if cmd != "C0644" {
			log.Println("Didn't receive c0644 command")
			return
		}
		size, err := strconv.Atoi(cmdParts[1])
		if err != nil {
			log.Println("Didn't receive valid  size: " + err.Error())
			return
		}
		if size < 1 {
			log.Println("Didn't receive valid  size")
			return
		}
		log.Printf("Receiving file '%s' of size %d", filename, size)
		s.Write([]byte{0x00})
		_, err = io.CopyN(os.Stdout, s, int64(size))
		if err != nil {
			log.Println("Didn't receive file correctly: " + err.Error())
			return
		}
		s.Write([]byte{0x00})
	})
	log.Fatal(ssh.ListenAndServe(":2222", nil))
}
