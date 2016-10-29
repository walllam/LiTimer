package logic

import (
	"sync"
	"time"
)

type runLog struct {
	startTime int64
	endTime   int64
	result    string
	status    byte
}

type timer struct {
	Id              int64
	BaseTime        int64
	IntervalMinutes int64
	timeout         int
	url             string
	uid             int64
	runStatus       byte
	clearTag        byte
}

const initListLen = 100

var timerList = make(map[int64]*timer, initListLen)
var mutex sync.Mutex

func LoadConfig() map[int64]*timer {
	mutex.Lock()
	defer mutex.Unlock()

	conn := CreateDBConn()
	rows, err := conn.Query("SELECT tp_id,base_time,interval_minute,timeout,url,uid FROM timer_process WHERE status=1")
	if err != nil {
		Addlog(err.Error())
		return nil
	}

	for _, timer := range timerList {
		timer.clearTag = 1
	}
	
	
	// 计算默认的基准时间
	timezone, _ := time.LoadLocation(TimeZoneSet)
	timestamp, _ := time.ParseInLocation("2006-01-02 15:04:05", "2016-01-01 00:00:00", timezone)
	defaultBaseTime := timestamp.Unix()

	for rows.Next() {
		var t timer
		err = rows.Scan(&t.Id, &t.BaseTime, &t.IntervalMinutes, &t.timeout, &t.url, &t.uid)
		// 将分钟转换成秒
		t.IntervalMinutes = t.IntervalMinutes * 60
		
		// 计算基准时间
		if t.BaseTime > 0 {
			t.BaseTime = t.BaseTime - (t.BaseTime % 60)
		} else {
			t.BaseTime = defaultBaseTime
		}
		
		t.BaseTime += t.Id % 60
		
		if err == nil {
			timer, ok := timerList[t.Id]
			if ok {
				if timer.runStatus == 0 {
					timer.IntervalMinutes = t.IntervalMinutes
					timer.timeout = t.timeout
					timer.BaseTime = t.BaseTime
					timer.url = t.url
					timer.uid = t.uid
				}
			} else {
				timerList[t.Id] = &t
			}

			timerList[t.Id].clearTag = 0
		} else {
			Addlog(err.Error())
		}
	}

	for id, timer := range timerList {
		if timer.clearTag == 1 {
			delete(timerList, id)
		}
	}

	return timerList
}

func insertLog(timer *timer, log *runLog) {
	mutex.Lock()
	defer mutex.Unlock()

	conn := CreateDBConn()

	stmtIns, err := conn.db.Prepare("UPDATE timer_process SET last_run_status=?,last_run_time=? WHERE tp_id=?")
	if err != nil {
		AddTimerlog(timer, "updateLogA", err)
		return
	}

	_, err = stmtIns.Exec(log.status, log.startTime, timer.Id)
	if err != nil {
		AddTimerlog(timer, "updateLogB", err)
		return
	}

	stmtIns.Close()

	stmtIns, err = conn.db.Prepare("INSERT INTO run_logs (tp_id,start_runtime,end_runtime,url,result,status)VALUES(?,?,?,?,?,?)")
	if err != nil {
		AddTimerlog(timer, "insertLogA", err)
		return
	}

	var result *string
	if len(log.result) > 1000 {
		s := log.result[0:1000]
		result = &s
	} else {
		result = &log.result
	}

	_, err = stmtIns.Exec(timer.Id, log.startTime, log.endTime, timer.url, *result, log.status)
	if err != nil {
		AddTimerlog(timer, "insertLogB", err)
		return
	}

	stmtIns.Close()
}
