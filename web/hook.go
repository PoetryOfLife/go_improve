package web

import (
	"context"
	"errors"
	"sync"
	"time"
)

type Hook func(ctx context.Context) error

func BuildCloseServerHook(servers ...Server) Hook {
	return func(ctx context.Context) error {
		wg := sync.WaitGroup{}
		doneCh := make(chan struct{})
		wg.Add(len(servers))
		for _, s := range servers {
			go func(svr Server) {
				//err := svr.shutdown(ctx)
				//if err != nil {
				//	fmt.Println(err)
				//}
				time.Sleep(time.Second)
				wg.Done()
			}(s)
		}
		go func() {
			wg.Done()
			doneCh <- struct{}{}
		}()
		select {
		case <-ctx.Done():
			return errors.New("")
		case <-doneCh:
			return nil
		}
	}
}
