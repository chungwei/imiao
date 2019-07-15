# 理解 make 和 new 的异同点  

先看定义  
```go
// The make built-in function allocates and initializes an object of type
// slice, map, or chan (only). Like new, the first argument is a type, not a
// value. Unlike new, make's return type is the same as the type of its
// argument, not a pointer to it. The specification of the result depends on
// the type:
//	Slice: The size specifies the length. The capacity of the slice is
//	equal to its length. A second integer argument may be provided to
//	specify a different capacity; it must be no smaller than the
//	length. For example, make([]int, 0, 10) allocates an underlying array
//	of size 10 and returns a slice of length 0 and capacity 10 that is
//	backed by this underlying array.
//	Map: An empty map is allocated with enough space to hold the
//	specified number of elements. The size may be omitted, in which case
//	a small starting size is allocated.
//	Channel: The channel's buffer is initialized with the specified
//	buffer capacity. If zero, or the size is omitted, the channel is
//	unbuffered.
func make(t Type, size ...IntegerType) Type

// The new built-in function allocates memory. The first argument is a type,
// not a value, and the value returned is a pointer to a newly
// allocated zero value of that type.
func new(Type) *Type
```

## 相同
- new 和 make 都可以用来分配空间，初始化类型

## 不同
- make 只能用于 slice/map/channel 三种类型；
- make 返回的是第一个参数的类型；new 返回的是该类型0值的指针

## 示例
```go
i := new(int)
fmt.Printf(`%d %#v `, *i, i) // 0 (*int)(0x416020)
	
m := make([]int, 3)
fmt.Printf(`%d %#v `, len(m), m) // 3 []int{0, 0, 0} 
```
