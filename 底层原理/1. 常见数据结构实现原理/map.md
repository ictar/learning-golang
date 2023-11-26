# 实现原理
源码见 `src/runtime/map.go`

## 数据结构
Golang的map使用**哈希表**作为底层实现，一个哈希表里可以有多个哈希表节点，也即**bucket**，而每个bucket就保存了map中的一个或一组键值对。`src/runtime/map.go:hmap` 定义了 map 的数据结构：

```go
type hmap struct {
	count     int // 当前保存的元素个数
	flags     uint8
	B         uint8  // 指示 bucket 数组的大小
	noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
	hash0     uint32 // hash seed

	buckets    unsafe.Pointer // bucket 数组指针，数组的大小为 2^B
	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

	extra *mapextra // optional fields
}
```

### bucket 数据结构

## 哈希冲突

## 负载因子

## 渐进式扩容

### 扩容的前提

### 增量扩容

### 等量扩容

## 查找过程

## 插入过程
