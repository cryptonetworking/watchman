package watchman

import (
	"fmt"
	"github.com/itsabgr/go-handy"
	"os"
	"time"
)

type fileOutput os.File

func (f *fileOutput) Close() error {
	if f == nil {
		return nil
	}
	(*os.File)(f).Sync()
	return (*os.File)(f).Close()
}

func (f *fileOutput) Put(t time.Time, s string, a ...any) error {
	_, err := fmt.Fprintln((*os.File)(f), t.UnixNano(), s, fmt.Sprint(a...))
	return err
}

func File(path string) Output {
	if path == "" {
		path = fmt.Sprintf("%s.%d.log", os.Args[0], time.Now().UnixNano())
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	handy.Throw(err)
	return (*fileOutput)(file)
}
