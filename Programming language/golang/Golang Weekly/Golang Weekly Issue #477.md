原文：[Golang Weekly Issue #477](https://golangweekly.com/issues/477)

---
[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/ili93rduppzbkoj5xhwk.jpg)](https://golangweekly.com/link/145458/web)  
---  

[组织 Go 模块](https://golangweekly.com/link/145458/web "go.dev") —— Go 开发人员通常会考虑如何组织项目中的文件和目录（我们的一些最受欢迎的链接都是关于这个主题的）。然而，Go 项目本身在这个话题上往往保持沉默，却让社区找出最佳实践，这在过去[导致了问题](https://golangweekly.com/link/145459/web)，所以很高兴看到他们现在发布了更官方的东西。

_The Go Team_


[走！为你的服务提供专家支持](https://golangweekly.com/link/145457/web "www.ardanlabs.com") —— 能差距、并使用 Go、Docker、K8s、Terraform 和 Rust 加快开发速度以及创建高性能软件？我们将帮助您最大化您的架构、结构、技术债务和人力资本。

_Ardan Labs Consulting sponsor_

[修复 Go 1.22 中的 `for` 循环](https://golangweekly.com/link/145460/web "go.dev") —— Go 1.21 包含了[备受期待的 `for`-loop 范围更改](https://golangweekly.com/link/145461/web)的 _预览_，预计将在 Go 1.22 中完全发布。这篇文章展示了将要发生的变化，并解释了如此重大的变化如何不会影响 Go 的向后兼容性，仅适用于专门声明其针对 Go 1.22 或更高版本的代码。for

_David Chase and Russ Cox_

  
_快速了解：_

* JetBrains 已启动 [GoLand 2023.3 的抢先体验计划](https://golangweekly.com/link/145462/web)了。与往常一样，可以免费尝试，并可以一睹下一个主要版本中的内容。

* 使用 GitHub 时，现在可以使用[对 Passkeys 的支持](https://golangweekly.com/link/145463/web)了。

* 有人 [建议将 runtime/trace 的“飞行记录”](https://golangweekly.com/link/145465/web)添加到 Go 的运行时中，就像 _Java Flight Recorder_ 一样。想法是以低开销的方式，维护跟踪数据的循环缓冲区，使其可以根据需要进行访问。

* [TinyGo 0.30](https://golangweekly.com/link/145497/web) 已经发布，其中包括切换到 LLVM 16、支持 Adafruit Gemma M0 以及修复一些错误，正好赶上 GopherCon。

* 是的，[GopherCon 2023](https://golangweekly.com/link/145464/web) _正在_ 加利福尼亚州圣地亚哥举行。如果您在那里，我们希望您度过愉快的时光！👋


[Failsafe：容错和弹性模式](https://golangweekly.com/link/145466/web "failsafe-go.dev") —— 一个用于构建容错应用程序的新的 Go 库，您可以用它将代码包装在各种弹性策略中，例如 Retry、CircuitBreaker、RateLimiter、Timeout 和 Fallback。它支持异步执行、[协作取消]((https://golangweekly.com/link/145467/web))和事件侦听器。
 
_Jonathan Halterman_

  
[GitHub Actions 可以做得更好](https://golangweekly.com/link/145468/web "blog.yossarian.net") —— _Actions_ 是一项有用的服务，但开发人员的体验还有很多不足之处。如果您在设置和调试工作流程时感到沮丧，您会发现很多值得肯定的地方。

_William Woodruff_

  
[如何（不）申请一份软件工作](https://golangweekly.com/link/145469/web "benhoyt.com") —— 不特定于 Go，但是是我们最喜欢的博主之一。

_Ben Hoyt_

  
[Temporal 101 和 102 Go 课程](https://golangweekly.com/link/145470/web "t.mp") —— 通过我们免费的自定进度 Go 培训课程来学习 Temporal 的开源关键概念和最佳实践。

_Temporal Technologies sponsor_

  
[使用 New Relic 检测 Go 应用程序的十大技巧](https://golangweekly.com/link/145471/web "newrelic.com") —— 虽然这是该系列的第五部分，但它在一个地方总结了所有的十个技巧。

_Steve Ng (New Relic)_


[用 Go 替换 Python，并将 Docker 镜像大小减少约 87%](https://golangweekly.com/link/145472/web)   
_Dragondrop_


[在 Go 中实现枚举](https://golangweekly.com/link/145473/web)   
_William Kennedy_

  
## 🛠 代码和工具 
  
[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/mvar82mt1hv137q2xono.jpg)](https://golangweekly.com/link/145474/web)  

[Ebitengine v2.6.0：2D 游戏引擎](https://golangweekly.com/link/145474/web "ebitengine.org") —— [Ebitengine](https://golangweekly.com/link/145475/web)（之前称为 Ebiten）是在 Go 中构建 2D 游戏的最著名方法，可以在多种平台上运行（甚至包括 Nintendo Switch）。v2.6 版本改进了对 Windows 支持，让您可以控制在哪个显示器显示您的游戏，提供鼠标光标直通功能[等等](https://golangweekly.com/link/145476/web)。

_Hajime Hoshi_

[Encore：用于 Go 开发的开发者生产力平台](https://golangweekly.com/link/145477/web "encore.dev") —— Encore 自动化开发任务和基础设施，以缩短反馈循环、提高质量和 2 倍的生产力。

_Encore sponsor_


[Sonnet：一个（更）高性能的 JSON 库](https://golangweekly.com/link/145478/web "github.com") —— _“与 Go 标准库完全兼容，编码器和解码器都通过了 Go 标准库的所有测试。”_ 作者有基准，但也[写了一篇关于他的方法的解释性文章](https://golangweekly.com/link/145479/web)。

_Sugawara Yuuta_

  
[Afero 1.10：Go 文件系统抽象系统](https://golangweekly.com/link/145480/web "github.com") —— 用于访问各种文件系统的单一一致 API。它还允许你创建完全不依赖于磁盘的模拟和测试文件系统。

_Steve Francia_

  
[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/higu0j1xd25wifna4f1w.jpg)](https://golangweekly.com/link/145481/web) 
  

[Flameshow：终端火焰图查看器](https://golangweekly.com/link/145481/web "github.com") —— 在您使用之前 `go install` 之前，请记住该工具是用 Python 编写的。尽管如此，您还是获得了一个功能强大的基于终端的火焰图渲染器，目前 _仅_ 支持 Go 的 pprof 数据。

_laixintao_

[automaxprocs：自动设置 `GOMAXPROCS` 以匹配 Linux 容器的 CPU 配额](https://golangweekly.com/link/145482/web "github.com")

_Uber Golang_



📰 分类广告
   
---  


🎟️ [CityJS 柏林](https://golangweekly.com/link/145483/web)：演讲者包括 Tejas Kumar、Christian Heilmann 和 Rachel Nabors。使用 `COMMUNITY` 折扣代码可节省 25%。

* * *

💻 [通过 Hired 找工作](https://golangweekly.com/link/145138/web) —— Hired 使找工作变得容易 —— 公司不再追逐招聘人员，而是预先向你提供薪资详细信息。立即创建免费的个人资料吧。

---

* [bitset 1.9](https://golangweekly.com/link/145485/web)  
↳ 用于紧凑存储和访问 _位_ 的数据结构。

* [Hertz 0.7](https://golangweekly.com/link/145486/web)  
↳ 用于构建 HTTP 微服务的框架。

* [mo 1.10](https://golangweekly.com/link/145487/web)  
↳ 泛型驱动的 monad 和 FP 抽象。

* [purego 0.5](https://golangweekly.com/link/145488/web)  
↳ 从 Go 调用 C 函数，无需 Cgo。

* [GoAWK 1.25.0](https://golangweekly.com/link/145489/web)  
↳ 支持 CSV 的 Go AWK 解释器。

* [Podinfo 6.5](https://golangweekly.com/link/145490/web)  
↳ Kubernetes 的 Go 微服务模板。

* [AWS Lambda Go API Proxy 0.16.0](https://golangweekly.com/link/145491/web)




## 🤔 PHP?  
   
[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/i9s3py5kkgzhxz0quwj4.jpg)](https://golangweekly.com/link/145492/web)  

[FrankenPHP 1.0 Beta：另一个 PHP 应用服务器](https://golangweekly.com/link/145492/web "dunglas.dev") -- 你没看错。这里仍然是 Go Weekly，但是，[FrankenPHP](https://golangweekly.com/link/145493/web) 是一个新的 PHP 应用程序服务器，_用 Go 编写_并构建在 Caddy 之上，简化了 PHP 应用程序的部署。如果这听起来很耳熟，那么你可能见到过 [RoadRunner](https://golangweekly.com/link/145494/web)，这是另一个采用 Go 驱动的 PHP 应用服务器，[它采用了非常不同的方法](https://golangweekly.com/link/145495/web)。

_Kevin Dunglas_