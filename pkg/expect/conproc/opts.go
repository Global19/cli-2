package conproc

import (
	"io/ioutil"
	"os"
	"sync"
	"time"

	"github.com/ActiveState/cli/pkg/expect"
)

type SendObserver func(msg string, num int, err error)

type ExpectObserver func(matchers []expect.Matcher, raw, pty string, err error)

// Options contain optional values for ConsoleProcess construction and usage.
type Options struct {
	DefaultTimeout time.Duration
	WorkDirectory  string
	RetainWorkDir  bool
	Environment    []string
	ObserveSend    SendObserver
	ObserveExpect  ExpectObserver
	CmdName        string
	Args           []string
}

func (opts *Options) Normalize() error {
	if opts.DefaultTimeout == 0 {
		opts.DefaultTimeout = time.Second * 20
	}

	if opts.WorkDirectory == "" {
		tmpDir, err := ioutil.TempDir("", "")
		if err != nil {
			return err
		}
		opts.WorkDirectory = tmpDir
	}

	if opts.ObserveSend == nil {
		opts.ObserveSend = func(string, int, error) {}
	}

	if opts.ObserveExpect == nil {
		opts.ObserveExpect = func([]expect.Matcher, string, string, error) {}
	}

	return nil
}

func (opts *Options) CleanUp() error {
	if !opts.RetainWorkDir {
		return os.RemoveAll(opts.WorkDirectory)
	}

	return nil
}

type expectObserverTransform struct {
	observeFn ExpectObserver
	mu        sync.Mutex
	rawDataFn func() string
}

func (t *expectObserverTransform) setRawDataFn(fn func() string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.rawDataFn = fn
}

func (t *expectObserverTransform) rawData() string {
	t.mu.Lock()
	defer t.mu.Unlock()

	return t.rawDataFn()
}

func (t *expectObserverTransform) observe(matchers []expect.Matcher, buf string, err error) {
	raw := t.rawData()
	t.observeFn(matchers, buf, raw, err)
}
