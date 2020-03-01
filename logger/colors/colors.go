package colors

import (
	"io"
)

// Mode_ColorEscSeq 支持划分的颜色转义序列。但是不输出非彩色转义序列。
// 如果要输出非彩色，请使用 Mode_NonColorEscSeq 转义序列，例如ncurses。 但是，它不支持分割颜色逃逸序列。
const (
	_                   int = iota
	Mode_ColorEscSeq    int = iota
	Mode_NonColorEscSeq int = iota
)

// NewColorWriter使用io.Writer w作为其初始内容创建并初始化一个新的winColorWriter。
// 在Windows控制台中，它通过转义序列更改文本的前景色和背景色。在其他系统的控制台中，它写入所有文字。
func NewColorWriter(w io.Writer) io.Writer {
	if _, ok := w.(*colorWriter); !ok {
		return &colorWriter{
			w:    w,
			mode: Mode_ColorEscSeq,
		}
	}
	return w
}

// Bold 加粗
func Bold(msg string) string {
	return "\x1b[1m" + msg + "\x1b[21m"
}

// Black 黑色
func Black(msg string) string {
	return "\x1b[30m" + msg + "\x1b[0m"
}

// Red 红色
func Red(msg string) string {
	return "\x1b[31m" + msg + "\x1b[0m"
}

// Green 绿色
func Green(msg string) string {
	return "\x1b[32m" + msg + "\x1b[0m"
}

// Yellow 黄色
func Yellow(msg string) string {
	return "\x1b[33m" + msg + "\x1b[0m"
}

// Blue 蓝色
func Blue(msg string) string {
	return "\x1b[34m" + msg + "\x1b[0m"
}

// Magenta 洋红色
func Magenta(msg string) string {
	return "\x1b[35m" + msg + "\x1b[0m"
}

// Cyan 青色 returns a cyan string
func Cyan(msg string) string {
	return "\x1b[36m" + msg + "\x1b[0m"
}

// White 白色 returns a white string
func White(msg string) string {
	return "\x1b[37m" + msg + "\x1b[0m"
}

// BlackBold 黑色加粗
func BlackBold(msg string) string {
	return "\x1b[30m\x1b[1m" + msg + "\x1b[21m\x1b[0m"
}

// RedBold 红色加粗
func RedBold(msg string) string {
	return "\x1b[31m\x1b[1m" + msg + "\x1b[21m\x1b[0m"
}

// GreenBold 绿色加粗
func GreenBold(msg string) string {
	return "\x1b[32m\x1b[1m" + msg + "\x1b[21m\x1b[0m"
}

// YellowBold 黄色加粗
func YellowBold(msg string) string {
	return "\x1b[33m\x1b[1m" + msg + "\x1b[21m\x1b[0m"
}

// BlueBold 蓝色加粗
func BlueBold(msg string) string {
	return "\x1b[34m\x1b[1m" + msg + "\x1b[21m\x1b[0m"
}

// MagentaBold 洋红色加粗
func MagentaBold(msg string) string {
	return "\x1b[35m\x1b[1m" + msg + "\x1b[21m\x1b[0m"
}

// CyanBold 青色加粗
func CyanBold(msg string) string {
	return "\x1b[36m\x1b[1m" + msg + "\x1b[21m\x1b[0m"
}

// WhiteBold 白色加粗
func WhiteBold(msg string) string {
	return "\x1b[37m\x1b[1m" + msg + "\x1b[21m\x1b[0m"
}

// \x1b[0m	All attributes off(color at startup)
// \x1b[1m	Bold on(enable foreground intensity)
// \x1b[4m	Underline on
// \x1b[5m	Blink on(enable background intensity)
// \x1b[21m	Bold off(disable foreground intensity)
// \x1b[24m	Underline off
// \x1b[25m	Blink off(disable background intensity)
// Escape sequence	Foreground colors
// \x1b[30m	Black
// \x1b[31m	Red
// \x1b[32m	Green
// \x1b[33m	Yellow
// \x1b[34m	Blue
// \x1b[35m	Magenta
// \x1b[36m	Cyan
// \x1b[37m	White
// \x1b[39m	Default(foreground color at startup)
// \x1b[90m	Light Gray
// \x1b[91m	Light Red
// \x1b[92m	Light Green
// \x1b[93m	Light Yellow
// \x1b[94m	Light Blue
// \x1b[95m	Light Magenta
// \x1b[96m	Light Cyan
// \x1b[97m	Light White
// Escape sequence	Background colors
// \x1b[40m	Black
// \x1b[41m	Red
// \x1b[42m	Green
// \x1b[43m	Yellow
// \x1b[44m	Blue
// \x1b[45m	Magenta
// \x1b[46m	Cyan
// \x1b[47m	White
// \x1b[49m	Default(background color at startup)
// \x1b[100m	Light Gray
// \x1b[101m	Light Red
// \x1b[102m	Light Green
// \x1b[103m	Light Yellow
// \x1b[104m	Light Blue
// \x1b[105m	Light Magenta
// \x1b[106m	Light Cyan
// \x1b[107m	Light White
