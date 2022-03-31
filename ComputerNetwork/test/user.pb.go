
type String struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *String) Reset()         { *m = String{} }
func (m *String) String() string { return proto.CompactTextString(m) }
func (*String) ProtoMessage()    {}
func (*String) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_069698f99dd8f029, []int{0}
}

func (m *String) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}