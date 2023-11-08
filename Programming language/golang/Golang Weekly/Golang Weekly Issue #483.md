原文：[Golang Weekly Issue #483](https://golangweekly.com/issues/483)

---

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/juzgfucfglgralmvavdb.jpg)](https://golangweekly.com/link/147341/web)   
  

[Charm 筹集了 600 万美金来构建下一代命令行](https://golangweekly.com/link/147341/web "charm.sh") -- 与 Go 社区的许多成员一样，我们一直是 [Charm](https://golangweekly.com/link/147342/web) 构建有用的基于 Go 的工具和 CLI 库（例如[Bubble Tea](https://golangweekly.com/link/147343/web)）的方法的忠实粉丝，所以这是很棒的消息。这主要也与筹集的资金 _无_ 关，而是对 Charm 在 CLI 开发领域的工作以及他们迄今为止发布的项目的总结。

_Toby Padilla and Christian Rocha (Charm)_ 


[为什么团队使用 Go 而不是 Rust 来开发桌面应用](https://golangweekly.com/link/147344/web "blog.moonguard.dev") -- [Krater](https://golangweekly.com/link/147373/web)，一个用于调试 Laravel 应用的桌面应用，最开始是使用 Tauri 和 Rust 进行开发的，但后来因为 Go 的开发者经历，从而更改为使用 Go 和 Wails 进行开发。

_Ronald Pereira_ 


[![](https://copm.s3.amazonaws.com/b59bec3e.png)](https://golangweekly.com/link/147340/web) 

[一个专为现代工作负载构建的数据存储](https://golangweekly.com/link/147340/web "www.dragonflydb.io") -- Dragonfly 是 Redis 的直接替代品。基于 Dragonfly 构建的应用程序具有现代应用程序所需的速度、可靠性和可扩展性，使您能够为用户提供令人难以置信的体验，同时降低成本和复杂性。

_Dragonfly sponsor_


[Go、容器和 Linux 调度器](https://golangweekly.com/link/147346/web "www.riverphillips.dev") -- 如果你在容器内运行 Go，那么不管容器运行时设置了何种限制，运行时都会为每个 CPU 核创建一个线程。这会导致显着的 GC 循环时间，但别担心，`GOMAXPROCS` 是你的朋友呀，很容易你就可以使它。

_River Phillips_ 


_快速了解：_


🎤 Go 团队的 Filippo Valsorda 和 Roland Shoemaker 加入了 _Go Time_ 播客，[讨论 Go 密码学库中的新增内容。](https://golangweekly.com/link/147347/web)


📗 [《使用 Go 实现家居自动化》](https://golangweekly.com/link/147348/web)是 Ricardo Gerardi 和 Mike Riley 即将出版的 _实用书架_ 电子书，重点介绍如何在 Raspberry Pi 上使用 Go 来构建运动探测器、温度传感器等。


[Go 1.21 中的泛型排序：两个性能的故事](https://golangweekly.com/link/147349/web "aead.dev") -- 对 Go 1.21 中引入的新泛型函数 `slices.Sort` 的性能分析表明，虽然它为数字提供了更快的排序方法，但在对 _字符串_ 进行排序时却意外地表现不佳。[GitHub 上的此问题](https://golangweekly.com/link/147350/web)可能会帮助您了解原因以及解决方法。

_Andreas Auernhammer_ 
  

[运行“Reflections on Trusting Trust”编译器](https://golangweekly.com/link/147345/web "research.swtch.com") -- Russ 讲述了 Ken Thompson（Unix 的共同创建者）在 80 年代早期是如何在 C 程序中创建后门的故事，而 Russ 试图在 Go 中重新创建后门。

_Russ Cox_ 

 
---  

📰 分类广告

[作为站点可靠性工程师，加入  Sticker Mule 的“强悍”团队吧！](https://golangweekly.com/link/147351/web) 我们的软件团队遍布 17 个国家/地区，我们正在寻找更多优秀的工程师加入我们的安全团队。

[领导一支工程师团队，来塑造](https://golangweekly.com/link/147352/web) [Dyte](https://golangweekly.com/link/147353/web) 实时视频的未来：领导一支工程团队负责关键后端系统和 API。优化、规模化并在塑造产品方向方面发挥关键作用。

💻 Hired 让求职变得简单 - 公司不再需要追逐招聘人员，而是预先向您提供薪资详细信息。[立即创建免费个人资料。](https://golangweekly.com/link/147354/web).

---
  
[我们是如何将 oapi-codegen 的依赖开销减少约 84%](https://golangweekly.com/link/147355/web "www.jvt.me") -- 如何通过利用模块修剪来减少模块依赖项的大小。_（有关实际发布的更多信息，请参见下文。）_

_Jamie Tanna_ 


[探索 Go 中的列压缩挑战](https://golangweekly.com/link/147356/web)   

_Jamie Brandon_  

  
---  

## 🛠 代码和工具


[oapi-codegen v2：OpenAPI 客户端和服务器代码生成器](https://golangweekly.com/link/147357/web "github.com") -- 一组流行的实用程序，用于为基于 OpenAPI 3.0 API 定义的服务生成 Go 样板代码，现在使用的依赖[比以前少得多](https://golangweekly.com/link/147355/web) 。如果您想了解如何使用它，[请从这里开始。](https://golangweekly.com/link/147358/web)

_DeepMap, Inc_


[Go v9.3 的 Redis 客户端 – 现在支持 JSON ](https://golangweekly.com/link/147359/web "github.com") -- 官方 Redis 客户端向前迈出了一大步，直接支持 JSON（使用 [RedisJSON 模块](https://golangweekly.com/link/147360/web)时 Redis 直接支持）。

_Redis Team_
  

[免费课程：时态云简介](https://golangweekly.com/link/147361/web "t.mp") -- 查看我们关于时态云的角色和功能的新课程，该课程专为新用户和有经验的用户而设计。

_Temporal Technologies sponsor_
  

[DoltgreSQL：带版本控制功能，兼容 Postgres 的数据库](https://golangweekly.com/link/147362/web "www.dolthub.com") -- [Dolt](https://golangweekly.com/link/147363/web) 是一个历史悠久的由 Go 支持的数据库，提供一种“数据领域的`git`”风格的体验，让你可以使用分支、克隆、合并等，它使用 MySQL 有线协议。DoltgreSQL 提供相同的功能，但它模拟 _Postgres_ 服务器。

_Daylon Wilkins_ 


[env 10.0:：将环境变量解析为结构体](https://golangweekly.com/link/147364/web "github.com") -- 简单，无依赖性，支持所有内置类型以及 `time.Duration`,`encoding.TextUnmarshaler` 和 `url.URL` - 您也可以为您想要支持的任何其他类型定义自定义解析器函数。

_Carlos Alexandro Becker_ 


[GoLeak 1.3：Goroutine 泄漏检测器](https://golangweekly.com/link/147365/web "github.com") -- 用它来检测测试结束时是否有任何意外的 goroutine 还在运行。

_Uber Engineering_ 


[Cobra 1.8](https://golangweekly.com/link/147366/web) - 用于构建全功能 CLI 应用程序的库。

[color 1.16](https://golangweekly.com/link/147367/web) - 使用 ANSI 码的彩色输出。

[GoReleaser 1.22.0](https://golangweekly.com/link/147368/web) - 为多个平台构建和发布 Go 二进制文件。

[validator 10.16](https://golangweekly.com/link/147369/web) - 基于标签的结构和字段的值验证。

[progressbar 3.14](https://golangweekly.com/link/147370/web) - 基础线程安全进度条。

[goverter 1.1](https://golangweekly.com/link/147371/web) - 轻松创建类型安全的转换器。

  
---  

> _💬 引用_
> 
> “最好的程序是在程序员应该做其他事情的时候编写的。”（"The best programs are the ones written when the programmer is supposed to be working on something else"）
> ___ 
> Melinda Varian
