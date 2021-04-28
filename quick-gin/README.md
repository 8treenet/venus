# QuickGin
#### QuickGin是为Gin打造的HTTP/2 H2C版本的扩展引擎。

###### HTTP/2（超文本传输协议第2版，最初命名为HTTP 2.0），简称为h2（基于TLS/1.2或以上版本的加密连接）或h2c（非加密连接），是HTTP协议的的第二个主要版本

## 安装
```sh
$ go get github.com/8treenet/venus
```

## 使用
```go
package main

import (
	quickgin "github.com/8treenet/venus/quick-gin"
	"github.com/gin-gonic/gin"
)

func main() {
	//创建QuickGin
	engine := quickgin.New()
	engine.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	//engine.Run()
	engine.RunH2C()//启动 
}
```


## 客户端请求
```go
package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/8treenet/freedom/infra/requests"
)

func main() {
	round := 30
	for i := 0; i < round; i++ {
		press()
	}
}

func press() {
	requestCount := 10000
	group := sync.WaitGroup{}
	begin := time.Now()

	for i := 0; i < requestCount; i++ {
		group.Add(1)
		go func() {
			result, _ := requests.NewH2CRequest("http://127.0.0.1:8080/ping").Get().ToString()
			if result != "pong" {
				panic("ping error")
			}
			group.Done()
		}()
	}

	group.Wait()
	ms := time.Now().Sub(begin).Milliseconds()
	fmt.Printf("Number of requests: %d, total time: %d ms. \n", requestCount, ms)
}

```

<img src="https://github.com/8treenet/venus/blob/master/quick-gin/example/client/client.jpg">