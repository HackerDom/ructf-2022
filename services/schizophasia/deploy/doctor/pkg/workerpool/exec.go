package workerpool

import (
	"context"
	"fmt"
	"github.com/usernamedt/doctor-service/pkg/logging"
	"github.com/usernamedt/doctor-service/pkg/setting"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var Pool *WorkerPool

func worker(ctx context.Context, wg *sync.WaitGroup, jobs <-chan Job, results chan<- Result) {
	defer wg.Done()
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}
			results <- job.execute(ctx)
		case <-ctx.Done():
			fmt.Printf("cancelled worker. Error detail: %v\n", ctx.Err())
			results <- Result{
				Err: ctx.Err(),
			}
			return
		}
	}
}

type WorkerPool struct {
	workersCount       int
	receivedJobsCount  uint64
	processedJobsCount uint64
	jobs               chan Job
	results            chan Result
	workerDone         chan struct{}
	queueStatDone      chan struct{}
	rpmStatDone        chan struct{}
}

func New(wcount int) *WorkerPool {
	return &WorkerPool{
		workersCount:  wcount,
		jobs:          make(chan Job, wcount),
		results:       make(chan Result, wcount),
		workerDone:    make(chan struct{}),
		queueStatDone: make(chan struct{}),
		rpmStatDone:   make(chan struct{}),
	}
}

func (wp *WorkerPool) Run(ctx context.Context) {
	var wg sync.WaitGroup

	wp.initQueueStatWriter()
	wp.initRpmStatWriter()
	for i := 0; i < wp.workersCount; i++ {
		wg.Add(1)
		// fan out worker goroutines
		//reading from jobs channel and
		//pushing calcs into results channel
		go worker(ctx, &wg, wp.jobs, wp.results)
	}

	wg.Wait()
	close(wp.workerDone)
	close(wp.queueStatDone)
	close(wp.rpmStatDone)
	close(wp.results)
}

func (wp *WorkerPool) NextResult() (Result, bool, bool) {
	select {
	case r, ok := <-Pool.results:
		atomic.AddUint64(&wp.processedJobsCount, 1)
		return r, ok, false

	case <-Pool.workerDone:
		return Result{}, false, true
	}
}

func (wp *WorkerPool) GenerateJobId() string {
	id := time.Now().UnixNano()
	return strconv.FormatInt(id, 10)
}

func (wp *WorkerPool) AddJob(job Job) {
	atomic.AddUint64(&wp.receivedJobsCount, 1)
	wp.jobs <- job
}

func (wp *WorkerPool) initQueueStatWriter() {
	statInterval := setting.AppSetting.QueueStatInterval
	if statInterval == 0 {
		logging.Infof("WorkerPool: not enabling queue statistics writer")
		return
	}

	logging.Infof("WorkerPool: init queue statistics writer (each %d seconds)", statInterval)
	timer := time.NewTicker(statInterval * time.Second)
	go func() {
		for {
			select {
			case <-timer.C:
				queueSize := wp.receivedJobsCount - wp.processedJobsCount
				if queueSize == 0 {
					continue
				}
				logging.Infof("WorkerPool: queue size is %d, ", queueSize)
			case <-wp.queueStatDone:
				logging.Info("WorkerPool: stopping queue statistics writer")
				timer.Stop()
				return
			}
		}
	}()
}

func (wp *WorkerPool) initRpmStatWriter() {
	statInterval := setting.AppSetting.RequestStatInterval
	if statInterval == 0 {
		logging.Infof("WorkerPool: not enabling request statistics writer")
		return
	}

	logging.Infof("WorkerPool: init request statistics writer (each %d seconds)", statInterval)
	timer := time.NewTicker(statInterval * time.Second)
	prevProcessed := uint64(0)
	prevReceived := uint64(0)
	go func() {
		for {
			select {
			case <-timer.C:
				currProcessed := wp.processedJobsCount
				currReceived := wp.receivedJobsCount
				logging.Infof("WorkerPool: processed %d jobs, received %d jobs / %d seconds",
					currProcessed-prevProcessed, currReceived-prevReceived, statInterval)
				prevProcessed = currProcessed
				prevReceived = currReceived
			case <-wp.rpmStatDone:
				logging.Info("WorkerPool: stopping request statistics writer")
				timer.Stop()
				return
			}
		}
	}()
}

func Setup() {
	Pool = New(setting.AppSetting.WorkersCount)
	go Pool.Run(context.Background())

	go func() {
		for {
			_, ok, finished := Pool.NextResult()
			if finished {
				logging.Info("Worker pool has finished working")
				return
			}

			if !ok {
				continue
			}
		}
	}()
}
