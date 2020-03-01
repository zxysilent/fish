// +build windows

package colors

// 参考 https://github.com/shiena/ansicolor
// 颜色属性由两个十六进制数字指定
//  - 第一个对应于背景，第二个对应于前景。
// 	- 当只传入一个值时，则认为是前景色
// https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/cmd

import (
	"bytes"
	"io"
	"strings"
	"syscall"
	"unsafe"
)

const (
	Csi_Outside     int = iota
	Csi_First       int = iota
	Csi_Second      int = iota
	Draw_Foreg      int = iota
	Draw_Backg      int = iota
	Op_NoConsole    int = iota
	Op_ChangedColor int = iota
	Op_Unknown      int = iota
)

type colorWriter struct {
	w             io.Writer
	mode          int
	state         int
	paramStartBuf bytes.Buffer
	paramBuf      bytes.Buffer
}

const (
	csiFirstChar      byte   = '\x1b'
	csiSecondeChar    byte   = '['
	separatorChar     byte   = ';'
	sgrCode           byte   = 'm'
	fgBlue            uint16 = 0x0001 // 蓝色
	fgGreen           uint16 = 0x0002 // 绿色
	fgRed             uint16 = 0x0004 // 红色
	fgIntensity       uint16 = 0x0008 // 8 前景强度
	bgBlue            uint16 = 0x0010 // 蓝色
	bgGreen           uint16 = 0x0020 // 绿色
	bgRed             uint16 = 0x0040 // 红色
	bgIntensity       uint16 = 0x0080 // 128 背景强度
	underscore        uint16 = 0x8000 // 32768 下划线
	fgMask            uint16 = fgBlue | fgGreen | fgRed | fgIntensity
	bgMask            uint16 = bgBlue | bgGreen | bgRed | bgIntensity
	winReset          string = "0"
	winIntensityOn    string = "1"
	winIntensityOff   string = "21"
	winUnderlineOn    string = "4"
	winUnderlineOff   string = "24"
	winBlinkOn        string = "5"
	winBlinkOff       string = "25"
	winFgBlack        string = "30"
	winFgRed          string = "31"
	winFgGreen        string = "32"
	winFgYellow       string = "33"
	winFgBlue         string = "34"
	winFgMagenta      string = "35"
	winFgCyan         string = "36"
	winFgWhite        string = "37"
	winFgDefault      string = "39"
	winBgBlack        string = "40"
	winBgRed          string = "41"
	winBgGreen        string = "42"
	winBgYellow       string = "43"
	winBgBlue         string = "44"
	winBgMagenta      string = "45"
	winBgCyan         string = "46"
	winBgWhite        string = "47"
	winBgDefault      string = "49"
	winFgLightGray    string = "90"
	winFgLightRed     string = "91"
	winFgLightGreen   string = "92"
	winFgLightYellow  string = "93"
	winFgLightBlue    string = "94"
	winFgLightMagenta string = "95"
	winFgLightCyan    string = "96"
	winFgLightWhite   string = "97"
	winBgLightGray    string = "100"
	winBgLightRed     string = "101"
	winBgLightGreen   string = "102"
	winBgLightYellow  string = "103"
	winBgLightBlue    string = "104"
	winBgLightMagenta string = "105"
	winBgLightCyan    string = "106"
	winBgLightWhite   string = "107"
)

// 坐标
type coord struct {
	X, Y int16
}

type winColor struct {
	code     uint16
	drawType int
}

var colorMap = map[string]winColor{
	winFgBlack:   {0, Draw_Foreg},
	winFgRed:     {fgRed, Draw_Foreg},
	winFgGreen:   {fgGreen, Draw_Foreg},
	winFgYellow:  {fgRed | fgGreen, Draw_Foreg},
	winFgBlue:    {fgBlue, Draw_Foreg},
	winFgMagenta: {fgRed | fgBlue, Draw_Foreg},
	winFgCyan:    {fgGreen | fgBlue, Draw_Foreg},
	winFgWhite:   {fgRed | fgGreen | fgBlue, Draw_Foreg},
	winFgDefault: {fgRed | fgGreen | fgBlue, Draw_Foreg},

	winBgBlack:   {0, Draw_Backg},
	winBgRed:     {bgRed, Draw_Backg},
	winBgGreen:   {bgGreen, Draw_Backg},
	winBgYellow:  {bgRed | bgGreen, Draw_Backg},
	winBgBlue:    {bgBlue, Draw_Backg},
	winBgMagenta: {bgRed | bgBlue, Draw_Backg},
	winBgCyan:    {bgGreen | bgBlue, Draw_Backg},
	winBgWhite:   {bgRed | bgGreen | bgBlue, Draw_Backg},
	winBgDefault: {0, Draw_Backg},

	winFgLightGray:    {fgIntensity, Draw_Foreg},
	winFgLightRed:     {fgIntensity | fgRed, Draw_Foreg},
	winFgLightGreen:   {fgIntensity | fgGreen, Draw_Foreg},
	winFgLightYellow:  {fgIntensity | fgRed | fgGreen, Draw_Foreg},
	winFgLightBlue:    {fgIntensity | fgBlue, Draw_Foreg},
	winFgLightMagenta: {fgIntensity | fgRed | fgBlue, Draw_Foreg},
	winFgLightCyan:    {fgIntensity | fgGreen | fgBlue, Draw_Foreg},
	winFgLightWhite:   {fgIntensity | fgRed | fgGreen | fgBlue, Draw_Foreg},

	winBgLightGray:    {bgIntensity, Draw_Backg},
	winBgLightRed:     {bgIntensity | bgRed, Draw_Backg},
	winBgLightGreen:   {bgIntensity | bgGreen, Draw_Backg},
	winBgLightYellow:  {bgIntensity | bgRed | bgGreen, Draw_Backg},
	winBgLightBlue:    {bgIntensity | bgBlue, Draw_Backg},
	winBgLightMagenta: {bgIntensity | bgRed | bgBlue, Draw_Backg},
	winBgLightCyan:    {bgIntensity | bgGreen | bgBlue, Draw_Backg},
	winBgLightWhite:   {bgIntensity | bgRed | bgGreen | bgBlue, Draw_Backg},
}

var (
	kernel32                 = syscall.NewLazyDLL("kernel32.dll")
	procSetConsoleAttribute  = kernel32.NewProc("SetConsoleTextAttribute")
	procGetConsoleBufferInfo = kernel32.NewProc("GetConsoleScreenBufferInfo")
	defaultAttr              *textAttributes
)

func init() {
	screenInfo := getConsoleBufferInfo(uintptr(syscall.Stdout))
	if screenInfo != nil {
		colorMap[winFgDefault] = winColor{
			screenInfo.WAttributes & (fgRed | fgGreen | fgBlue),
			Draw_Foreg,
		}
		colorMap[winBgDefault] = winColor{
			screenInfo.WAttributes & (bgRed | bgGreen | bgBlue),
			Draw_Backg,
		}
		defaultAttr = convertTextAttr(screenInfo.WAttributes)
	}
}

type smallRect struct {
	Left, Top, Right, Bottom int16
}

type consoleBufferInfo struct {
	DwSize              coord
	DwCursorPosition    coord
	WAttributes         uint16
	SrWindow            smallRect
	DwMaximumWindowSize coord
}

func getConsoleBufferInfo(hConsoleOutput uintptr) *consoleBufferInfo {
	var csbi consoleBufferInfo
	ret, _, _ := procGetConsoleBufferInfo.Call(hConsoleOutput, uintptr(unsafe.Pointer(&csbi)))
	if ret == 0 {
		return nil
	}
	return &csbi
}

func setConsoleAttribute(hConsoleOutput uintptr, wAttributes uint16) bool {
	ret, _, _ := procSetConsoleAttribute.Call(hConsoleOutput, uintptr(wAttributes))
	return ret != 0
}

type textAttributes struct {
	fgColor         uint16
	bgColor         uint16
	fgIntensity     uint16
	bgIntensity     uint16
	underscore      uint16
	otherAttributes uint16
}

func convertTextAttr(winAttr uint16) *textAttributes {
	fgColor := winAttr & (fgRed | fgGreen | fgBlue)
	bgColor := winAttr & (bgRed | bgGreen | bgBlue)
	fgIntensity := winAttr & fgIntensity
	bgIntensity := winAttr & bgIntensity
	underline := winAttr & underscore
	otherAttributes := winAttr &^ (fgMask | bgMask | underscore)
	return &textAttributes{fgColor, bgColor, fgIntensity, bgIntensity, underline, otherAttributes}
}

func convertWinAttr(textAttr *textAttributes) uint16 {
	var winAttr uint16
	winAttr |= textAttr.fgColor
	winAttr |= textAttr.bgColor
	winAttr |= textAttr.fgIntensity
	winAttr |= textAttr.bgIntensity
	winAttr |= textAttr.underscore
	winAttr |= textAttr.otherAttributes
	return winAttr
}

func changeColor(param []byte) int {
	screenInfo := getConsoleBufferInfo(uintptr(syscall.Stdout))
	if screenInfo == nil {
		return Op_NoConsole
	}
	winAttr := convertTextAttr(screenInfo.WAttributes)
	strParam := string(param)
	if len(strParam) <= 0 {
		strParam = "0"
	}
	csiParam := strings.Split(strParam, string(separatorChar))
	for _, p := range csiParam {
		c, ok := colorMap[p]
		if !ok {
			switch p {
			case winReset:
				winAttr.fgColor = defaultAttr.fgColor
				winAttr.bgColor = defaultAttr.bgColor
				winAttr.fgIntensity = defaultAttr.fgIntensity
				winAttr.bgIntensity = defaultAttr.bgIntensity
				winAttr.underscore = 0
				winAttr.otherAttributes = 0
			case winIntensityOn:
				winAttr.fgIntensity = fgIntensity
			case winIntensityOff:
				winAttr.fgIntensity = 0
			case winUnderlineOn:
				winAttr.underscore = underscore
			case winUnderlineOff:
				winAttr.underscore = 0
			case winBlinkOn:
				winAttr.bgIntensity = bgIntensity
			case winBlinkOff:
				winAttr.bgIntensity = 0
			default:
				// Op_Unknown code
			}
		} else {
			if c.drawType == Draw_Foreg {
				winAttr.fgColor = c.code
			} else if c.drawType == Draw_Backg {
				winAttr.bgColor = c.code
			}
		}
	}
	winTextAttribute := convertWinAttr(winAttr)
	setConsoleAttribute(uintptr(syscall.Stdout), winTextAttribute)
	return Op_ChangedColor
}

func parseEscapeSequence(command byte, param []byte) int {
	if defaultAttr == nil {
		return Op_NoConsole
	}
	if command == sgrCode {
		return changeColor(param)
	}
	return Op_Unknown
}

func (cw *colorWriter) flushBuffer() (int, error) {
	return cw.flushTo(cw.w)
}

func (cw *colorWriter) resetBuffer() (int, error) {
	return cw.flushTo(nil)
}

func (cw *colorWriter) flushTo(w io.Writer) (int, error) {
	var l1, l2 int
	var err error
	startBytes := cw.paramStartBuf.Bytes()
	cw.paramStartBuf.Reset()
	if w != nil {
		l1, err = cw.w.Write(startBytes)
		if err != nil {
			return l1, err
		}
	} else {
		l1 = len(startBytes)
	}
	paramBytes := cw.paramBuf.Bytes()
	cw.paramBuf.Reset()
	if w != nil {
		l2, err = cw.w.Write(paramBytes)
		if err != nil {
			return l1 + l2, err
		}
	} else {
		l2 = len(paramBytes)
	}
	return l1 + l2, nil
}

func (cw *colorWriter) Write(p []byte) (int, error) {
	var r, nw, first, last int
	if cw.mode != Mode_ColorEscSeq {
		cw.state = Csi_Outside
		cw.resetBuffer()
	}
	var err error
	for i, ch := range p {
		switch cw.state {
		case Csi_Outside:
			if ch == csiFirstChar {
				cw.paramStartBuf.WriteByte(ch)
				cw.state = Csi_First
			}
		case Csi_First:
			switch ch {
			case csiFirstChar:
				cw.paramStartBuf.WriteByte(ch)
				break
			case csiSecondeChar:
				cw.paramStartBuf.WriteByte(ch)
				cw.state = Csi_Second
				last = i - 1
			default:
				cw.resetBuffer()
				cw.state = Csi_Outside
			}
		case Csi_Second:
			if ('0' <= ch && ch <= '9') || ch == separatorChar {
				cw.paramBuf.WriteByte(ch)
			} else {
				nw, err = cw.w.Write(p[first:last])
				r += nw
				if err != nil {
					return r, err
				}
				first = i + 1
				result := parseEscapeSequence(ch, cw.paramBuf.Bytes())
				if result == Op_NoConsole || (cw.mode == Mode_NonColorEscSeq && result == Op_Unknown) {
					cw.paramBuf.WriteByte(ch)
					nw, err := cw.flushBuffer()
					if err != nil {
						return r, err
					}
					r += nw
				} else {
					n, _ := cw.resetBuffer()
					// Add one more to the size of the buffer for the last ch
					r += n + 1
				}
				cw.state = Csi_Outside
			}
		default:
			cw.state = Csi_Outside
		}
	}
	if cw.mode != Mode_ColorEscSeq || cw.state == Csi_Outside {
		nw, err = cw.w.Write(p[first:])
		r += nw
	}
	return r, err
}
