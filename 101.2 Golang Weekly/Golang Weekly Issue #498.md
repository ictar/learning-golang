原文：[Golang Weekly Issue #498](https://golangweekly.com/issues/498)

---

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/w4viyllz0nvvibkfy4hh.jpg)](https://golangweekly.com/link/152089/web)  
---  
  

📊 [我是如何使用 Go 让自己活着的](https://golangweekly.com/link/152089/web "www.bytesizego.com") —— 探讨如何使用开源、医疗设备和 Go 来帮助 1 型糖尿病患者监测血糖水平并对异常情况发出警报，包括在必要时向第三方发送消息。

_Matt Boyle_ 


[Go 1.22 中 `for` 循环的语义变化：注意影响](https://golangweekly.com/link/152090/web "go101.org") —— Tapir 喜欢 `for`/`range` 的新语义，但不喜欢更传统的 `for` 循环。这篇文章让我们了解了其变化，并深入研究了一些示例，展示了需要注意的事项。

_Tapir (Go 101)_ 

  
[![](https://copm.s3.amazonaws.com/8b542de2.png)](https://golangweekly.com/link/152088/web) 

[结对编程结合了面对面工作和远程工作的最佳部分](https://golangweekly.com/link/152088/web) —— Tuple“让与同事和朋友的结对编程再次变得有趣”。免费试用并了解为什么 Figma 的工程师们无法停止谈论 Tuple。

_Tuple sponsor_


[Go 2022-2024 及未来：Go 的工程总监谈 AI 的作用](https://golangweekly.com/link/152091/web "ajmani.net") —— 上周，我们很高兴 Go 的工程总监在博客中介绍了[他是如何进入管理领域](https://golangweekly.com/link/152092/web)，以及[Go 的早期发展](https://golangweekly.com/link/152093/web) — 现在他正在思考 Go 近年来是如何成熟的，并思考人工智能将如何改变现状。

_Sameer Ajmani_ 

  
_快速了解：_

  * ⚠️ [Go 1.22.1 和 1.21.8](https://golangweekly.com/link/152094/web) 预计将于今天晚些时候发布，其中包括对标准库进行的安全修复，这些修复基于一些尚未披露的漏洞。

  * ⭐️ [Charm](https://golangweekly.com/link/152095/web) 维护着各种流行的 Go [库](https://golangweekly.com/link/152096/web)，包括 [Bubble Tea](https://golangweekly.com/link/152097/web) 和 [Wish](https://golangweekly.com/link/152098/web)，而 Charm 的 COO 最近 [▶️ 谈到了公司是如何组建的，以及他们最近是如何筹集了 600 万美元的。](https://golangweekly.com/link/152099/web)

  * 恕我直言，[Caddy HTTP 服务器](https://golangweekly.com/link/152100/web)是  Go 最成功的使用案例之一，因此很高兴看到[这个“在 Y 分钟内尝试 Caddy”交互式指南](https://golangweekly.com/link/152101/web)，以直接在 Web 上了解关于其配置和 API 的更多信息。

  * 🧹 [fork-sweeper](https://golangweekly.com/link/152102/web) 是一个小型 Go 工具，用于删除 GitHub 上未使用的分叉。


[预分配切片内存对性能的影响](https://golangweekly.com/link/152103/web "oilbeater.com") —— 作者希望通过定量测量和自动检测工具，以数字的形式确定预分配内存是如何提高性能的。

_Oilbeater_ 


[通过持久执行构建可靠的应用程序](https://golangweekly.com/link/152104/web "pages.temporal.io") —— 了解持久执行的概念，它用于解决分布式系统中的各种问题。

_Temporal Technologies sponsor_


[使用范型的更简单的可组合 HTTP 处理器](https://golangweekly.com/link/152105/web)   
_Willem Schots_  

[当 Kubernetes 和 Go 不能很好地协同工作时](https://golangweekly.com/link/152106/web)   
_Emin Laletovic_  


[给 Go 小白的一些建议](https://golangweekly.com/link/152107/web)   
_Ewan Valentine_  
 


🛠 代码和工具  

📄 [pdfcpu 0.7：一个 PDF 处理和操作库](https://golangweekly.com/link/152108/web "pdfcpu.io") —— 您可以验证和优化 PDF、拆分 PDF、将 PDF 合并在一起、提取元素等等。已获得 Apache 2.0 许可。[GitHub 存储库。](https://golangweekly.com/link/152109/web)

_pdfcpu_


[fgprof 0.9.4：Go 的完整采样分析器](https://golangweekly.com/link/152110/web "github.com") —— [pprof](https://golangweekly.com/link/152111/web) 是分析 Go 代码的首选工具，但[它仅分析“CPU 上”时间](https://golangweekly.com/link/152112/web)。`fgprof` 也可以与 pprof 一起工作并测量“非 CPU”时间（例如花费在 IO 上的时间），以便您可以看得更全。

_Felix Geisendorfer_ 


[Nuk：Go 的内存 Arena 实现](https://golangweekly.com/link/152113/web "github.com") —— 一个在 Go 中提供内存 arena 的包的提案[于两年前提出](https://golangweekly.com/link/152114/web)，但结果是最初的实现被搁置了。Nuke 再次尝试了这个想法，甚至包括了并发 arena 的实现。

> 译注： Arena 指的是一种从一个连续的内存区域分配一组内存对象的方式

_Miguel Ángel Ortuño_ 


[Huma：构建由 OpenAPI 和 JSON Schema 支持的 API ](https://golangweekly.com/link/152115/web "huma.rocks") —— Huma 是一个微框架，用于利用像 OpenAPI 这样的通用标准来创建 HTTP REST 或 RPC API。

_Daniel Taylor_ 


---
📰 分类广告

🪝[Hookdeck](https://golangweekly.com/link/152116/web)：用于 Webhooks 和异步消息传递的托管可靠性和可观察性层。[大规模接收、转换、路由和交付事件](https://golangweekly.com/link/152116/web)。

[Go 调试的终极指南](https://golangweekly.com/link/152117/web) – 大多数高级 Go 工程师都是调试生产系统的专家。您可以通过马特·博伊尔的[这门全新的课程](https://golangweekly.com/link/152117/web)来提升您的技能。

---  
  

[TinyTest：与 TinyGo 配合使用的唯一（？）断言库](https://golangweekly.com/link/152118/web "github.com") —— 它不使用反射，而是使用泛型来比较值。_（当然，它也适用于常规 Go。）_

_Orsinium Labs_ 


[Polaris：一个新的最小化的 Go 工作流编排器: A New, Minimal Workflow Orchestrator for Go](https://golangweekly.com/link/152119/web "github.com") —— 如果您有一个多步骤工作流，其中每个步骤都依赖于其他步骤，那么像 Polaris 这样的系统可以帮助您将某些结构应用于这种情况。

_Harshad Manglani_ 

[EchoVault：基于 Go 的分布式内存数据存储](https://golangweekly.com/link/152120/web "github.com") —— 它使用 Redis 的 RESP 协议，因此可以通过现有的 Redis 客户端库使用它。

_EchoVault_ 

  
[Testify 1.9](https://golangweekly.com/link/152121/web) - 与标准库配合良好的测试断言和模拟。

[Gofeed 1.3](https://golangweekly.com/link/152122/web) -  RSS、Atom 和 JSON feed 解析库。

[Toxiproxy 2.8](https://golangweekly.com/link/152123/web) - TCP proxy to simulate chaotic network conditions.

[OpenFGA 1.5](https://golangweekly.com/link/152124/web) - 受 Google Zanzibar 启发的身份验证引擎。

[go-github 60.0](https://golangweekly.com/link/152125/web) - GitHub v3 API 的客户端库。

[clickhouse-go 2.20](https://golangweekly.com/link/152126/web) - ClickHouse 的 Go 驱动。

[Ginkgo 2.16](https://golangweekly.com/link/152127/web) - 成熟的测试框架。
