## 可用性

> 事故: 责任事故/非责任事故.
>
> ​	责任事故: 明明告诉你怎么做,你偏偏不那么做.
>
> ​	非责任事故: 

* 隔离
  * 服务隔离:读写分离(https://zhuanlan.zhihu.com/p/115685384)/动静分离
  * 轻重隔离:
  * 物理隔离:
* 超时控制
* 过载保护: 自保护,不管其他
  * [限速](https://pkg.go.dev/golang.org/x/time/rate ) - 基础库提供限速
  * [QPS](https://zhuanlan.zhihu.com/p/84012183) - QPS计算方法
  * 可控延迟算法: CoDel
  * [利特尔法则](https://zhuanlan.zhihu.com/p/65687548) - L = λW
  * [滑动窗口]() - 计算w
  * 计算系统吞吐? - 简单滑动平均值,计算拉莫纳
  * [机器学习](https://www.bilibili.com/video/BV164411S78V?p=6&t=12) - 吴恩达的机器学习,工程学习
* 限流: 服务级别的全局限流.
  * 分布式限流: 
* 熔断: 
  * 断路器: circuit breaker,
  * 客户端流控: 