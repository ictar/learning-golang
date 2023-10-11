原文：[Golang Weekly Issue #479](https://golangweekly.com/issues/479)

---

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/k1w7h25y2sdbiiatfdep.jpg)](https://golangweekly.com/link/146181/web)  
---  
  

[思考更好的 `encoding/json`](https://golangweekly.com/link/146181/web "github.com") —— 由 Joe Tsai 领导，但在几位 gopher 同伴的参与下，这是一次讨论，旨在通过新的 v2 版本（已经有[一个实验性实现可以使用](https://golangweekly.com/link/146182/web)）来启动对长期存在的 `encoding/json` 包进行可能性现代化的过程。
 

Joe Tsai et al. 

  
---  
  

[关于类型推断你一直想知道的一切（以及更多）](https://golangweekly.com/link/146183/web "go.dev") —— 类型推断允许 Go 的编译器在 _没有_ 显式注释的情况下确定类型。在这篇文章中，Go 核心团队的 Robert Griesemer 基于他最近在 GopherCon 上发表的演讲，针对一个经常被误解的主题，围绕您需要了解的所有内容进行了阐述。

Robert Griesemer 

  
---  
[![](https://copm.s3.amazonaws.com/5293d39d.png)](https://golangweekly.com/link/146180/web) 

[❤️ Postgres](https://golangweekly.com/link/146180/web "www.crunchydata.com") —— 您需要一个像您一样热爱 Postgres 的数据库提供商。我们将处理所有麻烦 - 监控、备份、HA、灾难恢复，因此您无需担心。想要惊人的支持吗？只要你有疑问，我们就在。

Crunchy Bridge sponsor

  
---  
  

[插入操作繁重的工作负载上的 Go 数据库驱动开销](https://golangweekly.com/link/146184/web "notes.eatonphil.com") —— 对于 Go 中的插入繁重的工作负载，Phil 建议切换数据库驱动程序 - 特别指出，现在 _“可能没有充分的理由再使用 `lib/pq` 来从 Go 访问 PostgreSQL ”_。他建议使用 [pgx](https://golangweekly.com/link/146185/web)。
 

Phil Eaton 

  
---  




**简而言之：**

📢 [Go 1.21.2 and 1.20.9](https://golangweekly.com/link/146186/web) 已发布，但是 [1.21.3 和 1.20.10](https://golangweekly.com/link/146187/web) 预计在 _今天_ 晚些时候发布。所有这些版本都包含安全修复。

📰 比尔·肯尼迪 (Bill Kennedy) 写信告诉我们，他的 [Ultimate Go Tour](https://golangweekly.com/link/146188/web) 现已被翻译成多种其他世界语言，其中 [波斯语版本](https://golangweekly.com/link/146189/web)是第一个可用的。其他十二种语言也将很快推出。

📗 [通过袖珍项目学习 Go](https://golangweekly.com/link/146190/web) 是 Manning 本月晚些时候出版的一本新书，但其中大部分内容已经以抢先体验的形式提供。

  
---  


  

[Web 服务器“Hello World”记住呢是：Go vs Node vs Nim vs Bun](https://golangweekly.com/link/146191/web "lemire.me") —— 标准免责声明适用：基准测试很困难，并且并不总是衡量您 _应该_ 关心的内容。尽管如此，这里还是有一个简单的示例，比较了使用 Go、Node、Bun 和 Nim 的 _最简单_ 的 HTTP 服务器。
 

Daniel Lemire 

  
---  
  

[可视化 Go GC 的一次尝试](https://golangweekly.com/link/146192/web "www.aadhav.me") —— Aadhav 在 GopherCon India '23 上的演讲提案可能已被拒绝，但他将他的失望变成了一篇关于他的实验的有趣博客文章。
 

Aadhav Vignesh 

  
---  


📈 [Statsviz](https://golangweekly.com/link/146193/web) 并不 _完全_ 相同，但它是一个简洁的工具，用于可视化与程序相关的运行时指标的实时图，并且刚刚发布了新版本。
  
---  
  

[添加图标到用 Go 构建到 Windows 可执行文件](https://golangweekly.com/link/146194/web "hjr265.me") —— 如果您构建了一个应用程序并希望它的 `.exe` `有一个更好的图标，那么，[rsrc](https://golangweekly.com/link/146195/web) 提供了一种方法来实现这一点。
 

Mahmud Ridwan 

  
---  


## 🛠 代码和工具 
  
---  
 [![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/km9znpmimnfvpjskmsbe.jpg)](https://golangweekly.com/link/146196/web)  


[Mods 1.0：来自 Charm 的命令行 AI](https://golangweekly.com/link/146196/web "github.com") —— 以 [Bubble Tea](https://golangweekly.com/link/146197/web) 闻名的）那些 _Charm_ 人员最近发布了一款 OpenAI 支持的命令行 AI 工具，可帮助您分析代码、编写文档等。v1.0 版本添加了流响应、保存的对话和 Azure OpenAI 支持。
 

Charm 

  
---  
  

[templ：一种使用 Go 来构建 HTML 的方式](https://golangweekly.com/link/146198/web "templ.guide") —— 创建渲染 HTML 片段的组件，并组合它们以创建屏幕、页面、文档和应用程序。[GitHub 存储库](https://golangweekly.com/link/146199/web)。
 

Adrian Hesketh 

  
---  
  

[Sqinn：无需 Cgo 即可访问 SQLite 数据库](https://golangweekly.com/link/146200/web "github.com") —— 反而，它使用通过 I/O 流/stdin/stdout 进行通信的中介。
 

Christoph Vilsmeier 

  
---  
  

[💻 Hired 使找工作变得容易](https://golangweekly.com/link/146208/web "hired.com") —— 公司不再追逐招聘人员，而是预先向你提供薪资详细信息。[立即创建免费的个人资料吧。](https://golangweekly.com/link/146208/web)
 

Hired sponsor

  
---  



[Viper 1.17](https://golangweekly.com/link/146201/web) - 一个非常强大且灵活的 Go 应用程序配置解决方案。v1.17 _需要_ Go 1.19 并添加对 `log/slog` 的支持。


[Requests 0.23.5](https://golangweekly.com/link/146202/web) - 这个方便的 HTTP 请求库添加了对更改默认 JSON 序列化器/反序列化器或基于每个请求进行设置的支持。


[Task v3.31.0](https://golangweekly.com/link/146203/web) - 想象一下如果 `make` 用 Go 重新实现。([主页。](https://golangweekly.com/link/146204/web))



[Roaring 1.6](https://golangweekly.com/link/146205/web) - “Roaring” 位图数据结构实现。



[Betteralign 0.3.1](https://golangweekly.com/link/146206/web) - 检测其元素可以更好对齐的结构。



[Ginkgo 2.13](https://golangweekly.com/link/146207/web) - 现代测试框架。


**🕰 ICYMI** ( _酒瓶新酒，常看常新_ )  


  * Adebayo Adams 分享了在 Go 应用程序中[有效处理日期和时间](https://golangweekly.com/link/146209/web)的一些基础知识。
  * [探索性编写一个 `slog` 处理器](https://golangweekly.com/link/146211/web)的全系列文章。首先是对 `log/slog` 包的介绍。
  * Shane 分享了使用 Rust 创建 100 万个线程与使用 Go 创建 100 万个 goroutine 的[比较](https://golangweekly.com/link/146212/web)。
