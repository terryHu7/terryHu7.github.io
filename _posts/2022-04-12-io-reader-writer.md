---
layout: post
title: "io.Reader & io.Writer"
subtitle: ''
author: "Terry"
header-style: text
tags:
  - Golang
---


```go
package main
import (
    "fmt"
)
func main() {
    var input string
    _, err := fmt.Scanln(&input)
    if err == nil {
        fmt.Println(input)
    }
}
```
  

```go
package main
import (
    "bufio"
    "fmt"
    "os"
)
func main() {
    r := bufio.NewReader(os.Stdin)
    input, _ := r.ReadString('\n')
    fmt.Println(input)
}
```
  

  
```go
package main
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)
func main() {
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        temp, _ := strconv.Atoi(strings.Split(input.Text()[1:len(input.Text())-1], ",")[0])
        fmt.Println(temp)
    }
}
```