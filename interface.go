package watchman

import (
	"io"
	"time"
)

type Logger interface {
	io.Closer
	Log(data ...any)
}

type Output interface {
	io.Closer
	Put(date time.Time, instance string, data ...any) error
}
