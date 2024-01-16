原文：[Golang Weekly Issue #491](https://golangweekly.com/issues/491)

---

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/wywka9utpyrjxehlzit4.jpg)](https://golangweekly.com/link/149953/web)  
  

[关于 Go 1.22 的初步想法](https://golangweekly.com/link/149953/web "www.dolthub.com") —— 1.22 将于下个月发布，但候选版本让我们看到了后续的大量更改和改进，包括 `for` 循环中循环变量的新默认行为、“rangefunc”实验 _（下一条有更多相关内容）_ ，甚至还有一些性能改进。

_Jason Fulghum (DoltHub)_ 
  

[Rangefunc 实验](https://golangweekly.com/link/149954/web "go.dev") —— Go 团队正在研究在 Go 中添加 range-over 函数迭代器，通过设置特定的环境变量，1.22 版本可以获得一个相关的初步实现。他们希望获得有关其工作方式方面的反馈，而此页面展示了详细信息。_（需要 1.22 RC 或更高版本才能立即尝试此操作。）_ 

_The Go Wiki_ 

  
[![](https://copm.s3.amazonaws.com/6e92e852.png)](https://golangweekly.com/link/149952/web) 

[Redis With Wings](https://golangweekly.com/link/149952/web "www.dragonflydb.io") —— Dragonfly 是一种简单、高性能且经济高效的内存数据存储，非常适合用作高性能 Go 应用程序的缓存或数据库。Dragonfly 完全兼容 Redis API，还没有 Redis 管理的复杂性。

_Dragonfly sponsor_
  

[Go 1.22 新功能：`reflect.TypeFor`](https://golangweekly.com/link/149955/web "blog.carlana.net") —— Carlana 继续我们的 Go 1.22 项目系列，还介绍了她参与的 `reflect.TypeFor`  实现的幕后花絮。

_Carlana Johnson_ 


_快速了解：_

  * [Go 1.21.6 和 1.20.13](https://golangweekly.com/link/149956/web) 已发布。

  * Anton Medvedev [回顾了五年以来](https://golangweekly.com/link/149957/web)在出色的 [Expr](https://golangweekly.com/link/149958/web) 表达式语言和表达式评估器方面的工作。

  * Alex Ellis 尝试[将 GitHub Actions 作为“分时超级计算机”](https://golangweekly.com/link/149959/web). 很有趣，当然，使用 Go。
  

[如何表示一个可能不存在、可能是 `null` 或者可能有值的 JSON 字段？](https://golangweekly.com/link/149960/web "www.jvt.me") —— 为什么在使用 `encoding/json` 时很难确定某个字段是否已发送或者是否显式为空？还提供了[一个新的库](https://golangweekly.com/link/149961/web)来帮助解决该问题。

_Jamie Tanna_ 


[在 Go 中使用 React 和 templ](https://golangweekly.com/link/149962/web "templ.guide") —— 一个教程，有关如何将 React 丰富的组件生态系统整合为嵌入到 Go 和 [templ](https://golangweekly.com/link/149963/web)服务器端渲染功能中的“交互岛”。

_Adrian Hesketh_ 


[在 Go 中“使用 Google 登录（Sign in with Google）”](https://golangweekly.com/link/149964/web "eli.thegreenplace.net") —— 快速了解如何为您的应用实现“使用 Google 登录”身份验证方法。

_Eli Bendersky_ 


▶ [Go 中的内存管理：好的、坏的和丑的](https://golangweekly.com/link/149965/web) —— 保持相当高的层次和可访问性。   
_Liam Hampton_  

[Go 的测试替代示例](https://golangweekly.com/link/149966/web)   
_Amin Rashidbeigi_  


🛠 代码和工具

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/ric7ksf0yj5mknzn72oz.jpg)](https://golangweekly.com/link/149967/web)  


[Gotraceui：Go 执行跟踪的 GUI 前端](https://golangweekly.com/link/149967/web "gotraceui.dev") —— 旨在成为比 `go tool trace` 更快、更易于访问且更强大的替代方案，提供了一个 [Gio UI](https://golangweekly.com/link/149968/web) 驱动的界面，专门针对 Go 跟踪的独特特征进行了调整。现在 [GitHub 存储库](https://golangweekly.com/link/149969/web)添加了火焰图 🔥。

_Dominik Honnef_ 


[尝试一款专为您打造的如闪电般迅速的结对工具](https://golangweekly.com/link/149970/web) —— Tuple 将改变您对结对编程的看法。与您的团队一起免费试用 14 天，无需卡片。

_Tuple sponsor_


[Avo 0.6：使用 Go 生成 x86 汇编](https://golangweekly.com/link/149971/web "github.com") —— 通过 Go 控制结构和虚拟寄存器，使汇编编写更容易。[两个数字相加的示例](https://golangweekly.com/link/149972/web)以简单的方式展示了它是如何工作的。[v0.6](https://golangweekly.com/link/149973/web) 增加了对更多 AVX 指令的支持，还改进了 Go 1.20+ 支持。

_Michael McLoughlin_ 


[Algnhsa 1.1：AWS Lambda `net/http` 服务器适配器](https://golangweekly.com/link/149974/web "github.com") —— 使用 API Gateway 或 ALB，在 AWS Lambda 上以无服务器的方式运行 Go web 应用程序，而无需更改现有的 HTTP 处理程序。

_Artem Krylysov_ 
  

[Harmony：使用 Discord API 的“和平”模块](https://golangweekly.com/link/149975/web "github.com") —— 当我几年前第一次链接它的时候，我写道 _“[Discord](https://golangweekly.com/link/149976/web) 是一款流行的聊天应用程序，尤其流行于游戏社区。”_ 想象一下现在需要被告知..😆

_Antoine Couchard_ 


---  

📰 分类广告

🪝 使用 [Hookdeck](https://golangweekly.com/link/149977/web) 事件网关接收、验证、转换、过滤、路由和交付 Webhooks 。[尝试接收 webhooks 快速入门](https://golangweekly.com/link/149977/web)。

🗓️ [GopherCon Europe](https://golangweekly.com/link/149978/web) 将于今年二月（下个月！）和六月分别在希腊雅典和德国柏林举行。Ardan Labs 还在希腊的活动中[举办了一些 Go 研讨会](https://golangweekly.com/link/149979/web)。
  
---  


[Steampipe：使用 SQL 查询云服务](https://golangweekly.com/link/149980/web "steampipe.io") —— 使用 SQL 查询甚至 _加入_ 来自 Airtable 、AWS、Heroku、Slack 和 Stripe 等提供商以及 Reddit 和 Hacker News 等社交媒体网站的[100 多个 API](https://golangweekly.com/link/149981/web)。

_Steampipe_ 


[go-github 58.0](https://golangweekly.com/link/149982/web) - GitHub v3 API 的客户端库。

[Permify 0.7](https://golangweekly.com/link/149983/web) - 受 Google Zanzibar 启发的授权。

[JiraCLI 1.5](https://golangweekly.com/link/149984/web) - 适用于 Jira 的功能丰富的交互式 CLI 工具。

[go-sqlite 1.1](https://golangweekly.com/link/149985/web) - SQLite 3 的低级接口。

[Kaf 0.2.7](https://golangweekly.com/link/149986/web) - 适用于 Apache Kafka 的现代 CLI。

[NATS.go 1.32](https://golangweekly.com/link/149987/web) - NATS 的 Go 客户端。
