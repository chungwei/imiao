# 理解Redis

## 前言
Redis是一个开源(BSD许可)的，内存中的数据结构存储系统，它可以用作数据库、缓存和消息中间件。
> Redis = Remote Dictionary Service

Redis以其超高的性能、完美的文档、简洁易懂的源码和丰富的客户端库支持在开源中间件领域广受好评，并成为互联网技术领域使用最为广泛的存储中间件之一。

> 赶紧试试 [官方提供的在线命令学习工具](http://try.redis.io/)


## 关于线程IO模型和通信协议
### 线程IO模型
Redis是单进程，同时对外提供服务使用的是单线程！  
Redis是单进程，同时对外提供服务使用的是单线程！  
Redis是单进程，同时对外提供服务使用的是单线程！  

因为Redis所有的数据都在内存中，所有的运算都是内存级别的运算，所以运算很快。
> 也正因为是单线程，一些时间复杂度较高的指令，一不小心就可能会导致卡顿

同时，由于使用了非阻塞IO，所以可以处理客户端的高并发连接。

> 【提问：Redis如何使用单进程单线程支撑高并发？】

### 通信协议
Redis的通信协议是RESP(即Redis Serialization Protocol)，它是一种直观的文本协议，优势在于实现异常简单，解析性能极好；不足在于数据传输比较浪费流量。

RESP将传输的结构数据分为5种类型，且结束时统一加上回车换行符号`\r\n`，分别是
1. 单行字符串 以 + 符号开头
2. 多行字符串 以 $ 符号开头，后跟字符串长度
3. 整数值 以 : 符号开头，后跟整数的字符串形式
4. 错误消息 以 - 符号开头
5. 数组 以 * 号开头，后跟数组的长度

## 关于事务
- 不具备原子性——不支持回滚
- 不具备严格的持久性
- 不支持严格的一致性
- 支持隔离性——单进程单线程


## 关于Redis基础数据结构


## 关于Redis持久化
详情请参看[理解Redis持久化的实现](https://github.com/chungwei/imiao/blob/master/存储/102-理解Redis持久化的实现.md)

## 关于Redis过期策略和缓存淘汰
详情请参看[理解Redis数据缓存淘汰的实现](https://github.com/chungwei/imiao/blob/master/存储/103-理解Redis数据缓存淘汰的实现.md)

## 经典案例分析和最佳实践


1. [如何发现 Redis 热点 Key ，解决方案有哪些？](https://mp.weixin.qq.com/s?__biz=MzUyNDkzNzczNQ==&mid=2247486805&idx=1&sn=55dd5c2d296b097470fbde17a4e3c3d6&chksm=fa24f23dcd537b2b0a7eef2e20f604980adc58fdc066324bf7eed35ab156e38f87c3166bce7f&scene=21#wechat_redirect)
2. [拼多多面试真题：如何用Redis统计独立用户访问量](https://zhuanlan.zhihu.com/p/69425231)
1. [如何正确访问Redis中的海量数据？服务才不会挂掉！](https://mp.weixin.qq.com/s?__biz=MzU0MzQ5MDA0Mw==&mid=2247485848&idx=1&sn=81c36b7cffb839db18c01bc7667516c6&chksm=fb0be30ccc7c6a1a46fbf602635be7fd2ee949c9fcef508c0b027ce5fc4dc9f8ad7b6b7d4417&mpshare=1&scene=24&srcid=&key=744487672bdc02d1f3059568c7737477f9e721fdac391008a42aab16b9c8fbcbd26a9cffbf469cd174444e8a3f09cf17a3276108faf5678b87807c8c1edf5429402f4457fd86b300acb17eb1634c8782&ascene=0&uin=Nzc3MzQ2MTgy&pass_ticket=6AHO5gBAReB43RtRd6u0irzM75OfuXjdb5Xs9RdAMSwoieMLW7ic1%2Bjsd1djK0Ay)
1. [一份完整的阿里云 Redis 开发规范，值得收藏！](https://mp.weixin.qq.com/s?__biz=MzI3NzE0NjcwMg==&mid=2650123676&idx=1&sn=ee018916b35d11f84e2d16c552f7e5dc&chksm=f36bb0bdc41c39abac4663cd2a38787ba20b70af1f97c9c82b38a9e6893f7a85e4bb8ef2e812&mpshare=1&scene=24&srcid=&key=67c94f9eddc6c14499bff146878410f33a5dd271c0ad3d54847781c6a261613a2a460bf6d735599164df1ebcf80732f758433b5cb35f1eca4008138116b8552a12c0c1cce132fd16204c7a71c879e3cf&ascene=0&uin=Nzc3MzQ2MTgy)
1. [Redis 分布式锁进化史解读 + 缺陷分析](https://mp.weixin.qq.com/s?__biz=MzA5ODM5MDU3MA==&mid=2650864919&idx=1&sn=f9d2218155e7c4a2c04d57970ede4160&chksm=8b661a52bc119344bffbc46292d162163525c9525b3c507bdab23e425494d6e58ab6a364e046&mpshare=1&scene=24&srcid=0220xpwaWSmhP1XEnhtWZxYc&key=51937ec95710ec6392c5232e6c2494e17f5ce044d1b68f3c350a03481b32139b460cb496c8292e10ef498ebcd29614411d76b5832c2bf45218210f70490eae05fa5a184c01e2a3c12aea6029b5da8e8d&ascene=0&uin=Nzc3MzQ2MTgy)
1. [缓存穿透，缓存击穿，缓存雪崩解决方案分析](https://blog.csdn.net/zeb_perfect/article/details/54135506)
1. [Redis系列十：缓存雪崩、缓存穿透、缓存预热、缓存更新、缓存降级](https://www.cnblogs.com/leeSmall/p/8594542.html)
1. [面试题之redis实现限制1小时内每用户Id最多只能登录5次](https://www.cnblogs.com/wujf/p/5206354.html)
2. 超级热 key：向全国推送一篇新闻,千万级用户同时打开这篇新闻,由于流量太大,导致缓存组件 memcached集群中命中这篇新闻的某台机器的网卡被打满,请问如何解决?

### 其他
1. [redis原理](https://zhuanlan.zhihu.com/p/73733011)
1. [Redis总结](https://zhuanlan.zhihu.com/p/55057324)
1. [吃透了这些Redis知识点，面试官一定觉得你很NB（干货 | 建议珍藏）](https://zhuanlan.zhihu.com/p/68694458)
1. [Redis的那些最常见面试问题](https://www.cnblogs.com/Survivalist/p/8119891.html)
1. [Redis 哨兵机制](http://wiki.jikexueyuan.com/project/redis/guard-mechanism.html)
1. [Codis作者黄东旭细说分布式Redis架构设计和踩过的那些坑们](https://mp.weixin.qq.com/s?__biz=MzAwMDU1MTE1OQ==&mid=208733458&idx=1&sn=691bfde670fb2dd649685723f7358fea&scene=2&from=timeline&isappinstalled=0&key=744487672bdc02d166e3f67ac368f073cf0db249cdf3047ec2fc414da5d3f26a580d715c66a7f6bcbeacfaf04bacb10cdf046278fb0f52322ab7139f0d609eaaefa50e82e47242fc96f00b8aedc077ab&ascene=0&uin=Nzc3MzQ2MTgy&devicetype=iMac+MacBookPro11%2C1+OSX+OSX+10.10.3+build(14D136)&version=12020810&nettype=WIFI&lang=zh_CN&fontScale=100&pass_ticket=6AHO5gBAReB43RtRd6u0irzM75OfuXjdb5Xs9RdAMSwoieMLW7ic1%2Bjsd1djK0Ay)
1. [Hello Redis，我有7个问题想请教你](https://mp.weixin.qq.com/s?__biz=MzI1NDQ3MjQxNA==&mid=2247489614&idx=1&sn=803113aa84dfc6796bc93b73903d845d&chksm=e9c5e1ffdeb268e9899349abd769e3459b8dc33f225379c8fa10b55abafa12b8f0b477b45b2c&mpshare=1&scene=24&srcid=&key=51937ec95710ec63deca9e2191d31f40f932318dc82b6ff31b346a33b5908f0cdfd64453e36343259eadee65380a3c392f7a91ec8f54666dad0ea01c3e14b49f9070796e0f7008684968518a3d369794&ascene=0&uin=Nzc3MzQ2MTgy&nettype=WIFI&lang=zh_CN&fontScale=100&pass_ticket=6AHO5gBAReB43RtRd6u0irzM75OfuXjdb5Xs9RdAMSwoieMLW7ic1%2Bjsd1djK0Ay)
1. [面试必备：什么是一致性Hash算法？](https://zhuanlan.zhihu.com/p/34985026)
1. [五分钟看懂一致性哈希算法](https://juejin.im/post/5ae1476ef265da0b8d419ef2)
1. 一致性hash的作用，一致性怎么体现？
1. Redis 哈希槽
2. 