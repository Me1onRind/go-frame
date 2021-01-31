package deamon

import (
	"context"
	"go-frame/global"
	"go-frame/internal/core/custom_ctx"
	"sync"
	"time"

	"github.com/coreos/etcd/clientv3/concurrency"

	"go.uber.org/zap"
)

func Campaign(ctx *custom_ctx.Context, wg *sync.WaitGroup) (success <-chan struct{}) {
	ip := "127.0.0.1"
	c, cancel := context.WithCancel(ctx)
	defer cancel()

	if wg != nil {
		wg.Add(1)
	}

	var notify = make(chan struct{}, 1)

	go func() {
		if wg != nil {
			defer wg.Done()
		}

		for {
			select {
			case <-ctx.Done():
				return
			default:
			}
			session, err := concurrency.NewSession(global.EtcdClient, concurrency.WithTTL(5))
			if err != nil {
				ctx.Logger().Error("New etcd seesion failed", zap.Error(err))
				time.Sleep(time.Second * 2)
				continue
			}

			election := concurrency.NewElection(session, "go-frame")
			if err = election.Campaign(c, ip); err != nil {
				ctx.Logger().Error("Etcd campaign failed", zap.Error(err))
				session.Close()
				time.Sleep(3 * time.Second)
				continue
			}

			shouldBreak := false
			for !shouldBreak {
				select {
				case notify <- struct{}{}:
				case <-session.Done():
					ctx.Logger().Error("Campaign has done")
				case <-c.Done():
					ctxTmp, cancel := context.WithTimeout(context.Background(), time.Second*1)
					_ = election.Resign(ctxTmp)
					session.Close()
					cancel()
					return
				}
			}
		}

	}()

	return notify
}
