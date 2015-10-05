JobQueue
========

JobQueue allows the scheduling of multiple concurrent FIFO job queues.

### How it works

A `Broker` maintains a map of active `JobQueue`s.

```
broker := jobqueue.NewBroker()
```

`Broker`s can be used to create multiple queues for performing work. They must have unique id's, and require a `jobqueue.Handler`.

```
jobId := broker.EnqueueCreate("queue-1", worker)
```

In order to be a worker, you must implement the `jobqueue.Handler` interface, which exposes one method, `JobHandler(*Job)`.

```
type MyWorker struct {
	// Some environment to do work in.
}

func (m *MyWorker) JobHandler(job *jobqueue.Job) {
	fmt.Println(job.UUID)
	time.Sleep(5 * time.Second)
}
```

Create as many `JobQueue`s and `JobHandler`s as you'd like.

```
for i := 0; i < 100; i++ {
	go func() {
		broker.EnqueueCreate("queue-1", worker1)
	}()
	go func() {
		broker.EnqueueCreate("queue-2", worker2)
	}()
}
```

### By Anology

- `Broker`: A grocery store managing multiple cash registers.
- `JobQueue`: An individual cash register. Capable of handling one customer/`Job` at a time. Multiple registers can be handling customers at the same time.
- `Job`: The "event" or "moment in time" a customer/`Job` is being handled.
- `Handler`: The groceries that need checked out, or the work that needs done.
