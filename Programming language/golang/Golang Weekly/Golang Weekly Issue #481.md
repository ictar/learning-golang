原文：[Golang Weekly Issue #481](https://golangweekly.com/issues/481)

---

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/sqcauotx321dimn9qpag.jpg)](https://golangweekly.com/link/146730/web)  


！[在 Go 中对错误使用感叹号？](https://golangweekly.com/link/146730/web "flak.tedunangst.com") —— 对 Go 中错误处理的冗长性批评是很常见的，但是有提出解决方案吗？很少有。Ted 提出了一种建议语法，即首先使用感叹号（“bangs”），然后使用脱字符号（`^`），可以通过一个叫做 [bango](https://golangweekly.com/link/146731/web) 的工具将其扩展为典型 `if err != nil` 方法。

_Ted Unangst_


> 🗣 即使按照 Hacker News 的标准，这篇文章也引发了[一场激烈的讨论](https://golangweekly.com/link/146732/web)，可以预见的是，意见涵盖了从“我喜欢 Go 的做法”到“使用 Rust 来代替”，以及关于 Go 为保持简单而做出的权衡的更多有启发性的观点。
  

[它们被称为切片，因为它们具有锋利的边缘 - 更多的 Go 陷阱](https://golangweekly.com/link/146733/web "www.dolthub.com") —— Nick最近一直在博客中讨论[各种 Go “陷阱”](https://golangweekly.com/link/146734/web)。这第三篇文章特别解决了一些关于切片的误解。

_Nick Tobey (Dolthub)_


[![](https://copm.s3.amazonaws.com/e6b1409e.png)](https://golangweekly.com/link/146729/web) 

[感受来自超过 420,000 名队友和工作的力量#LikeABosch](https://golangweekly.com/link/146729/web "www.bosch.com") —— 在 Bosch，我们通过高质量的技术和服务塑造未来，激励人们并改善他们的生活。正是我们的员工让我们变得卓越。我们的成功就是您的成功。让我们一起庆祝吧。[了解更多](https://golangweekly.com/link/146729/web)

_Bosch sponsor_


[使用差异模糊测试揭开一个 Go HTML 错误的面纱](https://golangweekly.com/link/146735/web "mionskowski.pl") —— 拿起一盏灯笼（或者只是一杯咖啡），进入使用 _差异模糊测试_ （将某个东西的两种实现相互进行比较）在标准库中查找错误的黑暗世界。Maciej 解释了如何报告错误，以及针对其（部分）解决方案涉及的拉扯。

_Maciej Mionskowski_


❓ Reddit 上有人问[Go 最好的 IDE 是什么？](https://golangweekly.com/link/146736/web) 与许多开源社区不同，商业选项（JetBrains _GoLand_）赢得了大部分人的喜爱。

🤔 Go 很“无聊”吗？有关系吗？Dev-tuber _ThePrimeagen_ [▶️ 对 Go 有一些话要说](https://golangweekly.com/link/146737/web)，他称赞它的简单性和实用性，但他说他不“喜欢”这门语言。


📗 [《x days of Go》](https://golangweekly.com/link/146739/web)的作者 Matt Boyle [接受了 Leanpub 关于这本书的简短采访](https://golangweekly.com/link/146740/web)。

🕸 [WebAssembly 2023 年状态结果](https://golangweekly.com/link/146764/web)已出炉，与 Go（成为 WASM 项目第四大流行语言）有一些关系。

📻 [RadioGoGo](https://golangweekly.com/link/146738/web) 是一款新的基于 Go 和 BubbleTea 的 TUI 网络广播应用程序。


## 🛠 代码和工具

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/slf0fo5ki0pybot4zbhe.jpg)](https://golangweekly.com/link/146741/web)  


[Viddy：`watch` 的现代替代方案](https://golangweekly.com/link/146741/web "github.com") —— 这里的想法是 `viddy` 可以定期执行命令，但您还可以获得“时间旅行”、有吸引力的输出和分页等细节。

_Takumasa Sakao_
  

[elem-go：创建和操作 HTML 元素的类型安全方法](https://golangweekly.com/link/146742/web "github.com") —— _“类似于 JSX，但适用于 Go，”_ 作者说，他还刚刚发布了[使用 htmx、Go Fiber 和 elem-go，构建一个计数器应用](https://golangweekly.com/link/146743/web)，一个展示其用例的教程。

_Chase Fleming_


[使用 OAuth 保护 Golang 应用程序](https://golangweekly.com/link/146744/web "fusionauth.io") —— 本教程将向您展示如何使用 OAuth 对 golang 应用程序中的用户进行身份验证。

_FusionAuth sponsor_


[Carbon：简单的语义化日期时间包](https://golangweekly.com/link/146745/web "github.com") —— 包含大量函数，使日期和时间的创建、解析和比较更容易、更流畅。

_gouguoyin_


[go-echarts 2.3：一个“可爱的”图表库](https://golangweekly.com/link/146746/web "github.com") —— 想要组合条形图、烛台图、饼图、折线图、热图或……更多？它使用 [Apache ECharts](https://golangweekly.com/link/146747/web)，因此比起渲染到图像，它更适合 Web 应用程序使用。

_go-echarts_
  

[go-quartz 0.8：简单的零依赖调度库](https://golangweekly.com/link/146748/web "github.com") —— 受到同名的Java [Quartz 调度程序](https://golangweekly.com/link/146749/web)的启发。

_Eugene R._


---


📰 分类广告


💻 Hired 让求职变得简单 - 公司不再需要追逐招聘人员，而是预先向您提供薪资详细信息。[立即创建免费个人资料](https://golangweekly.com/link/146750/web)。


🐘 如果您使用 Postgres，请查看我们的姊妹通讯，[Postgres Weekly](https://golangweekly.com/link/146751/web) – 下一期将于明天发布。

---
  

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/pif8ligg61fkapmxr2gu.jpg)](https://golangweekly.com/link/146752/web)  


[Circumflex / `clx`：基于终端的 _Hacker News_ 客户端](https://golangweekly.com/link/146752/web "github.com") —— 一个基于 [Bubble Tea](https://golangweekly.com/link/146753/web) 的终端客户端，对流行科技新闻进行聚合。

_Ben Sadeh_


[FSRouter：一个简单的文件系统路由器库](https://golangweekly.com/link/146754/web "github.com") —— 受到 Next.js 的新应用程序路由方法的启发。

_Antonio De Lucreziis_


[Bloom 3.6：用于实现 Bloom 过滤器的包](https://golangweekly.com/link/146755/web "github.com") —— Bloom 过滤器是内存高效的数据结构，用于确定具有已定义的潜在误报率的集合成员资格。

_Will Fitzgerald_
  

[oasdiff：检测 OpenAPI 规范中的重大更改](https://golangweekly.com/link/146756/web "github.com") —— 包括命令行工具和 Go 包，用于比较和检测 OpenAPI 规范中的重大更改。也许在 CI/CD 中很有用，可以确保在部署之前对修改进行审查和测试。

_Tufin_



[SCS 2.6](https://golangweekly.com/link/146757/web) - Web 应用程序的 HTTP 会话管理。


[fsnotify 1.7](https://golangweekly.com/link/146758/web) - Go 中的跨平台文件系统通知。

[lakeFS 1.0](https://golangweekly.com/link/146759/web) - 对你的数据湖的数据版本进行控制（或“git”）。

[golangci-lint 1.55](https://golangweekly.com/link/146761/web) - 运行 Go linter 的一种更快方法。


[Air 1.49](https://golangweekly.com/link/146760/web) - Go 应用程序的实时重新加载。


🕹 还有一个是为了好玩…… 
  
[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/xpiqmka4llzahajfuczl.jpg)](https://golangweekly.com/link/146762/web)    
  

[Gopher2600：Atari 2600 模拟器](https://golangweekly.com/link/146762/web "github.com") —— 一款功能基本齐全的模拟器，包括控制器支持、游戏录制，甚至 CRT 显示效果。我们几年前就链接到了它，但它持续更新，这对于这样一个项目来说是令人耳目一新且值得注意的。

_Stephen Illingworth_