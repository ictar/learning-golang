原文：[Golang Weekly Issue #487](https://golangweekly.com/issues/487)

---


[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/plvs5sdayu8acnhuz6ag.jpg)](https://golangweekly.com/link/148580/web)  


[errtrace：错误堆栈跟踪的替代方案](https://golangweekly.com/link/148580/web "github.com") —— 受 Zig 启发，errtrace 是一个 _实验性_ 包，通过 Go 程序跟踪错误的返回路径，即错误到达用户的代码路径，这可能和 _导致_ 错误的代码路径大不同。是不同的到导致错误的代码路径。幸运的是，如果您感到困惑，该存储库提供了[一个比较](https://golangweekly.com/link/148581/web)。

_Braces_ 


▶ [Russ Cox 谈谷歌的开源供应链安全](https://golangweekly.com/link/148582/web "research.swtch.com") —— 上周，Go 的 Russ Cox 在 ACM 软件供应链会议上发表了远程主题演讲。这场演讲不仅仅涉及 Go，还它自然地构成了 Russ 的许多示例的基础。_（45分钟。）_

_Russ Cox_ 

[![](https://copm.s3.amazonaws.com/5ddc7bfc.jpg)](https://golangweekly.com/link/148579/web) 

[仅需两行代码，即可将功能齐全的身份验证添加到你的 Go 应用中](https://golangweekly.com/link/148579/web) —— 简单、安全且经济实惠：皆可兼得。FusionAuth 是由开发人员为开发人员构建的身份验证。

_FusionAuth sponsor_

[通过实际例子，了解 Go 字符串操作优化](https://golangweekly.com/link/148583/web "medium.com") —— 又到了[_代码显现_](https://golangweekly.com/link/148584/web)之时，因此字符串争论变得至关重要：_“我将向你展示我是如何使用一个非常简单的程序，将其运行速度提高了近 5 倍，仅需很少的调整。”_

_Alex Bledea_ 

[将 Go 应用变成“静态” WASM 支持的站点](https://golangweekly.com/link/148585/web "skarlso.github.io") —— 作者用 Go 构建了一个工具，用于将 CRD（Kubernetes 中的自定义资源）转换为 YAML，并希望提供这个工具的在线版本。将其全部编译成一个由 WebAssembly 驱动的应用程序，该应用程序可以以静态方式托管，同时提供动态功能（不依赖于服务器），这并不容易，但它确实能用。

_Gergely Brautigam_ 
 

_快速了解：_
* 📘 [_Learn Concurrent Programming with Go_](https://golangweekly.com/link/148587/web) 是一本由 James Cutajar 撰写、Manning 出版的新书 - 印刷版现已可从出版商处获得，电子书版本和更广泛的发现要等到 1 月份。

* 🌉 _The New Stack_ 报告 [OpenTelemetry 的贡献者正在为 Go 开发一座桥梁](https://golangweekly.com/link/148588/web).

* 💣 是否想在 Windows 98 上用附带的游戏消磨时间？现有一个用 Go 编写的基于终端的[_扫雷_ 游戏实现](https://golangweekly.com/link/148589/web) in Go.

* 🎤 _Go Time_ 播客[▶️ 涵盖了事件驱动系统和架构](https://golangweekly.com/link/148586/web)。


[Ebitengine 的 2023 年回顾](https://golangweekly.com/link/148590/web "ebitengine.org") —— Ebitengine（前身是 Ebiten）是一种流行的 Go 2D 游戏引擎，可让您跨多种平台（包括移动设备和 Nintendo Switch）构建游戏体验。这篇文章通过对今年该项目发生的情况按月进行总结来庆祝该项目十周年，并分享了一些已发布的由 Ebitengine 驱动的游戏的视频。

_Hajime Hoshi_ 


[为什么 Go 堆那么复杂](https://golangweekly.com/link/148591/web "www.dolthub.com") —— 为什么经验丰富的 Gopher 会觉得标准库（`container/heap`）的堆令人困惑，以及社区是怎样想出各种更简单的堆实现的（当然，其中一个是基于泛型的）。注意：还有[一个长期存在的问题建议](https://golangweekly.com/link/148592/web)，它提出将切片支持的堆添加到 `container/heap`。

_Max Hoffman (DoltHub)_ 

  

▶ [用 NATS JetStream 取代 Kafka、RabbitMQ、Redis 等等？](https://golangweekly.com/link/148593/web "www.youtube.com") —— 一个简单介绍，涉及 _NATS JetStream_（NATS 分布式应用消息传递系统的子系统），以及如何使用其新的客户端 API，在 Go 中创建发布者和消费者应用。

_Jeremy Saenz_ 


[通过 Temporal Schedules 增强工作流程控制和可观察性](https://golangweekly.com/link/148594/web "t.mp") —— 现已推出：Temporal Schedules 提供了一种更可靠、更可扩展的解决方案，这正是传统 Cron 作业的短板。

_Temporal Technologies sponsor_


[使用 Go 和 Leaflet 可视化地图数据](https://golangweekly.com/link/148595/web "www.ardanlabs.com") —— 一种解析 GPX 文件（Strava 等应用程序使用和导出的格式）并将其绘制到基于 Web 的地图上的方法。

_Miki Tebeka (Ardan Labs)_ 

## 🛠 代码和工具 

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/ymz0im8zvrssta3bljit.jpg)](https://golangweekly.com/link/148596/web)  


[Sloc、Cloc 和 Code (scc) 3.2：快速准确的“代码计数器”](https://golangweekly.com/link/148596/web "github.com") —— 对于多种语言，它都可以计算代码行数并估算代码复杂度。

_Ben Boyter_ 


[golang-set：简单、经过充分测试的通用集合类型](https://golangweekly.com/link/148598/web "github.com") —— Go 的通用集合类型集合：_“除非 Go 内置了集合……否则使用这个。”_

_Ralph Caraveo_ 


[Gabs：在 Go 中解析、创建和编辑“Unknown”或者动态 JSON](https://golangweekly.com/link/148599/web "github.com") —— 一个封装器，用来导航 `encoding/json` 提供的 `map[string]interface{}` 对象的层次结构。

_Ashley Jeffs_ 


---  
📰 分类广告

💻 Hired 让求职变得简单 - 公司不再需要追逐招聘人员，而是预先向您提供薪资详细信息。[立即创建免费个人资料。](https://golangweekly.com/link/148600/web)

💻 在 [Bosch](https://golangweekly.com/link/148601/web)，你总是不断成长。提升自己的技能，适应无数新的角色、职位和机会。[了解更多。](https://golangweekly.com/link/148601/web)

  
---  
  

[Goiabada：一个身份验证和授权服务器](https://golangweekly.com/link/148602/web "goiabada.dev") —— 一个新的用户管理系统，实现 OAuth2 和 OpenID Connect 协议。

_Leonardo D'Ippolito_ 
  

[Gorse：通用开源推荐系统](https://golangweekly.com/link/148603/web "gorse.io") —— Gorse 使用 ML，基于所提供的数据提供一个推荐 RESTful API。它包括仪表板和 API 文档。

_zhenghaoz_

[go-i18n 2.3](https://golangweekly.com/link/148604/web) - 将程序翻译成多种（自然）语言

[GopherLua 1.1.1](https://golangweekly.com/link/148605/web) - Go 驱动的 Lua VM 和编译器。

[go-github 57.0](https://golangweekly.com/link/148606/web) - GitHub v3 API 的客户端库。

[Task 3.32](https://golangweekly.com/link/148607/web) - 任务运行器/用 Go 编写的替代品。

[Garble 0.11](https://golangweekly.com/link/148608/web) - Go 构建混淆。