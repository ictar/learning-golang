原文：[Golang Weekly Issue #504](https://golangweekly.com/issues/504)

---

[Dolt 如何使用 GitHub Actions 创建性能引导的优化构建](https://golangweekly.com/link/154095/web "www.dolthub.com") —— 了解 Dolt 构建过程的幕后情况，以及他们的团队是如何将配置文件引导优化 (PGO) 引入到管道中的。虽然存在一定的复杂性，但回报你的是一个可立即部署、经过性能调整的二进制文件。

_Dustin Brown (DoltHub)_ 

> 📈 Dolt 的人员之前曾写过[他们在构建中使用 PGO 的经验](https://golangweekly.com/link/154100/web)，以及所看到的小但显着的性能改进。  
  

[从零到生产：Go 的 Google 之旅](https://golangweekly.com/link/154096/web "i-admin.cetico.org") —— 作者在 Google 担任了九年的 SRE，并分享了一些他看到 Go 在其形成时期的成长经历和被采用的故事。以及一些我在其他地方没有看到的有趣背景。

_Yves Junqueira_ 


[![](https://copm.s3.amazonaws.com/5dec8b68.png)](https://golangweekly.com/link/154094/web) 

[Go 事件驱动：驯服混乱的不公平优势](https://golangweekly.com/link/154094/web "threedots.tech") —— 微服务承诺了简单性，但带来了更多复杂性？多米诺故障和性能瓶颈让您感到沮丧吗？不要落后。了解用于创建真正解耦和可扩展服务的经过充分验证的模式。

_Three Dots Labs sponsor_



[Go 1.22 中的新功能：`cmp.Or`](https://golangweekly.com/link/154097/web "blog.carlana.net") —— `cmp.Or`（你可能会说，这是“1.22 的隐藏宝石”）是一个函数，它返回第一个不是类型的零值的传入值 —— 这是缩短某些条件赋值的一种便捷方法。这是涵盖作者自己对 1.22 的贡献的系列的最后一篇。

_Carlana Johnson_ 


_快速了解：_


  * Reddit 上的人们开始讨论 [Google 实际使用了多少 Go](https://golangweekly.com/link/154098/web)。名为 _assbuttbuttass_ 和 _Deathmaster99_ 的所谓员工也参与了进来，确认谷歌确实在大量使用 Go。

  * 📅 如果您在英国曼彻斯特附近，那我想告诉您，5 月 1 日将举办一场[曼彻斯特 Golang 聚会](https://golangweekly.com/link/154129/web)。


[Render 如何利用 Go 泛型强制执行访问控制](https://golangweekly.com/link/154099/web "render.com") —— Render 是一个现代托管和部署平台，具有角色系统，供用户和管理员拥有指定的、已定义的权限。他们的工程师希望在 Go 代码库中获得 _编译时_ 保证，以减少危险缺陷和漏洞进入其角色系统的可能性。

_Shawn Moore (Render)_ 


📄 [CLI UX 最佳实践：提高进度的三种模式](https://golangweekly.com/link/154101/web) _Roman Shamin_


📄 [一条命令实现利用 Bazel 访问 Go gRPC 服务器](https://golangweekly.com/link/154102/web) - 唯一真正的缺点是必须使用 Bazel。 _Uros Popovic_


📄 [如何在 Go 中使用 Tailscale 进行 gRPC 身份验证](https://golangweekly.com/link/154103/web) _Khash Sajadi_


📄 [简介：使用 Delve 进行调试](https://golangweekly.com/link/154104/web) _Adam Gordon Bell_


## 🛠 代码和工具  
  
[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/itajasjhnwtuhwmjlg3h.jpg)](https://golangweekly.com/link/154105/web)  


[makefile-graph：将 Makefile 转换为图](https://golangweekly.com/link/154105/web "github.com") —— 可用作库和 CLI 工具，它解析 makefile 并生成表示各个目标之间关系的图表（然后由 [Graphviz](https://golangweekly.com/link/154106/web) 的 `dot` 渲染）。

_Marin Atanasov Nikolov_ 


[简化的状态机](https://golangweekly.com/link/154107/web "pages.temporal.io") —— 探索 Temporal 的简单性和耐用性，这是应对复杂状态机挑战的解决方案。立即解锁效率！

_Temporal Technologies sponsor_

  

[elem-go：一种创建和操作 HTML 元素的类型安全型方法](https://golangweekly.com/link/154108/web "github.com") —— 使用 Go 代码轻松创建 HTML 元素，并以类型安全的方式使用元素、属性和特性。[v0.25.0](https://golangweekly.com/link/154109/web) 版本引入了一个新的 _StyleManager_ 功能，用于以编程方式管理 CSS 样式 - 查看示例。

_Chase Fleming_ 


[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/wpjkhofc4yqjav3qehiz.jpg)](https://golangweekly.com/link/154110/web)   
  

[go-size-analyzer：分析已编译的 Go 二进制文件中的依赖大小](https://golangweekly.com/link/154110/web "github.com") —— 如果分析您的 `Makefile`（见上面）还不够诱人的话，那么分析您已编译的 Go 二进制文件怎么样？看看所有这些字节都去哪里了。

_Zxilly_ 


💎 [Gemfast：面向 Ruby 开发者的嵌入式“Gem”服务器](https://golangweekly.com/link/154111/web "github.com") —— 一个新的自托管 Rubygems（Ruby 的包管理系统）服务器，用 Go 编写，以便于部署，让 Rubyists 可以镜像和缓存来自官方注册表的 gem，以及提供他们自己的私人 gem。

_Gemfast_ 


[Gnet 2.5：高性能、非阻塞、事件循环网络库](https://golangweekly.com/link/154112/web "github.com") —— _“通过利用`epoll`和`kqueue`，从头开始​​构建，在许多特定场景中，它可以实现比 Go 的 `net` 更高的性能和更低的内存消耗。”_

_Andy Pan_ 


[Opengist：一个自托管 Pastebin 应用](https://golangweekly.com/link/154113/web "github.com") —— 一个由 Git 驱动的 [Gist](https://golangweekly.com/link/154114/web) 式 Pastebin 应用程序，其功能比您想象的要多。这是[一个现场演示](https://golangweekly.com/link/154115/web) – 它看起来确实非常 _GitHubby_。

_Thomas Miceli_ 
  

[Muffet：一个快速的递归网站链接检查器](https://golangweekly.com/link/154116/web "github.com") —— 用于递归地抓取和检查站点页面的 CLI 工具。

_Yota Toyama_ 


[spark-connect-go：Go 的 Apache Spark Connect 客户端](https://golangweekly.com/link/154117/web "github.com") —— 一个 _“高度实验性”_ 的Spark Connect for Apache Spark 客户端。

_Apache_ 


[q：Go 的量子计算模拟器](https://golangweekly.com/link/154118/web "github.com") —— 适合比我聪明得多的人。

_itsubaki_ 


[Micro 4.5](https://golangweekly.com/link/154119/web) -  Go 分布式系统抽象/平台。

[AWS Lambda for Go 1.47](https://golangweekly.com/link/154120/web) - 供 Go 开发人员在 AWS Lambda 上构建的示例和工具。

[env 11.0](https://golangweekly.com/link/154121/web) - 将 env 变量解析为结构的零依赖方法。

[Goose 3.20](https://golangweekly.com/link/154122/web) - 数据库迁移工具。支持SQL _和_ Go函数。

[Go-OpenAI 1.22](https://golangweekly.com/link/154123/web) -  OpenAI API 包装器库。

[Expr 1.16.5](https://golangweekly.com/link/154124/web) - 以 Go 为中心的表达语言。


---

📰 分类广告

🪝 [Hookdeck](https://golangweekly.com/link/154125/web)：用于事件驱动应用程序的无服务器队列。[了解更多](https://golangweekly.com/link/154125/web)。


[GoLand 2024.1 来咯。](https://golangweekly.com/link/154126/web) 带来了免费、本地运行的 AI 全线完成、性能和用户体验改进。[一探究竟。](https://golangweekly.com/link/154126/web)

---  
  

## 🔭 一个快速实验？ 
  
[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/v3o3woargcqoisrze93j.jpg)](https://golangweekly.com/link/154127/web)  
  
[gops：一种列出当前运行中的由 Go 构建的进程的方式](https://golangweekly.com/link/154127/web "github.com")

适用于任何 Go 构建的进程，但如果您的程序可以启动诊断代理，则可以获得额外的信息。

该工具本身非常简单，但有趣的“实验”在于是在您的系统上运行它，并查看您正在运行哪些您没有意识到是 Go 应用程序的进程！您可以轻松运行它，如下所示：

```
$ go install github.com/google/gops@latest$ gops
```