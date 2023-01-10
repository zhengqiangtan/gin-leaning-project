package poller

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type Poller struct {
	routineGroup *goroutineGroup // 并发控制
	workerNum    int             // 记录同时在运行的最大goroutine数

	sync.Mutex
	ready  chan struct{} // 某个goroutine已经准备好了
	metric *metric       // 统计当前在运行中的goroutine数量
}

func NewPoller(workerNum int) *Poller {
	return &Poller{
		routineGroup: newRoutineGroup(),
		workerNum:    workerNum,
		ready:        make(chan struct{}, 1),
		metric:       newMetric(),
	}
}

// 调度器
func (p *Poller) schedule() {
	p.Lock()
	defer p.Unlock()
	if int(p.metric.BusyWorkers()) >= p.workerNum {
		return
	}

	select {
	case p.ready <- struct{}{}: // 只要满足当前goroutine数量小于最大goroutine数量 那么就通知poll去调度goroutine执行任务
	default:
	}
}

func (p *Poller) Poll(ctx context.Context) error {
	for {
		// step01
		p.schedule() // 调度

		select {
		case <-p.ready: // goroutine准备好之后 这里就会有消息
		case <-ctx.Done():
			return nil
		}

	LOOP:
		for {
			select {
			case <-ctx.Done():
				break LOOP
			default:
				// step02
				task, err := p.fetch(ctx) // 获取任务
				if err != nil {
					log.Println("fetch task error:", err.Error())
					break
				}
				fmt.Println("获取任务：" + task) // 打印生成的task
				p.metric.IncBusyWorker()    // 当前正在运行的goroutine+1
				// step03
				p.routineGroup.Run(func() { // 执行任务
					if err := p.execute(ctx, task); err != nil {
						log.Println("execute task error:", err.Error())
					}
				})
				break LOOP
			}
		}
	}
}

func (p *Poller) fetch(ctx context.Context) (string, error) {
	time.Sleep(1000 * time.Millisecond)
	return "task" + strconv.Itoa(rand.Intn(100)), nil
}

func (p *Poller) execute(ctx context.Context, task string) error {
	defer func() {
		p.metric.DecBusyWorker() // 执行完成之后 goroutine数量-1
		p.schedule()             // 重新调度下一个goroutine去执行任务 这一步是必须的
	}()
	fmt.Println("执行task：", task)
	return nil
}
