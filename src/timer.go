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
	if len(os.Args) == 1 {
		// 启动守护进程
		configFile, _ := filepath.Abs("./timer.ini")
		_, err := os.Stat(configFile)
		if err != nil && os.IsNotExist(err) {
			fmt.Println("Not Exist Config File:", configFile)
			return
		}

		cmd := exec.Command(os.Args[0], configFile)
		cmd.Stdin = nil
		cmd.Stdout = nil
		cmd.Stderr = nil
		cmd.Start()
		fmt.Println("Start Timer Service...ok!")
		return
	}

	cfg, err := ini.Load(os.Args[1])
	if err != nil {
		logic.Addlog("timer.ini load error")
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
