### redis 使用

##### string
```
set hello word
get hello
del hello
```

##### 列表（List）-- 只有键没有值,相同元素可以重复出现
list-key val1 val2 val2...  (相同元素可以重复出现)
```
rpush list-key val1
rpush list-key val1
rpush list-key val2
lrange list-key 0 -1  -- 获取列表在给定范围内的所以值
lindex list-key 1 -- 使用Lindex 可以从列表里取出单个元素
lpop list-key -- 弹出元素
```

### 集合(set) 只有键，没有值，每个值都是唯一的
```
sadd set-key item1
sadd set-key item2
sadd set-key item3
sadd set-key item1
# 取出所有元素
SMEMBERS set-key
# 判断元素是否存在
SISMEMBER set-key item4
# 如果存在则移除
SREM set-key item3
```

##### 散列（hash）
微缩版的redis
```
hset hash-key sub-key value
hset hash-key sub-key1 value1
hset hash-key sub-key2 value1
#
hget hash-key sub-key
hgetall hash-key
#
hdel hash-key sub-key
```
##### 有序集合(ZSET)
键：成员
值：分值（必须为浮点数）
既可以根据成员访问元素（类似hash）,也可以根据分值及分值的排序访问元素
```
zadd zset-key 666 member1
zadd zset-key 666 member2
ZRANGE zset-key 0 -1
ZRANGE zset-key 0 -1 withscores
ZREM zset-key member1
```
