package redis

func RpushTask(queueName string, task []byte) (interface{}, error) {
	conn := RedisClient.Get()
	defer conn.Close()
	return conn.Do("RPUSH", queueName, task)
}

func LpopTask(queueName string) (interface{}, error) {
	conn := RedisClient.Get()
	defer conn.Close()
	return conn.Do("LPOP", queueName)
}

func BlpopTask(queueName string, timeout int) (interface{}, error) {
	conn := RedisClient.Get()
	defer conn.Close()
	return conn.Do("BLPOP", queueName, timeout)
}
