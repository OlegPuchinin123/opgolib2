package opgolib2

type ByteBuffer struct {
	M_buf []byte
}

func BB_FromString(s string) *ByteBuffer {
	return &ByteBuffer{M_buf: []byte(s)}
}

func NewByteBuffer(buf []byte) *ByteBuffer {
	return &ByteBuffer{M_buf: buf}
}

func (bb *ByteBuffer) ToString() string {
	return string(bb.M_buf)
}
