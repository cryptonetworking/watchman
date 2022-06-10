package watchman

import (
	"context"
	"fmt"
	"github.com/itsabgr/go-handy"
	"os"
	"runtime"
	"time"
)

func New(ctx context.Context, name string, output Output, cap int) Logger {
	aCtx, cancel := context.WithCancel(ctx)
	if name == "" {
		name = fmt.Sprintf("%s[%d]", os.Args[0], os.Getpid())
	}
	w := &watchman{
		instance: name,
		queue:    make(chan log, cap),
		output:   output,
		cancel:   cancel,
		ctx:      aCtx,
	}
	go w.start()
	return w
}

func (w *watchman) Close() error {
	w.cancel()
	return w.output.Close()
}

func (w *watchman) Log(data ...any) {
	w.queue <- log{
		time: time.Now(),
		data: data,
	}
}

func (w *watchman) start() error {
	runtime.LockOSThread()
	defer handy.Just(w.Close)
	for {
		select {
		case <-w.ctx.Done():
			return w.ctx.Err()
		case log := <-w.queue:
			err := w.output.Put(log.time, w.instance, log.data...)
			if err != nil {
				return err
			}
		}
	}
}

type watchman struct {
	queue    chan log
	instance string
	output   Output
	ctx      context.Context
	cancel   context.CancelFunc
}

type log struct {
	time time.Time
	data []any
}
