原文：[Golang Weekly Issue #488](https://golangweekly.com/issues/488)

---

🎄 我们正在准备下周的年度综述，就在圣诞假期之前，所以下周二请留意 :-)  

_你的编辑，Peter Cooper_  
  
 
[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/si8yp1brd4zpv4armazn.jpg)](https://golangweekly.com/link/148896/web)  

[最新的 Go 开发者调查结果](https://golangweekly.com/link/148896/web "go.dev") —— 官方的 Go 调查现在每年进行两次，有 4000 人参加了最新的一次调查（2023 年下半年）。结果并不令人震惊，但有一些新的东西，特别是对于 Go 团队本身：
  * Linux 和 macOS _迄今为止_ 是最常用的操作系统。
  * Go 团队认为结果在很大程度上验证了他们在 [`gonew` 项目模版](https://golangweekly.com/link/148897/web) 方法上所做的工作。
  * VS Code 是最受欢迎的 IDE，但 GoLand 紧随其后。
  * 大多数的 Go 开发者热衷于使用 AI 来 _帮助_ 他们构建软件以及（特别是）编写测试，但不那么热衷于用它来编写完整代码。
  * 目前有 21% 的开发人员维护开源 Go 模块。

_Todd Kulesza_


> 🎧 更想听到关于结果的讨论？ [▶️ Cup o' Go 最新一期](https://golangweekly.com/link/148898/web)干了这件事。  
  

[![](https://copm.s3.amazonaws.com/bcc68a3c.png)](https://golangweekly.com/link/148869/web) 

[冲！专家为你服务](https://golangweekly.com/link/148869/web "www.ardanlabs.com") —— 您是否需要帮助填补技能差距、加快开发速度并使用 Go、Docker、K8s、Terraform 和 Rust 创建高性能软件？我们将帮助您最大化您的架构、结构、技术债务和人力资本。

_Ardan Labs Consulting sponsor_


[CGo-Less SQLite 包到了 1.0 里程碑（Milestone）](https://golangweekly.com/link/148870/web "www.zombiezen.com") —— 三年工作已至顶峰，Ross 宣布他的 CGo-less SQLite 库是稳定了，该库是通过自动将原始 C 翻译成 Go 的方式生成的。是时候开始用用它了。[GitHub 存储库。](https://golangweekly.com/link/148871/web)

_Ross Light_ 


> 🚨 在其他版本中，[Go 1.21.5 和 1.20.12 发布了，](https://golangweekly.com/link/148899/web)主要是为了提供三个安全修复程序。


[Go Recipes：在 Go 项目中运行的便捷命令](https://golangweekly.com/link/148872/web "github.com") —— 自我们几年前首次链接这个便捷资源以来，它已经变得更大，并且包含大量方法（准确地说是 179 个），涵盖测试、基准测试、安全等领域、静态分析、代码生成、PR 建议等等。

_Nikolay Dubina_

[隆重推出 GoLand 2023.3，给 Gophers 的 IDE](https://golangweekly.com/link/148874/web "jetbrains.com") —— 现在提供 AI Assistant、开发容器支持以及多种提示、重构和快速修复，以提高生产力！

_JetBrains sponsor_


> 🤖 我们​​本来要在社论中提到 GoLand 2023.3，但由于他们抢先一步，我们将改为重点介绍 GoLand 新的 [▶️ 自动测试生成和执行功能](https://golangweekly.com/link/148875/web)。  


[用 Go 实现“使用 GitHub 登录”](https://golangweekly.com/link/148873/web "eli.thegreenplace.net") —— 全面、简单地介绍了在 Go 中，使用 OAuth 和 GitHub 的 OAuth 提供商，在您自己的应用程序中获得“使用 GitHub 登录”功能的 _三种_ 方法。

_Eli Bendersky_ 


▶ [我们为什么从 SvelteKit 切换到 Go + HTMX](https://golangweekly.com/link/148876/web "www.youtube.com") —— 这更多是关于使用 [templ](https://golangweekly.com/link/148877/web) HTML 模板框架，为您提供真正的服务器端渲染，简化使用 Go 创建网站的过程。

_Anthony GG_ 


🛠 代码和工具

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/ereaqqe2unngkxvkiprm.jpg)](https://golangweekly.com/link/148878/web)  


[FrankenPHP v1.0：PHP 应用服务器](https://golangweekly.com/link/148878/web "laravel-news.com") —— _Go Weekly_ 中的 PHP 新闻？[FrankenPHP](https://golangweekly.com/link/148879/web) 是一个非常棒的 PHP 应用服务器，它是用 Go 编写的，并且在背后也使用了 Caddy。它还可以当成 Go 的库使用，将可通过 HTTP 访问的 PHP 嵌入到 Go 应用程序中。[GitHub 存储库。](https://golangweekly.com/link/148880/web)

_Paul Redmond_ 


🐍 [Pygolo 0.2.0：使用 Go 潜入或扩展 Python](https://golangweekly.com/link/148881/web "gitlab.com") —— 双向可用。您的 Go 应用程序可以在同一进程中并行运行自己私有的 Python 运行时，并根据需要访问它，[如下所示。](https://golangweekly.com/link/148882/web)或者，您可以[在 Go 中编写组件](https://golangweekly.com/link/148883/web)，然后将其加载到 Python 中。不管怎样：Python 和 Go 和谐共处。

_Domenico Andreoli_ 
   
---
📰 分类广告

[加入 Sticker Mule 的“强悍”团队，担任站点可靠性工程师！](https://golangweekly.com/link/148884/web) 我们的软件团队遍布 17 个国家/地区，我们正在寻找更多优秀的工程师加入我们的安全团队。

👉 [免费的 Temporal 101 & 102 Go 课程](https://golangweekly.com/link/148885/web)：通过我们的 Go 培训（自定进度）来学习 Temporal 的开源关键概念和最佳实践。

💻 Hired 让求职变得简单 - 公司不再需要追逐招聘人员，而是预先向您提供薪资详细信息。[立即创建免费个人资料。](https://golangweekly.com/link/148886/web)

---  


[pdfcpu 0.6：使用 Go 处理 PDF](https://golangweekly.com/link/148887/web "github.com") —— PDF 处理库 _和_ CLI。v0.6 添加了一些新功能，包括基本的 PDF 2.0 支持和一些新的页面布局命令。[项目主页。](https://golangweekly.com/link/148888/web)

_pdfcpu Contributors_ 


∿ [Microwave：一个非常简单的信号生成器](https://golangweekly.com/link/148889/web "github.com") —— 由于缺乏典型的独立信号生成器，因此，作者构建它是来帮助教授电子课程。

_Nikola Ubavić_


[gocron 2.0](https://golangweekly.com/link/148890/web) - 以预定的时间间隔运行 Go 函数。

[OSV-Scanner 1.5](https://golangweekly.com/link/148891/web) - 依赖项漏洞扫描器。

[Tcell 2.7](https://golangweekly.com/link/148892/web) - 文本终端的基于单元格的视图。

[OpenFGA 1.4](https://golangweekly.com/link/148893/web) - 受 Zanzibar 启发的授权引擎。

[Vale 2.30](https://golangweekly.com/link/148894/web) - 自然语言/散文的 linter。

[gRPC-Go 1.60](https://golangweekly.com/link/148895/web) - Go gRPC 实现。