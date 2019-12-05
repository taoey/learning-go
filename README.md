# Lets-GO

## 前言
从零开始进行GO语言开发，相关文档及示例代码



## 有关go的相关学习资料

### 社区

- [Go语言中文网](https://studygolang.com/)
- [web框架 iris中文手册](https://studyiris.com/doc/index.html)
- [Go by Example 英文网](https://gobyexample.com/)
- [Go By Example 中文网](https://books.studygolang.com/gobyexample/)

### 推荐书籍

| 名称                                                         | 备注                                                         |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| [Mastering GO](https://github.com/hantmac/Mastering_Go_ZH_CN) | 适用于go的深入了解，包括其内部机制等等                       |
| [Go 语言标准库](https://www.kancloud.cn/wizardforcel/golang-stdlib-ref/121475) |                                                              |
| [over-golang](https://github.com/overnote/over-golang)       | github开源书籍，涉及Go语法、Go并发思想、Go与web开发、Go微服务设施等，适用于深入学习go |
| [研磨设计模式](https://github.com/senghoo/golang-design-pattern) | go语言版设计模式                                             |
| [the way to go](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/directory.md) |                                                              |
| [All programming languages books](https://github.com/KeKe-Li/book) |                                                              |
| [看云-golang](https://www.kancloud.cn/uvohp5na133/golang/933968) | 有很多错误及面试总结，适合各阶段go语言开发者                 |
| [Go语言标准库](https://books.studygolang.com/The-Golang-Standard-Library-by-Example/) |                                                              |

### 推荐文章/资料

| 文章                                                         | 评价                                                 |
| ------------------------------------------------------------ | ---------------------------------------------------- |
| [crawlab的golang后端内存分析及优化-基于go pprof](https://juejin.im/post/5d5be347f265da03b94ff66b) | ioutil.ReadAll()读取大文件有坑                       |
| [记一次golang程序CPU高的排查过程](https://juejin.im/post/5d5189446fb9a06b1a567e93) | time.NewTicker使用需谨慎                             |
| [深入理解Go-垃圾回收机制](https://juejin.im/post/5d78b3276fb9a06b1829e691) | 非常深入                                             |
| [go 垃圾回收：三色算法](https://juejin.im/post/5d398417f265da1b904c26b6) | 垃圾回收算法专场                                     |
| [Go map原理剖析](https://juejin.im/post/5d9c650a518825091b2c2679) | 读Map源码已经成为一种习惯                            |
| [Golang高效编程](https://juejin.im/post/5d958b9be51d4577f4608b2b) | 有很多值的借鉴的内容                                 |
| [[译] Go 实现百万 WebSocket 连接](https://juejin.im/post/5d48f1cd6fb9a06b233ca719) | 现在还看不太懂，反正感觉挺NB的                       |
| [golang框架解析-iris](https://mp.weixin.qq.com/s?__biz=MzA5MDEwMDYyOA==&mid=2454619020&idx=1&sn=c74e06ce6ce6805c9fbeb357b484e284&chksm=87aae577b0dd6c61c8aa7057873ebba5567057ca816fb8a17b664dcf696f0b9544866de7c6c5&mpshare=1&scene=1&srcid=1005Rk3kOZl8R1xRL4qWZtLc&sharer_sharetime=1570257270847&sharer_shareid=06041e0e5e8bc247cd43fed6c5ced62a#rd) | iris 生命周期这个图非常好                            |
| [beego框架代码分析](https://mp.weixin.qq.com/s?__biz=MzA5MDEwMDYyOA==&mid=2454618967&idx=1&sn=6cafd61e5a57ab7950901ea9ac3c0e44&chksm=87aae5acb0dd6cba7e999db9a43eaa7c30f9f22ad1cd67d8008e2757a21241f853d2ee0af5eb&scene=21#wechat_redirect) |                                                      |
| [一文理清 Go 引用的常见疑惑](https://mp.weixin.qq.com/s/o-iE3ny3-GOIhcWsUbVgVA) | 传值？传址？引用传递？                               |
| [Golang实现请求限流的几种办法](https://blog.csdn.net/micl200110041/article/details/82013032) |                                                      |
| [Go channel 实现原理分析](https://www.jianshu.com/p/d841f251d3bc) |                                                      |
| [go 学习笔记之10 分钟简要理解 go 语言闭包技术](https://mp.weixin.qq.com/s/GJnvPgW7IONK9LVw-i34hQ) |                                                      |
| [如何在golang http服务端程序中读取2次Request Body？](https://www.zhihu.com/question/329045911/answer/714781838) | 一般情况下是这个需求是不需要的                       |
| [os.open竟然在文件不存在时返回err==nil，亏我那么信任它](https://studygolang.com/topics/10068) | 在windows下直接使用`con`作为文件名，文件直接创建失败 |
| [go 学习笔记之学习函数式编程前不要忘了函数基础](https://mp.weixin.qq.com/s/dprkCOvPZHr6fi_qC91dVw) | 你是gopher？居然不知道啥是函数式编程，好吧，你走吧   |
| [gopherchina](https://github.com/gopherchina/conference)     |                                                      |
| [Golang的反射reflect深入理解和示例](https://studygolang.com/articles/12348?fr=sidebar) | 写的挺好的一片文章                                   |

### 工具

- [jsonTOGo ](https://mholt.github.io/json-to-go/)：好用的json转go struct工具
- [gopm.io](https://gopm.io/)：科学下载第三方包



### 相关库

| 名称                                            | 功能               | 备注                                                         |
| ----------------------------------------------- | ------------------ | ------------------------------------------------------------ |
| [excelize](https://xuri.me/excelize/zh-hans/)   | excel文件操作      |                                                              |
| [sprig](https://github.com/Masterminds/sprig)   | 有用的 Go 模板函数 |                                                              |
| [go-qrcode](https://github.com/skip2/go-qrcode) | 二维码生成         | [Go语言生成二维码是如此简单](https://www.flysnow.org/2017/09/29/go-qrcode.html) |
| [go-micro](https://github.com/micro/go-micro)   | go的微服务框架     | [micro-in-cn/tutorials](https://github.com/micro-in-cn/tutorials) : Micro 中文示例、教程、资料，源码解读 |

