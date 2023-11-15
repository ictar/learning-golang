原文：[Golang Weekly Issue #484](https://golangweekly.com/issues/484)

---


[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/uxvjyguivlozmfkfhd5e.jpg)](https://golangweekly.com/link/147678/web)  


[Go 的十四年](https://golangweekly.com/link/147678/web "go.dev") —— 回顾了 Go 富有成果的一年，其中包括引入[配置文件引导优化](https://golangweekly.com/link/147728/web) (PGO)、[增强的覆盖配置文件](https://golangweekly.com/link/147729/web)以改进测试、更好的工具链管理、第一个通用标准库、[slog](https://golangweekly.com/link/147730/web) 等等。希望明年同样美好。

_Russ Cox (The Go Team)_ 



💡 [嘿！嗬！让我们走着（Go）！](https://golangweekly.com/link/147680/web) 是 2009 年在 Google 开源博客上发布的第一篇文章，它宣布了该项目。  

 
[![](https://copm.s3.amazonaws.com/0770979a.png)](https://golangweekly.com/link/147677/web) 

[开发友好的 SDK 和 TF 提供商，以促进 API 采用率](https://golangweekly.com/link/147677/web "www.speakeasyapi.dev") —— 通过 Speakeasy 文档齐全、强类型、轻量级和可定制的 SDK、Terraform 提供程序和始终同步的文档，从而提供一流的开发人员体验。查看 Speakeasy SDK 与其他 OSS 生成器的比较。

_Speakeasy sponsor_
  

[Go 1.22+ 版本提供的 ServeMux 和 Chi Router 的对比](https://golangweekly.com/link/147682/web "www.calhoun.io") —— Go 今年最有趣的积极提案之一是围绕[增强内置 HTTP 路由器](https://golangweekly.com/link/147684/web)的 — 现在可能会出现在 Go 1.22 或更高版本中。人们开始权衡是使用它还是坚持现有的选项——乔恩计划在 _新_ 项目中尝试它，但对于那些已经使用 Chi 或 `gorilla/mux` 的项目，则不做更改。

_Jon Calhoun_ 


_快速了解：_

  * [Go 1.21.4 嗬 1.20.11](https://golangweekly.com/link/147686/web) 现已发布。

  * Chris Siebenmann 思考[域的过期会如何](https://golangweekly.com/link/147688/web)影响通过域名命名的 Go 包的使用。

  * Reddit 上的一些人正在思考[如何才能让现代 Go 再次在 Windows XP 上运行](https://golangweekly.com/link/147690/web)。

  * 📗 Mads Ravn 读完了 Thorsen Ball 精彩的《用 Go 编写解释器》一书，并[分享了他的想法](https://golangweekly.com/link/147692/web)。

  * 🎸 Go + guitars + effects pedals + 深度学习 = [Waveny](https://golangweekly.com/link/147694/web)。


[通过 Go，榨干 SQLite（同时还能享受乐趣）](https://golangweekly.com/link/147696/web "www.terlici.com") —— 这是一个简单的实验，但却体现了保持应用程序简单的现代理念，而且如果使用得当，即使是一台普通的单机也可以做很多事情。

_Stefan Fidanov_ 


---


📰 分类广告

💻 Hired 让求职变得简单 - 公司不再需要追逐招聘人员，而是预先向您提供薪资详细信息。[立即创建免费个人资料](https://golangweekly.com/link/147698/web)。

[加入 Sticker Mule 的“强悍”团队，担任站点可靠性工程师！](https://golangweekly.com/link/147700/web) 们的软件团队遍布 17 个国家/地区，我们正在寻找更多优秀的工程师加入我们的安全团队。

[领导一支工程师团队塑造](https://golangweekly.com/link/147701/web) [Dyte](https://golangweekly.com/link/147703/web) 实时视频的未来：领导一支工程团队负责关键后端系统和 API。优化、规模化并在塑造产品方向方面发挥关键作用。

---  

[优化结构内存使用](https://golangweekly.com/link/147705/web "prog-bytes.hashnode.dev") —— 结构中字段的排序会显着影响内存使用，因此优化是有用的。[fieldalignment](https://golangweekly.com/link/147706/web) 帮你找到可以从该技术中受益的结构。

_Satyarth Ojha_ 


[Go 中的“Ungrammar”和弹性解析](https://golangweekly.com/link/147708/web "eli.thegreenplace.net") —— 您可能听说过 _抽象_ 语法树 (AST)，但还有 _具体_ 语法树 (CST) 这么一个东西，它直接反映输入语言的语法。[Ungrammar](https://golangweekly.com/link/147709/web) 是 Rust 中用于定义 CST 的 DSL，但 Eli 已在 Go 中重新实现了它。

_Eli Bendersky_ 


[`git rebase`：会出现什么问题？](https://golangweekly.com/link/147710/web)   
_Julia Evans_  

[Go 类型标签](https://golangweekly.com/link/147711/web)   
_Boris Smidt_  

  
--- 

## 🛠 代码和工具 

[cpuid：CPU 功能识别库](https://golangweekly.com/link/147713/web "github.com") —— 了解有多少个物理和逻辑核心、CPU 系列以及类似的详细信息。目前支持x86/x64和arm64。不使用任何cgo。

_Klaus Post_ 


[SCS 2.7：HTTP 会话管理中间件](https://golangweekly.com/link/147714/web "github.com") —— 通过中间件加载和保存会话数据，可与各种存储（Postgres、Redis、SQLite 等）配合使用，并且易于扩展和自定义。哦，它已经存在 _很多年_ 了。

_Alex Edwards_ 


[SAML 和 SCIM 的 API](https://golangweekly.com/link/147715/web "workos.com") —— 尚不支持 SSO 或用户配置？加入数百家使用 WorkOS 的公司 — 立即让您的应用程序为商用做好准备。

_WorkOS sponsor_


[httpretty：漂亮地打印出 HTTP 请求，用 Go 编写](https://golangweekly.com/link/147712/web "github.com") —— 受 `curl` 的 `--verbose` 功能启发，这在调试时非常方便。

_Henrique Vicente_ 


[jsondiff：基于 RFC 6902（JSON 补丁）的 JSON 比较库](https://golangweekly.com/link/147716/web "github.com") —— [RFC 6902](https://golangweekly.com/link/147717/web) 定义了一个 JSON 结构，用于表达对其他 JSON 文档的修补操作。

_William Poussier_


[SIPGO 0.15](https://golangweekly.com/link/147718/web) - 用 Go 编写 SIP 服务。

[Miniredis 2.31.0](https://golangweekly.com/link/147719/web) - 用于单元测试的纯 Go Redis 服务器。

[go-quartz 0.9](https://golangweekly.com/link/147720/web) - 无依赖调度库。

[Dynamo 1.21](https://golangweekly.com/link/147721/web) - 富有表现力的 Amazon DynamoDB 库。

[Buf 1.28](https://golangweekly.com/link/147722/web) - 用于使用协议缓冲区的工具。

[go-nfs](https://golangweekly.com/link/147723/web) - NFSv3 协议实现。
  
---

🎲 双陆棋，下棋，走..
  
[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/cred8ykei8kffg6shrzj.jpg)](https://golangweekly.com/link/147724/web)  


[bgammon：浏览器中的西洋双陆棋（感谢有 Go）](https://golangweekly.com/link/147724/web "bgammon.org") —— 它不仅是经典游戏的在线版本，它还是开源的，用 Go 编写，也使用了 Ebitengine。如果您想查看代码，这里是[后端](https://golangweekly.com/link/147725/web)和[前端](https://golangweekly.com/link/147726/web)。

_Trevor Slocum_
