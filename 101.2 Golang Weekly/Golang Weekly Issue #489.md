原文：[Golang Weekly Issue #489](https://golangweekly.com/issues/489)

---

![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/zjoj9ab7g6p3ldymbuzx.jpg)  


2023 年最棒的 Go Newsletter！

欢迎来到 2023 年最后一期！我们将于 1 月 9 日星期二回归，但在此之前，我们想回顾一下今年最受欢迎的东西，再加上一些额外的福利，希望您度过一个愉快的假期！  

_Peter Cooper, your editor_

 

[Rust vs Go：一次实际的比较](https://golangweekly.com/link/149166/web "www.shuttle.rs") —— 尽管它们之间存在大量差异，Rust 和 Go 仍然经常被拿来比较，而且经常引起争议。某个主要是 Rust 开发人员的人通过专门关注用两种语言构建的小型 HTTP 服务的实用性来探讨该主题。并不完美，但总体来说还算公平。

_Matthias Endler (Shuttle)_ 

比较 Go 和 Rust 的文章今年特别受欢迎：
  * Krater 团队解释了[为什么他们使用 Go 而不是 Rust](https://golangweekly.com/link/149167/web)来构建桌面应用程序。
  * Vercel 写了一篇文章来解释他们为什么[将 Turborepo 工具从 Go 迁移到 Rust](https://golangweekly.com/link/149168/web)。
  * John Arundel [从自己的角度对 Rust 和 Go 进行了比较](https://golangweekly.com/link/149169/web)。


[![](https://copm.s3.amazonaws.com/1ea6f5f1.png)](https://golangweekly.com/link/149165/web) 

[冲！通过 Ardan Labs Consulting 释放您的技术潜力](https://golangweekly.com/link/149165/web "www.ardanlabs.com") —— 你还在因技能差距、开发速度或复杂的技术挑战而苦苦挣扎？Ardan Labs 专注于 Go、Rust、Docker 和 K8s，以加速您的软件开发、优化架构和管理技术债务。让我们为您的团队助力！

_Ardan Labs Consulting sponsor_


[实验项目模版](https://golangweekly.com/link/149170/web "go.dev") —— 启动项目的主题仍然还是 Go 团队推出的新 _实验_ 工具，用于根据预定义模板在 Go 中创建新项目——他们将在 2024 年继续开发这一工具。

_Cameron Balahan (Go Team)_ 
  

[在 2023 年，要如何启动一个 Go 项目](https://golangweekly.com/link/149171/web "boyter.org") —— 同一位作者在 2018 年写了一篇类似的文章，但此后事情发生了很大的变化。虽然这是针对较新的 Gophers，但老玩家可能会从中受益。_“较早的指南会提到设置你的 $GOPATH。到了 2023 年，你可以轻松地忽略这一点。”_

_Ben E. C. Boyter_ 


[组织 Go 模块](https://golangweekly.com/link/149172/web "go.dev") —— Go 开发人员通常会考虑如何在一个典型 Go 项目中组织文件和目录（事实上，最新的 Go 调查将此作为 Go 开发人员的主要关注点）。虽然Go 项目倾向于让社区找出该领域的最佳实践，但很高兴看到他们发布了更官方的东西。

_The Go Team_ 


[Go 1.20 新增功能](https://golangweekly.com/link/149173/web "blog.carlmjohnson.net") —— 这是由三部分组成的系列中的第一部分，讨论了对核心语言功能的调整，例如接口、泛型、`unsafe` 以及新的切片到数组转换技术。Carlana 也兑现了编写整个系列的承诺，[第二部分](https://golangweekly.com/link/149174/web)涵盖了主要的标准库更改，[第三部分](https://golangweekly.com/link/149175/web)涵盖了更小的更改。

_Carlana Johnson_  


> **其他新的或者即将到来的想法：** 除了在今年发布的 Go 1.20 和 1.21 中引入的新功能之外，还讨论了各种提案和想法，包括[向 Go 添加协程](https://golangweekly.com/link/149176/web)、[重大 `encoding/json` 更新](https://golangweekly.com/link/149177/web)、[Go 1.22 中 `for` 循环](https://golangweekly.com/link/149178/web)的未来、[增强 `ServeMux` 路由](https://golangweekly.com/link/149179/web)以及不太重要的，新增的[ Go 源文件的一个奇怪的备用文件扩展名](https://golangweekly.com/link/149180/web)。  
  

[用 Go 实现分布式键值存储](https://golangweekly.com/link/149181/web "notes.eatonphil.com") —— Phil 花了几个月的时间来熟悉 [Raft 共识算法](https://golangweekly.com/link/149182/web)，并深入研究如何将它与 Go 一起作为分布式键值存储的核心基础。

_Phil Eaton_ 


[最小的 Go 二进制文件 - 5KB?](https://golangweekly.com/link/149183/web "totallygamerjet.hashnode.dev") —— 我一开始觉得这看起来很蠢，尤其是它的日期是 4 月 1 日，但作者有一个有趣的用例：他们想使用 Go 的汇编器，但 _不想_ 使用Go 的运行时。

_Over Engineered_ 


🛠 代码和工具 


[Conc：Go 中更好的结构化并发](https://golangweekly.com/link/149184/web "github.com") —— Go 的并发很好，但 Conc 的目标是通过提供各种概念（思考池、并发映射和迭代以及 panic 捕获）和技术的抽象来使其更加安全和简单，从而使得 goroutine 泄漏和未处理 panic 成为过去。

_Sourcegraph_  
  

[从未让我们失望的 Go 库：您需要了解的 22 个库](https://golangweekly.com/link/149185/web "threedots.tech") —— 虽然还有其他列表（例如[Awesome Go](https://golangweekly.com/link/149186/web)），但这个列表经过了严格策划，仅包含他们在生产中使用过的那些库。

_Robert Laszczak (Three Dots Labs)_ 


[害怕结对编程吗？Tuple 可能会改变您的想法](https://golangweekly.com/link/149187/web) —— .我们不会成为您第一次配对约会的第三者！我们让开，这样配对魔法就发生了。14 天免费试用。

_Tuple sponsor_


[Service Weaver：谷歌用于编写分布式 Go 应用的框架](https://golangweekly.com/link/149188/web "opensource.googleblog.com") —— 来自 Google 的一个框架，可让您_“以单体模块的方式编写您的 (Go) 应用程序，并将其部署为一组微服务”，从而获得两全其美的效果，即：_“单体应用的开发速度，以及微服务的可扩展性、安全性和容错能力。”_ 如果您喜欢更具技术性且不那么功利的东西，Robert Grandl [在这里有一个快速介绍](https://golangweekly.com/link/149189/web)。

_Google Open Source_ 


[River：用于 Go 和 Postgres 的快速健壮的 Job 队列](https://golangweekly.com/link/149190/web "brandur.org") —— River 于上个月推出，是一个开源作业队列，_“用于构建快速、密封的应用程序”_，它用 Go 编写并利用泛型。

_Brandur Leach_


[betteralign：让你的 Go 程序使用更少的内存……可能吧](https://golangweekly.com/link/149191/web "github.com") —— 一种工具，用于检测一种结构，这种结构的字段如果已排序会使用更少的内存，然后可以选择对这些字段进行排序。

_Dinko Korunic_ 


[NilAway：实用 Nil Panic 检测](https://golangweekly.com/link/149192/web "www.uber.com") —— Nil panic 是一个常见且难以检测的问题，但 Uber 创建了一个静态分析工具（基于 Java 世界的类似工具），该工具易于设置并与构建工具集成。

_Uber_ 


---  

📰 分类广告


💻 Hired 让求职变得简单 - 公司不再需要追逐招聘人员，而是预先向您提供薪资详细信息。[立即创建免费个人资料](https://golangweekly.com/link/149193/web)。


📅 [GopherCon Europe](https://golangweekly.com/link/149197/web) 将于明年 2 月 _与_ 6 月分别在希腊雅典和德国柏林举行。Ardan Labs 还将在明年 2 月份在希腊举行的活动中[举办一些 Go 研讨会](https://golangweekly.com/link/149198/web)。

🐘 PostgreSQL 用户？请访问我们的姐妹通讯 [Postgres Weekly](https://golangweekly.com/link/149194/web)，每周了解 Postgres 世界的最新动态。

---

📺 热门视频

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/udezqn3ypbzknpzaelrs.jpg)](https://golangweekly.com/link/149195/web)  

▶ [我讨厌 Go 的十件事](https://golangweekly.com/link/149195/web "www.youtube.com") —— 以负面的语气结束本期似乎很奇怪，但这是点击次数最多的视频！出于显而易见的原因，这类事情总是很受欢迎，但乔纳森显然在他的选择中加入了一些思考，并详细演示了这些问题。他还承认自己是 Go 的忠实粉丝，这一点在他的[▶️ 我喜欢 Go 编程语言的 10 个理由](https://golangweekly.com/link/149196/web)中得到了证明！他频道的其余部分也值得一看。

_Jonathan Hall (Boldly Go)_