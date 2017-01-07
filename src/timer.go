package main

import (
	"fmt"
	// "reflect"
	"./logic"
	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

/*
模块：
配置的载入
主循环(每秒一次)
执行请求
记录和更新结果
*/

func main() {
	iniFile := filepath.Dir(os.Args[0]) + "/timer.ini"
	if len(os.Args) == 1 {
		// 启动守护进程
		_, err := os.Stat(iniFile)
		if err != nil && os.IsNotExist(err) {
			fmt.Println("Not Exist Config File:", iniFile)
			return
		}
		
		cmd := exec.Command(os.Args[0], "daemon")
		cmd.Stdin = nil
		cmd.Stdout = nil
		cmd.Stderr = nil
		cmd.Start()
		fmt.Println("Start Timer Service...ok!")
		return
	}

	cfg, err := ini.Load(iniFile)
	if err != nil {
		logic.Addlog("load error:"+iniFile)
	}

	logic.TimeZoneSet = cfg.Section("Timezone").Key("default").String()

	mysqlConfig := cfg.Section("MySQL")
	host := mysqlConfig.Key("host").String()
	port := mysqlConfig.Key("port").String()
	username := mysqlConfig.Key("username").String()
	password := mysqlConfig.Key("password").String()
	database := mysqlConfig.Key("database").String()

	conn := logic.CreateDBConn()
	if conn.Init(host, port, username, password, database) {
		logic.Addlog("Load config ok. Start service...")
	} else {
		logic.Addlog("error Init")
		return
	}

	startRunTimer()
}

// 启动定时器的主循环
func startRunTimer() {
	for {
		timerList := logic.LoadConfig()

		currentTime := time.Now().Unix()
		if timerList != nil {
			for _, timer := range timerList {
				if ((currentTime - timer.BaseTime) % timer.IntervalMinutes) == 0 {
					go logic.Request(timer)
				}
			}
		}

		time.Sleep(time.Second)
	}
}
