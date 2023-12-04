原文：[Golang Weekly Issue #486](https://golangweekly.com/issues/486)

---

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/v4wreqe0hse2bzqpvdiv.jpg)](https://golangweekly.com/link/148241/web)  

[GoFakeIt 6.25.0：随机虚假数据生成器](https://golangweekly.com/link/148241/web "github.com") —— 具有[超过 260 个函数](https://golangweekly.com/link/148242/web)，用于生成姓名、电子邮件、位置、颜色、用户代理、早餐项目（！）等。v6.25.0 引入了一个模板选项，因此您可以一次性创建更复杂的东西（如上所述）。[仓库](https://golangweekly.com/link/148243/web)。

_Brian Voelker_ 


> 💡 使用这样的工具是为了好的，顺便说一句 —— 不要生成[假的会议演讲者……](https://golangweekly.com/link/148244/web) 🤦  （译注：左边的新闻说的是，DevTernity会议因为有两个假的演讲者所以取消了）
  

[Python 很容易。Go 很简单。简单 != 容易](https://golangweekly.com/link/148245/web "preslav.me") —— A high level post touching on the symbiosis that the author and his team have found for using both Python _and_ Go which, naturally, provoked [an extensive discussion on Hacker News](https://golangweekly.com/link/148246/web), if you're not bored of reading what everyone else thinks about programming languages yet.一篇高水平的文章，涉及作者和他的团队发现的同时使用 Python _和_ Go 的共生关系（这自然会在 Hacker News 上引发广泛的讨论），如果你尚未厌倦阅读其他人对编程语言的思考的内容的话。

_Preslav Rachev_ 


[冲呀！通过 Ardan Labs Consulting，解锁您的技术潜力](https://golangweekly.com/link/148240/web "www.ardanlabs.com") —— 因技能差距、开发速度或复杂的技术挑战而苦苦挣扎？Ardan Labs 专注于 Go、Rust、Docker 和 K8s，以加速您的软件开发、优化架构和管理技术债务。让我们为您的团队添柴加薪！

_Ardan Labs Consulting sponsor_


[使用 Go 制作游戏（绝对适合初学者）](https://golangweekly.com/link/148247/web "threedots.tech") —— Miłosz 使用备受喜爱的 [Ebitengine](https://golangweekly.com/link/148248/web) 来创建一个简单的 _Asteroids_ 游戏（如果你不像我那么老的话，请[参阅维基百科](https://golangweekly.com/link/148249/web)）。有很多内容需要介绍，但如果您想了解游戏的制作方式，所涉及的步骤都很小并且很容易就能理解。

_Miłosz Smołka_ 


_快速了解：_
  * ❄️ 在Go 编译器和运行时团队[最新的会议记录](https://golangweekly.com/link/148250/web)中，我们了解到 **Go 1.22 已进入“代码冻结”状态**（这是 Go 发布前三个月的典型情况）。[发行说明](https://golangweekly.com/link/148251/web)的工作才刚刚开始，但最终版本要到二月份才能发布。

  * 🧗 Reddit 子版块 `/r/climbing` 上，有人为他们的配偶制作了[一个 Go gopher“粉笔袋”](https://golangweekly.com/link/148252/web)。

  * 🐦 如果您仍在使用 Twitter/X，请务必查看 [Golang Insiders 社区](https://golangweekly.com/link/148253/web)。目前它已有近 3000 名开发人员，实力雄厚，素质较高。
  
  * 🔠 当然，[Go 有它自己的官方字体](https://golangweekly.com/link/148254/web)，但是 GitHub 最新版本的 [Monaspace](https://golangweekly.com/link/148255/web)，看起来确实不错……（译注：我去看了下，表示赞同😁）
  
  * ƛ 如果您在 _AWS Lambda_ 上运行 Go 函数，那么这是个好消息 —— 面对大量请求，[函数的扩展速度现在提高了 12 倍](https://golangweekly.com/link/148256/web)。



[解密 Go 中的函数参数](https://golangweekly.com/link/148257/web "www.alexedwards.net") —— 不同类型的函数参数如何（以及为什么）表现不同。

_Alex Edwards_ 


[`etcd` 和并发 STM](https://golangweekly.com/link/148258/web "george.macro.re") —— 深入了解 `etcd` 与 Go 集成时所面临的并发问题，特别是围绕事务操作。

_George MacRorie_ 


[结构体尾部的大小为零的字段是如何具有非零大小的](https://golangweekly.com/link/148259/web "i.hsfzxjy.site") —— 大小为零的类型是否始终占用零字节空间？并不是……

_hsfzxjy_ 


[如何解析时间或者日期](https://golangweekly.com/link/148260/web "www.willem.dev") —— 很好地通过交互式示例进行了展示

_Willem Schots_ 


## 🛠 代码和工具

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/meem5grbcpzgollu7icj.jpg)](https://golangweekly.com/link/148261/web)  

[Pagoda 0.10：快速全栈“入门套件”](https://golangweekly.com/link/148261/web "github.com") —— 如果您不介意固定模式，那么，Pagoda 可以作为构建全栈 Go 应用程序的框架和模块的基础。Echo 和 Ent 在后端与 Postgres 和 Redis 进行交互，前端使用 [htmx](https://golangweekly.com/link/148262/web) 和 [Alpine.js](https://golangweekly.com/link/148263/web)。

_Mike Stefanello_ 


[准备好使用您喜欢的配对工具了吗？Tuple 让配对成为一种乐趣](https://golangweekly.com/link/148264/web) —— 您值得拥有专门构建的工具。无缝控制，低CPU占用率，超清晰的音频和视频。14 天免费试用。

_Tuple sponsor_


[htmx-go：一种使用 Go 的 HTMX 的类型安全方式](https://golangweekly.com/link/148265/web "github.com") —— [htmx](https://golangweekly.com/link/148262/web) 是一套越来越流行的 HTML“强大工具” — 本质上是一种以声明性方式，通过 JavaScript 支持的功能来丰富 HTML 的方法。该库提供了函数和结构，从而更加容易地使用 Go 中的 HTMX。

_Angelo Fallaria_ 


[NutsDB 1.0：纯 Go 编写的持久键值存储](https://golangweekly.com/link/148266/web "nutsdb.github.io") —— 一种简单、快速、可嵌入的 K/V 存储，支持完全可序列化的事务和数据结构，例如列表、集合和排序集。[GitHub 仓库](https://golangweekly.com/link/148267/web)。

_nutsdb_ 


[SDNS：轻量级递归 DNS 服务器](https://golangweekly.com/link/148268/web "sdns.dev") —— 用 Go 编写的 DNS 服务器，专注于隐私。包括 DNSSEC 支持。[GitHub 仓库](https://golangweekly.com/link/148269/web)。

_Yasar Alev_ 

   
--- 
📰 分类广告

💻 Hired 让求职变得简单 - 公司不再需要追逐招聘人员，而是预先向您提供薪资详细信息。[立即创建免费个人资料](https://golangweekly.com/link/148270/web)。

🐘 Postgres 用户？请看看 [Postgres Weekly](https://golangweekly.com/link/148271/web)，我们每周也会汇总 Postgres 领域的最新动态。。

---  


[cast 1.6.0](https://golangweekly.com/link/148272/web) - 安全且轻松地从一种类型转换为另一种类型。

[mpb 8.7](https://golangweekly.com/link/148273/web) - 在终端应用中渲染进度条。

[fq 0.9](https://golangweekly.com/link/148274/web) - 想象一下是 `jq`，但是二进制格式的。

[glog 1.2](https://golangweekly.com/link/148275/web) - Go 的分级执行日志。

[Mockery 2.38](https://golangweekly.com/link/148276/web) - 模拟代码自动生成器。

[Afero 1.11](https://golangweekly.com/link/148277/web) - 文件系统抽象系统。
