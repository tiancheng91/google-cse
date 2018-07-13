google cse for golang
=====================

cse.google.com api for golang

google 自定义搜索


## Example

```go
package main

import cse "github.com/tiancheng91/google-cse"

func main() {
    agent := cse.New("008063188944472181627:xqha3yefaee", "zh_CN")

    # 参数: 查询词, 页数, 页大小
    result = agent.Search("123456", 1, 20)
}
```