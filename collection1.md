### 基础
- [为什么 Go 有两种声明变量的方式，有什么区别，哪种好？](https://mp.weixin.qq.com/s/ADwEhSA1kFOFqzIyWvAqsA)
- [Go 为什么不在语言层面支持 map 并发？](https://mp.weixin.qq.com/s/70ttSbuQ3RpgKkaaIj3D1g)
- [泛型之后，interface不再是那个interface](https://colobu.com/2022/01/08/the-interface-is-not-that-interface-in-go-1-18/)
- [for 和 range 的性能比较](https://geektutu.com/post/hpg-range.html)：range在迭代大value的数据时性能很差，因为range给的item是value的拷贝；
- [切片(slice)性能及陷阱](https://geektutu.com/post/hpg-slice.html)：大切片使用copy替代re-slice防止不被gc；
- [Reflect反射的性能](https://geektutu.com/post/hpg-reflect.html): 避免使用反射，比如官方json包就是用的反射，可以用[easyjson](https://github.com/mailru/easyjson)代替
- [字符串拼接性能](https://geektutu.com/post/hpg-string-concat.html): 尽量使用`strings.Builder`来进行字符串拼接;


### 实现
- [go-zero实现limit](https://mp.weixin.qq.com/s/s8T1430LS-l3ks9cL7sZyw)
- [go-zero实现令牌桶](https://mp.weixin.qq.com/s/ulGRw4qkWbGKdF83VaIb7A)
- [最简单的服务响应时长优化方法，没有之一](https://mp.weixin.qq.com/s/Ec1nuR5Q_QgaC3FqeX1gLg)

### 并发编程
- [读写锁与互斥锁的性能比较](https://geektutu.com/post/hpg-mutex.html): 读写锁的读写性能明显比互斥锁更好；*这篇文章里面有一段snyc的源码注释:互斥锁如何实现公平*，不清楚的话看下面这篇；
- [sync.mutex 源代码分析](https://colobu.com/2018/12/18/dive-into-sync-mutex/)
- [如何退出协程 goroutine (超时场景)](https://geektutu.com/post/hpg-timeout-goroutine.html): 子协程在读channel需要考虑退出的问题
- [如何退出协程 goroutine (其他场景)](https://geektutu.com/post/hpg-exit-goroutine.html)：子协程在读channel需要考虑退出的问题
- [控制协程的并发数量](https://geektutu.com/post/hpg-concurrency-control.html): 1. 使用buffer channel；2. 使用协程池；
- [sync.Once提升性能](https://geektutu.com/post/hpg-sync-once.html): 提到了最常用的字段作为结构体的第一个字段(热路径,不需要计算偏移);
    
    > 结构体第一个字段的地址和结构体的指针是相同的，如果是第一个字段，直接对结构体的指针解引用即可。如果是其他的字段，除了结构体指针外，还需要计算与第一个值的偏移(calculate offset)。在机器码中，偏移量是随指令传递的附加值，CPU 需要做一次偏移值与指针的加法运算，才能获取要访问的值的地址。因为，访问第一个字段的机器代码更紧凑，速度更快。


### debug、性能分析与编译
- [benchmark基准测试](https://geektutu.com/post/hpg-benchmark.html)
- [pprof 性能分析](https://geektutu.com/post/hpg-pprof.html)
- [减小 Go 代码编译后的二进制体积](https://geektutu.com/post/hpg-reduce-size.html)
- [死码消除与debug模式](https://geektutu.com/post/hpg-dead-code-elimination.html)



## wasted
```markdown
IM系统相关
- [如何设计一个亿级消息量的 IM 系统](https://xie.infoq.cn/article/19e95a78e2f5389588debfb1c)
- [如何保证IM实时消息的“时序性”与“一致性”？](http://www.52im.net/thread-714-1-1.html)

```

- [你所应该知道的A/B测试基础](http://blog.leapoahead.com/2015/08/27/introduction-to-ab-testing/)
- [让代码审查扮演更好的角色](http://blog.leapoahead.com/2016/10/04/code-review-one-step-further/)
- [“函数是一等公民”背后的含义](http://blog.leapoahead.com/2015/09/19/function-as-first-class-citizen/)
- [从开源项目中获得的docker经验](http://blog.leapoahead.com/2015/10/07/docker-lessons-learned-md/)

- [公司内部做的一个分享，有缘人可见](https://mp.weixin.qq.com/s/6AP8E9ilsX5MAMeYW4ia6Q)

- [如何获取客户端真实 IP？从 Gin 的一个 "Bug" 说起](https://mp.weixin.qq.com/s/C-Xf6haLrOWkmBm2lRTdEQ)

- [美团二面：怎么做MySQL优化？](https://mp.weixin.qq.com/s/S7WOWlZY6SUC5dW9UjGpjA)

- [几行代码为老板省百万-某高并发服务Go GC及UDP Pool优化](https://mp.weixin.qq.com/s/YAz5NyiNWJCMlGsJRTAaxw)

- [你了解微服务的超时传递吗？](https://mp.weixin.qq.com/s/vZlFS7lkKplil5cyHirSmA)

- [为什么Go服务容器化之后延迟变高](https://zhuanlan.zhihu.com/p/404699097?utm_source=wechat_session&utm_medium=social&utm_oi=615808713324498944&utm_campaign=shareopn)

- [网络代理 Envoy 开源五周年，创始人 Matt Klein 亲述开源心路历程及经验教训](https://cloudnative.to/blog/envoy-oss-5-year/)

## others
- [Bilibili 毛剑：Go 业务基础库之 Error ](https://mp.weixin.qq.com/s?__biz=MzA4ODg0NDkzOA==&mid=2247487124&idx=1&sn=0f6141c2ccd9a0abc4baf26e04f0fd4c&source=41#wechat_redirect)
- [毛剑：Bilibili 的 Go 服务实践（上篇）](https://mp.weixin.qq.com/s?__biz=MzA4ODg0NDkzOA==&mid=2247487505&idx=1&sn=c9de6535528d2102bee364937201f6e6&source=41#wechat_redirect)
- [毛剑：Bilibili 的 Go 服务实践（下篇）](https://mp.weixin.qq.com/s?__biz=MzA4ODg0NDkzOA==&mid=2247487504&idx=1&sn=9b8663676ee689e0fcd4b990ecf99f3d&source=41#wechat_redirect)
- [Gopher China 2019 讲师专访-bilibili架构师毛剑 ](https://www.sohu.com/a/303913388_657921)
- [Gopher China 2019](https://www.bilibili.com/video/BV1c4411g77Y?p=5)
- [Gopher China 2019 PPT](https://github.com/gopherchina/conference/blob/master/README.md)
- [bilibili技术总监毛剑：B站高可用架构实践](https://zhuanlan.zhihu.com/p/139258985)