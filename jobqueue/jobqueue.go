package jobqueue

import (
	"sync"
	"time"

	"code.google.com/p/go-uuid/uuid"
)

// JobQueue represents a scheduler that
// allows a single job to run at a time.
type JobQueue struct {
	ID     string
	Config interface{}
	done   chan int
	broker *Broker
	wg     sync.WaitGroup
	clear  sync.Once
}

// Wait queues up a job and halts
// until the queue is empty.
func (j *JobQueue) Wait() {
	j.wg.Add(1)
	// Delete queue if not in use.
	j.clear.Do(func() {
		go func() {
			j.wg.Wait()
			if job := j.broker.Get(j.ID); job != nil {
				j.broker.Lock()
				defer j.broker.Unlock()
				close(job.done)
				delete(j.broker.JobQueues, j.ID)
			}
		}()
	})
	j.done <- 1
}

// Done flushes the current queue of it's job.
func (j *JobQueue) Done() {
	<-j.done
	j.wg.Done()
}

// Job represents a scheduled event.
type Job struct {
	UUID      string
	StartTime time.Time
	QueueId   string
}

func NewJob(qid, jid string) *Job {
	return &Job{
		UUID:      jid,
		StartTime: time.Now(),
		QueueId:   qid,
	}
}

type JobHandlerFunc func(*Job)

// Handler represents the interface
// a job's worker must implement.
type Handler interface {
	JobHandler(*Job)
}

// Broker maintains a list of active job queues.
type Broker struct {
	sync.RWMutex
	JobQueues map[string]*JobQueue
}

func NewBroker() *Broker {
	return &Broker{
		JobQueues: make(map[string]*JobQueue),
	}
}

func (b *Broker) NewJobQueue(id string) *JobQueue {
	var wg sync.WaitGroup
	var once sync.Once
	return &JobQueue{
		ID:     id,
		done:   make(chan int, 1),
		broker: b,
		wg:     wg,
		clear:  once,
	}
}

// Enqueue adds a new Handler to a JobQueue.
func (b *Broker) Enqueue(jq *JobQueue, handler Handler) string {
	jid := uuid.New()
	go func() {
		jq.Wait()
		defer jq.Done()
		handler.JobHandler(NewJob(jq.ID, jid))
	}()
	return jid
}

// EnqueueCreate adds a new Handler to a JobQueue
// while creating a new JobQueue if it dones't exist.
func (b *Broker) EnqueueCreate(id string, handler Handler) string {
	active := b.Get(id)
	if active == nil {
		active = b.NewJobQueue(id)
		b.Set(active)
	}
	return b.Enqueue(active, handler)
}

// Get returns a JobQueue by id.
func (b *Broker) Get(pid string) *JobQueue {
	b.RLock()
	defer b.RUnlock()
	return b.JobQueues[pid]
}

// Set adds a JobQueue to the Broker.
func (b *Broker) Set(job *JobQueue) {
	b.Lock()
	defer b.Unlock()
	b.JobQueues[job.ID] = job
}
