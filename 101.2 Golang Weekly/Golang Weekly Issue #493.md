原文：[Golang Weekly Issue #493](https://golangweekly.com/issues/493)

---


[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/bga4uwbeurqonvvpvpk2.jpg)](https://golangweekly.com/link/150549/web)  


[最新的 Go 开发者调查现已开放](https://golangweekly.com/link/150549/web "go.dev") —— 2024 年第一个官方 Go 开发者调查已经发布（您可以[在此处获取](https://golangweekly.com/link/150550/web)），Go 团队已准备好接受您的反馈了。截止日期是2月11日，所以不要等太久。我们分享的事情最终 _确实_ 会影响 Go 的未来，所以这是值得做的。如果你想了解它的变化趋势，这里是上次调查的[结果](https://golangweekly.com/link/150551/web)。

_The Go Team_ 

[Go 1.22 候选版本 2 已发布](https://golangweekly.com/link/150552/web "groups.google.com") —— 第二个 RC 版本已经发布，这提醒了我们距离 Go 1.22 最终版本只有几周的时间了。[发行说明草案](https://golangweekly.com/link/150553/web)已进行了一些更新，如果您想了解更多内容，可以很好地看看有关 1.22 版本所获得的改进的介绍。

_Cherry and Carlos (Go Team)_ 


[![](https://copm.s3.amazonaws.com/9bd98a7f.png)](https://golangweekly.com/link/150548/web) 

[帮助在希腊雅典举办的 GopherConEU 成为一场令人难忘的活动吧！](https://golangweekly.com/link/150548/web "events.ardanlabs.com") —— 参加 2024 年 2 月 6 日至 8 日在希腊雅典举行的 GopherConEU。Bill Kennedy 和 Miki Tebeka 参加了 GopherConEU 培训研讨会，分享前沿的 Go 工程最佳实践、库、框架、性能优化等。

_Ardan Labs & GopherConEU sponsor_


[Go 1.22 中的新增功能：`slices.Concat`](https://golangweekly.com/link/150554/web "blog.carlana.net") —— `slices.Concat` 的实现者讨论了它的实现以及它是如何解决使用切片时的常见性能问题的。如果您在循环中操作切片，那么就应该读读本文。

_Carlana Johnson_ 


_快速了解：_

  * 🗣 Reddit 上出现了一个关于 [Go 开发者在日常工作中实际做什么](https://golangweekly.com/link/150555/web)的长帖子。

  * Slack 的 Claire Adams 撰写了[有关 Slack 如何大规模运行 cron 作业](https://golangweekly.com/link/150556/web)的文章。所有主要部分都是用 Go 编写的。

  * 🎸 Mat Ryer 发送的是 [▶️ 这个音乐提醒](https://golangweekly.com/link/150557/web)，其中，`filepath` 的 [`Walk` 函数](https://golangweekly.com/link/150558/web) _不_ 遵循符号链接。我当然永远不会忘记…… 😆

  * Datadog 的 Felix Geisendorfer [r揭开了一些 Go 内存指标的神秘面纱](https://golangweekly.com/link/150559/web) - 当然，他们的产品可以帮助您可视化这些指标，但这仍然是一篇有用的文章。


▶ [Go 中 JSON 的未来](https://golangweekly.com/link/150560/web "www.youtube.com") —— Tailscale 的 Joe Tsai 介绍了 Go 中 JSON 的解析和生成，展示了当前方法的问题并深入研究了 [`encoding/json/v2`](https://golangweekly.com/link/150561/web)。

_Joe Tsai_ 
  

[使用 elem-go 和 goldmark 构建 Go 静态站点生成器](https://golangweekly.com/link/150562/web "dev.to") —— 您 _可以_ 使用一些库来滚动您自己的静态站点生成器而不是使用 [Hugo](https://golangweekly.com/link/150563/web) 之类的东西。还有[一个模板应用程序](https://golangweekly.com/link/150564/web)可以让你快速入门。

_Chase Fleming_ 

🛠 代码和工具
  
[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/zr20llcgzhcd0gfcjvjf.jpg)](https://golangweekly.com/link/150565/web)  


[Memos：隐私第一的 Go 笔记应用](https://golangweekly.com/link/150565/web "www.usememos.com") —— 一个完整的、MIT许可的笔记 Web 应用程序，您可以自行托管。它使用 [Echo](https://golangweekly.com/link/150566/web) 作为框架，并使用 SQLite 存储数据。[GitHub 存储库。](https://golangweekly.com/link/150567/web)

_Memos Contributors_ 


[gofumpt 0.6：更严格的 `gofmt`](https://golangweekly.com/link/150568/web "github.com") —— —如果您喜欢规则但 但你对来说 `gofmt` 不够严格，那么，`gofumpt` 则有[更严格的规则](https://golangweekly.com/link/150569/web)来保持代码库干净且受控。 v0.6 基于 Go 1.21 的 `gofmt`，并且需要 Go 1.20+ 版本

_Daniel Marti_ 
  

[尝试一款专为您打造的闪电般快速的结对工具](https://golangweekly.com/link/150570/web) —— Tuple 将改变您对结对编程的思考方式。与您的团队一起免费试用 14 天，无需卡片。

_Tuple sponsor_


[goja：纯用 Go 写的 ECMAScript/JavaScript 引擎](https://golangweekly.com/link/150571/web "github.com") —— 想要将基于 JS 的脚本功能添加到您的 Go 应用程序中吗？它为您提供了一个不涉及引入外部引擎的选项。_有关如何使用它的示例，请查看下一项_

_Dmitry Panov_ 

[k6：一个用 Go 写的负载测试库](https://golangweekly.com/link/150572/web "github.com") —— 一个功能齐全、可配置的负载生成工具，使用 [goja](https://golangweekly.com/link/150571/web) 引擎（上面），支持用 JavaScript 编写测试脚本。_（AGPL 许可。）_

_Grafana Labs_ 


[Grape：一个新的零依赖 Go HTTP 库](https://golangweekly.com/link/150573/web "github.com") —— _" A thin wrapper around the standard library, providing helper functions to facilitate faster and easier development. Adding only a single dependency to your projects."_

_Hossein Yazdani_ 
 
---
📰 分类广告


🪐 使用 Go 和开源构建您的权限系统。[SpiceDB 是 Google Zanzibar 的开源实现](https://golangweekly.com/link/150574/web)。

---

🧊 [tinygo-wasm-webgl-demo](https://golangweekly.com/link/150575/web "github.com") —— 一个小型存储库，展示如何使用 TinyGo、WebAssembly 和 WebGL 创建在浏览器中运行的简单 3D 图形演示。

_Andriy Semenets_ 


[oapi-codegen 2.1](https://golangweekly.com/link/150576/web) - 根据 OpenAPI 规范生成样板代码。

[gRPC-Go 1.61](https://golangweekly.com/link/150577/web) - gRPC 的 Go 实现，用于基于 HTTP/2 的 RPC。


[Roaring 1.9](https://golangweekly.com/link/150578/web) -  Roaring 位图数据结构实现。

[Expr 1.16](https://golangweekly.com/link/150579/web) - 表达式语言和表达式求值。 

[uuid 1.6](https://golangweekly.com/link/150580/web) - Google 的 RFC 4122 UUID 包。

[Lego 4.15](https://golangweekly.com/link/150581/web) - Let's Encrypt/ACME 客户端和库。

[BlueRPC](https://golangweekly.com/link/150582/web) - 维护 Go 和 TypeScript 之间的类型安全。