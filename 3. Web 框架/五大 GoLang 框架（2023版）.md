原文：[Top 5 GoLang Frameworks (2023)](https://masteringbackend.com/posts/top-5-golang-frameworks)

---


# 五大 GoLang 框架（2023版）

[作者：Solomon Eseme](/authors/kaperskyguru) | 最近更新：2023 年 9 月 23 日星期六

自 2007 年谷歌发布到技术市场开始，GoLang 就获得了广泛的使用。根据[2022 年 Stackoverflow 的调查](https://survey.stackoverflow.co/2022/#most-loved-dreaded-and-wanted-language-love-dread)，Golang 成为了软件工程行业第三大最受欢迎和第八大最受喜爱的编程语言。

这种使用率的增加，使得 Golang 成为致力于构建创新以及具有安全意识应用程序的软件工程师必须学习的语言。

在我们深入讨论之前，让我们先解释一下 Golang 是什么，以及为什么你应该考虑学习它。Golang 是谷歌于 2007 年创建的一种通用、编译型开源编程语言，用于大规模构建更快、更可靠、更高效的软件。

此外，它是由谷歌支持的开源软件，简单易学，几乎可以立即上手。它支持内置并发、强大的标准库以及不断发展的社区和生态。

本文将探讨排名前 5 的 Golang 框架，让您深入了解当前的需求以及 2023 年可以开始学习的内容。（译注：抱歉，2023年底才看到 😛）

## **排名前 5 的 Golang 框架**

下面，我们将根据受欢迎程度、使用次数和 Github 星数列出 2023 年最值得学习的 5 大 Golang 框架。综合上述标准，我们得到了以下 5 大 Golang 框架：

1. Gin 框架
2. FastHTTP
3. Echo
4. Beego
5. Fiber

让我们直接深入探索这些 Golang 框架：

### **Gin 框架**

Gin 框架是最快的全功能 Go Web 框架。它是一个用 Go 编写的框架，类似 Martini。此外，Gin 的性能比 Martini 的性能快了 40 倍，广受需要性能和生产力的开发人员和团队的喜爱。

截至撰写本文时，Gin 框架的一些增长统计数据显示，Github Stars 超过 **61.9k**，GitHub 使用量约**79.6k**，累计总积分达**141.6k**，这使得 Gin 框架位居“2023 年排名前 5 的 Golang 框架“的图表榜首。

![gin-golang-1.png](https://strapi-images-aws-s3.s3.us-west-2.amazonaws.com/gin_golang_1_28355da10a.png)
<center>Gin Go 框架统计数据</center>


#### **Gin 框架的特点**

以下是该框架的一些显着特征。

##### ***快***

Gin 框架的开发速度比 Martini 的性能快 40 倍，它包含基于基数树的路由、较小的内存占用，并且无反射，从而实现可预测的 API 性能。

##### ***中间件支持***

中间件在任何后端开发语言和工具中都至关重要，因为它会拦截请求并在将请求发送到服务器之前执行某些操作。

Gin 框架包括中间件支持，开发人员可以将其配置为通过一系列中间件和最终操作来处理传入的 HTTP 请求。例如，记录器、授权、GZIP，以及最终在数据库中发布消息。

##### ***错误管理***

软件开发中的[错误处理](https://medium.com/backenders-club/error-handling-in-node-js-ef5cbfa59992)是一个重要的概念，应该在代码库成为大型企业级应用程序量级之前，在早期开发阶段决定和处理。

Gin 框架提供了一种处理 HTTP 请求期间错误的便捷方法。最终，中间件可以将这些错误写入日志文件、数据库，并通过网络发送它们。

##### ***JSON 验证***

Gin 可以解析和验证请求的 JSON 内容，例如检查必填值是否存在。

##### ***路由分组***

Gin 框架允许开发人员更好地组织他们的路由。将需要授权与不要求授权，以及不同的 API 版本分开。此外，组可以无限嵌套，而不会降低性能。

### **FastHTTP**

FastHTTP 不是一个框架，而是 Go 的快速 HTTP 包。它提供高性能。热路径中的零内存分配，并且速度比 Golang 的原始 net/http 包快 10 倍。

FastHTTP 是 Golang 中的快速 HTTP 实现，专为某些高性能边缘场景而设计，例如每秒处理数千个中小型请求并要求一致的低毫秒响应。

截至撰写本文时，FastHTTP 包的一些增长统计数据（基于 Github）包括超过**18.2k 的 Github Stars**和约**17.2k 的 GitHub 使用**，累计总点数为**35.5k**，这使得 FastHTTP 包位居“2023 年排名前 5 的 Golang 框架“的图表榜眼。

![unnamed-10.png](https://strapi-images-aws-s3.s3.us-west-2.amazonaws.com/unnamed_10_842a2b552d.png)
<center>FastHTTP Go 库的统计信息。</center>

#### **FastHTTP 包的特点**

以下是 FastHTTP 包的一些最重要的功能。

1. FastHTTP 包针对速度进行了优化，因为它可以在现代硬件上轻松处理超过 100Kqps 和超过 1M 的并发 keep-alive 连接。
2. 最重要的是，它还针对低内存使用进行了优化。
3. FastHTTP API 旨在扩展现有的客户端和服务器实现。
4. 它提供默认的安全实现和反 DoS 限制（例如并发连接数、请求读取超时等）。
5. 它支持 RequestCtx.Hijack，以便轻松地 ‘Connection: Upgrade’

### **Echo 框架**

Echo 框架是一个极简、高性能、可扩展的 Go Web 框架。它允许开发人员构建可扩展的 API，支持优化的路由器、中间件、数据渲染、模板等。

截至撰写本文时，Echo 框架的一些增长统计数据（基于Github）包括超过**23.1k 的 Github Stars**和约**9.1k 的 Github 使用**，累计总点数为**32.2k**，这使得 Echo 框架成为了“2023 年排名前 5 的 Golang 框架“的图表的探花。

![unnamed-11.png](https://strapi-images-aws-s3.s3.us-west-2.amazonaws.com/unnamed_11_1424eb20c2.png)
<center>Echo Go 框架的统计信息。</center>

#### **Echo 框架的特点**

以下是 Echo 框架的主要功能。

##### ***优化路由器***

它包含一个高度优化的 HTTP 路由器，带零动态内存分配，可以智能地确定路由的优先级。

##### ***中间件***

Echo 框架支持许多内置中间件，让您可以直接使用或者基于这些中间件定义您自己的中间件。可以在根、分组或路由级别设置中间件。

##### ***数据渲染***

Echo 允许 API 发送各种类型的 HTTP 响应，包括 JSON、XML、HTML、文件、附件、内联、流或 Blob。

##### ***数据绑定***

它还支持 HTTP 请求负载的数据绑定，包括 JSON、XML 或表单数据。

##### ***可扩展***

Echo 有自定义的中央 HTTP 错误处理，带一个易于扩展的 API。

### **Beego 框架**

Beego 框架是开源的，用以以 Go 的方式构建和开发您的应用程序。它是一个 RESTful HTTP 框架，用于快速开发 Go 应用程序，包括 API、Web 应用和具有集成的 Go 特定功能（如接口和结构嵌入）的后端服务。

截至撰写本文时，Beego 框架的一些增长统计数据（基于 Github）包括超过**28.k 个 Github Stars**和约**120 个 Github 使用**，累计总积分为**29.1k**，这使得 Beego 框架位居“2023 年排名前 5 的 Golang 框架“的图表的第四位。

![unnamed-12.png](https://strapi-images-aws-s3.s3.us-west-2.amazonaws.com/unnamed_12_c0ed71d7c8.png)
<center>Beego Go 框架的统计信息。</center>

#### **Beego 框架的特点**

以下是使得 Beego 框架独一无二的一些功能：

##### ***易用***

对于核心 Go 开发人员来说，Beego 框架是使用核心的 Go 模块，以 Go 的方式创建的，从而使 Go 开发人员可以轻松适应和构建可扩展的企业级应用程序。

借助 RESTful 支持、MVC 模型，使用 bee 工具快速构建具有代码热编译、自动化测试、自动化打包和部署等功能的应用。

##### ***智能***

Beego 框架支持智能路由和监控。它可以监控你的 QPS、内存、CPU 使用率和 goroutine 状态。它使您可以完全控制您的在线应用程序。

##### ***模块化***

该框架具有强大的内置模块，包括会话控制、缓存、日志记录、配置解析、性能监控、上下文处理、ORM 支持和请求模拟。您所得到的是一个为任何应用程序奠定强大的基础的框架。

##### ***高性能***

使用原生 Go HTTP 包来处理请求，并具有 goroutine 的高效并发。您的 Beego 应用程序可以处理大量流量，就像 Beego 在许多产品中所做的那样。

### **Fiber 框架**

Fiber 是一个构建在 Fasthttp 之上的 Golang Web 框架。它旨在简化过程以实现快速开发，同时考虑到零内存分配和性能。该框架是一个用 Go 编写的[受 Express 启发的](https://masteringbackend.com/posts/expressjs-5-tutorial-the-ultimate-guide/)Web 框架。

截至撰写本文时，Fiber 框架的一些增长统计数据（基于 Github）包括超过**21.7k 个 Github Stars**和约**2.4k 个 Github 使用**，累计总点数为**24.0k**，使 Fiber 框架位居“2023 年排名前 5 的 Golang 框架“的图表的第五位。

![unnamed-13.png](https://strapi-images-aws-s3.s3.us-west-2.amazonaws.com/unnamed_13_8a98037825.png)
<center>Fiber Go 框架的统计信息。</center>

#### **Fiber 框架的特点**

以下是 Fiber 框架的一些独特功能：

##### ***稳健的路由***

Fiber 框架拥有最好的路由系统之一；路由系统来自 Express 框架，用起来简单且流畅。  

##### ***提供静态文件***

通过定义静态路由，轻松提供静态 HTML、CSS 和 JavaScript 文件。您还可以在同一条路线上提供多个目录的内容！

##### ***极致性能***

Fiber 框架使用 FastHTTP 库，该库被称为 Golang 的快速 HTTP 库之一。使用 Fiber 框架，您的应用程序将享受无与伦比的速度。

##### ***API 就绪***

[类 Express 框架](https://masteringbackend.com/posts/expressjs-5-tutorial-the-ultimate-guide/)非常适合用来开发 API，因为它结合了 [Express](https://masteringbackend.solomoneseme.com/posts/expressjs-5-tutorial-the-ultimate-guide/) 的简单性和创建生产就绪 API 的强大功能。

#### ***灵活的中间件支持***

从[几个现有中间件](https://docs.gofiber.io/contrib/#-middleware-implementations)中选择，或者创建你自己的中间件！使用它们在应用程序中的某些请求到达控制器之前验证和操作它们。

## **福利**

您还可以对以下列表中的 Golang 框架和库进行更多的研究，它们承诺提供出色的用例和实用性。

1. Revel
2. Martini
3. Mango
4. Iris
5. Kit
6. Goji
7. Buffalo
8. Webgo

## **总结**

为您的项目选择哪一个框架仅取决于项目类型、编程语言和团队实力。

这篇文章粗略地概述了这些框架及其一年来的趋势。它还让您了解在选择最好的 Golang 框架之一进行学习时会发生什么。

去原文评论区告诉原作者（译注：或者你也可以告诉我😄）你的想法吧。
