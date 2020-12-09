###  基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。



## ANSWER

```go
package main

import (
	"context"

	"golang.org/x/sync/errgroup"
)

func f() error {
	return nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	go g.Go(f)

	if err := g.Wait(); err != nil {
		
	}
}

```

