原文：[Golang Weekly Issue #497](https://golangweekly.com/issues/497)

---

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/l7teswvwc8wjknycbimx.jpg)](https://golangweekly.com/link/151790/web)  
  

[关于切片的强大的通用函数](https://golangweekly.com/link/151790/web "go.dev") —— 核心库为切片提供了一些很棒的通用函数，了解每个函数（修改切片、创建新切片等）的底层逻辑对于避免错误至关重要，尤其是 1.22 中对它们进行了一些更改。还建议进行一些进一步阅读。

_Valentin Deleplace (The Go Blog)_ 


[Go 的“枚举（Enum）”糟糕透了](https://golangweekly.com/link/151791/web "www.zarl.dev") —— 虽然标题感觉有点像是骗流量，而且 Go 并没 _有_ 真正的枚举，但这篇文章围绕枚举的用例提出了一些有趣的观点，解决方法如何比它们自身更烦人，并提供了新方法来让事情变得更好一点。

_Steven McCutcheon_ 


[![](https://copm.s3.amazonaws.com/1ea6f5f1.png)](https://golangweekly.com/link/151789/web) 

[走！通过 Ardan Labs Consulting 释放您的技术潜力](https://golangweekly.com/link/151789/web "www.ardanlabs.com") —— 因技能差距、开发速度或复杂的技术挑战而苦苦挣扎？Ardan Labs 专注于 Go、Rust、Docker 和 K8s，以加速您的软件开发、优化架构和管理技术债务。让我们为您的团队加油！

_Ardan Labs Consulting sponsor_


[解决 Go 中的 10 亿行挑战](https://golangweekly.com/link/151792/web "www.bytesizego.com") —— 所谓的[_10 亿行挑战_](https://golangweekly.com/link/151793/web)旨在了解 Java 从文本文件中聚合 1,000,000,000 个（最小、平均和最大）值的速度有多快。在这里，Shraddha 在 Go 中进行了此项尝试，每一步的时间都从 6 分多钟缩短到了 14 秒。

_Shraddha Agrawal_ 


_快速了解：_

  * Sameer Ajmani 是 Google Go 团队的工程总监，他在文章中讲述了自己从一个不想成为经理的人，[转型到管理层](https://golangweekly.com/link/151794/web)的经历。他还撰写了有关 [Go 在 2012 年至 2016 年间的发展](https://golangweekly.com/link/151795/web)的文章。

  * Go 官方语言服务器的最新版本 – [gopls 0.15.0](https://golangweekly.com/link/151796/web) – 已经登陆，支持“零配置” gopls 工作区，并改进了更复杂场景中的初始行为。它也是适用于 Go 1.18 的最终版本。

  * [Git 2.44](https://golangweekly.com/link/151797/web) 是 2024 年第一个新的 Git 版本，Taylor Blau 介绍了其增强功能。
  

[TinyGo 0.31.0：现在支持 Go 1.22 啦](https://golangweekly.com/link/151798/web "github.com") —— 版本号的一小步，对于小型 Go 类型来说却是巨大的飞跃，因为 _“用于小地方的 Go 编译器”_ 获得了 Go 1.22 支持，一个 macOS arm64 原生二进制版本构建， Nim flake 支持，并升级到 LLVM 17。

_Ron Evans_ 


▶ [实际了解下如何在 Go 1.22 中构建 REST API](https://golangweekly.com/link/151799/web "www.youtube.com") —— 使用新的 Go 1.22 方法在短短几分钟内构建快速 REST API。

_TutorialEdge_ 


[使用两行代码将功能齐全的身份验证添加到您的 Go 应用程序中](https://golangweekly.com/link/151800/web) —— 简单、安全且经济实惠：三个都要。FusionAuth 是由开发人员为开发人员构建的身份验证。

_FusionAuth sponsor_


[使用 Glee 更快地写博客：为什么工具从 Python 转向 Go 了](https://golangweekly.com/link/151801/web "journal.hexmos.com") —— [Glee](https://golangweekly.com/link/151802/web) 是一个基于 ~~Python~~ _Go_ 的工具，用于将 Markdown 文件转换为 [Ghost 支持](https://golangweekly.com/link/151803/web)的博客帖子。这篇文章简要介绍了他们将以前基于 Python 的工具重写为 Go 的原因（这并不奇怪）。

_Lince Mathew_ 


🎨 [在 Go 中，为终端输出增加颜色](https://golangweekly.com/link/151804/web "www.dolthub.com") —— 首先演示了它完全没有使用任何依赖项，然后演示 [fatih/color](https://golangweekly.com/link/151805/web) 如何使其更整洁。

_Stephanie You_ 


🛠 代码和工具  
  
[gemini-cli：通过命令行访问 Gemini 模型](https://golangweekly.com/link/151806/web "eli.thegreenplace.net") —— Eli 展示了他为使用 Google Gemini 模型进行数据分析而构建的新的用 Go 编写的 CLI 工具。

_Eli Bendersky_ 


[goquery：类 jQuery 的 HTML/DOM 操作方法](https://golangweekly.com/link/151807/web "github.com") —— 一种类似 jQuery 的方法，用于在 Go 中处理 HTML。就像 jQuery 本身一样，它已经存在很多年了，但最新版本 ([v1.9](https://golangweekly.com/link/151808/web)) 添加了通用的 `Map` 功能并且要求 Go 1.18+。

_Martin Angers and Contributors_ 


[Podinfo：一个 Kubernetes 的 Go 微服务模版](https://golangweekly.com/link/151809/web "github.com") —— 一个由 Go 驱动的小型应用程序，用于展示在 Kubernetes 中运行微服务时的最佳实践，包括运行状况检查、正常关闭、结构化日志记录、构建多架构容器映像。

_Stefan Prodan_ 

  
---
📰 分类广告


🪝 使用 Hookdeck 可靠地大规模接收、验证、转换、过滤、路由和交付 Webhooks。[尝试接收 webhooks 快速入门](https://golangweekly.com/link/151810/web)。

📧 [Forward Email](https://golangweekly.com/link/151811/web) 是一个完全开源的电子邮件转发服务。他们还使用量子安全、单独加密的 SQLite 邮箱。
 
---  


[Risor 1.4：用于 Go 应用的可嵌入脚本](https://golangweekly.com/link/151812/web "risor.io") —— Risor 是一种快速、轻量级的纯 Go 脚本语言，针对 devops 和云集成用例，支持与现有 Go 库进行互操作（[在此处了解更多信息](https://golangweekly.com/link/151813/web)）。v1.4 版本添加了对 `go` 和 `defer` 关键字、通道、HTTP 服务器等的支持。

_Curtis Myzie_ 


[Mailpit 1.14：电子邮件和 SMTP 测试工具](https://golangweekly.com/link/151814/web "github.com") —— Mailpit 内置于 Go，是面向所有开发人员的通用电子邮件测试工具和 API。您可以使用它查看通过其模拟 SMTP 服务器“发送”的电子邮件、存储电子邮件，甚至转发它们。如果您想将电子邮件提取到客户端，v1.14 版本引入了可选的[POP3 服务器](https://golangweekly.com/link/151815/web)。


_Ralph Slooten_ 


[gologin 2.5](https://golangweekly.com/link/151816/web) - 众多身份验证提供商的可链接登录 `http.Handler`。

[NutsDB 1.0.4](https://golangweekly.com/link/151817/web) - 纯 Go 编写的可嵌入、持久键/值存储。

[Excelize 2.8.1](https://golangweekly.com/link/151808/web) - 用于处理 Excel 电子表格的库。


[Sarama 1.43](https://golangweekly.com/link/151818/web) - Apache Kafka 的 Go 库。


[zap 1.27](https://golangweekly.com/link/151819/web) - 快速、结构化、分级的日志记录。


[Air 1.51](https://golangweekly.com/link/151820/web) - Go 应用程序的实时重载。
