###  基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。



## ANSWER

```go
package main

import (
	"context"

	"golang.org/x/sync/errgroup"
)

func Start() error {
	return nil
}

func Shutdown()error {
    return nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background()) // v0
	// v1

	g, ctx := errgroup.WithContext(ctx)  
     // v2
	
    
	g.Go(Start)
    g.Go(Shutdown)

	if err := g.Wait(); err != nil {
        cancel()
	}
}

```

