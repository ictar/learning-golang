原文：[Golang Weekly Issue #490](https://golangweekly.com/issues/490)

---

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/evhh5zncify0tu3yftig.jpg)](https://golangweekly.com/link/149673/web)  


[Rob Pike：“我们做对了什么，我们又做错了什么”](https://golangweekly.com/link/149673/web "commandcenter.blogspot.com") —— Rob 在 11 月份的 GopherConAU 演讲的书面版本（实际上是脚本）（[这里是 ▶️ 47 分钟的视频](https://golangweekly.com/link/149674/web)），他在其中讲述了构建 Go 的故事，包括团队的目标以及进行决策的方式和原因。

_Rob Pike_ 


[Go 1.22 的交互式发行说明](https://golangweekly.com/link/149675/web "antonz.org") —— 不要惊慌，Go 1.22 直到下个月才会发布最终版本，但这是一组简洁的 _交互式_ 说明，展示了 1.22 的各种功能。这非常值得一看以加快速度。

_Anton Zhiyanov_ 


[![](https://copm.s3.amazonaws.com/9bd98a7f.png)](https://golangweekly.com/link/149672/web) 

[帮助希腊雅典的 GopherConEU 成为一场令人难忘的活动！](https://golangweekly.com/link/149672/web "events.ardanlabs.com") —— 参加 2024 年 2 月 6 日至 8 日在希腊雅典举行的 GopherConEU。Bill Kennedy 和 Miki Tebeka 参加了 GopherConEU 培训研讨会，分享前沿的 Go 工程最佳实践、库、框架、性能优化等。

_Ardan Labs & GopherConEU sponsor_

  

[Go 中嵌入的快乐](https://golangweekly.com/link/149676/web "cybernetist.com") —— _“虽然 Python 仍然是 AI/ML 领域的主导语言，但有很多 Go 模块可以让您使用数据和向量嵌入做一些有趣的事情。”_ 这是一篇简洁的文章，其中包含一个很酷的交互式 3D 可视化效果。" 

_Milos Gajdos_ 


📊 [不同 Go SQLite 驱动的基准测试](https://golangweekly.com/link/149677/web "github.com") —— 与任何基准测试结果一样，请保持批判性的眼光，但作者发现这七个选项之间存在显着差异。

_Christoph Vilsmeier_ 


_快速了解：_

  * Go 编译器团队最近一次每周会议的[会议记录](https://golangweekly.com/link/149678/web)，围绕 Go 调查中的一些让与会者感到惊讶的有趣观点，此外还有关于 ARM 中国 Go 团队的传言。

  * 在 Hacker News，有人问 [Go 是新的 BASIC 吗？](https://golangweekly.com/link/149679/web) 不..


  * 提醒您，AWS Lambda 上的 `go1.x`  运行时[现已弃用](https://golangweekly.com/link/149680/web)。


[约束泛型设计中的复杂性](https://golangweekly.com/link/149681/web "blog.merovius.de") —— 一篇基于 GopherConAU 演讲（[此处为 ▶️ 视频](https://golangweekly.com/link/149682/web)）的深入文章，通过一点计算机科学的方式，解释了为什么联合元素不能调用方法作为后备，以及未来可能会出现什么样子的解决方法。

_Axel Wagner_ 


[`CompareAndSwap` 是如何不总是“比较交换”的](https://golangweekly.com/link/149683/web "lu.sagebl.eu") —— 该 `sync/atomic` 包提供了对原始原子操作的访问，这些操作在处理并发性时非常有用（[更多信息请参见此处](https://golangweekly.com/link/149684/web)）。如果架构不支持其中一项功能，会发生什么情况？

_Luciano Calamia_ 


[Uber 的 Go 风格指南](https://golangweekly.com/link/149685/web "github.com") —— 想知道大型组织是如何编写 Go 的吗？了解 Uber 的指南、性能说明和风格。

_Uber_


[学习如何使用 Temporal 和 Go，将 Cron 转换为时间表](https://golangweekly.com/link/149686/web "t.mp") —— 在您选择的界面中将 Cron 作业转换为临时时间表非常简单。请参阅我们的 Go SDK 中的示例。

_Temporal Technologies sponsor_


[关于使用 Go 编写一个简单的 JSON 解析器](https://golangweekly.com/link/149687/web "buildwithgo.substack.com") —— 这实际上是学习之旅中所涉及的一些部分。

_Öner Çiller_ 


[将 Go Lambda 函数与 AWS CDK 捆绑](https://golangweekly.com/link/149688/web)   
_Matt Bacchi_ 

[接口不合适](https://golangweekly.com/link/149689/web)   
_Preslav Rachev_ 


🛠 代码和工具  

[astjson 简介：快速转换和合并 JSON 对象](https://golangweekly.com/link/149690/web "wundergraph.com") —— Wundergraph 的 GraphQL 网关产品需要快速聚合和合并来自多个来源的数据。他们构建了 astjson 来提高性能，并取得了一些显着的成果，甚至超越了其他（基于 Rust 的）工具。

_Jens Neuse (WunderGraph)_ 
  

[Tarmac：语言无关框架](https://golangweekly.com/link/149691/web "tarmac.gitbook.io") —— 冒着听起来像是 _万能_ 框架的风险，Tarmac 采用了一种有趣的无服务器方法来构建应用程序框架。它由 Go 驱动，但使用 WebAssembly 来允许您根据需要使用 Go、Rust、Zig 等语言来编写函数，而 Tarmac 则使它们与现实世界、数据库等保持连接。

_Benjamin Cane_ 


[在你的域上托管 Go 模块的新方法](https://golangweekly.com/link/149692/web "petersanchez.com") —— 介绍  _GoHome_，这是一个小服务器，用来管理在您自己的域上托管 Go 模块。

_Peter Sanchez_


[wazero 1.6](https://golangweekly.com/link/149693/web) - Go 的零依赖 WebAssembly 运行时。一个大版本，现在具有（可选）优化编译器。

[rqlite 8.16.0](https://golangweekly.com/link/149694/web) - SQLite 支持的分布式数据库，现在支持自动备份和自动恢复到 S3 兼容存储。

[golang-set 2.6](https://golangweekly.com/link/149695/web) - 健壮的通用集合类型。

[Mock 0.4](https://golangweekly.com/link/149696/web) - 集成 `testing` 的 Mocking（模拟）。

[Gobot 2.3](https://golangweekly.com/link/149697/web) - 机器人/无人机/物联网框架。

[dae 0.5](https://golangweekly.com/link/149698/web) - 基于 eBPF 的高性能透明代理。

[Expr 1.15.8](https://golangweekly.com/link/149699/web) - 表达式语言和评估。

[Otto 0.3](https://golangweekly.com/link/149700/web) - Go 驱动的 JavaScript 解释器。

[go-redis 9.4](https://golangweekly.com/link/149701/web) - Redis 客户端库。


  
---  


🏓 一场古老的比赛？  
  
[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/v1704795836/ucrxpe09wbm5t9kvmjip.png)](https://golangweekly.com/link/149702/web)  

[一个由 Go 驱动的打砖块克隆版本](https://golangweekly.com/link/149702/web "github.com") ——  对于那些在 80 年代不玩游戏的人来说，可以想象一下蝙蝠和球类游戏 😄 这是用 Go 编写的，使用 [Ebitengine](https://golangweekly.com/link/149703/web)，并采用实体组件系统 (ECS，entity component system) 架构。

