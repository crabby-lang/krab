package terminal

import (
	"strings"
	"sync"
)

// Buffer manages terminal output with a **scrollback**
type Buffer struct {
	lines			[]string
	maxLines		int
	mu				sync.RWMutex
}

// NewBuffer creates a new terminal buffer
func NewBuffer(maxLines int) *Buffer {
	return &Buffer {
		lines:		make([]string, 0, maxLines),
		maxLines:	maxLines,
	}
}

// AppendLine adds a new line to the buffer
func (b *Buffer) AppendLine(line string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.lines = append(b.lines, line)

	if len(b.lines) > b.maxLines {
		b.lines = b.lines[len(b.lines)-b.maxLines:]
	}
}

func (b *Buffer) Append(text string) {
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		if line != "" || len(lines) > 1 {
			b.AppendLine(line)
		}
	}
}

func (b *Buffer) Lines() []string {
	b.mu.RLock()
	defer b.mu.RUnlock()

	result := make([]string, len(b.lines))
	copy(result, b.lines)
	return result
}

func (b *Buffer) String() string {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return strings.Join(b.lines, "\n")
}

func (b *Buffer) LineCount() int {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return len(b.lines)
}

func (b *Buffer) Clear() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.lines = make([]string, 0, b.maxLines)
}
