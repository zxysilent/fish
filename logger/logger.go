package logger

import (
	"fmt"
	"io"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/zxysilent/fish/logger/colors"
)

type level int

const (
	levelSucc level = iota
	levelInfo
	levelWarn
	levelError
	levelFatal
	levelNull
)

var (
	instNo   uint64                 //序列😊
	instance *fishLogger            //实例🍿
	maxLevel = levelNull            //所有都显示👁‍🗨
	Flog     = NewLogger(os.Stdout) //默认日志👍
)

// fishLogger 🐟
type fishLogger struct {
	locker sync.Mutex
	writer io.Writer
}

// NewLogger 创建logger
func NewLogger(w io.Writer) *fishLogger {
	return &fishLogger{writer: colors.NewColorWriter(w)}
}

func (lv level) String() string {
	switch lv {
	case levelSucc:
		return "SUCC  "
	case levelInfo:
		return "INFO  "
	case levelWarn:
		return "WARN  "
	case levelError:
		return "ERROR "
	case levelFatal:
		return "FATAL "
	default:
		return "NULL  "
	}
}

func colorString(lv level) string {
	switch lv {
	case levelSucc:
		return colors.GreenBold(lv.String())
	case levelInfo:
		return colors.BlueBold(lv.String())
	case levelWarn:
		return colors.YellowBold(lv.String())
	case levelError:
		return colors.RedBold(lv.String())
	case levelFatal:
		return colors.MagentaBold(lv.String())
	default:
		return colors.WhiteBold(lv.String())
	}
}

// write logs🐟
func (flog *fishLogger) write(lv level, format string, args ...interface{}) {
	if lv > maxLevel {
		return
	}
	flog.locker.Lock()
	defer flog.locker.Unlock()
	ags := make([]interface{}, 0, 4)
	ags = append(ags, time.Now().Format("2006/01/02 15:04:05"), colorString(lv), atomic.AddUint64(&instNo, 1))
	ags = append(ags, args...)
	fmt.Fprintf(flog.writer, "%s %s▶ %05d "+format+"\n", ags...)
}

// Succ 成功
func (flog *fishLogger) Succ(msg string) {
	flog.write(levelSucc, msg)
}

// Succf 成功
func (flog *fishLogger) Succf(format string, args ...interface{}) {
	flog.write(levelSucc, format, args...)
}

// Info 信息ℹ
func (flog *fishLogger) Info(msg string) {
	flog.write(levelInfo, msg)
}

// Infof 信息ℹ
func (flog *fishLogger) Infof(format string, args ...interface{}) {
	flog.write(levelInfo, format, args...)
}

// Warn 警告⚠
func (flog *fishLogger) Warn(msg string) {
	flog.write(levelWarn, msg)
}

// Warnf 警告⚠
func (flog *fishLogger) Warnf(format string, args ...interface{}) {
	flog.write(levelWarn, format, args...)
}

// Error 错误❌
func (flog *fishLogger) Error(msg string) {
	flog.write(levelError, msg)
}

// Errorf 错误❌
func (flog *fishLogger) Errorf(format string, args ...interface{}) {
	flog.write(levelError, format, args...)
}

// Fatal 致命错误
func (flog *fishLogger) Fatal(msg string) {
	flog.write(levelFatal, msg)
	os.Exit(1009)
}

// Fatalf 致命错误
func (flog *fishLogger) Fatalf(format string, args ...interface{}) {
	flog.write(levelFatal, format, args...)
	os.Exit(1009)
}
