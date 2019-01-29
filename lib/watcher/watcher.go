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
	isRunning  bool
}

func (w *Watcher) Must(e error, logstring string) {
	w.logMan.Must(e, logstring)
}

func (w *Watcher) Start(action func(name string, status FileStatus)) {

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

	for w.isRunning == true {
		time.Sleep(w.delay)
		err := filepath.Walk(w.searchPath, pathIterator)
		w.Must(err, "Unable to walk through the root directory")
	}
}

func (w *Watcher) Stop() {
	w.isRunning = false
}

func NewWatcher(c conf.Configuration, waitDelay time.Duration) *Watcher {
	w := Watcher{
		searchPath: c.RootDirectory,
		delay:      waitDelay,
		paths:      make(map[string]time.Time),
		logMan:     c.GetLogger(),
		isRunning:  true,
	}
	writeTimings := func(p string, info os.FileInfo, err error) error {
		w.paths[p] = info.ModTime()
		return nil
	}
	err := filepath.Walk(w.searchPath, writeTimings)
	w.Must(err, "Unable to create the Watcher")
	return &w
}
