package types

// 定义变量：var 变量名 类型
// - 变量会自动初始化为零值
var x int // 0
// - 如果显式提供初始值，则可以省略类型，由编译器推断
var y = false // bool
// - 统一作用域内不同重复定义。

// - 可一次定义多个变量，包括不同初始值定义不同类型
var (
	i, j int // 相同类型
	a, s = 100, "abc" // 不同类型
)
func main() {
	// - 未使用的局部变量当作错误
	var x int 

	
}