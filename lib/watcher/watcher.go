package watcher

import (
	"os"
	"path/filepath"
	"time"

	"github.com/dgplug/dexer/lib/conf"
	"github.com/dgplug/dexer/lib/logger"
)

type FileStatus int

const (
	create FileStatus = iota
	modified
	erased
)

type Watcher struct {
	paths      map[string]time.Time
	searchPath string
	delay      time.Duration
	logMan     *logger.Logger
}

func (w *Watcher) Must(e error, logstring string) {
	w.logMan.Must(e, logstring)
}

func (w *Watcher) start(action func(name string, status FileStatus)) {
	time.Sleep(w.delay)

	for file := range w.paths {
		if _, err := os.Stat(file); err != nil {
			delete(w.paths, file)
			action(file, erased)
		}
	}

	pathIterator := func(p string, info os.FileInfo, err error) error {
		lastWrite := info.ModTime()

		if _, ok := w.paths[p]; !ok {
			w.paths[p] = lastWrite
			action(p, create)
			return nil
		}

		if w.paths[p] != lastWrite {
			w.paths[p] = lastWrite
			action(p, modified)
		}

		return nil
	}

	err := filepath.Walk(w.searchPath, pathIterator)
	w.Must(err, "Unable to walk through the root directory")
}

func NewWatcher(c conf.Configuration, delayInMilliSeconds time.Duration) *Watcher {
	w := Watcher{
		searchPath: c.RootDirectory,
		delay:      delayInMilliSeconds,
		paths:      make(map[string]time.Time),
		logMan:     c.GetLogger(),
	}
	writeTimings := func(p string, info os.FileInfo, err error) error {
		w.paths[p] = info.ModTime()
		return nil
	}
	err := filepath.Walk(w.searchPath, writeTimings)
	w.Must(err, "Unable to create the Watcher")
	return &w
}
