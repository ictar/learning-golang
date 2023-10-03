原文：[Golang Weekly Issue #476](https://golangweekly.com/issues/476)

---
[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/mbrcasvunkncywnj4iz8.jpg)](https://golangweekly.com/link/145105/web)  
---  
[fx 30：用 Go 写的终端 JSON 查看器](https://golangweekly.com/link/145105/web "fx.wtf") -- `fx` 长期以来一直是查看 JSON 问价的一个有用工具，但是 Anton 借机[从头开始重写了它](https://golangweekly.com/link/145106/web)，这个新版本带有新的主题 TUI 外观、正则搜索、更好的封装、模糊搜索，并支持 _“甚至是最大的 JSON 文件”_。这是一个很棒的工具，也是一个展示了 Go 所擅长的领域的很好的例子。

_Anton Medvedev_ 
  

[Go 对 WASI 的支持](https://golangweekly.com/link/145107/web "go.dev") -- Go 最近的版本引入了针对 WebAssembly 和 WebAssembly 系统接口（[WASI](https://golangweekly.com/link/145108/web)）的功能，使得 WASM 可以和外部系统进行交互。这篇文章涵盖了相关的基础知识，并分享了那些 Go 仍处于新生阶段支持的地方。

_Brandhorst-Satzkorn, Fabre, Gryski, et al._ 

 
[![](https://copm.s3.amazonaws.com/04d191ad.png)](https://golangweekly.com/link/145104/web) 

[仅用两行代码，为你的 Go 应用增加功能齐全的身份验证](https://golangweekly.com/link/145104/web) -- 简单、安全且经济实惠：选择三个。FusionAuth 是由开发人员为开发人员构建的身份验证工具。

_FusionAuth sponsor_

  
_快速了解：_
* 如果你使用 ChatGPT 来帮助自己编写 Go 代码，那么请注意，它的训练截止日期已更改为 2022 年 1 月。我注意到这样做的一个副作用是，它现在认识到 Go 有泛型，并将泛型编写代码。
* 如果你是 JetBrains 的 _GoLand_ IDE 的粉丝，并且也在使用 Rust，那么请注意，他们已经推出了 [RustRover](https://golangweekly.com/link/145109/web)，这是一个针对 Rust 开发人员的类似 IDE。
* 你知道[X/Twitter 上有一个 Go _社区_](https://golangweekly.com/link/145111/web)吗？目前已经约 1600 位成员了。
* [Awesome Slog](https://golangweekly.com/link/145112/web) 尝试将有关 Go 的[新结构化日志记录（slog）功能](https://golangweekly.com/link/145113/web)的链接和资源整合在一起。
* 🐘 上周很晚的时候，[PostgreSQL 16 发布了](https://golangweekly.com/link/145110/web)。


[用 Go 构建语言服务器：Lama2 的 VSCode LSP](https://golangweekly.com/link/145114/web "journal.hexmos.com") -- [Lama2](https://golangweekly.com/link/145115/web) （不要与 Meta 的 Llama 2 LLM 混淆）是一个纯文本驱动的 API 管理器和 REST API 客户端。这篇文章背后的团队使用 Go 为 Lama2 创建了一个 LSP，并解释了它是如何组合在一起的以及为什么 Go 很适合。

_Athreya aka Maneshwar_

[使用数据库为状态机提供动力](https://golangweekly.com/link/145116/web "blog.lawrencejones.dev") -- 如何使用 Go 和数据库构建状态机，以通过并发安全代码实现清晰的状态和转换，以及一致的审计跟踪。

_Lawrence Jones_

[Temporal 入门开发人员指南](https://golangweekly.com/link/145117/web "t.mp") -- 使用此 Go SDK 指南了解 Temporal OSS 是如何为您的服务和应用程序提供持久执行的。

_Temporal Technologies sponsor_

▶ [使用 Templ 讨论 Go 模板](https://golangweekly.com/link/145118/web)   
_The Go Time Podcast podcast_


## 🛠 代码和工具
[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/lsxcky7uea7ldgcdji5m.jpg)](https://golangweekly.com/link/145119/web)  
 
  
[PTerm：一个用于 Go 的 TUI 框架](https://golangweekly.com/link/145119/web "github.com") -- 一个扩展库，包含用于改进命令行应用程序输出的[组件](https://golangweekly.com/link/145120/web)，包括图表、列表、树、精美的日志记录、旋转器等。[文档站点。](https://golangweekly.com/link/145121/web)

_pterm project_

[static-server：一个用于静态内容的 Go HTTP 服务器](https://golangweekly.com/link/145122/web "eli.thegreenplace.net") -- Eli 让它变得非常简单，只需运行一个 `go run` `命令即可开始为本地测试提供当前目录。他在这篇文章中详细解释了如何做以及为什么这么做。

_Eli Bendersky_

[Encore：用于 Go 后端开发的开发人员生产力平台](https://golangweekly.com/link/145123/web "encore.dev") -- Encore 可自动执行开发任务，以缩短反馈循环、帮助协作、提高质量以及提供 2 倍的生产力。

_Encore sponsor_

[Maroto 1.0：在 Go 中生成整个 PDF](https://golangweekly.com/link/145124/web "maroto.io") -- 使用特别流畅的基于网格的方法。这里是一个创建发票 PDF 以及输出结果的[完整代码示例](https://golangweekly.com/link/145125/web)。

_Johnathan Fercher da Rosa_

[u：给 Go 的类型添加 “Unset” 状态](https://golangweekly.com/link/145126/web "github.com") -- 当您想知道是否已设置某个值时。

_Lea Anthony_  

[Goyave：一个优雅的 Go REST 框架](https://golangweekly.com/link/145127/web "goyave.dev") -- 提供结构、路由、配置系统，重要的是，它还提供相当[完备的文档](https://golangweekly.com/link/145128/web)和示例代码。[GitHub repo。](https://golangweekly.com/link/145129/web)

_Jeremy Lambert_


* [Resty 2.8](https://golangweekly.com/link/145130/web)  
↳ 简单的 HTTP / REST 客户端库。

* [tproxy 0.8](https://golangweekly.com/link/145131/web)  
↳ 用于代理和分析 TCP 连接的 CLI 工具。

* [zap 1.26](https://golangweekly.com/link/145132/web)  
↳ 快速、结构化的分级日志库。

* [Skipper 0.18](https://golangweekly.com/link/145133/web)  
↳ 用于服务组合的 HTTP 路由器和反向代理。

* [Lux 0.20](https://golangweekly.com/link/145134/web)  
↳ 视频下载库和 CLI 工具。

* [gocron 1.34](https://golangweekly.com/link/145135/web)  
↳ Go 任务调度包重获新生。

* [compress 1.17](https://golangweekly.com/link/145136/web)  
↳ 实现多种压缩算法。

* [rueidis 1.0.18](https://golangweekly.com/link/145137/web)  
↳ 高性能的 Redis 客户端库。


### 工作   

[通过 Hired 找工作](https://golangweekly.com/link/145138/web) —— Hired 使找工作变得容易 —— 公司不再追逐招聘人员，而是预先向你提供薪资详细信息。立即创建免费的个人资料吧。

_Hired_
