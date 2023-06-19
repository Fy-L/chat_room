package bytes

// IO Writer
type Writer struct {
	n   int
	buf []byte
}

//创建指定size的writer
func NewWriterSize(n int) *Writer {
	return &Writer{buf: make([]byte, n)}
}

//buff 的长度
func (w *Writer) Len() int {
	return w.n
}

//buff 的容量
func (w *Writer) Size() int {
	return len(w.buf)
}

//重置buff
func (w *Writer) Reset() {
	w.n = 0
}

//获取buff
func (w *Writer) Buffer() []byte {
	return w.buf[:w.n]
}

//窥视buff,并改变游标位置
func (w *Writer) Peek(n int) []byte {
	var buf []byte
	w.grow(n)
	buf = w.buf[w.n : w.n+n]
	w.n += n
	return buf
}

//将p 写到缓冲区
func (w *Writer) Write(p []byte) {
	w.grow(len(p))
	//讲p复制到w.buf,并更新w.n
	w.n += copy(w.buf[w.n:], p)
}

//扩容
func (w *Writer) grow(n int) {
	var buf []byte
	//如果 +n 之后还没有w.buf长，则无需扩容
	if w.n+n < len(w.buf) {
		return
	}
	//定义新的buf，2倍w.buf 长度 + n 为size
	buf = make([]byte, 2*len(w.buf)+n)
	//复制w.buf到buf
	copy(buf, w.buf[:w.n])
	//更新buf
	w.buf = buf
}
