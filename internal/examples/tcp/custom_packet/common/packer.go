package common

import (
	"encoding/binary"
	"io"

	"github.com/DarthPestilane/easytcp"
)

// CustomPacker treats packet as:
//
// totalSize(4)|idSize(2)|id(n)|data(n)
//
// | segment     | type   | size    | remark                |
// | ----------- | ------ | ------- | --------------------- |
// | `totalSize` | uint32 | 4       | the whole packet size |
// | `idSize`    | uint16 | 2       | length of id          |
// | `id`        | string | dynamic |                       |
// | `data`      | []byte | dynamic |                       |
type CustomPacker struct{}

func (p *CustomPacker) bytesOrder() binary.ByteOrder {
	_ = "STUB: not implemented"
	return *new(binary.ByteOrder)
}

func (p *CustomPacker) Pack(msg *easytcp.Message) ([]byte, error) {
	_ = "STUB: not implemented"
	// format: totalSize(4)|idSize(2)|id(n)|data(n)
	return nil, nil
}

// write totalSize
// write idSize
// write id
// write data

func (p *CustomPacker) Unpack(reader io.Reader) (*easytcp.Message, error) {
	_ = "STUB: not implemented"
	// format: totalSize(4)|idSize(2)|id(n)|data(n)
	return nil, nil
}

// read totalSize
// read idSize

// read id
// read body

// ID is a string, so we should use a string-type id to register routes.
// eg: server.AddRoute("string-id", handler)
