# 基本使用方式
slice 又称动态数组，依托数组实现，可以方便的进行扩容、传递等，实际使用中比数组更灵活。
```go
var array [10]int
// array[start:end]，不包括end。如果 slice 根据数组创建，则与数组共享存储空间
var slice1 = array[5:6]

// 读取长度：O(1)
len(slice1)
// 读取容量：O(1)
cap(slice1)

// 不根据数组创建 slice
var slice2 []int
// 追加数据。append函数执行时会判断切片容量是否能够存放新增元素，如果不能，则会重新申请存储空间，新存储空间将是原来的2倍或1.25倍(取决于扩展原空间大小)
slice2 = append(slice2, 1, 2, 3)

// 通过 make 创建 slice
slice3 := make([]int, 2)
// slice3[low:high:max]操作意思是对order进行切片，新切片范围是[low, high),新切片容量是max
slice3[2:3:10]
```

# 实现原理
源码见 `src/runtime/slice.go`

## 数据结构
`src/runtime/slice.go:slice` 定义了 slice 的数据结构：
```go
type slice struct {
	array unsafe.Pointer // 指向底层数组
	len   int // 切片长度
	cap   int // 底层数组容量
}
```

## 创建 slice
### 使用 make 创建 slice
使用make来创建Slice时，可以同时指定长度和容量，创建时底层会分配一个数组，数组的长度即容量。
![](./img/1683619294136.jpg)

### 使用数组或切片创建 slice
使用数组或切片来创建Slice时，Slice将与原数组或切片共用一部分内存。
![](./img/1683619400035.jpg)

- `slice := array[start:end]` 这种新生成的切片并没有指定切片的容量，实际上新切片的容量是从start开始直至array的结束。
- `slice := array[start:end:cap]`, 其中cap即为新切片的容量，当然容量不能超过原切片实际值

注意：数组和切片操作可能作用于同一块内存

## 扩容（`append`）
使用append向Slice追加元素时，如果Slice空间不足，将会触发Slice扩容。扩容实际上**重新一配一块更大的内存**，将原Slice数据拷贝进新Slice，然后**返回新Slice**，扩容后再将数据追加进去。

使用`append()` 向Slice添加一个元素的实现步骤如下:
1. 假如Slice容量够用，则将新元素追加进去，Slice.len++，返回原Slice
2. 原Slice容量不够，则将Slice先扩容，扩容后得到新Slice
3. 将新元素追加进新Slice，Slice.len++，返回新的Slice。

## copy()
使用`copy()`内置函数拷贝两个切片时，会将源切片的数据逐个拷贝到目的切片指向的数组中，拷贝数量取两个切片长度的最小值。**copy过程中不会发生扩容。**

# 总结
- 创建切片时可跟据实际需要预分配容量，尽量避免追加过程中扩容操作，有利于提升性能;
- 切片拷贝时需要判断实际拷贝的元素个数
- 谨慎使用多个切片操作同一个数组，以防读写冲突
- **通过函数传递切片时，不会拷贝整个切片**，因为切片本身只是个结构体而已