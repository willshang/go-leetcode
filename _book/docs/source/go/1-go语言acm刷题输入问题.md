# ACM题目输入问题

- 一般校招技术会有笔试，通常在nowcoder上，而题目通常是采用ACM模式，掌握ACM模式的输入很重要

| No.  | 描述                         |      |
| ---- | ---------------------------- | ---- |
| 01   |                              |      |
| 02   |                              |      |
| 03   | 第一行输入N，第二行输入N个数 | 完成 |

## 0.处理多个case

### 0.1 牛客网

- 对于传统ACM的OJ模式题目，你的程序需要stdin（标准输入）读取输入，然后stdout（标准输出）来打印结果

- https://www.nowcoder.com/discuss/276

```go
package main
import (
    "fmt"
)
func main() {
  a:=0
  b:=0
  for {
        n, _ := fmt.Scan(&a,&b)
        if n == 0 {
                break
        } else {
                fmt.Printf("%d\n",a+b)
        }
  }
}
```

## 1.字符串输入

### 1.1 单个字符串输入

- 适用于单个字符串输入

#### 1.1.1 使用bufio.NewReader(os.Stdin)+ReadLine()

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	str := string(data)
	fmt.Println(str)
}
```

#### 1.1.2 使用bufio.NewReader(os.Stdin)+ReadBytes()

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadBytes('\n')
	str := string(data)
	fmt.Println(str)
	return
}
```

#### 1.1.3 使用bufio.NewReader(os.Stdin)+ReadString()

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	str := string(data)
	fmt.Println(str)
	return
}
```

#### 1.1.4 使用bufio.NewScanner(os.Stdin)+Scan()+Text()

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	str := scan.Text()
	fmt.Println(str)
}
```

### 1.2 第一行输入2个字符串a和b

```go
package main

import (
	"fmt"
)

func main() {
	var a, b string
	for {
		n, _ := fmt.Scanf("%s %s", &a, &b)
		if n == 0 {
			break
		}
		// 下面是处理逻辑
	}
}
```

#### 1.2.1 xx

## 2.数字输入

### 2.1 单个数字输入

#### 2.1.1 使用fmt.Scan

```go

```

### 输入N，然后输入N个数

- 第一行输入N， 第二行是N个数

```go
package main

import (
	"fmt"
)

func main() {
	var n, m int
	for {
		a, _ := fmt.Scan(&n)
		if a == 0 {
			break
		}
		nums := make([]int, 0)
		for i := 0; i < n; i++ {
			_, _ = fmt.Scan(&m)
			nums = append(nums, m)
		}
        // 下面是处理逻辑
	}
}
```

