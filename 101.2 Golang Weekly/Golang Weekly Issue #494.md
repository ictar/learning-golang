原文：[Golang Weekly Issue #494](https://golangweekly.com/issues/494)

---

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/n0h8x9mkafk85pkagqlj.jpg)](https://golangweekly.com/link/150854/web)  


[在 Dolt 的 SQL 基准上测试配置文件引导优化](https://golangweekly.com/link/150854/web "www.dolthub.com") —— Go 1.20 中引入了[配置文件引导优化（Profile-guided optimization）](https://golangweekly.com/link/150855/web)。 Dolters（？）创建了一个基准测试，并在不更改任何代码的情况下看到了针对读取和写入延迟的微小（尽管不是微不足道的）改进。如果您的应用程序能够从一些边际性能改进中获益，从而换取一点实验时间，那么几乎肯定是值得尝试的。

_Zach Musgrave (DoltHub)_ 

[根据配置文件引导优化获得最佳效果](https://golangweekly.com/link/150856/web "andrewwphillips.github.io") —— 与上述内容相关，PGO 是 Go 1.21 的一个漂亮补充（在 1.20 中首次可见），但如果您想充分利用它，那么这里有一些东西值得一看。

_Andrew Phillips_ 

[![](https://copm.s3.amazonaws.com/ccc372d1.png)](https://golangweekly.com/link/150853/web) 

[❤️ Postgres](https://golangweekly.com/link/150853/web "www.crunchydata.com") —— 您需要一个像您一样热爱 Postgres 的数据库提供商。我们将处理所有麻烦 - 监控、备份、HA、灾难恢复，因此您无需担心。想要惊人的支持吗？只要您有疑问，我们就会在。

_Crunchy Bridge sponsor_


[Go 的联合创始人 Rob Pike：《Go 的是与非》](https://golangweekly.com/link/150857/web "thenewstack.io") —— 对去年年底 Rob 在 GopherCon AU 上的演讲进行的高水平、新闻性的审视——如果您不想 [▶️ 观看演讲](https://golangweekly.com/link/150858/web)本身，那么值得一读。他介绍了 Go 中他认为正确的部分、遗漏的部分，以及对 Go 吉祥物如何使用的一些疑虑（最初是由 Rob 的妻子 [Renee French](https://golangweekly.com/link/150859/web) 创建的）。

_David Cassel (The New Stack)_ 


> 💡 在 [Rob 自己的演讲文章](https://golangweekly.com/link/150860/web)中有更多详细信息，我们在几期之前对此进行了专题报道。T  
  

_快速了解：_

  * 🚨 最新的[官方 Go 开发者调查](https://golangweekly.com/link/150861/web) 还有五天时间就关闭啦 —— 请立即发表您的看法。

  * 📅 [GopherCon 2024 CFP](https://golangweekly.com/link/150862/web) 开放至 2 月 26 日。

  * 🔐 密码学专家、前 Go 团队成员 Filippo Valsorda 分享了[为 Go 创建一个名为 `mlkem768` 的后量子密码学库背后的细节](https://golangweekly.com/link/150863/web)。

  * 🎧 Dominic St-Pierre 的 [go podcast()](https://golangweekly.com/link/150864/web) 已经 _上_ 线一年了，他决定通过邀请嘉宾来扩展其简短的风格，如本周的剧集：[▶️ 与 Matt Boyle 一起在 Go 中进行调试。](https://golangweekly.com/link/150865/web)

  * 在 Go 播客领域的其他地方，Neil S Primmer 和 Benji Vesterby 继续在 _Go Time_ 中与主持人 Angelica Hill 一起[▶️ 谈论去年的 GopherCon“夺旗”活动](https://golangweekly.com/link/150866/web)。
  

[Go 中的依赖管理简史](https://golangweekly.com/link/150867/web "www.bytesizego.com") —— 快速浏览依赖管理的六个时期，从 _无_ 到有（Go 模块）。如果你早期没用过，那么你可能不知道这一切。

_Matt Boyle_ 
  

[Go 的上下文控制](https://golangweekly.com/link/150868/web "zenhorace.dev") —— 看看 Go 中处理上下文时一些容易违反的规则。演示的修复将帮助您识别自己代码中的违规行为，并且作为福利，提醒您，_聪明_ 很少会 _更好_。

_Horace_ 

[Saga 模式让事情变得简单](https://golangweekly.com/link/150869/web "pages.temporal.io") —— Sagas 是一种常见的开发蓝图，但可能很难构建、测试和维护它们 - 了解我们可以如何提供帮助。

_Temporal Technologies sponsor_


[“我要避免将 `any` 用作实际类型”](https://golangweekly.com/link/150870/web "utcc.utoronto.ca") —— 泛型引入了 `any` 类型，从概念上讲，该类型类似于`interface{}`，但将 `any` 用作泛型外部的类型来代替 `interface{}` 意味着一些不同的东西。

_Chris Siebenmann_ 


[通过 LangChainGo，在 Go 中使用 Gemini 模型](https://golangweekly.com/link/150871/web "eli.thegreenplace.net") —— 在本文中， [_Gemini_](https://golangweekly.com/link/150872/web)指的是 Google 的多模式 AI 模型，您可以通过 Google Cloud 使用该模型。

_Eli Bendersky_ 


[类型断言与类型开关](https://golangweekly.com/link/150873/web) —— 一个小 _备忘录_。   
_Redowan Delowar_  


🛠 代码和工具

[Golte：在 Go HTTP Handler 中渲染 Svelte 组件](https://golangweekly.com/link/150874/web "github.com") —— 一个可以与您选择的路由器配合使用的库，其中布局可以被视为中间件，页面可以被视为处理程序。

_Nicholas Thai_ 


[Goldmark 1.7：用 Go 写的 Markdown 解析器](https://golangweekly.com/link/150875/web "github.com") —— 纯 Go，易于扩展，并且兼容 [CommonMark](https://golangweekly.com/link/150876/web)（GitHub 风味的 Markdown 就是基于此）。也是一个由 WebAssembly 驱动的[Goldmark 游乐场](https://golangweekly.com/link/150877/web)。

_Yusuke Inuzuka_ 

    
---  

📰 

🪐 [使用受 Zanzibar 启发的 SpiceDB 构建应用权限：](https://golangweekly.com/link/150878/web) 可调一致性、动态策略评估、强大的可观察性[等等](https://golangweekly.com/link/150878/web)。

📢 100% PostgreSQL，分布在 3 个区域，在完全托管的云中提供多主、基于延迟的 DNS 路由 - [无需注册](https://golangweekly.com/link/150879/web)。

---

[Gofeed：用于 RSS、Atom 和 JSON Feed 的解析器](https://golangweekly.com/link/150880/web "github.com") —— 一个用于解析 RSS、Atom 和 _JSON Feed_ 的成熟、高度依赖且强大的选项。

_mmcdole_ 


[IntegreSQL：管理隔离的 Postgres 数据库以进行测试](https://golangweekly.com/link/150881/web "github.com") —— 提供 RESTful JSON API，用于管理 Postgres 模板并启动（和管理数据库池）数据库以进行集成测试

_all about apps GmbH_ 


[Inbucket：带 SMTP 和 POP3 的一次性 Web 邮件服务器](https://golangweekly.com/link/150882/web "inbucket.org") —— 用于测试电子邮件的自托管服务 — 它接受来自任何地址的邮件，并可通过 Web、REST 和 POP3 接口使用它。

_James Hillyerd_ 


🧊 [Tetra3D 0.15.0](https://golangweekly.com/link/150884/web) - 用于基于 [Ebitengine](https://golangweekly.com/link/150885/web) 构建的游戏的 3D 渲染器。

📧 [Listmonk 3.0](https://golangweekly.com/link/150886/web) - 自托管电子邮件通讯系统 （[主页](https://golangweekly.com/link/150887/web)）。

[go-i18n 2.4](https://golangweekly.com/link/150888/web) - 将程序翻译成多种（自然）语言。

[Wish 1.3](https://golangweekly.com/link/150889/web) - Charm 的工具，用于制作 Go 支持的 SSH 应用程序。

[GoReleaser 1.24](https://golangweekly.com/link/150890/web) - 为多个平台构建/发布二进制文件。

[FrankenPHP 1.1](https://golangweekly.com/link/150891/web) - Go 支持的 PHP 现代应用服务器。

[Bubbles 0.18](https://golangweekly.com/link/150892/web) - 用于 [Bubble Tea](https://golangweekly.com/link/150893/web) 的 TUI 组件。

[Carbon 2.3.8](https://golangweekly.com/link/150894/web) - 友好的时间操作库。

[Pet 0.6](https://golangweekly.com/link/150895/web) - 简单的基于 CLI 的代码片段管理器。

  
---  


🀄️ 模式识别……  
  
[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/vwqw1owlyc9qr3e3a0sa.jpg)](https://golangweekly.com/link/150896/web)  

[TileEx：一个平铺模式（Tile Pattern）提取器](https://golangweekly.com/link/150896/web "github.com") —— 一个有趣的 Go 小项目，可以接收包含模式的图像，然后从中提取实际的重复平铺元素。代码也少得出奇。

_Sarthak Shah_
