package logger

import (
	"os"
	"time"
)

type Logger struct {
	FileName string
}

func (l *Logger) Write(p []byte) (n int, err error) {
	file, err := os.OpenFile(l.FileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return 0, err
	}
	n, err = file.Write(p)
	file.Close()
	return n, err
}

func (l *Logger) Must(e error, logstring string) {
	if e != nil {
		l.Write([]byte(e.Error()))
		panic(e)
	}

	l.Write([]byte(formatLog(logstring) + "\n"))
}

func formatLog(logstring string) string {
	tm := time.Now()
	logtime := tm.Format("2/Jan/2006:15:04:05 -0700")
	return "[" + logtime + "] " + logstring
}

func NewLogger(fname string) *Logger {
	temp := Logger{
		FileName: fname,
	}
	return &temp
}
