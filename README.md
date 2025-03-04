# Quick Start

```shell
go get github.com/hyphennn/glambda@latest
```

灵感来源：字节跳动 code.byted.org/lang/gg

超简易版本的 lang/gg 包， 用于简化一些常用的泛型、闭包工具，并且不引入任何第三方依赖以及过多的性能消耗

此包在作者离职时仍未开源，也未找到一个类似的仓库，因此基于 lang/gg 的一些思想做了此包，但未使用其 Iter 的做法以及一些 Rust
化的思想，个人觉得没啥意义。

另：想做一些 Stream API，但调整了很久，还是发现意义不大，因此也没动了，能力有限。

随缘更新，如果运气好有人用并不幸发现有 bug，欢迎 issue/email，我会第一时间修复