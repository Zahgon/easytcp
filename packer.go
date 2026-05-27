package easytcp

import (
	"encoding/binary"
	"io"
)

//go:generate mockgen -destination ./packer_mock.go -package easytcp . Packer

// Packer is a generic interface to pack and unpack message packet.
type Packer interface {
	// Pack packs Message into the packet to be written.
	Pack(msg *Message) ([]byte, error)

	// Unpack unpacks the message packet from reader,
	// returns the message, and error if error occurred.
	Unpack(reader io.Reader) (*Message, error)
}

var _ Packer = &DefaultPacker{}

// NewDefaultPacker create a *DefaultPacker with initial field value.
func NewDefaultPacker() *DefaultPacker { _ = "STUB: not implemented"; return nil }

// 1MB
// big endian as default

// DefaultPacker is the default Packer used in session.
// Treats the packet with the format:
//
// dataSize(4)|id(4)|data(n)
//
// | segment    | type   | size    | remark                  |
// | ---------- | ------ | ------- | ----------------------- |
// | `dataSize` | uint32 | 4       | the size of `data` only |
// | `id`       | uint32 | 4       |                         |
// | `data`     | []byte | dynamic |                         |
// .
type DefaultPacker struct {
	// MaxDataSize represents the max size of `data`
	MaxDataSize int
	byteOrder   binary.ByteOrder
}

func (d *DefaultPacker) bytesOrder() binary.ByteOrder {
	_ = "STUB: not implemented"
	return *

	// SetByteOrder sets the byte order
	new(binary.ByteOrder)
}

func (d *DefaultPacker) SetByteOrder(order binary.ByteOrder) { _ = "STUB: not implemented"; return }

// Pack implements the Packer Pack method.
func (d *DefaultPacker) Pack(msg *Message) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// write dataSize

// write id
// write data

// Unpack implements the Packer Unpack method.
// Unpack returns the message whose ID is type of int.
// So we need use int id to register routes.
func (d *DefaultPacker) Unpack(reader io.Reader) (*Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
