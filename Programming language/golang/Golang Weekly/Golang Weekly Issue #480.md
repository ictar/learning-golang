原文：[Golang Weekly Issue #480](https://golangweekly.com/issues/480)

——-

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/qav2a2wupfp5rz4lzfwu.jpg)](https://golangweekly.com/link/146357/web)  
  

[Go 1.22 中更好的 HTTP 服务器路由](https://golangweekly.com/link/146357/web "eli.thegreenplace.net") —— 早在 5 月份，我们就链接到了一个有关增强 `http.ServeMux` 的路由功能的[讨论](https://golangweekly.com/link/146358/web)。在 7 月份的时候，它[成为了一项提案](https://golangweekly.com/link/146359/web)，现在 Eli Bendersky 给出了新的多路复用器可以提供的功能的实际示例，并将其与 `gorilla/mux` 进行了比较。Hacker News 也举办了[一场扩展的讨论](https://golangweekly.com/link/146360/web)，讨论了多条路由匹配时发生 panic 的利弊，以及使用魔术字符串与特定于动词的方法。Go 1.22 预计将于 2024 年初发布，因此在此之前预计会看到更多有关此主题的信息。

_Eli Bendersky_


[来吧，专家为您服务](https://golangweekly.com/link/146356/web "www.ardanlabs.com") —— 您是否需要帮助填补技能差距、加快开发速度并使用 Go、Docker、K8s、Terraform 和 Rust 创建高性能软件？我们将帮助您最大化您的架构、结构、技术债务和人力资本。

_Ardan Labs Consulting sponsor_


[为什么说 Gokrazy 真的很酷](https://golangweekly.com/link/146361/web "xeiaso.net") —— 您知道有一个针对 Raspberry Pi 的最小化且以 Go 为中心的 Linux 实现吗？[gokrazy](https://golangweekly.com/link/146362/web) 允许您将 Go 程序部署为此类设备上的“设备”（想想像 Alpine Linux 这样的最小的东西，但仅适用于 Go）。

_Xe Iaso_
  

[重试：一次常见重试方法的交互式探索](https://golangweekly.com/link/146363/web "encore.dev") —— 一篇包含视觉示例的精彩文章，探索不同方式的重试请求，以展示出为什么某些方法比其他方法更好，并在一些实现理想策略的 Go 代码中得出结论。

_Sam Rose_


🚨 [Go 1.21.3 和 1.20.10 发布](https://golangweekly.com/link/146364/web)。主要修复是针对[一个广泛讨论的 HTTP/2 漏洞](https://golangweekly.com/link/146365/web)（恶意客户端可以轻松淹没 HTTP/2 服务器）。


👥 Garrit Franke 演示了[如何组织多个 git 身份](https://golangweekly.com/link/146366/web)，这或许是保持工作和个人使用分离的一种方式。


🎤 _Go Time_ 播客讨论了[他们在 GopherCon 的经历](https://golangweekly.com/link/146367/web)，这个是上个月在圣地亚哥举行的 GopherCon。


▶️ 受欢迎的 Go YouTuber Anthony GG 解释了[他是如何构建他新的 Go 项目的](https://golangweekly.com/link/146368/web)。


🏢 Go 团队一直在维护[一系列 Go 案例学习](https://golangweekly.com/link/146369/web)，涵盖了诸如American Express、Dropbox、Cloudflare 和 Uber 等公司。

📗 [Go101.org 书籍](https://golangweekly.com/link/146370/web) 已全部更新至 Go 1.21 标准。

## 🛠 代码和工具 

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/xqtda9aevpj72lhin88c.jpg)](https://golangweekly.com/link/146371/web)  

[Lip Gloss 0.9：“我的天，表格是如何转变的”](https://golangweekly.com/link/146371/web "github.com") —— _Lip Gloss_ 提供了一个具有“流畅”风格的 API，用于对程序的文本输出进行样式化，并且刚刚添加了对绘制 _表格_ 的支持（见上图）。发布的帖子（链接）包含一个创建您自己的表的快速教程，它很有用。[GitHub 存储库](https://golangweekly.com/link/146372/web)。

_Charm_

[Go OpenAI 1.16.0：在 Go 中使用 OpenAI 的 API](https://golangweekly.com/link/146373/web "github.com") —— 提供对 ChatGPT、GPT-3、GPT-4、DALL-E、Whisper 和 OpenAI 嵌入（Embeddings）的访问。他们在自述文件中[维护了一系列很好的使用示例](https://golangweekly.com/link/146374/web)，并且刚刚添加了一个展示如何在嵌入（Embeddings）之间进行语义相似性比较的示例。

_Sasha Baranov_

▶ [Gosh：在命令行编写 Go](https://golangweekly.com/link/146377/web "www.youtube.com") —— _“我编写了一个名为 [gosh](https://golangweekly.com/link/146378/web) 的工具，我认为它填补了 Go 工具包中的一个缺失。许多语言都提供了一种编写代码并直接在命令行执行它的方法，但是 Go 没有，所以我写了 gosh。”_

_Nick Wells_

[与 OpenTelemetry Spans 关联的 `log/slog` `Handler` 接口](https://golangweekly.com/link/146379/web "github.com") —— 如果您同时使用 `slog` 和 OpenTelemetry，那这个可能适合您。它将各种信息添加到日志记录中，以帮助与 OTel 的 [spans](https://golangweekly.com/link/146380/web) 进行关联。

_Remy Chantenay_
  

[TruffleHog：到处找寻泄漏的凭证](https://golangweekly.com/link/146381/web "github.com") —— 就像猪会寻找松露一样，这个由 Go 驱动的“猪”将通过 git 存储库、S3、文件系统及其他地方来寻找那些你不愿意暴露在外的秘密和其他类似的宝贝儿。

_Truffle Security_

[Algernon：一个独立的小型纯 Go Web 服务器](https://golangweekly.com/link/146382/web "github.com") —— 支持 Lua、Markdown、HTTP/2、QUIC、Redis、MySQL 和 Postgres。

_Alexander F. Rødseth_

  
[Air 1.47](https://golangweekly.com/link/146383/web) - Go 应用程序的实时重新加载。（甚至还有[▶️ 一个截屏视频，展示了](https://golangweekly.com/link/146384/web)如何使用它以及使用它可以拥有的好处。）

[pgroll 0.3](https://golangweekly.com/link/146385/web) - 勇于探索 Postgres 的零停机、可逆模式迁移。

[GoCV 0.35](https://golangweekly.com/link/146386/web) - OpenCV 4 视觉库的绑定。

[Fiber 2.50.0](https://golangweekly.com/link/146387/web) - 受 Express.js 启发的 Go Web 框架。

[Mockery 2.36](https://golangweekly.com/link/146388/web) - 为接口生成模拟。

[Slogor](https://golangweekly.com/link/146389/web) - 一位坚持己见的多彩 `slog` 处理器。

[DynamiteMC](https://golangweekly.com/link/146390/web) - 实验性 Minecraft 服务器实现。

[官方 Go Stripe 库 76.0](https://golangweekly.com/link/146391/web)
