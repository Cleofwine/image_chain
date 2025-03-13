package log

import (
	"io"
	"sync"

	"gopkg.in/natefinch/lumberjack.v2"
)

type fileRotateWriter struct {
	data map[string]io.Writer
	sync.RWMutex
}

func (frw *fileRotateWriter) getWriter(logPath string) io.Writer {
	frw.RLock()
	defer frw.RUnlock()
	w, ok := frw.data[logPath]
	if !ok {
		return nil
	}
	return w
}
func (frw *fileRotateWriter) setWriter(logPath string, w io.Writer) io.Writer {
	frw.Lock()
	defer frw.Unlock()
	frw.data[logPath] = w
	return w
}

var _fileRotateWriter *fileRotateWriter

func init() {
	_fileRotateWriter = &fileRotateWriter{
		data: map[string]io.Writer{},
	}
}

func GetRotateWriter(logPath string) io.Writer {
	if logPath == "" {
		panic("日志文件路径不能为空")
	}
	writer := _fileRotateWriter.getWriter(logPath)
	if writer != nil {
		return writer
	}
	writer = &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    1,
		MaxBackups: 15,
		MaxAge:     7,
		LocalTime:  true,
	}
	return _fileRotateWriter.setWriter(logPath, writer)
}
