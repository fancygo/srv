package logger

//这边是抄的我们游戏的 = =0

import (
	"fmt"
	"runtime"
)

const (
	LOG_ALL  = 0
	LOG_NONE = 1

	LOG_TRACE  = 9
	LOG_DEBUG  = 10
	LOG_NOTICE = 11
	LOG_INFO   = 12
	LOG_WARN   = 13
	LOG_ERROR  = 14
	LOG_FATAL  = 16

	LOG_MAX = LOG_FATAL + 1
)

var stdTags [LOG_MAX][2]string
var mSystem = runtime.GOOS

const (
	colorBlack = (iota + 30)
	colorRed
	colorGreen
	colorYellow
	colorBlue
	colorMagenta
	colorCyan
	colorWhite
)

func colorSeq(color int) string {
	return fmt.Sprintf("\033[%dm", color)
}

func colorSeqBold(color int) string {
	return fmt.Sprintf("\033[%d;1m", color)
}

/*
func colorInit() {

	stdTags = [LOG_MAX][4]string{}
	stdTags[LOG_DEBUG] = [4]string{"[DEBUG]", "", "", ""}
	stdTags[LOG_NOTICE] = [4]string{"[NOTICE]", colorSeq(colorBlue), colorSeqBold(colorBlue), "\033[0m"}
	stdTags[LOG_INFO] = [4]string{"[INFO]", colorSeq(colorGreen), colorSeqBold(colorGreen), "\033[0m"}
	stdTags[LOG_WARN] = [4]string{"[WARN]", colorSeq(colorYellow), colorSeqBold(colorYellow), "\033[0m"}
	stdTags[LOG_ERROR] = [4]string{"[ERROR]", colorSeq(colorRed), colorSeqBold(colorRed), "\033[0m"}
	stdTags[LOG_TRACE] = [4]string{"[TRACE]", colorSeq(colorCyan), colorSeqBold(colorCyan), "\033[0m"}
	stdTags[LOG_FATAL] = [4]string{"[FATAL]", colorSeq(colorMagenta), colorSeqBold(colorMagenta), "\033[0m"}

}
*/

func colorInit() {

	stdTags = [LOG_MAX][2]string{}
	stdTags[LOG_DEBUG] = [2]string{"[DEBUG]", "\033[0m"}
	stdTags[LOG_NOTICE] = [2]string{"[NOTICE]", "\033[0m"}
	stdTags[LOG_INFO] = [2]string{"[INFO]", "\033[0m"}
	stdTags[LOG_WARN] = [2]string{"[WARN]", "\033[0m"}
	stdTags[LOG_ERROR] = [2]string{"[ERROR]", "\033[0m"}
	stdTags[LOG_TRACE] = [2]string{"[TRACE]", "\033[0m"}
	stdTags[LOG_FATAL] = [2]string{"[FATAL]", "\033[0m"}

}

func colorString(level int) (tag string, cPre string, cSuf string) {
	if level < LOG_TRACE || level >= LOG_MAX {
		return
	}

	v := stdTags[level]
	tag = v[0]
	//cPre = v[1]
	//cSuf = v[3]
	cPre = ""
	cSuf = ""

	return
}
