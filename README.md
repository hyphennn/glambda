Quick Start

```shell
go get github.com/hyphennn/glamda@latest
```

```go
package main

import (
	"fmt"

	"github.com/hyphennn/glamda/lslice"
)

func main() {
	s := []int{1, 1, 4, 5, 1, 4}
	ss := lslice.Map(s, func(f int) int {
		return f + 1
	})
	fmt.Println(ss) //[2 2 5 6 2 5]
}
```

Dedicated to providing a go lamda expression library that does not introduce any additional dependencies and only incurs minimal additional performance costs



Inspiration source: ByteDance code.byted.org/lang/gg



This package was not yet open source at the time of the author's resignation, and also cannot find a similar repository, so some ideas from lang/gg were used to create this package. If lang/gg plans to open source it in the future, please feel free to contact me so that ownership of this repository can be transferred

致力于提供一个不引入任何额外依赖以及只带来极小额外性能成本的 go lamda 表达式库，涵盖面向 go 的一些 lamda 常用表达式

灵感来源：字节跳动 code.byted.org/lang/gg

此包在作者离职时仍未开源，也未找到一个类似的仓库，因此基于 lang/gg 的一些思想做了此包，若后续 lang/gg 计划开源，欢迎联系我，可以转移此仓库的所有权

注：所有代码均为手写，个人理解不存在侵权问题，如果 lang/gg 愿意开源自然更好
