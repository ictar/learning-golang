原文：[Golang Weekly Issue #478](https://golangweekly.com/issues/478)

——-

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/w6bzpvbuo9o4ffcuaxbb.jpg)](https://golangweekly.com/link/145855/web)  

  
[Rust vs Go：实际比较](https://golangweekly.com/link/145855/web "www.shuttle.rs") —— 尽管 Rust 和 Go 有很多差异，但人们[经常](https://golangweekly.com/link/145865/web)对它们[进行](https://golangweekly.com/link/145864/web)[比较](https://golangweekly.com/link/145866/web)。某个主要是 Rust 开发的开发人员的人再次尝试进行比较，重点关注两者构建 HTTP 服务的实用性。并不完美，但仍然相当公平。

_Matthias Endler (Shuttle)_


[解构类型参数](https://golangweekly.com/link/145856/web "go.dev") —— 如果 `func Clone[S ~[]E, E any](s S) S ` 这个定义让你不寒而栗，别怕，Go 官方博客上的这篇文章会将其中所涉及的想法分解为更容易理解的形式。

_Ian Lance Taylor_


[![](https://copm.s3.amazonaws.com/95876832.png)](https://golangweekly.com/link/145828/web) 

[最后召集：2023年学习事件驱动的 Go 的最后一次机会！](https://golangweekly.com/link/145828/web "threedots.tech") —— 如今的就业市场对于普通开发人员来说很艰难。不要错过这个提高您的技能并在其他候选人中脱颖而出的机会。通过构建现实生活中的事件驱动系统来提升您的职业水平。通过编写代码来学习——无需观看视频。

_Three Dots Labs sponsor_

  
_快速了解：_

* 📅 Rob Pike 已被确认担任 [GopherConAU 2023](https://golangweekly.com/link/145829/web) 上的演讲者了，该会议将于 11 月 8 日至 10 日在澳大利亚悉尼举行。

* [pcz](https://golangweekly.com/link/145857/web) 是一个好奇的、实验性的 Go 的“重新想象”，它使用（未经修改的）官方工具链，但提供自己的 stdlib 并针对替代用例。它甚至还有一个 Web SDK，可以实现[类似的功能](https://golangweekly.com/link/145858/web)。

* 在 Reddit 上，InfluxDB 的创建者 Paul Dix [分享了一些关于 InfluxDB 为什么从 Go 转向 Rust 的见解。](https://golangweekly.com/link/145859/web)

* 👾 [GoBC 1.0](https://golangweekly.com/link/145860/web) 是完全用 Go 写的 Game Boy 模拟器。

* HashiCorp 的 Mitchell Hashimoto 提出了[围绕变更集合重新调整 GitHub pull 请求](https://golangweekly.com/link/145830/web)的案例。



[通过测试学习 Go：无需使用 Mocks](https://golangweekly.com/link/145838/web "quii.gitbook.io") —— 备受欢迎的[通过测试学习 Go](https://golangweekly.com/link/145839/web)资源增加了一个新章节，介绍了各种伪造依赖项的方法、每种方法的优缺点以及示例。好的测试依赖于好的设计，因此这里的范围远远超出了测试桩和虚假数据。

_Learn Go with Tests_

  

[一个开发者的首选 Go 栈](https://golangweekly.com/link/145840/web "jtarchie.com") —— 一位开发人员的首选库集合。想法可能会有所不同，但分享是件好事。

_JT Archie_


[枚举的编译时安全](https://golangweekly.com/link/145842/web "vladimir.varank.in") —— 一种提高 Go 中枚举的编译时安全的可能方法。_优雅吗？_ 尽量吧。

_Vladimir Varankin_
  

[使用 AWS Lambda 和 Go，生成待命日历](https://golangweekly.com/link/145861/web "www.dolthub.com") —— 一个有趣的业余项目，展示了如何使用 AWS SDK 与 Go 和各种 AWS 服务来提供方便的 ICS 源，同时演示 Go 是如何成为构建无服务器函数的良好选择。

_Jason Fulghum_


[感受超过 420,000 名队友和工作的力量#LikeABosch](https://golangweekly.com/link/145854/web "www.bosch.com") —— 是我们的员工让我们变得卓越。而我们的成功就是您的成功。让我们一起庆祝吧。[了解更多](https://golangweekly.com/link/145854/web)。

_Bosch sponsor_


[使用去中心化身份令牌保护 Go API](https://golangweekly.com/link/145843/web)

_Robert Kimani (The New Stack)_ 

  
## 🛠 代码和工具
  
[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/btymans62elszce87rzf.jpg)](https://golangweekly.com/link/145862/web)  
  

[Giu 0.7：基于 _Dear ImGui_ 的 Go 跨平台 GUI](https://golangweekly.com/link/145862/web "github.com") —— 另一种创建 GUI 应用的方式。[Dear ImGui](https://golangweekly.com/link/145863/web) 是一个流行的 GUI 库（用于 C++），它生成针对高级用户的特殊 UI而不是典型的 UI。

_Allen Dang_

[go-jsonschema：根据 JSON 模式生成 Go 数据类型](https://golangweekly.com/link/145846/web "github.com") —— 生成与模式相对应的数据类型和结构，以及根据模式的验证规则验证输入 JSON 的反序列化代码。

_Claudio Beatrice_


📰 分类广告

--- 

📑 [Go SDK 开发者指南](https://golangweekly.com/link/145832/web)中了解 Temporal OSS 是如何为您的服务和应用程序提供持久执行的。

* * *

💻 Hired 使找工作变得容易 —— 公司不再追逐招聘人员，而是预先向你提供薪资详细信息。[立即创建免费的个人资料吧。](https://golangweekly.com/link/145833/web)


[Goph 1.4：原生 Go SSH 客户端](https://golangweekly.com/link/145847/web "github.com") —— 支持使用密码、私钥、带密码的密钥进行连接、进行文件传输等。

_Mohamed El Bahja_

[Gotify Server：用于 Web 应用程序的实时消息服务器](https://golangweekly.com/link/145848/web "gotify.net") —— 一种由 Go 驱动的服务，您可以自行托管，通过 REST API 发送消息并通过 WebSocket 将消息传递给最终用户，从而在 Web 应用程序中提供简化的实时消息传递。[GitHub repo.](https://golangweekly.com/link/145849/web)

_Gotify_

  
[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/gq7lhqkj4pasanmmldh8.jpg)](https://golangweekly.com/link/145850/web)  
  

[Repo Trends：查看和分析 GitHub 问题随时间变化的趋势](https://golangweekly.com/link/145850/web "www.repotrends.com") —— 一个设计精美的工具，可以显示给定 GitHub 存储库上未解决问题和拉取请求的可视化数据。例如，查看[ `golang/go` 的统计数据](https://golangweekly.com/link/145851/web)。

_Steve Sanders_

[Boxes 和 Glue：一个受 TeX 启发的排版库](https://golangweekly.com/link/145844/web "github.com") —— 一个 PDF 排版库，遵循 _“TeX 算法的精神”_ 将元素布局到页面上。[示例代码](https://golangweekly.com/link/145845/web)将帮助您理解其想法。

_speedata GmbH_

[Countdown 1.5：终端倒计时器](https://golangweekly.com/link/145852/web "github.com") —— 运行 `countdown 168h` 来倒计时下一份时事通讯，大概也许？:-)

_Anton Medvedev_

* [FastHTTP 1.50.0](https://golangweekly.com/link/145834/web) - 针对[特定场景](https://golangweekly.com/link/145835/web)的面向性能的 `net/http` 替代方案

* [Go Imagick 3.5](https://golangweekly.com/link/145836/web) - Go 绑定到 ImageMagick 的 MagickWand C API。

* [msgpack 5.4](https://golangweekly.com/link/145837/web) - Go 的 MessagePack 编码。