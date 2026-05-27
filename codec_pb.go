package easytcp

// ProtobufCodec implements the Codec interface.
type ProtobufCodec struct{}

// Encode implements the Codec Encode method.
func (p *ProtobufCodec) Encode(v interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Decode implements the Codec Decode method.
func (p *ProtobufCodec) Decode(data []byte, v interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
