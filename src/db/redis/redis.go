package redis

func Store(cid string, field string, value string) (interface{}, error) {
	conn := RedisClient.Get()
	defer conn.Close()
	return conn.Do("HMSET", cid, field, value)
}

func Get(cid string, field string) (interface{}, error) {
	conn := RedisClient.Get()
	defer conn.Close()
	return conn.Do("HGET", cid, field)
}

func Delete(cid string) (interface{}, error) {
	conn := RedisClient.Get()
	defer conn.Close()
	return conn.Do("DEL", cid)
}

//func RpushTask(queueName string, task []byte) (interface{}, error) {
//	conn := RedisClient.Get()
//	defer conn.Close()
//	return conn.Do("RPUSH", queueName, task)
//}
//
//func LpopTask(queueName string) (interface{}, error) {
//	conn := RedisClient.Get()
//	defer conn.Close()
//	return conn.Do("LPOP", queueName)
//}
//
//func BlpopTask(queueName string, timeout int) (interface{}, error) {
//	conn := RedisClient.Get()
//	defer conn.Close()
//	return conn.Do("BLPOP", queueName, timeout)
//}
