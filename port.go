// Package port provides a simple API to write Erlang ports in Go.
//
// For more information on Erlang ports see http://www.erlang.org/doc/tutorial/c_port.html.
package port

import (
	"errors"
	"io"
)

// Port is an abstraction over different types of ports.
//
// While using Read() on a line (or packet) based port it may happen that the size
// of the line (packet) is bigger than len(p). In this case port skips the line (packet)
// and return n=0 and err=ErrTooBig.
// It is up to the caller to choose whether to consider this as a fatal error
// or to continue reading.
//
// Line and stream based ports work as simple writers, writing everything as is and
// Write() will return the number of bytes written.
// Each Write() call in packet based ports, however, writes one single packet and returns
// the number of bytes written NOT including packet length (1, 2 or 4 bytes).
type Port interface {
	io.ReadWriter
	// ReadOne reads either one packet, one line (ending with '\n') or a byte
	// from a packet, line or stream-based port accordingly.
	ReadOne() (data []byte, err error)
}

var (
	// ErrBadSizeLen is returned by Packet(r, w, sizeLen) function when invalid
	// value of sizeLen is used. Valid values are 1, 2 and 4.
	ErrBadSizeLen = errors.New("port: bad 'packet size' length")
	// ErrSizeOverflow is returned by packet based port's ReadOne() function
	// when the size of the packet overflows int type.
	// It is also returned by packet based port's Write() function if the packet
	// is too big.
	ErrSizeOverflow = errors.New("port: packet size overflows integer type")
	// ErrToBig is an error which Read(p) function may return if the buffer p is too small
	// to receive whole line (or packet).
	ErrTooBig = errors.New("port: packet does not fit the buffer")
)
