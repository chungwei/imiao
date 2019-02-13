# Go 切片的使用和原理

> 原文请参看 [Go官方博客](https://blog.golang.org/go-slices-usage-and-internals)  

## 介绍
切片是 golang 提供的一种基本数据结构，实用且效率高，切片类似于其他语言的数组，但又稍有区别。  
本文讲带你了解切片的原理和使用。

## 数组
切片是在数组的基础上抽象出来的数据结构，所以在理解切片之前，我们先了解一下数组。  

数组定义的是固定长度和元素类型的数据集合。例如，`[4]int` 声明了包含4个int类型的数组，其长度是固定的，`[4]int` 和 `[5]int` 完全不一样。数组有索引，索引从0开始，`s[n]` 表示第n个元素，如以下示例  
```go
var a [4]int
a[0] = 1
i := a[0]
// i == 1
```

数组的值不需要进行显式初始化，实际上在声明完成后，数组里每个元素已经是0值，
```
var a [4]int
fmt.Printf("%v", a) // [0 0 0 0]
```
在内存中，[4]int是一块用于存储4个int类型的连续内存块，如下图  

![IMAGE](./resources/go-slices-usage-and-internals_slice-array.png)

Go数组是值类型，一个数组变量是指整个数组，而不是像C语言只是指向数组第一个元素的指针。这意味着给数组赋值或使用数组传值实际上是对其副本的操作，(为避免值复制可以传递一个指向数组的指针，但要注意的这是一个指向数组的指针而不是数组的值)。对于数组，换一个角度来理解，可认为它是一个固定长度的结构体，只是用索引取代了成员。  

数组可以像如下声明
```
b := [2]string{"Penn", "Teller"}
```
或者，也可以这么定义
```
b := [...]string{"Penn", "Teller"}
```

以上两种方式，变量b的类型都是[2]string  

## 切片
数组是适用于很多场景，但不够灵活，所以在Go的代码里用到数组的并不多，相反，切片的使用很频繁。切片在数组的基础上做了封装，提高了更强大、更便利的功能。

切片的声明，如`[]T`，其中的`T`是切片元素的类型，跟数组不一样之处，切片是不需要指定长度的。  
切片可以像如下声明
```
letters := []string{"a", "b", "c", "d"}
```
也可以使用内置函数`make`
```
func make（[] T，len，cap）[] T
```
其中，`T`是切片元素的类型，`make`方法需要3个参数，类型、长度和容量(可选的)。调用该方法后，实际上是分配了一个数组的内存空间，并返回一个引用该数组的切片。

示例如
```
var s []byte
s = make([]byte, 5, 5)
// s == []byte{0, 0, 0, 0, 0}

```
如果容量参数`cap`被忽略，默认为`cap` 等于`len`，所以你可以看到以上代码更简洁的版本如下
```
s := make([]byte, 5)
```
要获取切片的容量和长度，可以使用内置方法`cap` 和 `len`，如
```
len(s) == 5
cap(s) == 5
```
接下来的我们继续讨论一下长度和容量的差异。  
一个0值的切片实际上是`nil`，其长度和容量都是0。  
除了上面提到的定义方式，切片还可以通过已存在的切片、数组的其中一部分来完成定义。例如，`b[1:4]` 定义了一个新的切片，它包含切片`b`的第1，2，3个元素(从0开始计算)，在新的切片里，这3个元素的索引是0-2
```
b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
// b[1:4] == []byte{'o', 'l', 'a'}, sharing the same storage as b
```
上述例子中的切片索引是可以省略的，默认是从0到切片的长度(即len(b))，
```
// b[:2] == []byte{'g', 'o'}
// b[2:] == []byte{'l', 'a', 'n', 'g'}
// b[:] == b
```

此外，还可以使用数组来定义切片，如
```
x := [3]string{"Лайка", "Белка", "Стрелка"}
s := x[:] // a slice referencing the storage of x
```

## 切片原理
切片是数组段的描述符，它包含指向数组段的指针、数组段的长度和数组段的容量。如下图  
> 容量是指数组的最大长度

![IMAGE](./resources/go-slices-usage-and-internals_slice-struct.png)

如上述我们定义的切片`s = make([]byte, 5)`，其结构如下图

![IMAGE](./resources/go-slices-usage-and-internals_slice-1.png)

切片的长度是指引用的数组段的元素个数，而容量是底层数组的长度。接下来我们将通过几个示例来明确长度和容量的差异。  

还是我们上面的切片 `s`，我们看一下其数据结构和底层数组的关系  
```
s = s[2:4]
```
![IMAGE](./resources/go-slices-usage-and-internals_slice-2.png)

从上图可以看出，切片的赋值并没有复制原来切片的数据，而是创建了一个新的指针指向，这样就可以保证对切片的操作跟数组一样高兴。但是，修改新切片的值，则会影响到原来的切片：
```
d := []byte{'r', 'o', 'a', 'd'}
e := d[2:] 
// e == []byte{'a', 'd'}
e[1] = 'm'
// e == []byte{'a', 'm'}
// d == []byte{'r', 'o', 'a', 'm'}
```
再次提到我们上面的切片 `s`，目前它的长度小于容量，我们可以把它的长度增长到等于其容量：
```
s = s[:cap(s)]
```
![IMAGE](./resources/go-slices-usage-and-internals_slice-3.png)

但要注意的是，切片的长度不能大于容量，如果尝试做这样的操作，就意味着索引越界，最终将会引起`panic`；类似的，切片也不能访问小于0的索引。

## 切片的增长
要提高切片的容量，我们必须是新建一个更大容量的切片，并把原切片的值复制到新切片中，其中的技术细节，类似于其他语言中如何实现动态数组。下面的例子，是将原切片`s`容量翻倍增长的具体实现
```
t := make([]byte, len(s), (cap(s)+1)*2) // +1 in case cap(s) == 0
for i := range s {
        t[i] = s[i]
}
s = t
```
上述代码中的循环段，实际上就是内置方法`copy`的实现，通过名字可以知道，该方法的功能就是将原切片的值复制到目标切片中，最终返回的是被复制的元素个数。
```
func copy(dst, src []T) int
```
上面的`copy`方法，支持两个不同长度切片之间的复制，但复制的元素个数是以二者间较小者为准。此外，该方法可以正确处理在原切片和目标切片之间共享底层数组，以及切片重叠等。

使用`copy`方法，上面的代码段可以简化为
```
t := make([]byte, len(s), (cap(s)+1)*2)
copy(t, s)
s = t
```

还有一个常见的操作是追加数据到切片末尾。下面的方法实现了切片元素的追加功能，如果原切片容量不足可自动扩容，最后返回更新后的切片：
```
func AppendByte(slice []byte, data ...byte) []byte {
    m := len(slice)
    n := m + len(data)
    if n > cap(slice) { // if necessary, reallocate
        // allocate double what's needed, for future growth.
        newSlice := make([]byte, (n+1)*2)
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0:n]
    copy(slice[m:n], data)
    return slice
}
```

使用`AppendByte`场景如：
```
p := []byte{2, 3, 5}
p = AppendByte(p, 7, 11, 13)
// p == []byte{2, 3, 5, 7, 11, 13}
```

类似`AppendByte`方法的功能非常有用，因为它可以自由控制切片的增长，对用户比较透明。在不同的场景中，可能会需要不同的内存容量分配，或设置不同的容量。  

但实际应用场景中，自由控制切片的增长并不是必须的，所以Go提供了内置方法`append`，如
```
func append(s []T, x ...T) []T
```

该方法将一个元素追加到切片末尾，并在必要时会自动对切片扩容。

```
a := make([]int, 1)
// a == []int{0}
a = append(a, 1, 2, 3)
// a == []int{0, 1, 2, 3}
```

如果是追加一个切片到另一个切片末尾，则使用`...`即可，如
```
a := []string{"John", "Paul"}
b := []string{"George", "Ringo", "Pete"}
a = append(a, b...) // equivalent to "append(a, b[0], b[1], b[2])"
// a == []string{"John", "Paul", "George", "Ringo", "Pete"}
```

一个`nil`的切片实际上是长度为0的切片，所以我们可以把切片直接用于循环语句中，如
```
// Filter returns a new slice holding only
// the elements of s that satisfy fn()
func Filter(s []int, fn func(int) bool) []int {
    var p []int // == nil
    for _, v := range s {
        if fn(v) {
            p = append(p, v)
        }
    }
    return p
}
```

## 可能的陷阱
像前面提到的，切片的赋值并不会复制底层的数组，这个底层的数组一直在原来的内存块中直到没有被引用，最终`gc`回收，但有时一个大数组，实际上只有其中的一小段数据被使用，其他用不到，就造成了内存的浪费。  

例如，下面的方法，先把数据加载到内存中，从中找到连续的数字，并将找到的数字放入切片，最后返回。
```
var digitRegexp = regexp.MustCompile("[0-9]+")

func FindDigits(filename string) []byte {
    b, _ := ioutil.ReadFile(filename)
    return digitRegexp.Find(b)
}
```
当上例中的文件数据被加载到内存中，只要该方法返回的切片没有被`gc`回收，那么整个文件的数据就一直在内存中。  

为了避免这个问题，我们应该把要返回的切片数据赋值给一个新的切片，再把新的切片返回。如
```
func CopyDigits(filename string) []byte {
    b, _ := ioutil.ReadFile(filename)
    b = digitRegexp.Find(b)
    c := make([]byte, len(b))
    copy(c, b)
    return c
}
```
上例中更简单的实现是直接使用`append`方法，大家一起加油，多动手练习哈！

## 进一步阅读
[Effective Go](https://golang.org/doc/effective_go.html) 还有更多的关于[切片](https://golang.org/doc/effective_go.html#slices)和[数组](https://golang.org/doc/effective_go.html#arrays)的文章，同时，Go[语言规范](https://golang.org/doc/go_spec.html)中也有[切片的定义](https://golang.org/doc/go_spec.html)，切片和数组之间的[关联](https://golang.org/doc/go_spec.html#Length_and_capacity)、[辅助类](https://golang.org/doc/go_spec.html#Making_slices_maps_and_channels)、[方法](https://golang.org/doc/go_spec.html#Appending_and_copying_slices)等。


