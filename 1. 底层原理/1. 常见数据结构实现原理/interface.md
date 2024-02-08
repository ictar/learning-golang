
> 以下基于 Go 1.2x

Go 中描述 interface{} 的底层结构体有 `iface` 和 `eface` （定义在 `src/runtime/runtime2.go` 中）

## 参考：
- [interface · 深入解析Go - tiancaiamao](https://tiancaiamao.gitbooks.io/go-internals/content/zh/07.2.html)
- (深入研究Go interface 底层实现)[https://halfrost.com/go_interface/]