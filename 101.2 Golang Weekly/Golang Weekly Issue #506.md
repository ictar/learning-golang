原文：[Golang Weekly Issue #506](https://golangweekly.com/issues/506)

---


[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/g9yzviwurtennzrzhhij.jpg)](https://golangweekly.com/link/154746/web)  


[Go 标准库 `math/rand/v2` 的演进](https://golangweekly.com/link/154746/web "go.dev") —— 生成随机数花费的比您想象的要多得多。 Go 的初始 RNG 有多个缺陷，但修复它会破坏可重复性要求。因此，核心团队创建了一个“版本 2”包，以保持 Go 的兼容性承诺，并为未来此类“版本 2”包总体制定原则。

_Russ Cox (The Go Team)_ 


[Go 1.22 中的安全随机性](https://golangweekly.com/link/154747/web "go.dev") —— 如果您对 Go 最近 _如何_ 改进随机性感兴趣，那么这就是为您准备的，因为作者更深入地探讨了某些用例的安全要求、Go 如何启用不安全代码以及执行 ChaCha(Rand8) 是如何不仅带来了更好、更快、更安全的随机数生成，而且还带来了 _“Go 1.22 无需更改任何代码即可使您的程序更加安全”。_

_Russ Cox and Filippo Valsorda (The Go Team)_ 

 
[![](https://copm.s3.amazonaws.com/bcc68a3c.png)](https://golangweekly.com/link/154745/web) 

[冲呀！专家为您服务](https://golangweekly.com/link/154745/web "www.ardanlabs.com") —— 您是否需要帮助填补技能差距、加快开发速度并使用 Go、Docker、K8s、Terraform 和 Rust 创建高性能软件？我们将帮助您最大化您的架构、结构、技术债务和人力资本。

_Ardan Labs Consulting sponsor_


[Borgo：一种可编译为 Go 的新语言](https://golangweekly.com/link/154748/web "borgo-lang.github.io") —— Hacker News 上的人们[对此感到非常兴奋！](https://golangweekly.com/link/154749/web) Borgo 的作者旨在创建一种 _“比 Go 更具表现力，但比 Rust 更简单”_ 的语言。结果是类型安全性更高，当然值得一看。

_Marco Sampellegrini_ 


_快速了解：_

  * 📘 我们与此没有任何关系，但看到 Packt 和 Humble Bundle 联手制作了一本 [Go 编程书籍“Humble Bundle”](https://golangweekly.com/link/154775/web)，在那里你可以买到很多 Go 书籍，好吧，价格你来定。

  * 🚀 有趣的是，Go 通过为 [Thruster（一个用于_Ruby on Rails_ 应用程序的新 HTTP/2 服务器）](https://golangweekly.com/link/154751/web) 提供支持而进入 Ruby 生态系统的核心。


[将结构用于通用参数列表](https://golangweekly.com/link/154752/web "www.emoses.org") —— 该示例对新算法与旧算法的运行进行了比较，并比较结果以确保它们具有相同的答案。将其视为用于测试重构的功能标志，以及如何使用结构+泛型清理代码。

_Evan Moses_ 


[Go 中 EBPF 的应用介绍](https://golangweekly.com/link/154753/web "edgedelta.com") —— 我们已经链接了几次有关 eBPF 的故事，本文将介绍使用 eBPF 测量资源（例如跟踪数据包）的基础知识和两个示例。

_Ozan Sazak (Edge Delta)_ 


📄 [使用 SQLite 构建高可用搜索引擎](https://golangweekly.com/link/154754/web) - 多亏有了 Go 支持的 [rqlite](https://golangweekly.com/link/154755/web)。 _Philip O'Toole_

📄 [约束 Go 类型参数指针](https://golangweekly.com/link/154756/web) _Merovius_

📄 [Go Web 应用的基础](https://golangweekly.com/link/154757/web) _Willem Schots_


## 🛠 代码和工具  
  
[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/v1pu4yuxwkyhzdvgejjv.jpg)](https://golangweekly.com/link/154758/web)  


[Nimble Terminal Charts：Bubble Tea 图表](https://golangweekly.com/link/154758/web "github.com") —— 提供了一种使用 Bubble Tea 渲染条形图、折线图、散点图、时间序列图和其他图表的方法。这里有很多例子。

_Neomantra Corp_ 


[v8go：通过 V8，在 Go 中执行 JavaScript](https://golangweekly.com/link/154759/web "github.com") —— 如果你觉得很熟悉，这就对了，因为它是 [v8go 主项目](https://golangweekly.com/link/154760/web)的一个分支，但添加了对 Android、新版本的 V8、JS 符号等的支持。

_Tommie and Chapman_ 


[Hookdeck 事件网关](https://golangweekly.com/link/154761/web) —— 一个无服务器队列，用于跨事件驱动架构可靠地发送、接收、转换、过滤和路由事件。

_Hookdeck sponsor_


[mactop：Apple Silicon 的一个基于终端的监控工具](https://golangweekly.com/link/154762/web "github.com") —— 一款特定于 Mac 类`htop`工具，可显示实时 CPU 和 GPU 使用情况，以及有关电源和内存使用情况的指标。

_Carsen Klock_ 

 
[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/p3lauuktk6ouimbkygib.jpg)](https://golangweekly.com/link/154763/web)  


[Logdy：基于 Web 的日志查看器](https://golangweekly.com/link/154763/web "logdy.dev") —— 基于 Web 的实时日志查看器。使用自动生成的过滤器将任何内容流式传输到 Web UI，然后使用 TypeScript 解析任何格式。[现场演示。](https://golangweekly.com/link/154764/web)

_Peter Osinski_ 


[GoWrap：用于生成接口装饰器的 CLI 工具](https://golangweekly.com/link/154765/web "github.com") —— 提供各种模板，可以轻松地将指标、跟踪、后备、池和其他功能添加到现有代码中。

_Max Chechel_ 


---
📰 分类广告

🐷 使用 Go 的 Web 和移动应用程序开发人员更喜欢使用 [Porkbun](https://golangweekly.com/link/154766/web) 作为域名。现在只需 5 美元即可从[Porkbun](https://golangweekly.com/link/154766/web) 获取 .dev、.app 或 .foo 域名。

_Frontend Masters_ 已通过[在 AWS 上构建可扩展的 Go 应用程序](https://golangweekly.com/link/154767/web)扩展到 Go 世界，这是一门新的 Go 课程，专注于让 Go 应用程序在 Amazon 平台上运行。

---  


[gocron 2.5](https://golangweekly.com/link/154768/web) - 以预定的时间间隔运行 Go 函数。

[blake3 1.3](https://golangweekly.com/link/154769/web) - AVX-512 加速的 BLAKE3 哈希实现。

[HaxMap 1.4](https://golangweekly.com/link/154770/web) - 快速且内存高效的并发哈希表。

[fzf 0.51](https://golangweekly.com/link/154771/web) - 流行的命令行模糊查找器。

[GoBGP 3.26](https://golangweekly.com/link/154772/web) - BGP 的 Go 实现。

[Gin 1.10](https://golangweekly.com/link/154773/web) - 流行的 HTTP web 框架。

