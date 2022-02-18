package kredis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/guojiangli/picasso/pkg/klog/baselogger"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

// GlobalKeyPrefix is global redis key preifx
const GlobalKeyPrefix = "kcron:"

// RedisDriver is redisDriver
type RedisDriver struct {
	Client *redis.Client
	Key    string
	alive  bool
	logger baselogger.Logger
}

// NewDriver return a redis driver
func NewDriver(c *redis.Client) (*RedisDriver, error) {
	return &RedisDriver{
		Client: c,
	}, nil
}

// Ping is check redis valid
func (rd *RedisDriver) Ping() error {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("fail redis pool ping %v", err)
		}
	}()
	if err := rd.Client.Ping(context.TODO()).Err(); err != nil {
		return err
	}
	return nil
}

func (rd *RedisDriver) getKeyPre(serviceName string) string {
	return GlobalKeyPrefix + serviceName + ":"
}

// DoHeartBeat set heart beat
func (rd *RedisDriver) DoHeartBeat(nodeID string, timeout time.Duration) {
	// 每间隔timeout/2设置一次key的超时时间为timeout
	if rd.alive {
		return
	}
	tickers := time.NewTicker(timeout / 2)
	rd.alive = true
	key := nodeID
	c := func() {
		rd.alive = false
		tickers.Stop()
	}
	defer c()
	for range tickers.C {
		res, err := rd.Client.Expire(context.TODO(), key, timeout).Result()
		if err != nil || !rd.alive || res == false {
			return
		}
	}
}

func (rd *RedisDriver) IsCheckAlive() bool {
	return rd.alive
}

// GetServiceNodeList get a serveice node  list
func (rd *RedisDriver) GetServiceNodeList(serviceName string) ([]string, error) {
	mathStr := fmt.Sprintf("%s*", rd.getKeyPre(serviceName))
	return rd.scan(mathStr)
}

// RegisterServiceNode  register a service node
func (rd *RedisDriver) RegisterServiceNode(serviceName string, lifeTime time.Duration) (nodeID string) {
	nodeID = uuid.New().String()
	key := rd.getKeyPre(serviceName) + nodeID
	_, err := rd.Client.Set(context.TODO(), key, nodeID, lifeTime).Result()
	if err != nil {
		return ""
	}
	rd.alive = false
	return key
}

func (rd *RedisDriver) scan(matchStr string) ([]string, error) {
	cursor := 0
	ret := make([]string, 0)
	iter := rd.Client.Scan(context.TODO(), uint64(cursor), matchStr, 0).Iterator()
	for iter.Next(context.TODO()) {
		if err := iter.Err(); err != nil {
			return ret, err
		}
		ret = append(ret, iter.Val())
	}
	return ret, nil
}
