package mycache

import (
	"fmt"
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"time"
)

var globalCache cache.Cache

func init() {
	var err error
	// memory
	// globalCache, err := cache.NewCache("memory", `{"interval":60}`)
	// redis
	globalCache, err = cache.NewCache("redis", `{"key":"collectionName", "conn": "127.0.0.1:6379", "dbNum":"0"}`)
	if err != nil {
		fmt.Println(err)
	}
}

func InitCache() {
	fmt.Println("wwwwwwwwwww")

	globalCache.Put("astaxie", 1, 1000*time.Second)
	globalCache.Put("11111", 1111, 1000*time.Second)
	globalCache.Get("astaxie")
	// fmt.Println(bm.Get("astaxie"))
	fmt.Println(globalCache.IsExist("astaxie"))
	globalCache.Delete("astaxie")
	fmt.Println("wwwwwwwwwww")
}
