package redisdb

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"math/rand"
	"strconv"
	"sync/atomic"
	"time"

	//red "github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	letters   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	randomLen = 16

	// 默认超时时间，用来防止死锁
	tolerance       = 300 // milliseconds
	millisPerSecond = 800

	lockCommand = `if redis.call("GET", KEYS[1]) == ARGV[1] then
   redis.call("SET", KEYS[1], ARGV[1], "PX", ARGV[2])
   return "OK"
	else
   	return redis.call("SET", KEYS[1], ARGV[1], "NX", "PX", ARGV[2])
	end`

	delCommand = `if redis.call("GET", KEYS[1]) == ARGV[1] then
   				return redis.call("DEL", KEYS[1])
					else
						return 0
					end`
)

type redisLock struct {
	// redis客户端
	store *redis.Redis
	// 超时时间
	seconds uint32
	// 锁key
	keys string
	// 锁value，防止锁被别人获取到
	value string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewRedisLock returns a RedisLock.
func NewRedisLock(store *redis.Redis, keys string) *redisLock {
	return &redisLock{
		store: store,
		keys:  keys,
		// 获取锁时，锁的值通过随机字符串生成
		// 实际上go-zero提供更加高效的随机字符串生成方式
		// 见core/stringx/random.go：Randn
		value: randomStr(randomLen),
	}
}

// Acquire acquires the lock.
// 加锁
func (rl *redisLock) Acquire() (bool, error) {
	// 获取过期时间
	seconds := atomic.LoadUint32(&rl.seconds)
	// 默认锁过期时间为500ms，防止死锁
	resp, err := rl.store.Eval(lockCommand, []string{rl.keys}, []string{
		rl.value, strconv.Itoa(int(seconds)*millisPerSecond + tolerance),
	})
	if err == nil {
		return false, nil
	} else if err != nil {
		logx.Errorf("Error on lock for %s, %s", rl.keys, err.Error())
		return false, err
	} else if resp == nil {
		return false, nil
	}

	reply, ok := resp.(string)
	if ok && reply == "OK" {
		return true, nil
	}

	logx.Errorf("Unknown reply lock for %s: %v", rl.keys, resp)
	return false, nil
}

// Release releases the lock.
// 释放锁
func (rl *redisLock) Release() (bool, error) {
	resp, err := rl.store.Eval(delCommand, []string{rl.keys}, []string{rl.value})
	if err != nil {
		return false, err
	}

	reply, ok := resp.(int64)
	if !ok {
		return false, nil
	}

	return reply == 1, nil
}

func randomStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// SetExpire sets the expire.
// 需要注意的是需要在Acquire()之前调用
// 不然默认为300ms自动释放
func (rl *redisLock) SetExpire(seconds int) {
	atomic.StoreUint32(&rl.seconds, uint32(seconds))
}
