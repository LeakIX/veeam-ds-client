package veeam_ds_client

import (
	"bytes"
	"compress/flate"
	"encoding/binary"
	"errors"
	"io"
	"log"
	"net"
	"time"
)

// TODO : implement the actual compression, uncompressed works until 4096 bytes

type CompressedConn struct {
	tcpConn net.Conn
	buffer  *bytes.Buffer
}

const (
	IsCompressed uint32 = 2147483648
)

func WrapConnection(conn net.Conn) net.Conn {
	return &CompressedConn{
		tcpConn: conn,
		buffer:  bytes.NewBuffer([]byte{}),
	}
}

func (conn *CompressedConn) Read(b []byte) (n int, err error) {
	if conn.buffer.Len() != 0 {
		return conn.buffer.Read(b)
	}
	conn.buffer.Reset()
	var frameCompressedSize uint32
	var frameUncompressedSize uint32
	// First packet is total frame size
	err = binary.Read(conn.tcpConn, binary.LittleEndian, &frameUncompressedSize)
	err = binary.Read(conn.tcpConn, binary.LittleEndian, &frameCompressedSize)
	isCompressed := frameUncompressedSize&IsCompressed == IsCompressed
	if !isCompressed {
		log.Printf("frame size C/U : %d/%d",
			frameCompressedSize, frameUncompressedSize)
		written, err := io.CopyN(conn.buffer, conn.tcpConn, int64(frameCompressedSize))
		if err != nil {
			return n, err
		}
		if uint32(written) != uint32(frameUncompressedSize) {
			return 0, errors.New("could not retrieve full packet")
		}
	} else {
		frameUncompressedSize ^= IsCompressed
		log.Printf("frame size C/U : %d/%d",
			frameCompressedSize, frameUncompressedSize)
		compressedReader := flate.NewReader(conn.tcpConn)
		written, err := io.CopyN(conn.buffer, compressedReader, int64(frameUncompressedSize))
		if err != nil {
			return n, err
		}
		if uint32(written) != uint32(frameUncompressedSize) {
			return 0, errors.New("could not retrieve full packet")
		}
	}
	log.Println(conn.buffer.String())
	return conn.buffer.Read(b)
}

func (conn *CompressedConn) Write(b []byte) (n int, err error) {
	payloadBuffer := bytes.NewBuffer(nil)
	err = binary.Write(payloadBuffer, binary.LittleEndian, uint32(0))
	err = binary.Write(payloadBuffer, binary.LittleEndian, uint32(len(b)))
	_, err = payloadBuffer.Write(b)
	_, err = payloadBuffer.WriteTo(conn.tcpConn)
	return len(b), err
}

func (conn *CompressedConn) Close() error {
	return conn.tcpConn.Close()
}

func (conn *CompressedConn) LocalAddr() net.Addr {
	return conn.tcpConn.LocalAddr()
}

func (conn *CompressedConn) RemoteAddr() net.Addr {
	return conn.tcpConn.RemoteAddr()
}

func (conn *CompressedConn) SetDeadline(t time.Time) error {
	return conn.tcpConn.SetDeadline(t)
}

func (conn *CompressedConn) SetReadDeadline(t time.Time) error {
	return conn.tcpConn.SetReadDeadline(t)
}

func (conn *CompressedConn) SetWriteDeadline(t time.Time) error {
	return conn.tcpConn.SetWriteDeadline(t)
}
