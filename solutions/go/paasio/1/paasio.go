package paasio

import (
	"io"
	"sync"
)

// Define readCounter and writeCounter types here.
type readCounter struct {
	mu    sync.Mutex
	r     io.Reader
	bytes int64
	ops   int
}

type writeCounter struct {
	mu    sync.Mutex
	w     io.Writer
	bytes int64
	ops   int
}

// For the return of the function NewReadWriteCounter, you must also define a type that satisfies the ReadWriteCounter interface.
// Embedding both counters promotes their methods (Read/ReadCount from readCounter,
// Write/WriteCount from writeCounter), so readWriteCounter satisfies ReadWriteCounter.
type readWriteCounter struct {
	*readCounter
	*writeCounter
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{
		w: writer,
	}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{
		r: reader,
	}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
		readCounter:  &readCounter{r: readwriter},
		writeCounter: &writeCounter{w: readwriter},
	}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	n, err := rc.r.Read(p)
	rc.mu.Lock()
	rc.bytes += int64(n)
	rc.ops++
	rc.mu.Unlock()
	return n, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	return rc.bytes, rc.ops
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	n, err := wc.w.Write(p)
	wc.mu.Lock()
	wc.bytes += int64(n)
	wc.ops++
	wc.mu.Unlock()
	return n, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	return wc.bytes, wc.ops
}
