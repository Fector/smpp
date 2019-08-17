package io

import (
	"bufio"
	"net"
)

type Writer struct {
	*bufio.Writer
}

func NewWriter(conn *net.TCPConn) *Writer {
	return &Writer{
		Writer: bufio.NewWriter(conn),
	}
}

func (w Writer) writeInt(val uint32) error {
	return nil
}

func (w Writer) writeString(val string) error {
	return nil
}

func (w Writer) writeOctString(val string) error {
	return nil
}
