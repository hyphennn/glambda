Quick Start

```shell
go get github.com/hyphennn/glambda@latest
```

```go
package main

import (
	"fmt"

	"github.com/hyphennn/glambda/gstream"
)

func main() {
	s := []int{1, 1, 4, 5, 1, 4, 114514, 0, 0, 0}
	s2 := gstream.AsSliceStream(s).Filter(func(i int) bool {
		return i%2 == 0
	}).Convert(func(i int) int {
		return i + 1
	}).Sort(func(t1, t2 int) bool {
		return t1 < t2
	}).Collect()
	fmt.Println(s2) // [1 1 1 5 5 114515]
}
```

致力于构建一个 Go Stream API 以及一些常用的泛型、闭包工具，并且不引入任何第三方依赖以及过多的性能消耗

Source of inspiration: Bytedance code.byted.org.lang/gg

This package was not open source when the author resigned, and a similar repository was not found, so this package was
made based on some ideas from the Java Stream API and lang/gg

致力于构建一个 Go Stream API 以及一些常用的泛型、闭包工具，并且不引入任何第三方依赖以及过多的性能消耗

灵感来源：字节跳动 code.byted.org/lang/gg

此包在作者离职时仍未开源，也未找到一个类似的仓库，因此基于 Java Stream API 以及 lang/gg 的一些思想做了此包
