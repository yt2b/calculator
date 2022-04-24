package buffer

type Buffer struct {
	input string
	idx   int
	Ch    byte
}

func NewBuffer(input string) *Buffer {
	b := &Buffer{input: input}
	if len(b.input) > 0 {
		b.Ch = b.input[b.idx]
	}
	return b
}

func (b *Buffer) Read() {
	if b.idx+1 >= len(b.input) {
		// 終端に到達
		b.Ch = 0
		return
	}
	b.idx += 1
	b.Ch = b.input[b.idx]
}
