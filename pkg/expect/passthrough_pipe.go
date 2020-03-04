package expect

import (
	"context"
	"errors"
	"io"
	"time"
)

type errPassthroughTimeout struct {
	error
}

func (errPassthroughTimeout) Timeout() bool { return true }

// bufsize is the size of the PassthroughPipe channel
const bufsize = 1024

// PassthroughPipe pipes data from a io.Reader and allows setting a read
// deadline. If a timeout is reached the error is returned, otherwise the error
// from the provided io.Reader returned is passed through instead.
type PassthroughPipe struct {
	rdr      io.Reader
	deadline time.Time
	ctx      context.Context
	cancel   context.CancelFunc
}

// NewPassthroughPipe returns a new pipe for a io.Reader that passes through
// non-timeout errors.
func NewPassthroughPipe(r io.Reader) *PassthroughPipe {
	ctx, cancel := context.WithCancel(context.Background())

	p := PassthroughPipe{
		rdr:      r,
		deadline: time.Now(),
		ctx:      ctx,
		cancel:   cancel,
	}

	return &p
}

// SetReadDeadline sets a deadline for a successful read
func (p *PassthroughPipe) SetReadDeadline(d time.Time) {
	p.deadline = d
}

// Close releases all resources allocated by the pipe
func (p *PassthroughPipe) Close() error {
	p.Drain()
	p.cancel()
	return nil
}

// Drain flushes the pipe by consuming all the data written to it
func (p *PassthroughPipe) Drain() {
	buf := make([]byte, 1<<5)
	for {
		n, err := p.rdr.Read(buf)
		if n == 0 || err != nil {
			return
		}
	}
}

type chunk struct {
	size int
	err  error
}

// Read reads from the PassthroughPipe and errors out if no data has been written to the pipe before the read deadline expired
func (p *PassthroughPipe) Read(buf []byte) (n int, err error) {
	cs := make(chan chunk)
	done := make(chan struct{})
	defer close(done)

	go func() {
		defer close(cs)

		select {
		case <-done:
			return
		default:
		}

		n, err := p.rdr.Read(buf)

		select {
		case <-done:
			return
		default:
			cs <- chunk{n, err}
		}
	}()

	select {
	case c := <-cs:
		return c.size, c.err

	case <-p.ctx.Done():
		return 0, errors.New("passthrough closed")

	case <-time.After(p.deadline.Sub(time.Now())):
		return 0, &errPassthroughTimeout{errors.New("passthrough i/o timeout")}
	}
}
