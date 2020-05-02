# Golang并发编程

## 一、并发模型

golang中的并发编程思想：

- 通过通信共享内存（CSP：Communication   Sequential [sɪˈkwenʃl] Process（通信顺序进程））
- 通过锁共享内存

Go 语言的并发性哲学可以这样总结 ：**追求简洁，尽量使用 channel ，并且认为goroutine 的使用是没有成本的**。

<img src="golang并发编程.assets/clip_image001.png" alt="在 Go 语 言 的 FAQ 中 ， 有 如 下 陈 述 ：  为 了 尊 重 mutex, sync 包 实 现 了 mutex, 但 是 我 们 希 望 Go 语 言 的 编 程 风  榕 将 会 激 励 人 们 尝 试 更 高 等 级 的 技 巧 。 尤 其 是 考 虑 构 建 你 的 程 序 ， 以 便 一  次 只 有 一 个 goroutine 负 责 某 个 特 定 的 数 据 。  不 要 通 过 共 享 内 存 进 行 通 信 。 相 反 ， 通 过 通 信 来 共 享 内 存 。 有 数 不 清 的 关  于 Go 语 言 核 心 团 队 的 文 章 、 讲 座 和 访 谈 ， 相 对 于 使 用 像 sync.Mutex 这 样  的 原 语 ， 他 们 更 加 拥 护 CSPO " style="zoom:90%;" />

**在Channel和传统的锁之间如何选择？**

![是  是  否  这 是 一 个 对 性 能 要 求 很 高  的 临 界 区 吗 ？  否  你 想 要 转 让 数 据 的  所 有 权 吗 ？  你 是 否 试 图 在 保 护 某 个  结 构 的 内 部 状 态 ？  你 是 否 试 图 协 调 多 个 逻 辑  片 段 ？  使 传 统 的 锁  使 用 channel ](golang并发编程.assets/clip_image001-1583578821362.png)





## 二、goroutine特性





## 三、channel相关

- channel如果关闭后，channel中的数据仍然可以读取出来，但是如果再对此channel进行写操作，会引发panic错误

- 如何判断一个channel中是否含有数据？

- 使用close函数关闭通道
- 数据不发送完，不应该关闭通道
- 已经关闭的channel，不能再向其写数据。如果向已经关闭的通道写数据，会报错。panic：send on closed channel
- 已经关闭的channel，可以从中读到数据0。读到0就说明：写端已经关闭了

















