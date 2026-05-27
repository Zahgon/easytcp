//go:build !jsoniter
// +build !jsoniter

package easytcp

var _ Codec = &JsonCodec{}

// JsonCodec implements the Codec interface.
// JsonCodec encodes and decodes data in json way.
type JsonCodec struct{}

// Encode implements the Codec Encode method.
func (c *JsonCodec) Encode(v interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// Decode implements the Codec Decode method.
		nil
}

func (c *JsonCodec) Decode(data []byte, v interface{}) error { _ = "STUB: not implemented"; return nil }
