package terminal

import (
	"fmt"
	"io"
	"strings"
	"sync"
)

// Renderer controls the rendering of the buffer
type Renderer struct {
	Buf  *Buffer
	prev *Buffer
	mu   sync.Mutex
}

// NewRendererFrom creates a new renderer from giver buffer size
func NewRendererFrom(w, h int) *Renderer {
	return &Renderer{
		Buf: NewBuffer(w, h),
	}
}

// Resize resizes the buffer
func (r *Renderer) Resize(w, h int) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.Buf = NewBuffer(w, h)
	r.prev = nil
}

// FlushFull flushes the buffer to the given writer
func (r *Renderer) FlushFull(out io.Writer) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, err := fmt.Fprint(out, "\x1b[2J"); err != nil {
		return err
	}
	if _, err := fmt.Fprint(out, "\x1b[H"); err != nil {
		return err
	}

	var sb strings.Builder
	for y := 0; y < r.Buf.H; y++ {
		sb.WriteString(fmt.Sprintf("\x1b[%d;1H", y+1))
		sb.WriteString(r.Buf.RowString(y))
	}
	_, err := io.WriteString(out, sb.String())
	if err != nil {
		return err
	}

	r.prev = r.Buf.Copy()
	return nil
}
