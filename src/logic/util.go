package logic

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"time"
)

const OUTPUT_LOG = false

var TimeZoneSet = "Asia/Shanghai"

func Addlog(v interface{}) {
	var str string
	vtype := fmt.Sprintf("%s", reflect.TypeOf(v))
	switch vtype {
	case "int":
		str = fmt.Sprintf("%d", v)
	default:
		str = fmt.Sprintf("%s", v)
	}

	tm := time.Now()
	timezone, _ := time.LoadLocation(TimeZoneSet)
	tm = tm.In(timezone)
	s := fmt.Sprintf("[%s]%s\n", tm.Format("2006-01-02 15:04:05"), str)
	if OUTPUT_LOG {
		fmt.Print(s)
	} else {
		saveToFile(s)
	}
}

func AddTimerlog(timer *timer, action string, err error) {
	errstr := ""
	if err != nil {
		errstr = err.Error()
	}

	Addlog(fmt.Sprintf("[%d][%s]%s:%s", timer.Id, action, timer.url, errstr))
}

var logfile *os.File

func saveToFile(logtext string) {
	var err error
	if logfile == nil {
		filename := "./run.log"
		logfile, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
			return
		}
	}

	var n int
	n, err = io.WriteString(logfile, logtext)
	if OUTPUT_LOG {
		if err == nil {
			fmt.Printf("写入 %d 个字节\n", n)
		} else {
			fmt.Printf("error: %s\n", err.Error())
		}
	}
}
