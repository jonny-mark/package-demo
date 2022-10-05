## 📗 Prometheus监控
![结构图](https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fimg-blog.csdnimg.cn%2Fimg_convert%2Fb20ec7fa1b2f27f85eb54fc7c2a443e7.png&refer=http%3A%2F%2Fimg-blog.csdnimg.cn&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1666881564&t=fa347edd8cedb5d4754502a2579eee61)

主要知识点
```shell
指标分类
Counter:累加指标计数器,单调递增,如http请求数.服务重启时才会清零
Gauge:一个动态的数值指标,比如CPU使用率,内存使用率等
Historygram:Historygram是直方图,适合需要知道数值分布范围的场景,比如http请求的响应时长,http请求的响应包体大小等.
Summary:相比Historygram是按百分位聚合好的直方图,需要知道百分比分布范围的场景,如http请求的响应时长场景,Historygram可以统计响应时长为1ms的有多少个
```
```shell
初始化时,client_golang注册了2个Prometheu收集器
进程收集器 - 用于收集基本的 Linux 进程信息,比如 CPU,内存,文件描述符使用情况,以及启动时间等
Go收集器 - 用于收集有关 Go 运行时的信息,比如GC,gouroutine和OS线程的数量的信息
详情参考url客户端指标
```



链接<br>
[中文文档](https://www.prometheus.wang)

[客户端指标](https://juejin.cn/post/7027660874856792095)

