package easytcp

// MsgpackCodec implements the Codec interface.
type MsgpackCodec struct{}

// Encode implements the Codec Encode method.
func (m *MsgpackCodec) Encode(v interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// Decode implements the Codec Decode method.
		nil
}

func (m *MsgpackCodec) Decode(data []byte, v interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
