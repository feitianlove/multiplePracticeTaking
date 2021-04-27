package RedisLock

import (
	"context"
	redis "github.com/go-redis/redis/v8"
	"time"
)

type RedisConf struct {
	ListenPort   string
	IdleTimeout  int
	MinIdleConns int
	MaxConnAge   int
}

var redisClient *redis.Client

type RedisLock struct {
	lockKey    string
	counterKey string
}

func NewRedisClient(conf RedisConf) (*redis.Client, error) {
	redisServer := redis.NewClient(&redis.Options{
		Addr:         conf.ListenPort,
		Password:     "", // no password set
		DB:           0,  // use default DB
		MinIdleConns: conf.MinIdleConns,
		MaxConnAge:   time.Millisecond,
		IdleTimeout:  time.Microsecond,
	})
	_, err := redisServer.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return redisServer, nil
}

func init() {
	conf := RedisConf{
		ListenPort:   "",
		IdleTimeout:  0,
		MinIdleConns: 0,
		MaxConnAge:   0,
	}
	client, err := NewRedisClient(conf)
	if err != nil {
		panic(err)
	}
	redisClient = client
}

func NewRedisLock(lockKey, counterKey string) RedisLock {
	return RedisLock{
		lockKey:    lockKey,
		counterKey: counterKey,
	}
}
func (l RedisLock) Lock() error {
	// lock
	resp := redisClient.SetNX(context.Background(), l.lockKey, 1, time.Second*5)
	lockSuccess, err := resp.Result()

	if err != nil || !lockSuccess {
		return err
	}

	getResp := redisClient.Get(context.Background(), l.counterKey)
	cntValue, err := getResp.Int64()
	if err == nil || err == redis.Nil {
		cntValue++
		resp := redisClient.Set(context.Background(), l.counterKey, cntValue, 0)
		_, err := resp.Result()
		if err != nil {
			// log err
			return err
		}
	}
	return nil
}
func (l RedisLock) UnLock() error {
	delResp := redisClient.Del(context.Background(), l.lockKey)
	unlockSuccess, err := delResp.Result()
	if err == nil && unlockSuccess > 0 {
		return nil
	} else {
		return err
	}
}
