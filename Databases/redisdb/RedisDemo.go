package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接 standalone连接NewClient
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "root", // no password set
		DB:       0,      // use default DB
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

// https://www.liwenzhou.com/posts/Go/go_redis/
// https://pkg.go.dev/github.com/go-redis/redis/v8#pkg-examples
// 测试用例：https://juejin.cn/post/6910987709070704654
func main() {
	initClient()
	//test_set_or_get()
	test_zset()

}

func test_set_or_get() {
	// set test
	err := rdb.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}

	val, err := rdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed, err:%v\n", err)
		return
	}
	fmt.Println("score", val)

	val2, err := rdb.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	} else {
		fmt.Println("name", val2)
	}
}

func test_zset() {

	zsetKey := "language_rank"
	languages := []redis.Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 99.0, Member: "C/C++"},
	}
	fmt.Println("------------------------------------------------------------------")

	// ZADD
	num, err := rdb.ZAdd(zsetKey, languages...).Result()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Printf("zadd %d succ.\n", num)

	fmt.Println("------------------------------------------------------------------")
	// 把Golang的分数加10
	newScore, err := rdb.ZIncrBy(zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	fmt.Println("------------------------------------------------------------------")
	// 取分数最高的3个
	ret, err := rdb.ZRevRangeWithScores(zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Printf("zrevrange failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	fmt.Println("------------------------------------------------------------------")
	// 取95~100分的
	op := redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = rdb.ZRangeByScoreWithScores(zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

//------------------------------------------------------------------
//zadd 5 succ.
//------------------------------------------------------------------
//Golang's score is 100.000000 now.
//------------------------------------------------------------------
//Golang 100
//C/C++ 99
//Java 98
//------------------------------------------------------------------
//Python 95
//JavaScript 97
//Java 98
//C/C++ 99
//Golang 100

func getKeysByPrefix() {

	//ctx := context.Background()

	//vals, err := rdb.Keys(ctx, "prefix*").Result()

	//res, err := rdb.Do(ctx, "set", "key", "value").Result()

}
