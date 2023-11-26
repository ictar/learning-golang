原文：[Golang Weekly Issue #485](https://golangweekly.com/issues/485)

---

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/uddrl5pkdig9rvlgwmqy.jpg)](https://golangweekly.com/link/148029/web)  


[River：用于 Go + Postgres 的快速稳健的作业队列](https://golangweekly.com/link/148029/web "brandur.org") —— 一个新的测试版开源作业队列，“_用于构建快速、密封的应用程序_”，用 Go 编写并利用泛型。它还使用 [pgx](https://golangweekly.com/link/148030/web) 来与 Postgres 交互。这篇文章很好地宣传了 River 的优势。

_Brandur Leach_ 

[NilAway：实用的 Nil Panic 检测器](https://golangweekly.com/link/148005/web "www.uber.com") —— Nil Panic 是 Go 中常见且难以检测的问题，但 Uber 创建了一个静态分析工具（基于 Java 世界中的类似工具），它易于设置并与构建工具集成，所以有什么理由不该立刻尝试呢😅_（人们报告该工具可能会返回误报，所以要警惕。）_

_Uber_   


[Ardan Labs 在线培训享六五折](https://golangweekly.com/link/148031/web "www.ardanlabs.com") —— 通过我们的 Go、Rust、Docker 和 Kubernetes 在线高质量培训（享六五折），让您的技能更上一层楼。现在就了解经验丰富的工程专业人士所知道的知识！在销售结束前，结账时使用促销代码 `BLACKFRIDAY2023`。

_Ardan Labs Training sponsor_


[Google 的 Go 可读代码风格指南](https://golangweekly.com/link/148006/web "google.github.io") —— 适用于 Google 的内容不一定适用于其他人，但看看他们对编写惯用的 Go 代码的看法仍然是一件很有意思的事情。这个全面的资源并不意味着是 _绝对的_，而是旨在拥有一个 “Go 代码风格”，以帮助 Google 新的和现有的 Go 开发人员避免常见的陷阱，并保持代码库的一致性。

_Google_ 

[改变我的生产力的十二个 Go 技巧](https://golangweekly.com/link/148007/web "blog.devtrovert.com") —— 一个有趣的想法集锦，但如果您与其他人一起工作，第 6 个和第 9 个可能用起来不会那么顺利，最好避免 — 我们将把它留给读者自己来找出原因。

_Phuong Le_ 


_快速了解：_
  * GitHub 看看 [Git 最新版本的新增功能：Git 2.43。](https://golangweekly.com/link/148008/web)
  * 修复部署在 _Google Cloud Run_ 上[运行的一个缓慢 Go 应用](https://golangweekly.com/link/148009/web)的一个简单故事。
  * 如果您订阅了 ChatGPT Plus，那么看看[Go 专家 Moss](https://golangweekly.com/link/148010/web) —— 据称，这是一个所谓的自定义“GPT” Prompt，可以回答您所有的 Go 问题。_（我们不确定它来自哪里，您的体验可能会有所不同！）_


[用 Go 编写的检索增强生成](https://golangweekly.com/link/148011/web "eli.thegreenplace.net") —— 训练自己的 LLM（大型语言模型）面临着数据过时和成本方面的挑战，但是如果您可以通过按需获取外部数据来推动 LLM 从而获得更实时的答案，且无需进行重新训练，结果又会怎样呢？这就是所谓的“RAG”的作用。[Simon Vellei 还展示了如何使用 Redis 来做到这一点。](https://golangweekly.com/link/148012/web)

_Eli Bendersky_ 

[制作轻量级 Markdown 编辑器](https://golangweekly.com/link/148033/web "www.ersin.nz") —— [Wails](https://golangweekly.com/link/148034/web)为 Go 带来了一种 Electron 式的跨平台桌面应用程序构建方法，因此，如果您不喜欢 Electron（或其他桌面框架）_但_ 喜欢Go，Wails 可能是您的选择。 😅

_Ersin Buckley_ 
 
---

📰 分类广告

🕐 新！在我们即将[于 11 月 30 日举行的网络研讨会](https://golangweekly.com/link/148014/web)上了解 [Temporal Schedules](https://golangweekly.com/link/148013/web) 是如何比 Cron 作业更可靠、可扩展且更灵活的。

💻 Hired 让求职变得简单 - 公司不再需要追逐招聘人员，而是预先向您提供薪资详细信息。[立即创建免费个人资料](https://golangweekly.com/link/148015/web)。
  
---  


## 🛠 代码和工具

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/drtnxtdyr9kmcxhmnbpq.jpg)](https://golangweekly.com/link/148016/web)  

[Ants 2.9：高性能固定容量 Goroutine 池](https://golangweekly.com/link/148016/web "github.com") —— 如果您的程序旨在生成大量的 Goroutine，Ants 可以在限制范围内自动管理和回收它们。

_Andy Pan_
  

[Telebot 3.2：一个电报机器人框架](https://golangweekly.com/link/148017/web "github.com") —— 承诺提供“同类最佳”API，用于处理机器人命令路由、内联查询请求、回调等。该文档有很多简单的示例。v3.2 包括Bot API 6.3-6.5 中新 Telegram 机器人功能的[更新](https://golangweekly.com/link/148018/web)。

_Various Contributors_ 


[Ergo 2.12：现代 IRC 服务器/`ircd`](https://golangweekly.com/link/148019/web "github.com") —— 以前叫做 Oragono（我们于 7 年前首次链接！），这是一个由 Go 驱动的兼容 IRCv3 的聊天服务器，并且仍在不断更新以跟上 IRCv3 规范流程。

_Lingamneni, Oaks, Huber, Latt, et al._ 

  
[pREST 1.4](https://golangweekly.com/link/148020/web) - 想象是 [PostgREST](https://golangweekly.com/link/148021/web)，但是是在 Go 中的。

[Prisma Client Go 0.29.0](https://golangweekly.com/link/148022/web) - Go 的类型安全数据库访问。[Prisma](https://golangweekly.com/link/148023/web) 生态的一部分。

[goconst 1.7](https://golangweekly.com/link/148024/web) -  查找可能是常量的重复字符串。

[Orchestra 0.1](https://golangweekly.com/link/148025/web) - 管理长时间运行的 Go 进程。

[go-sse 0.7](https://golangweekly.com/link/148026/web) - 服务器发送事件（SSE）库。

[goqu 9.19.0](https://golangweekly.com/link/148027/web) - SQL 构建器和查询库。

[bearclaw 2.2](https://golangweekly.com/link/148028/web) - 小型静态站点生成器。