原文：[Golang Weekly Issue #500](https://golangweekly.com/issues/500)

---


[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/dklgdc36philzy7d9ebf.jpg)](https://golangweekly.com/link/152677/web)  


[更强大的 Go 执行跟踪](https://golangweekly.com/link/152677/web "go.dev") —— 在过去的两个版本中，Go 团队致力于修复和增强执行跟踪，从而使开销下降了 10 倍，并提高了跟踪的可扩展性。这些增强功能催生了两个实验性功能：持续跟踪执行的能力（即所谓的“[飞行​​记录](https://golangweekly.com/link/152678/web)”）和跟踪读取器 API。

_Michael Knyszek_ 

  
[![](https://copm.s3.amazonaws.com/58dfcb2f.png)](https://golangweekly.com/link/152676/web) 

[学习构建可扩展且有弹性的 Go 后端](https://golangweekly.com/link/152676/web "threedots.tech") —— Go 事件驱动的下一版本即将开始。来学习最热门的技能之一并从人群中脱颖而出吧。

_Three Dots Labs sponsor_


[使用 Go 测量系统性能](https://golangweekly.com/link/152679/web "lemire.me") —— 您可以使用任何编程语言运行基准测试，但 Go 提供了一些功能，让您可以相对轻松地测量 CPU 和内存性能、内存使用情况，甚至深入研究 Go 的一些优化。这里有很多例子。

_Daniel Lemire_ 

  

[_十亿行挑战_（又来啦！）- 从 95 秒到 1.96 秒](https://golangweekly.com/link/152680/web "r2p.dev") —— Gunnar Morling 的 [十亿行挑战](https://golangweekly.com/link/152681/web)（从 10 亿行文件中读取和汇总统计数据）已经流行起来了！我们之前链接到了 [Shraddha Agrawal](https://golangweekly.com/link/152682/web) 和 [Ben Hoyt](https://golangweekly.com/link/152683/web) 的解决方案，但这里有一个 _非常_ 深入的研究，涉及许多简洁的底层细节。

_Renato Pereira_ 



_快速了解：_

  * 🔒 EFF 一直在[想 Caddy 和 Traefik 是否终有一天应该替代 Certbot](https://golangweekly.com/link/152684/web)（值得注意的是，[Caddy](https://golangweekly.com/link/152685/web) 和 [Traefik](https://golangweekly.com/link/152686/web) 都是基于 Go 的项目。）

  * 🤖 [Mechanoid](https://golangweekly.com/link/152687/web) 是 Ron Evans（因 TinyGo 和 [Gobot](https://golangweekly.com/link/152688/web) 而有名）的一个新项目，旨在将 [WASM 引入到物联网的嵌入式系统](https://golangweekly.com/link/152689/web) - 内置 Go，但与 Gobot 不同，您可以使用 Go、Rust、Zig 和其他编译为 WASM 的语言。


[使用非阻塞读取操作调试 Go 错误](https://golangweekly.com/link/152690/web "scratchdata.com") —— 一种在不同操作系统上表现不同的场景，导致 Go 中的一个快速修补。但是，如果您等不及，有一个简单的解决方法。

_Scratch Data_ 


[Go 中的 For 循环已经更多内容](https://golangweekly.com/link/152691/web "www.ardanlabs.com") —— `for` 循环可能看起来像是初学者的主题，但要理解它们的完整用途，就不仅仅如此。如何处理多个循环变量？什么时候应该使用标签分隔符？范围循环语义如何影响更新循环结构？Miki 展示了一些例子。

_Miki Tebeka (Ardan Labs)_ 


[如何为你的项目编写一个能获得“4000 星”的 GitHub README](https://golangweekly.com/link/152692/web "www.daytona.io") —— Go 支持的 [Daytona](https://golangweekly.com/link/152693/web) 项目的创建者分享了一些关于创建良好 README 的技巧。

_Ivan Burazin_


[让我们来将一个 Go 程序嵌入到 Linux Kernel 中](https://golangweekly.com/link/152694/web) —— 可能不是您一般情况下想要做的事情，但是..   
_Richard Weinberger_  


[使用 GitHub Actions，将 Go 二进制文件添加到发布版本中](https://golangweekly.com/link/152695/web)   
_Rob Allen_  


[我是怎样构建自己的 Go 包索引的](https://golangweekly.com/link/152696/web)   
_Ozan Sazak_  



🛠 代码和工具  
  
[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/qpurw0zcab1itzkqunsm.jpg)](https://golangweekly.com/link/152697/web)  

[Valgo 0.3：富有表现力的验证器库](https://golangweekly.com/link/152697/web "github.com") —— 基于泛型构建的类型安全且可扩展的验证器库。_“Valgo 与其他验证库的不同之处在于，规则是用函数编写的，而不是用结构标签编写的。这在验证数据的地点和方式方面提供了更大的灵活性和自由度。”_

_Cohesive Stack_ 


[Hugot：Go 中的 Hugging Face Transformer 管道](https://golangweekly.com/link/152698/web "github.com") —— 该项目的目标是轻松地用 Go 运行 [Hugging Face](https://golangweekly.com/link/152699/web) Transformer 管道，而无需调用 Python 或外部 API（请注意，它不是 _纯_ Go 的 – 它依赖于 ONNX 运行时）。到目前为止，它仅支持 ONNX 模型和一些管道。

_Knights Analytics_ 


[Failsafe：容错和弹性模式](https://golangweekly.com/link/152700/web "failsafe-go.dev") —— 用于构建容错应用程序的库，您可以用它将代码包装在弹性策略（例如重试、CircuitBreaker、RateLimiter、超时和回退）中。我们几个月前就链接到了它，但从[v0.6](https://golangweekly.com/link/152701/web)开始，它可以[与 HTTP 客户端集成。](https://golangweekly.com/link/152702/web)

_Jonathan Halterman_ 


[Hookdeck：The Amazon EventBridge 的替代方案](https://golangweekly.com/link/152703/web) —— 使用工程团队的事件网关，在 EDA 中接收、转换、过滤、路由和发送消息。

_Hookdeck sponsor_


[Konf 1.0：一个灵活的配置加载器](https://golangweekly.com/link/152704/web "github.com") —— 如果您不想让您的应用程序和任何一个配置设置源紧密耦合在一起，那么 Konf 可能适合您 - 它可以与诸如文件和环境变量这样的本地源配合使用，而且还可以使用 S3、AWS AppConfig 和 GCP Secret Manager 等。

_Kuisong Tong_ 


[Codoworks Go Boilerplate：适用于生产的 RESTful API Boilerplate](https://golangweekly.com/link/152705/web "github.com") —— 一款基于 [Echo](https://golangweekly.com/link/152706/web) 的固定样板应用程序，可快速启动和运行 CRUD API。

_Codoworks_ 


[wazero 1.7](https://golangweekly.com/link/152707/web) - Go 的零依赖 WebAssembly 运行时。对新优化编译器的最终版本进行了重大升级。

[fx 33.0](https://golangweekly.com/link/152708/web) - 流行的 JSON 查看器和处理器获得了终端自动完成的能力。


[Toxiproxy 2.9](https://golangweekly.com/link/152709/web) - 用于模拟混乱网络条件的 TCP 代理。


[Gonum 0.15](https://golangweekly.com/link/152710/web) - Go 的数字库（代数、概率等）


[Resty 2.12](https://golangweekly.com/link/152711/web) - 简单的 HTTP a和 REST 客户端库。

[SCS 2.8](https://golangweekly.com/link/152712/web) - Go 中的 HTTP 会话管理。

[go-ora 2.8.10](https://golangweekly.com/link/152713/web) - 纯 Go Oracle 数据库驱动器。


[Bloom 3.7](https://golangweekly.com/link/152714/web) - Bloom 过滤器实现。
