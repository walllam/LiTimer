package logic

import (
	"io/ioutil"
	"net/http"
	"time"
)

func Request(timer *timer) {
	var result string
	var status byte = 0
	var log runLog

	log.startTime = time.Now().Unix()
	if timer.runStatus == 0 {
		timer.runStatus = 1

		client := &http.Client{Timeout: time.Duration(timer.timeout) * time.Second}
		resp, err := client.Get(timer.url)
		if err != nil {
			result = err.Error()
		} else {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				result = err.Error()
			} else {
				result = string(body)
				if len(result) >= 17 {
					if result[:17] == "__run_successed__" {
						status = 1
					}
				}
			}
		}

		timer.runStatus = 0
	} else {
		status = 0
		result = "The previous request is still in progress"
	}

	log.endTime = time.Now().Unix()
	log.status = status
	log.result = result

	insertLog(timer, &log)
}
