原文：[Golang Weekly Issue #496](https://golangweekly.com/issues/496)

---
  

[Go 1.22 中的 HTTP 路由增强功能](https://golangweekly.com/link/151496/web "go.dev") —— Go 团队对 Go 1.22 中新的 HTTP 路由增强功能的官方看法： _“Go 1.22 为 `net/http` 包的路由器带来了两项增强功能：方法匹配和通配符。这些功能让您可以将常见的路线表达为模式而不是 Go 代码。”_ 然而，到目前为止，这些功能还没有受到普遍欢迎。

_Jonathan Amsterdam_ 


💡 Willem Schots 的 [路由中的 URL 路径参数](https://golangweekly.com/link/151497/web) 教程提供了一个精彩的示例介绍，介绍了一些可能性。  


[Go （再次）进入了 TIOBE 指数的前十名](https://golangweekly.com/link/151498/web "www.infoworld.com") —— 这个故事一直在社交媒体上流传，尽管同一位作者[一年前写过类似的故事](https://golangweekly.com/link/151499/web)。_六_ 年前我们还报道过 Go 进入了 TIOBE 的前 10 名（在[第 168 期](https://golangweekly.com/link/151500/web)），所以不要对此过度解读（特别是考虑到[统计数据的获取方式](https://golangweekly.com/link/151501/web)）。尽管如此，第八位是 Go 目前为止所达到的 _最高_ 排名，所以就这样。

_Paul Krill (Infoworld)_ 

[结对 —— 一种被低估的与其他开发者协作的方式](https://golangweekly.com/link/151495/web) —— Tuple“让与同事和朋友的结对编程再次变得有趣”。免费试用并了解为什么 Figma 的工程师们会忍不住谈论 Tuple。

_Tuple sponsor_


[在 Go 中调用 C](https://golangweekly.com/link/151502/web "ericchiang.github.io") —— 一篇简便的文章，包含各种示例，包括将数组、字符串和其他类型从 Go 传递到 C。了解如何正确调用 C 代码应该是您工具箱中的一项，这也可以提高您对 Go 的理解。

_Eric Chiang_ 


🔥 [Fuego：受 Nest 启发的 API/Web Go 框架](https://golangweekly.com/link/151503/web "go-fuego.github.io") —— _“唯一从代码生成 OpenAPI 文档的 Go 框架。受到 Nest 的启发，专为 Go 开发人员打造。”_ 这是[一个“hello world”示例](https://golangweekly.com/link/151504/web)，带有完整的文档。或者跳到[GitHub 存储库](https://golangweekly.com/link/151505/web)。

_Fuego, Inc._ 


_快速了解：_

  * 🇮🇹 [GoLab 会议](https://golangweekly.com/link/151506/web)几个月前在意大利举行，他们刚刚上传了[▶️ 很多谈话视频](https://golangweekly.com/link/151507/web) – 这里有宝藏，但我们仍在浏览它们。

  * Vercel 已经将他们的 Turborepo 应用程序从 Go 过渡到 Rust 一段时间了，他们在博客中介绍了[过渡过程的最后阶段](https://golangweekly.com/link/151508/web)以及 _“达到 Go 三明治的极限”_。

  * [▶️ 迪斯尼使用一个由 Go 支持的服务](https://golangweekly.com/link/151509/web)，在现场活动期间捕获超过十亿个表情符号。

  * [扩展 Go 到 192 个核心](https://golangweekly.com/link/151510/web)的故事。


[Kubernetes CPU 限制和 Go](https://golangweekly.com/link/151511/web "www.ardanlabs.com") —— Go 的运行时本质上并不知道它是否是在 Kubernetes 环境中运行，因此如果您这样做了，那么您可能需要考虑 CPU 限制以及 `GOMAXPROCS`。

_William Kennedy_ 


[给 Go 的 `goto` 一个机会（`retry`）？](https://golangweekly.com/link/151512/web "ammar.io") —— 如果说我们中的许多人从小就被教育为“绝不使用 `goto`”，这肯定只是有点夸张了，但许多开发人员都强烈反对它的使用。尽管如此，Go 有它，在标准库中也用到了它，并且在很多情况下它都很有用。

_Ammar Bandukwala_ 
  

[如何用 Go 来监控 Reddit 上的关键字](https://golangweekly.com/link/151513/web "kwatch.io") —— Reddit 的 JSON API 使得扫描变得非常容易，但是结果因人而异（YMMV）。

_Arthur at KWatch_ 

▶ [利用 WebAssembly 和 Extism，为你的 Go 应用附加超能力](https://golangweekly.com/link/151514/web)   
_Philippe Charriere_  


[了解最近的一个针对 `reflect.TypeFor` 的优化](https://golangweekly.com/link/151515/web)   
_Chris Siebenmann_  


🛠 代码和工具  
  
[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/qay4qanip72movts5gqt.jpg)](https://golangweekly.com/link/151516/web)  


[Gofakeit v7：随机数据生成库](https://golangweekly.com/link/151516/web "github.com") —— 您可以要求它生成诸如随机名称、电子邮件地址、电话号码、职位这样内容，或者让它利用带注释的数据类型来[填充结构](https://golangweekly.com/link/151517/web)。v7 版本与 Go 1.22 的 `math/rand/v2` 集成，可以简化在发行说明中显示的某些情况下的使用。（[主要文档](https://golangweekly.com/link/151518/web)。）

_Brian Voelker_ 


[ObjectBox Go 1.8：用于 Go 对象的快速嵌入式数据库](https://golangweekly.com/link/151519/web "github.com") —— 将对象存储在 ObjectBox 中并使用[强大的查询](https://golangweekly.com/link/151520/web)再次查找它们。v1.8 版本添加了对内存数据库的支持。

_ObjectBox Ltd._ 


[WorkOS，B2B SaaS 的现代身份平台](https://golangweekly.com/link/151521/web "workos.com") ——  WorkOS 提供易于使用的 API，用于身份验证、用户身份以及 SSO 和 SCIM 等复杂的企业功能。

_WorkOS sponsor_


[AWS Lambda Web Adapter：在 AWS Lambda 上更轻松地运行 HTTP Web 应用](https://golangweekly.com/link/151522/web "github.com") —— Go Weekly 中的一个 Rust 项目？🫣 不用担心，这是一个方便的适配器，它通过让应用程序坚持其自身的 HTTP 服务方式，并利用适配器弥补差距，从而可以更轻松地在 AWS 的无服务器平台上运行以 _任何_ 语言构建的应用程序。

_Amazon Web Services Labs_ 


[go-redis 9.5](https://golangweekly.com/link/151523/web) - Redis 客户端库。_（请注意，当与 Redis v7.2.0 之前的旧版本一起使用时，[此版本可能会失败](https://golangweekly.com/link/151524/web)）。_

[testfixtures 3.10](https://golangweekly.com/link/151525/web) - 用于 Go 的 Ruby on Rails 式测试装置。

[go-fault 1.0.2](https://golangweekly.com/link/151526/web) - 使用 HTTP 中间件注入故障。

[Vale 3.1](https://golangweekly.com/link/151527/web) - 自然语言/散文的 linter。

[Miniflux 2.1](https://golangweekly.com/link/151528/web) - 极简主义并且有自己想法的提要阅读器

[MongoDB Go Driver 1.14](https://golangweekly.com/link/151529/web) - MongoDB 的官方 Go 驱动程序。

[eBPF 0.13](https://golangweekly.com/link/151530/web) - 用于 eBPF 程序的 纯 Go 库。

[go-resiliency 1.6](https://golangweekly.com/link/151531/web) - Go 的弹性模式。

[Mockery 2.42](https://golangweekly.com/link/151532/web) - 模拟代码自动生成器。



> 🕰️ ICYMI（_老点的链接，但仍然值得一看_）
>
> 🤖 能够在本地运行大型语言模型 (LLM) 是 AI 领域最令人兴奋的发展之一，而 [Ollama 是一个由 Go 驱动的工具](https://golangweekly.com/link/151533/web)，可以用来实现这一目标，并且在该领域越来越受欢迎。[官方主页]](https://golangweekly.com/link/151534/web)。
>
> Kyle Passarelli 着眼于[如何优化测试函数以提高可读性](https://golangweekly.com/link/151535/web)，以及这样做的好处。
>
> 这里介绍了如何[用 Go 构建一个实时 HTTP 音频流服务器](https://golangweekly.com/link/151536/web)，包括使用多个循环、自动收录器和缓冲区向客户端广播的能力。
