# ACM题目输入问题

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

### 1.2 多个字符串输入

#### 1.2.1 xx

## 2.数字输入

### 2.1 单个数字输入

#### 2.1.1 使用fmt.Scan

```go

```

