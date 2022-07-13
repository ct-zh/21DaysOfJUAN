```
A: k8s启动grpc的服务遇到过负载不均衡的事情了没，当时不知道，前几天才看了一篇文章才明白
B: 长链接么
A: 是的，http2就是长连接，所以才在群上找各位大佬请教一下。
C: https://pandaychen.github.io/2019/07/11/GRPC-SERVICE-DISCOVERY/
D: 可是，既然http了，那应该是client到server的流量吧？ 不是前面挂一个ng，openrestry就成了吗。。
A: 短连接是可以，长链接会有问题
C: 要用支持HTTP2的负载均衡中间件才行。client端自己负载，需要知道server的endpoint ，比如xds 协议，  或者headless服务 拿到全部endpoint，或者server mesh  挂个sidecar 。。。
D: 南北流量和东西流量的策略本就应该是不一样的吧。
```