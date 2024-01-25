原文：[Golang Weekly Issue #492](https://golangweekly.com/issues/492)

---

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60/t09miiozkfvedpu1ur45.jpg)](https://golangweekly.com/link/150277/web)  


[2024 年 Rust 和 Go 的对比？](https://golangweekly.com/link/150277/web "bitfieldconsulting.com") —— 去年的一篇长期流行的帖子已针对 2024 年的情况进行了调整（尽管优点和缺点仍然非常相似），它对“两种都很棒”的语言进行了比较，包括它们都擅长什么、它们的主要差异以及一些权衡考虑。

_John Arundel_ 
  

[在大型代码库中调试 Go 编译器性能](https://golangweekly.com/link/150278/web "incident.io") —— 代码库越大，编译所需的时间就越长。不过，Incident 的人员深入研究了依赖关系图，并找出了一些瓶颈，从而使他们的构建速度更快。

_Isaac Seymour_ 


[![](https://copm.s3.amazonaws.com/9bd98a7f.png)](https://golangweekly.com/link/150276/web) 

[帮助在希腊雅典举办的 GopherConEU 成为一场令人难忘的活动吧！](https://golangweekly.com/link/150276/web "events.ardanlabs.com") —— 参加 2024 年 2 月 6 日至 8 日在希腊雅典举行的 GopherConEU。Bill Kennedy 和 Miki Tebeka 参加了 GopherConEU 培训研讨会，分享前沿的 Go 工程最佳实践、库、框架、性能优化等。

_Ardan Labs & GopherConEU sponsor_


▶ [Russ Cox 谈 Go 的变化](https://golangweekly.com/link/150279/web "www.youtube.com") —— Russ 在 GopherCon 2023 上的演讲现已上线，很好地概述了 Go 对变化的态度以及它如何随着时间的推移不断变得更好。_（24 分钟。）_ 

_Russ Cox_ 


[忽略 Go 二进制文件中的开发依赖项](https://golangweekly.com/link/150280/web "rednafi.com") —— 将开发包排除在生产构建之外，并保持二进制文件较小的技巧。我们想象这个问题的修复最终会出现在 Go 工具链中，但在那之前……

_Redowan Delowar_ 


💡 在这方面，你可能会对 Alex Edwards 的 [《如何使用 `go run` 来管理工具依赖关系》](https://golangweekly.com/link/150281/web) 感兴趣。  
  

_快速了解：_

  * 📈 在[最新的 TIOBE Index 更新](https://golangweekly.com/link/150282/web)中，Go 在语言排行榜上继续攀升。它现在排名第 11 位。

  * 📗 [Jon Bodner 的 _Learning Go_](https://golangweekly.com/link/150312/web) 的第二版已出版。 现可在 O'Reilly 购买该书，或者[从 Amazon](https://golangweekly.com/link/150313/web) 等网站预定印刷版本。

  * [`sqlrange` 实验](https://golangweekly.com/link/150283/web) 使用 Go 1.22 的范围函数以及 `database/sql`……

  * 有[一项新提案](https://golangweekly.com/link/150284/web)，即添加一个 `go:wasmexport` 指令，作为一种标记要导出到 WebAssembly 主机的函数的方式。

  * 👋 AWS 宣布将于 2025 年 7 月[终止对 AWS SDK for Go (v1) 的支持](https://golangweekly.com/link/150286/web)，并于今年 7 月进入维护模式。 [v2 仍在等待中……](https://golangweekly.com/link/150287/web)


[Go 1.22 中的高效 HTTP 路由](https://golangweekly.com/link/150288/web "www.talkativedev.com") —— 了解 Go 下一版本中内置 HTTP 服务器多路复用器的[预期更改](https://golangweekly.com/link/150289/web)。如果你不喜欢它，你可以随时[转向 Gin……](https://golangweekly.com/link/150290/web)

_Cheikh Seck_ 
  

[使用 `deadcode` 查找不可达函数](https://golangweekly.com/link/150291/web "go.dev") —— 在圣诞节到来之际，我们不知何故错过了 Go 官方博客上的这篇文章！核心团队的一名成员调查 Go 应用程序中死的、无法访问的代码，此外，Go 团队还提供的一个工具可以帮助您自信地删除这些代码。

_Alan Donovan_ 



[从 Slow 到 SIMD：一个关于 Go 优化的故事](https://golangweekly.com/link/150292/web "sourcegraph.com") —— 关于 Sourcegraph 的人员是如何需要加快一些基本数学以处理嵌入向量的故事，考虑的初始步骤，以及 SIMD 设法为性能表带来什么 _（剧透：很多）_。

_Camden Cheek (Sourcegraph)_ 


▶ [是时候从 Docker 切到 Podman 了吗？](https://golangweekly.com/link/150293/web "www.youtube.com")

_Christian Lempa_ 


🛠 代码和工具

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/o3u1xsgb3kk90hmybggs.jpg)](https://golangweekly.com/link/150294/web)  

[Tint：`slog.Handler` 输出有色日志](https://golangweekly.com/link/150294/web "github.com") —— 有色，如彩色。

_lmittmann_ 


[ffmpeg-go：FFmpeg 的 Go 绑定](https://golangweekly.com/link/150295/web "github.com") —— 在 C 库 API 上进行了小小的封装，因此，您可以获得很多的功能，但还是要编写大量的代码才能使用它。这是[视频转码的一个例子](https://golangweekly.com/link/150296/web)。仅适于 Linux 和 macOS。

_Chandler Newman_ 


[《构建 Temporal Cloud》网络研讨会系列](https://golangweekly.com/link/150297/web "t.mp") —— 了解我们如何构建 Temporal Cloud，从而为所有工作负载提供世界一流的延迟、性能和可用性。

_Temporal Technologies sponsor_


[XLSX 3.3.5：一个用于读写 XLSX (Excel) 文件的库](https://golangweekly.com/link/150298/web "github.com") —— 您知道 Excel 电子表格文件基本上是 XML 吗？这使得它们非常易于操作和使用，但像这样的库让操作变得更加容易。

_Geoffrey J. Teale_ 


[Goyek 2.1：一个用于任务自动化的 Go 库](https://golangweekly.com/link/150299/web "github.com") —— 一个简单的跨平台任务自动化库，其中，任务以类 [Cobra](https://golangweekly.com/link/150300/web) 风格进行定义，而操作则采用单用测试的感觉进行定义。[v2.1](https://golangweekly.com/link/150301/web) 增加了对并行任务执行的支持。

_Robert Pająk_ 


[SQLBoiler 4.16：一个“数据库优先” ORM](https://golangweekly.com/link/150302/web "github.com") —— 首先在数据库级别创建模式，然后根据底层实际情况生成 Go 的 ORM 代码。 

_Volatile Technologies Inc._ 


[f-license：许可证密钥生成以及验证服务器](https://golangweekly.com/link/150303/web "github.com") —— 一个用来为其他应用创建及验证许可证密钥的系统。

_Furkan Senharputlu_ 


[go-quartz 0.10](https://golangweekly.com/link/150304/web) - 极简、无依赖的调度库。

[go-elasticsearch 8.12.0](https://golangweekly.com/link/150305/web) - 官方 Elasticsearch 客户端库。

[Carbon 2.3.6](https://golangweekly.com/link/150306/web) - 友好的时间操作库。

[Kivik 4.2](https://golangweekly.com/link/150307/web) - （类）CouchDB 数据库的通用接口。

[betteralign 0.3.4](https://golangweekly.com/link/150308/web) - 优化结构布局。

[Ginkgo 2.15.0](https://golangweekly.com/link/150309/web) - 成熟的测试框架。


[Go Feature Flag 1.21.0](https://golangweekly.com/link/150310/web)
