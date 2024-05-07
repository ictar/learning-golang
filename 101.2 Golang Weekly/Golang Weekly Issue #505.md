原文：[Golang Weekly Issue #505](https://golangweekly.com/issues/505)

---


✍️ 这一周，如果说有什么与 Go 相关的重大新闻，那并没有，但我们仍然遇到了一个大事儿，因为有很多我们之前没有抽出时间来看一看的东西 ;-) 这就来咯......！ 
__  
 _你的编辑，Peter Cooper_  
  

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/rbhch76corsxudmueok8.jpg)](https://golangweekly.com/link/154399/web)  
  
  
[Sonic：一个新的开源低延迟网络和 I/O 库](https://golangweekly.com/link/154399/web "www.talos.com") —— 不要与同名[博客工具](https://golangweekly.com/link/154400/web)或者 [JSON  库](https://golangweekly.com/link/154401/web)混淆，_这个_ Sonic 是一个 Go 的异步网络和 I/O 库，由一家专门从事贸易基础设施新开源的。_“Sonic 是网络包的替代品。它消除了使用多个 goroutine 来处理同一进程中的多个连接和读/写的需求。”_ [GitHub 存储库](https://golangweekly.com/link/154402/web)。

_Talos_ 



[![](https://copm.s3.amazonaws.com/224af046.png)](https://golangweekly.com/link/154398/web) 

[Stytch：AuthN、AuthZ、欺诈预防的 Auth0 替代方案](https://golangweekly.com/link/154398/web) —— 使用 Stytch 通过 SSO、RBAC 和 SCIM 进行企业级多租户 B2B 身份验证。使用预构建的 UI、无头或直接与 API 集成。此外，Stytch 还内置了设备指纹识别功能，可以检测机器人并防止滥用。[免费开始使用](https://golangweekly.com/link/154398/web)。

_Stytch sponsor_


[我更喜欢传递结构指针的两个原因](https://golangweekly.com/link/154403/web "preslav.me") —— 将结构传递给函数时，如何决定是否使用指针呢？ Preslav 有一些受领域驱动设计原则启发的可靠标准，可以为您提供一种新的、更易于维护的方法。

 
_Preslav Rachev_ 


_快速了解：_

  * Daniel Lemire 一直在根据拉取请求的数量来了解 GitHub 上的语言流行度 [🐦，而 Go 的表现非常好](https://golangweekly.com/link/154404/web)。

  * Google 最近解雇了 Python 和 Flutter 团队的成员后，Reddit 上开始讨论，[如果 Google 对 Go 核心团队也采取同样的做法会发生什么？](https://golangweekly.com/link/154405/web)

  * Twitter/𝕏 上的 [🐦 _Golang Insiders_ 社区](https://golangweekly.com/link/154406/web)已经拥有超过 5000 名成员 —— 如果您仍然活跃在那里，那么值得加入。

  * [cpmulator](https://golangweekly.com/link/154407/web) 是一个用 Go 编写的 CP/M 模拟器（具有[约 50 年历史](https://golangweekly.com/link/154408/web)的操作系统），可以运行原始的 ZORK 文本冒险游戏！


[Go 不是 Java](https://golangweekly.com/link/154409/web "blog.vertigrated.com") ——如果你想吐槽一下，这可能会解决你的问题。这篇文章还认为，根据面向对象的创建者的说法，Go 可能比 Java 或 C++ 更加面向对象……（David Wickes 最近[也写过这个主题](https://golangweekly.com/link/154410/web)）。

_Jarrod Roberson_ 


[LLM 的令牌：用 Go 进行字节对编码](https://golangweekly.com/link/154411/web "eli.thegreenplace.net") —— 令牌是 LLM 的基础，因此了解它们的编码和解码方式可以帮助您更好地理解 LLM 的工作原理。

_Eli Bendersky_ 


📄 [用 Go 写一个 LSP 服务器](https://golangweekly.com/link/154412/web) _Ewen Le Bihan_

📄 [掌握 Go 中的 Map：你需要了解的一切](https://golangweekly.com/link/154413/web) _Ivan Sharapenkov_

📄 [Go 中 HTTP 请求多路复用的背后](https://golangweekly.com/link/154414/web) _Akshay Kumar_

📄 [在控制流中存储数据](https://golangweekly.com/link/154415/web) _Russ Cox_

📄 [使用 WebAssembly 运行 Elixir 中的 Go 代码](https://golangweekly.com/link/154416/web) _Yasoob Khalid_

📺 [针对 CPU 缓存优化 Go 代码](https://golangweekly.com/link/154417/web) _William Moran_



## 🛠 代码和工具  

[Konf 1.1：灵活的配置加载器](https://golangweekly.com/link/154418/web "github.com") —— 从各种位置加载配置设置，从文件和环境变量等本地源到 S3、AWS Parameter Store、AWS AppConfig 和 GCP Secret Manager 等平台。 v1.1 添加了对通过 AWS SNS、GCP PubSub 和 Azure 服务总线进行更改通知的支持。

_Kuisong Tong_ 

[Gohalt：一个通用节流库](https://golangweekly.com/link/154419/web "github.com") —— 你可以用它来构建节流管道、速率限制器等。我们正在挖掘这个库上的 gopher 标志。现在支持 Go 1.22。

_Kostiantyn Masliuk_ 
  

[使用 OAuth 保护 Golang 应用程序](https://golangweekly.com/link/154420/web "fusionauth.io") —— 本教程将向您展示如何使用 OAuth 对 Golang 应用程序中的用户进行身份验证。敲码快乐！

_FusionAuth sponsor_


[gorush：基于 Gin 构建的推送通知服务器](https://golangweekly.com/link/154421/web "github.com") —— 支持 APNS（Apple 推送通知服务）、Firebase 和 HMS 推送服务器（华为）。

_Bo-Yi Wu_ 


[Gotenberg 8.5：用于创建 PDF 文件的 Docker 支持的无状态 API](https://golangweekly.com/link/154422/web "gotenberg.dev") —— Go 支持的系统，提供开发人员友好的 API，用于将多种文档格式转换为 PDF 文件。[GitHub 存储库](https://golangweekly.com/link/154423/web)。

_Gotenberg Inc._ 


[Awesome Ebitengine：Ebitengine 资源精选列表](https://golangweekly.com/link/154424/web "github.com") —— [Ebitengine](https://golangweekly.com/link/154425/web) 是一个流行且强大的 Go 2D 图形和游戏开发 API，这个列表包含用它编写的游戏、与之一起使用的库等等。

_Artem Sedykh_ 


[quic-go 0.43](https://golangweekly.com/link/154426/web) - 纯 Go QUIC 实现。发布的一小步，却是[新文档站点](https://golangweekly.com/link/154427/web)的巨大飞跃！

[Redka 0.3](https://golangweekly.com/link/154428/web) - 使用 SQLite 重新实现 Redis。 v0.3 开始支持 _排序集_。

[pdfcpu 0.8](https://golangweekly.com/link/154429/web) - Go PDF 处理器。现在支持 PDF 2.0 加密。

[Traefik 3.0](https://golangweekly.com/link/154430/web) - HTTP 反向代理和负载均衡器。最大的基于 Go 的项目之一。

[Gum 0.14](https://golangweekly.com/link/154431/web) - _Charm_ 提供的 Go 支持的 shell 脚本实用程序 。


---  

📰 分类广告

[查看 GoLand 中的全行代码补全：](https://golangweekly.com/link/154432/web) 本地运行的捆绑 AI 补全。这是体验人工智能体验的完美方式！

🪝 [Hookdeck](https://golangweekly.com/link/154433/web)：用于事件驱动应用的无服务器队列。[了解更多](https://golangweekly.com/link/154433/web)。

  
---  


## 🎁 还有一个哦，讨个吉利……

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/nankexhaou3cmcb3r7f4.jpg)](https://golangweekly.com/link/154434/web)  

  
[goread 1.6.5：基于终端的 RSS/Atom 阅读器](https://golangweekly.com/link/154434/web "github.com") —— 使用 [Bubble Tea](https://golangweekly.com/link/154435/web) TUI 框架构建。如果想测试下，可以访问 `https://golangweekly.com/rss` ;-)
 
_Adam Piaseczny_ 

