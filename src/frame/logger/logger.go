package logger

//自己实现的log处理

import (
	"fmt"
	"log"
	"os"
)

var gLogger *log.Logger

func init() {
	//file, err := os.Create("log/game.log")
	colorInit()
	file, err := os.OpenFile("log/game.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatalln("fail to create game.log")
	}
	gLogger = log.New(file, "", log.LstdFlags|log.Lshortfile)
}

func Debug(s string, v ...interface{}) {
	str := fmt.Sprintf(s, v...)
	//lTag, cPre, cSuf := colorString(LOG_INFO)
	lTag, _, _ := colorString(LOG_DEBUG)
	lstr := lTag + str
	//sstr := cPre + lstr + cSuf
	gLogger.Output(2, lstr)
}

func Info(s string, v ...interface{}) {
	str := fmt.Sprintf(s, v...)
	fmt.Println(str)
	//lTag, cPre, cSuf := colorString(LOG_INFO)
	lTag, _, _ := colorString(LOG_INFO)
	lstr := lTag + str
	//sstr := cPre + lstr + cSuf
	gLogger.Output(2, lstr)
}

func Notice(v ...interface{}) {
	str := fmt.Sprintln(v)
	lTag, _, _ := colorString(LOG_NOTICE)
	lstr := lTag + str
	gLogger.Output(2, lstr)
}

func Error(v ...interface{}) {
	str := fmt.Sprintln(v)
	//lTag, cPre, cSuf := colorString(LOG_INFO)
	lTag, _, _ := colorString(LOG_ERROR)
	lstr := lTag + str
	//sstr := cPre + lstr + cSuf
	gLogger.Output(2, lstr)
}
