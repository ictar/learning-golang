原文：[Golang Weekly Issue #482](https://golangweekly.com/issues/482)

---

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/luno0yils24gbtdfhxnb.jpg)](https://golangweekly.com/link/147069/web)   
  

[Awesome Go：数以千计的分类 Go 资源](https://golangweekly.com/link/147069/web "awesome-go.com") —— 
一个持续频繁更新的有用资源，非常值得在 Go 新闻出人意料的安静的一周里重新访问下。 😬 如果你有一个 Go 项目，你也可以[贡献你自己的项目](https://golangweekly.com/link/147070/web)到列表中。

_Awesome Go_
  

▶ 👻 [万圣节 Go Gopher Lofi 音乐合集](https://golangweekly.com/link/147072/web "www.youtube.com") —— 如果您想要一些万圣节主题音乐，那么 Ardan Labs 在这段 YouTube 视频中提供了相关内容。[▶️ 去年](https://golangweekly.com/link/147073/web)他们也这么做过。

_Ardan Labs_


[![](https://copm.s3.amazonaws.com/1ea6f5f1.png)](https://golangweekly.com/link/147068/web) 

[去！通过 Ardan Labs Consulting 释放您的技术潜力](https://golangweekly.com/link/147068/web "www.ardanlabs.com") —— 因技能差距、开发速度或复杂的技术挑战而苦苦挣扎吗？Ardan Labs 专注于 Go、Rust、Docker 和 K8s，以加速您的软件开发、优化架构和管理技术债务。让我们为您的团队加油！

_Ardan Labs Consulting sponsor_



[AWS SDK for Go 现已符合 Go 的发布政策](https://golangweekly.com/link/147071/web "aws.amazon.com") —— 从*今天*起，AWS SDK for Go 将符合 Go 的发布政策，支持 Go 两个最新的通用版本，并额外提供六个月的支持。这也意味着 AWS SDK for Go v1/v2 从今天起终止对 Go 1.5 至 1.18 的支持。

_Aaron Todd (AWS)_

  
[使用 Tqla 进行动态 SQL 模板化](https://golangweekly.com/link/147074/web "blog.vaunt.dev") —— Tqla（发音为*tequila*）提供了一种使用 `text/template` 生成动态 SQL 查询的替代方法，同时防止 SQL 注入。

_Ethan Lewis (Vaunt)_
  

[标准库中的同步构造](https://golangweekly.com/link/147075/web "hjr265.me") —— 讨论的构造是除 `sync.Mutex` 之外的构造，包括 Counter、Map、Once 和 Condition，每个构造都适合常见的同步用例。

_Mahmud Ridwan_


▶ [指针就是 Go 的西兰花 🥦](https://golangweekly.com/link/147076/web "www.youtube.com") —— 演讲者说这是她第一次技术演讲，但她自信地讨论了有关指针的主题，以及在短短七分钟内如何了解指针在 Go 中的工作原理。

_Beth Knight_


[Go 有子类型吗？](https://golangweekly.com/link/147094/web)   

_Bob Nystrom_


--- 

📰 分类广告

🚀 [开源数据可观察性](https://golangweekly.com/link/147077/web)：Streamdal 就是对你的数据进行 `tail -f` 操作，提供 UI，并通过我们的 SDK 支持 Golang。


📢 查看您是否有资格获得 2,400 美元的 Temporal Cloud 积分以及通过我们新的[Temporal Cloud for Startups 计划](https://golangweekly.com/link/147078/web)获得支持和服务。


💻 Hired 让求职变得简单 - 公司不再需要追逐招聘人员，而是预先向您提供薪资详细信息。[立即创建免费个人资料](https://golangweekly.com/link/147079/web)。
  
---

## 🛠 代码和工具
  

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/oxo5mkovcchnsbq9nvyi.jpg)](https://golangweekly.com/link/147080/web)  
  

[Warrant 1.0：受 _Google Zanzibar_ 启发的集中式授权服务](https://golangweekly.com/link/147080/web "github.com") —— An open-source, self hostable implementation of [Google Zanzibar](https://golangweekly.com/link/147095/web)（一项全球分布式服务，为 Google 的许多产品和服务提供授权。） 的一个开源的自托管实现。

_Warrant_


[TtlMap：条目具有时效的 Map](https://golangweekly.com/link/147081/web "github.com") —— 条目具有 TTL（生存时间）的映射（map），条目会在特定时间间隔后过期并被删除。

_John Taylor_
  

[Bluetuith：基于 TUI 的蓝牙管理器，适用于 Linux](https://golangweekly.com/link/147082/web "darkhz.github.io") —— 用 Go 构建，仅限 Linux，旨在替代 Blueman 等蓝牙管理器。

_darkhz_
  

[Beelzebub：一个安全的“低代码”蜜罐框架](https://golangweekly.com/link/147083/web "beelzebub-honeypot.com") —— 它可以伪装成机器人正在寻找的众多常见目标，例如 WordPress 安装或 httpd 的易受攻击版本，所有这些都通过 YAML 文件定义。[GitHub 存储库。](https://golangweekly.com/link/147084/web)

_Mario Candela_

[Spago 1.1](https://golangweekly.com/link/147092/web) - 用于自然语言处理工作的纯 Go ML 库。

[Goldmark 1.6](https://golangweekly.com/link/147085/web) - 符合 CommonMark 0.30 的 Markdown 解析器。

[sqlc 1.23](https://golangweekly.com/link/147096/web) - 从 SQL 生成类型安全代码。现在支持 `pgvector`！

[Gobot 2.2](https://golangweekly.com/link/147087/web) - 机器人、无人机和物联网框架。

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/c9xwsfcu2zjuxjmk8xve.jpg)](https://golangweekly.com/link/147088/web)   


[Walk 1.7](https://golangweekly.com/link/147088/web) - 终端文件管理器（上图）。

[Wish 1.2](https://golangweekly.com/link/147089/web) - Charm 的工具，用于制作 Go 驱动的 SSH 应用程序。

[Chroma 2.10](https://golangweekly.com/link/147090/web) - 通用语法荧光笔，纯 Go 编写。

[Toxiproxy 2.7](https://golangweekly.com/link/147086/web) - 用于模拟混沌网络条件的代理。


> _💬 引用_
> 
> “不要在了解问题之前就开始考虑解决方案。你的目标应该是主要在问题领域内解决问题，而不是在解决方案领域内。”
> ___  
> Oz Nova
