原文：[Golang Weekly Issue #475](https://golangweekly.com/issues/475)

---

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/ehdmc5luqjaalihbbm3n.jpg)](https://golangweekly.com/link/144735/web)
---  

[Go 1.21 中的配置引导优化（Profile-Guided Optimization (PGO)）](https://golangweekly.com/link/144735/web "go.dev") -- [PGO](https://golangweekly.com/link/144736/web) 最近已经成为了博文的一个热门话题，但这里有一个用它来改进代码的最接近官方使用的示例，包括更深入探讨该过程实现的两项主要优化。如果没有其他办法让你想去尝试下，请考虑一下：_“在 Go 1.21 中，启用 PGO 后，工作负载通常会提高 2% 到 7% 的 CPU 使用率。”_

_Michael Pratt (The Go Team)_

> 如果你喜欢该主题的_另一篇_文章，Landon Clipp 还提供了[另一篇关于 PGO 的使用介绍](https://golangweekly.com/link/144764/web)，这篇文章提供了更多“面向用户的视角”。
  

[Go! Experts at Your Service](https://golangweekly.com/link/144734/web "www.ardanlabs.com") —— 你是否需要有人帮你填补技能差距、加快开发速度并使用Go、Docker、K8s、Terraform 和 Rust 创建高性能软件？我们将帮助您最大化您的架构、结构、技术债务和人力成本。

_Ardan Labs Consulting sponsor_


[为不断增长的 Go 生态扩展 `gopls`](https://golangweekly.com/link/144765/web "go.dev") —— Go 官方博客为我们提供了本周值得深入研究的_两篇_文章。这次，焦点集中在被各种 IDE 用来增强对 Go 对支持的 [`gopls` Go 语言服务器](https://golangweekly.com/link/144766/web)。这篇文章着眼于最近的一些进展和加速，此外，如果你是 `gopls` 用户，那么邀请你参与[这项调查](https://golangweekly.com/link/144767/web)。

_Robert Findley 和 Alan Donovan_

 
_快速了解：_

* [Go 1.21.1 和 1.20.8 已发布。](https://golangweekly.com/link/144768/web) 带有一些安全修复的次要版本。

* 我想趁此机会再次链接到 Maria Letta 的 [The Free Gophers Pack](https://golangweekly.com/link/144737/web) —— 这是一组（基于 Renee French 原作）有用的 Go gopher 插图，我们在今天的主图中用了它。
  
* 谈到 Go gopher，早在 2016，Renee French 就进行了一场 [▶️ 涉及 Go gopher 的有趣的演讲](https://golangweekly.com/link/144738/web)，涉及到它是如何组合在一起的，以及 Renee 最喜欢的一些改编版本。

* 如果你必须写些 C++ 代码，[coost](https://golangweekly.com/link/144739/web) 值得一看。它是一个小型 Boost 式库，使 C++ 更容易编写并且具有更多的 Go 风格（包括 Go 风格的协程）。

* _IEEE Spectrum_ 发布了[2023 年“顶级编程语言”名单](https://golangweekly.com/link/144740/web)，Go 排在第八位，表现相当不错。（注：前三名依次是 Python，Java，C++）

* 📅 🤖 几周后，我将参加在旧金山举行的 [人工智能工程师峰会（AI Engineer Summit）](https://golangweekly.com/link/144769/web) —— 如果你热衷于在软件开发过程中使用 AI 和机器学习工具及技术这一快速发展的领域，那么请查看下。即使你无法参加，也可以或许免费的远程门票进行远程观看。

▶ [“这将让每个人都了解 Go 接口”](https://golangweekly.com/link/144741/web "www.youtube.com") —— 对于任何经历困惑的人，Anthony GG 以其独特的风格承担了这项（解惑的）任务。

_Anthony GG_

 
[我在软件开发过程中改变了想法](https://golangweekly.com/link/144743/web "henrikwarne.com") —— _“任何不经常改变想法的人都大大低估了我们生活的世界的复杂性。”_ - Jeff Bezos

_Henrik Warne_


[免费课程：Temporal 102 with Go](https://golangweekly.com/link/144742/web "t.mp") —— 这是我们的必要的 101 课程的实用后续课程，通过它，了解如何利用 Temporal 中的完整开发周期。

_Temporal Technologies sponsor_


[使用 Tailscale 进行内部工具身份验证](https://golangweekly.com/link/144770/web)  
_Khash Sajadi_

[使用 Go，为 Telegraf 实现一个 ClickHouse 输出插件](https://golangweekly.com/link/144744/web)  
_David Wołosowicz_

[关于 Go 中常见文件操作的实用之旅](https://golangweekly.com/link/144745/web)  
_Adebayo Adams_

 
## 🛠 代码和工具

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/xowkecrguksufntfsd0g.jpg)](https://golangweekly.com/link/144746/web) 

[Wails 2.6：使用 Go、JS 和 CSS 创建桌面应用](https://golangweekly.com/link/144746/web "wails.io") —— 你曾羡慕过 JavaScript 开发人员能够使用 Electron 来构建桌面应用吗？Wails 给 Go 带来了类似的选择。v2 版本已经成熟稳定了，但是 [Wails v3 即将推出](https://golangweekly.com/link/144747/web)，并且有望成为一次重大更新。[GitHub repo.](https://golangweekly.com/link/144748/web)

_Lea Anthony_

[Goxygen 0.7：为 JS 项目快速生成 Go 后端](https://golangweekly.com/link/144749/web "github.com") —— 该工具可以在前端使用 Angular、React 或 Vue 建立一个新的基于 Go 的项目，并使用 Docker 和 Docker Compose 文件使其全部正常工作。v0.7 版本引入了对 Go 1.21 的支持。

_Sasha Shpota_
 
[[博客] 如何破解 Kubernetes（以及如何保护它）](https://golangweekly.com/link/144750/web "goteleport.com") —— 本综述涵盖了集群可能受到攻击的七种主要方式以及相应的对策。

_Teleport | goteleport.com sponsor_

[Participle 2.1：一个简单的解析器包](https://golangweekly.com/link/144751/web "github.com") —— 旨在提供一种简单且惯用的方式来在 Go 中定义解析器，使用熟悉的方法，即使用结构体字段标签来定义语法，因此，如果你之前使用过 `encoding/json`，那么你就已经上手了。

_Alec Thomas_


[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/bvx2zuwg3zxmabc9gxmn.jpg)](https://golangweekly.com/link/144752/web) 


[Lip Gloss 0.8：良好终端布局的样式定义](https://golangweekly.com/link/144752/web "github.com") —— 提供“流畅”风格的 API，以有吸引力的方式对程序的文本输出进行样式化，正如所期望的 _Charm_ 项目一样。

_Charm_


[Spotify 2.4.0：Spotify Web API 的 Go 封装器](https://golangweekly.com/link/144753/web "github.com") —— 有很多[示例](https://golangweekly.com/link/144754/web)，但你可以执行所有明显的操作，例如搜索曲目、播放列表和控制播放曲目。

_Zac Bergquist_

[Enmime 1.0：MIME 解码和编码包](https://golangweekly.com/link/144755/web "github.com") —— 专注于生成和解析 MIME 编码的电子邮件。

_James Hillyerd_

 
_快速发布：_
* [Chroma 2.9](https://golangweekly.com/link/144756/web)  
↳ 纯 Go 通用语法高亮器。

* [Lingua 1.4](https://golangweekly.com/link/144757/web)  
↳ 自然语言检测库。

* [sqlc 1.21](https://golangweekly.com/link/144758/web)  
↳ 从 SQL 生成类型安全代码。

* [Benthos 4.21](https://golangweekly.com/link/144759/web)  
↳ 神奇的流处理使操作变得平凡。

* [OpenAPI Client and Server Code Generator 1.15](https://golangweekly.com/link/144760/web)  
↳ 根据 OpenAPI 3 规范生成样板。

* [go-github 55.0](https://golangweekly.com/link/144761/web)  
↳ GitHub v3 API 客户端库。

* [lemonsqueezy-go 1.0.3](https://golangweekly.com/link/144773/web)  
↳ [Lemon Squeezy](https://golangweekly.com/link/144774/web) SaaS 平台的 Go API 客户端。


### 工作  
  
[通过 Hired 找工作](https://golangweekly.com/link/144762/web) —— Hired 使找工作变得容易 —— 公司不再追逐招聘人员，而是预先向你提供薪资详细信息。立即创建免费的个人资料吧。

_Hired_
