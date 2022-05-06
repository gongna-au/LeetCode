package main

/*
Redis 近LRU 介绍
上文完成了咱们自己的 LRU 实现，
现在现在聊一聊 Redis 中的近似 LRU 。
由于真实LRU需要过多的内存（在数据量比较大时），
所以 Redis 是使用一种随机抽样的方式，来实现一个近似 LRU的效果。
说白了，LRU 根本只是一个预测键访问顺序的模型。
*/

func main() {

}
