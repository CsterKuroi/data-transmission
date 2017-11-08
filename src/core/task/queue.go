package task

import (
	"time"

	"uniswitch-agent/src/db/redis"

	"github.com/astaxie/beego/logs"
)

func EnqueueTask(taskStr []byte) {
	redis.RpushTask("send-queue", taskStr)
}

func DequeueTask() {
	for {
		task, err := redis.BlpopTask("send-queue", 0)
		if err != nil {
			logs.Error("Redis BLPOP ERROR : ", err)
			time.Sleep(10 * time.Second)
			continue
		}
		if task == nil {
			logs.Error("Get Nil Task Data!")
			continue
		}
		taskStr := string(task.([]interface{})[1].([]byte))
		logs.Info(taskStr)
		//TODO do send

		//TODO sucess update

	}
}
