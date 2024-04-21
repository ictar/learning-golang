原文：[Golang Weekly Issue #503](https://golangweekly.com/issues/503)

---

如果你想知道为什么上周没有，那是因为我们在复活节（较晚）休息了一个星期，但现在我们回来啦 :-)

__  
 _您的编辑，Peter Cooper_  
  

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/fl3amghpvdtupcwl8wlf.jpg)](https://golangweekly.com/link/153796/web)  


[从 1.0 版本到 1.22 的 Go 性能](https://golangweekly.com/link/153796/web "benhoyt.com") —— 两年前，我们很喜欢作者对[从 Go 1.2 到 1.18 的分析。](https://golangweekly.com/link/153797/web)这是一次更新，涵盖了从 Go 1.0 到今天的 1.22 版本。

_Ben Hoyt_ 


📊 [2024 年第一季度 Go 开发者调查结果](https://golangweekly.com/link/153798/web "go.dev") —— 开发者情绪和信任度仍然很高，以 AI 为中心的调查显示，Gophers 认为 Go 是适合这一波新用例的良好语言。然而，新的库和工具还有更多的空间可以让这一切变得更容易，特别需要与流行的 Python 库等价的 Go 库。

_The Go Team_ 


[![](https://copm.s3.amazonaws.com/2f2027c5.png)](https://golangweekly.com/link/153795/web) 

[WorkOS：现代 SaaS 应用的企业级身份验证](https://golangweekly.com/link/153795/web "workos.com") —— WorkOS 支持基本身份验证和 SSO 等复杂的企业功能。它提供灵活且易于使用的 API，帮助 Vercel、Loom 和 Webflow 等公司实现 Enterprise Ready。最重要的是，WorkOS 用户管理免费支持多达 100 万个 MAU。

_WorkOS sponsor_


[Go _是_ 一种面向对象的编程语言吗？](https://golangweekly.com/link/153799/web "blog.gypsydave5.com") —— 这是一个有争议且经常被讨论的话题（！），但在您将其视为纯粹的标题党之前，请阅读 David 的理由，了解他为什么认为 _“Go 是一种非常面向对象的编程语言。我认为比 Java 更要面向对象……”_

_David Wickes_ 


_快速了解：_

  * 🤖 一个奇怪的实验：[让大语言模型为空的 Go 函数写测试](https://golangweekly.com/link/153800/web)。

  * 📘  Miki Tebeka 所著的《[Effective Go Recipes](https://golangweekly.com/link/153801/web)》现已出版，它由 _Pragmatic Bookshelf_ 出版。

  * [Go 1.22.2](https://golangweekly.com/link/153802/web) 在我们发出最后一期后不久发布了，尽管改动不大，只包含一般错误修复和针对 `net/http` 的一个安全修复。

  * [GoLand 2024.1](https://golangweekly.com/link/153803/web)，JetBrains 的热门商用 Go IDE 也已经发布了。

  * 无聊？那[使用 Fyne 和 Go 实现的扫雷？](https://golangweekly.com/link/153804/web)这个游戏咋样？现在，你也可以通过 `go run github.com/mevdschee/fyne-mines@latest` 来玩一玩。


[Microsoft 推出了一个 Go 博客](https://golangweekly.com/link/153805/web "devblogs.microsoft.com") —— 更多 Go 博客为我们提供了更多可链接的内容，因此让我们祝贺 Microsoft 😆 吧。它将以 Go 与 Azure 的使用，以及 Microsoft 对 Go 生态系统的总体贡献为导向。

_George Adams (Microsoft)_ 


[Go 的版本控制工作流程](https://golangweekly.com/link/153806/web "t.mp") —— 了解何时以及如何应用版本控制，然后在这个免费的实践培训课程中练习使用我们的 Go SDK。


_Temporal Technologies sponsor_


📄 [将 C 库调用包装在防御性 Go 例程中](https://golangweekly.com/link/153807/web) _christoofar_

📄 [Go 的 I/O 基础：第二部分](https://golangweekly.com/link/153808/web) _Andrei Boar_

📄 [一个伟大的技术博客的要素是什么](https://golangweekly.com/link/153809/web) _Phil Eaton_


🛠 代码和工具  

**编者注：** 今天本节满是数据库！ 😅


[go-mysql-server：一个与存储无关的兼容 MySQL 的关系数据库](https://golangweekly.com/link/153810/web "github.com") —— 一个纯 Go SQL 引擎和服务器，使用 MySQL 协议和 SQL 方言。您将获得一个内存存储系统，但随后您可以实现自己的后端来存储或查询不同的数据源。

_DoltHub_ 


[trdsql：对基于文本的数据进行 SQL 查询的工具](https://golangweekly.com/link/153811/web "github.com") —— 一种 CLI 工具，可以对 CSV、LTSV、JSON、YAML 和 TBLN 文件执行 SQL 查询（采用 Postgres 或 MySQL 语法）。它还可以作为 Go 应用程序中的[库](https://golangweekly.com/link/153812/web)。

_Noboru Saito_ 
  

[Redka：利用 SQLite 重新实现 Redis](https://golangweekly.com/link/153813/web "github.com") —— 您很少会看到一个 Go 项目在短短几天内就能获得如此多的 GitHub star，但在最新的 Redis 许可失败之后，人们对替代方法产生了兴趣。

_Anton Zhiyanov_ 


---

📰 分类广告


[GoLand 2024.1 已发布！](https://golangweekly.com/link/153814/web) 借助免费的本地 AI 全线完成、性能和用户体验改进。[一探究竟！](https://golangweekly.com/link/153814/web)

🐘 虽然我们正以数据库为主题，但如果您是 Postgres 用户，那么请看我们的姐妹通讯 [Postgres Weekly](https://golangweekly.com/link/153815/web)。这是[最新一期。](https://golangweekly.com/link/153816/web)

---

[Huma 2.13.0：构建由 OpenAPI 和 JSON Schema 支持的 API](https://golangweekly.com/link/153817/web "huma.rocks") —— Huma 是一个微框架，用于利用 OpenAPI 等通用标准创建 HTTP REST 或 RPC API。最新版本不再需要除 Go 标准库之外的任何硬依赖项。

_Daniel Taylor_ 


⚙️ [Gorabbit](https://golangweekly.com/link/153818/web) - 一个官方 Go RabbitMQ 插件的高级包装器。_Kardinal_

⚙️ [mircroq](https://golangweekly.com/link/153819/web) - 一个最小的基于 Websocket 的事件代理。 _Chaitanya Munukutla_

⚙️ [go-generics-cache](https://golangweekly.com/link/153820/web) - 使用泛型的内存中键/值存储/缓存。_Kei Kamikawa_

[go-xmpp 0.2](https://golangweekly.com/link/153821/web) - 用于 XMPP 消息传递标准的 Go 库。

[Decimal 1.4](https://golangweekly.com/link/153822/web) - 任意精度的定点十进制数。

[Wish 1.4](https://golangweekly.com/link/153823/web) - Charm 的工具，用于制作 Go 支持的 SSH 应用程序。

[Echo 4.12](https://golangweekly.com/link/153824/web) - 高性能的极简 Go web 框架

[Otto 0.4](https://golangweekly.com/link/153825/web) -  Go 驱动的 JavaScript 解释器。

[VHS 0.7.2](https://golangweekly.com/link/153826/web) - CLI“家庭录像机”。_“将终端 GIF 编写为集成测试和演示 CLI 工具的代码。”"_

[Ebiten 2.7.1](https://golangweekly.com/link/153827/web) - 2D 游戏引擎。

