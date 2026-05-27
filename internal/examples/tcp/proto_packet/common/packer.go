package common

import (
	"encoding/binary"
	"io"

	"github.com/DarthPestilane/easytcp"
)

// CustomPacker treats packet as:
//
// totalSize(4)|id(4)|data(n)
//
// | segment     | type   | size    | remark                |
// | ----------- | ------ | ------- | --------------------- |
// | `totalSize` | uint32 | 4       | the whole packet size |
// | `id`        | uint32 | 4       |                       |
// | `data`      | []byte | dynamic |                       |
type CustomPacker struct{}

func (p *CustomPacker) Pack(msg *easytcp.Message) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// write totalSize
// write id
// write data

func (p *CustomPacker) Unpack(reader io.Reader) (*easytcp.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// read totalSize
// read id

// read data

func (*CustomPacker) byteOrder() binary.ByteOrder {
	_ = "STUB: not implemented"
	return *new(binary.ByteOrder)
}
