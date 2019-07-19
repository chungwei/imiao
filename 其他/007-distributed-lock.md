# 分布式锁的原理和实现

## 为什么需要锁
资源是有限的，在分布式服务中，我们既要保证效率，又要保证正确性。
> 分布式的 CAP 理论：任何一个分布式系统都无法同时满足一致性（Consistency）、可用性（Availability）和分区容错性（Partition tolerance），最多只能同时满足两项。

## 需要什么样的锁
1. 互斥
2. 高性能
2. 高可用
3. 可重入，避免死锁

## 常见分布式锁的实现
1. MySQL
2. Redis
3. Zookeeper

## 乐观锁 & 悲观锁


## 参考资料
1. [再有人问你分布式锁，这篇文章扔给他](https://juejin.im/post/5bbb0d8df265da0abd3533a5)
2. [分布式锁机制原理及实现方式](https://segmentfault.com/a/1190000016351095)
3. [分布式锁看这篇就够了](https://zhuanlan.zhihu.com/p/42056183)
2. [分布式锁简单入门以及三种实现方式介绍](https://blog.csdn.net/xlgen157387/article/details/79036337)
3. [如何实现靠谱的分布式锁？](https://www.infoq.cn/article/how-to-implement-distributed-lock)
4. [Redis RedLock 完美的分布式锁么？](https://juejin.im/post/59f592c65188255f5c5142d2)
5. [How to do distributed locking](http://martin.kleppmann.com/2016/02/08/how-to-do-distributed-locking.html)
6. [Is Redlock safe?](http://antirez.com/news/101)
7. [RedLock算法-使用redis实现分布式锁服务](https://www.jianshu.com/p/fba7dd6dcef5)
8. [Redlock（redis分布式锁）原理分析](https://www.cnblogs.com/rgcLOVEyaya/p/RGC_LOVE_YAYA_1003days.html)
9. [面试必备之乐观锁与悲观锁](https://juejin.im/post/5b4977ae5188251b146b2fc8)
10. [深入理解乐观锁与悲观锁](https://www.hollischuang.com/archives/934)
