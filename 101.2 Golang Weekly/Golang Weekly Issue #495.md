原文：[Golang Weekly Issue #495](https://golangweekly.com/issues/495)

---

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/heblx3cal5dvoa65cbfz.jpg)](https://golangweekly.com/link/151202/web)  
---  
  

[Go 1.22 发布啦](https://golangweekly.com/link/151202/web "go.dev") —— 秉承 _大多数_ Go 的偶数版本都在 2 月发布的传统，Go 1.22 来了！理论上，升级就像更新 `go.mod` 中的版本一样简单（只需小心任何[`net/http.ServeMux` 损坏……](https://golangweekly.com/link/151203/web)），然后您将能够享受各种改进啦：
 
  * 性能改进较小，毕竟通过改进[配置文件引导优化](https://golangweekly.com/link/151204/web)可以获得更大的潜在性能提升。
 
  * `for` 循环中定义的变量在每次迭代时都会重新创建。
 
  * `for` 循环现在可以针对整数 `range` 了。
 
  * [range-over 函数迭代器](https://golangweekly.com/link/151205/web) 现在出现在标志后面。
 
  * [`http/ServeMux`](https://golangweekly.com/link/151206/web) 及其路由能力得到了很大的增强。这里是[一个快速的比较。](https://golangweekly.com/link/151207/web)
 
  * [`math/rand/v2`](https://golangweekly.com/link/151208/web) - [这里解释了](https://golangweekly.com/link/151209/web)动机。
 
  * 一个让我印象深刻的可爱增强功能是：_“当 `io.Copy` 从 `TCPConn` 复制到 `UnixConn` 时，它现在将尽可能使用 Linux 的 `splice(2)` 系统调用，并使用新方法`TCPConn.WriteTo.`”_ 

要了解 _所有内容_，[完整的 Go 1.22 发型说明](https://golangweekly.com/link/151210/web) 正是你要 _去_ 的地方。

_Eli Bendersky and the Go team_ 


> 💡 我们一个月前就链接到它了，但值得重新审视 Anton Zhiyanov 的 [Go 1.22 交互式（非官方）发行说明](https://golangweekly.com/link/151211/web)。这是一种无需离开浏览器即可使用一些新代码功能的巧妙方法。  
  

[![](https://copm.s3.amazonaws.com/bcc68a3c.png)](https://golangweekly.com/link/151201/web) 

[冲鸭！专家为您服务](https://golangweekly.com/link/151201/web "www.ardanlabs.com") —— 您是否需要帮助填补技能差距、加快开发速度并使用 Go、Docker、K8s、Terraform 和 Rust 创建高性能软件？我们将帮助您最大化您的架构、结构、技术债务和人力资本。

_Ardan Labs Consulting sponsor_


[13 年后，我是如何用 Go 编写 HTTP 服务的](https://golangweekly.com/link/151212/web "grafana.com") —— _Go Time_ 播客的主持人 Mat Ryer 几年前写了一篇文章，介绍了[他在 _8_ 年后如何编写 HTTP 服务](https://golangweekly.com/link/151213/web) — 现在他回来了，进行了一次涉及 13 年的重演，涵盖了自 2018 年以来发生了变化的内容，并且总体上更加深入。

_Mat Ryer (Grafana Lab)_ 


[减少 Go 依赖](https://golangweekly.com/link/151214/web "dgt.hashnode.dev") —— [Huma](https://golangweekly.com/link/151215/web) 中减少依赖的一个案例研究，一个使用 OpenAPI 来创建 HTTP REST API 的 Go 框架，这可能会启发您为自己的项目考虑使用类似的选项。

_Daniel Taylor_ 



_快速了解_

  * Kashyap Kondamudi 观察到 [对于每一次 `Read` 调用，Go 只能读取 1GB 数据。](https://golangweekly.com/link/151216/web) 事实上，这没什么大不了的，但知道这件事也挺有趣的。

  * ▶️ Matt Boyle 提供了[一个免费的“ _22 分钟 Go 1.22_ ”课程](https://golangweekly.com/link/151217/web)。

  * John Arundel 表示他[更新了他写的所有的 Go 书籍](https://golangweekly.com/link/151218/web) 到支持 Go 1.22。

  * 有没有想过[SSH 最后是怎样选择端口 22 的？](https://golangweekly.com/link/151219/web) _（又是那个数字！）_

  * 📺 GitHub 联合创始人 Scott Chacon 在 FOSDEM 2024 上发表了[▶️ 关于 Git 的精彩演讲](https://golangweekly.com/link/151220/web)，涵盖了许多您可能不知道的有关 Git 功能的内容。我 _真的_ 很喜欢它。

  * Google [已向 Rust 基金会捐赠 100 万美元。](https://golangweekly.com/link/151221/web)


[从 Go 的标准库开始](https://golangweekly.com/link/151222/web "matthewsanabria.dev") —— Go 有一个很棒的标准库，故而 _“你应该从它开始”_，Matthew 说。他还花时间将一些库与第三方替代方案进行了比较。

_Matthew Sanabria_ 


[我希望 Go 有一个 `retry` 块](https://golangweekly.com/link/151223/web "xeiaso.net") —— _“我有点希望 Go 有某种语言级别的构造，用于‘由多个可能失败的部分组成的操作，当其中一个以非永久性方式失败时，程序将等待一段时间，然后再重试。”_

_Xe Iaso_ 


[WorkOS，B2B SaaS 的现代身份平台](https://golangweekly.com/link/151224/web "workos.com") —— WorkOS 提供易于使用的 API，用于身份验证、用户身份以及 SSO 和 SCIM 等复杂的企业功能。

_WorkOS sponsor_
  

[`jsonfile`：修补的快速技巧](https://golangweekly.com/link/151225/web "crawshaw.io") —— Tailscale 的 CTO 非常擅长使用 JSON 文件作为小型 _临时_ 数据库。他解释了原因、优点和缺点，并以 [jsonfile.go.](https://golangweekly.com/link/151226/web) 的形式分享了一个实现。

_David Crawshaw_ 


🐭🧀 [使用 `chromedp` 和 Go 喂养一只饥饿的老鼠](https://golangweekly.com/link/151227/web "www.pacenthink.io") —— [chromedp](https://golangweekly.com/link/151228/web) 是一个包，通过[Chrome 的 DevTools 协议](https://golangweekly.com/link/151229/web) 远程控制 Chrome 实例。本文的开发者试图用它来自动玩基于浏览器的游戏。

_Bhupesh Varshney_ 


🛠 代码和工具  
  
[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60/lumxdamlgnkcidxyibpi.jpg)](https://golangweekly.com/link/151230/web)  

[gdu / go DiskUsage() 5.26.0：带终端界面的磁盘使用分析器](https://golangweekly.com/link/151230/web "github.com") —— 当然是用 Go 编写的。这是一个很棒的工具。快速分析您的磁盘空间使用情况 - 与该领域的其他工具相比，它的基准测试良好。

_Daniel Milde_ 


[errcheck 1.7：它检查你是否进行了错误检查](https://golangweekly.com/link/151231/web "github.com") —— 检查错误是 Go 体验的一个基本组成部分，此工具将帮助您检查是否已检查了错误！
 
_Kamil Kisiel_ 
 
---  

📰 分类广告


🪝[Hookdeck](https://golangweekly.com/link/151232/web)：用于 Webhooks 和异步消息传递的托管可靠性和可观察性层。[大规模接收、转换、路由和交付事件](https://golangweekly.com/link/151232/web)。

---

[Wire 0.6.0：Go 的编译时依赖注入](https://golangweekly.com/link/151233/web "github.com") —— Wire 的任务是在进行依赖注入的时候，[简化初始化代码的管理](https://golangweekly.com/link/151234/web)。三年来首次发布！

_Google_ 


[sh 3.8：Shell 代码解析器、格式化器和解释器](https://golangweekly.com/link/151235/web "github.com") —— 支持 sh、bash 和 Korn/mksh 格式。现在至少需要 Go 1.21。

_Daniel Marti_


[FastHTTP 1.52](https://golangweekly.com/link/151236/web) - 针对[特定大容量场景](https://golangweekly.com/link/151237/web)，面向性能的 `net/http` 替代方案。

[Garble 0.12](https://golangweekly.com/link/151238/web) - Go 构建的混淆。现在支持 Go 1.22。

[Wails 2.8](https://golangweekly.com/link/151239/web) -  使用 Go + Web 技术构建桌面应用程序。

[go-github 59.0](https://golangweekly.com/link/151240/web) - GitHub v3 API 客户端库。

[fq 0.10](https://golangweekly.com/link/151241/web) - 想象一下是 `jq`，但是是二进制格式。

[go-imap 2.0 Beta 1](https://golangweekly.com/link/151242/web) - IMAP4 客户端库。
