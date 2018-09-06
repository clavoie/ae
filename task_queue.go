package ae

import (
	"net/url"
	"time"

	"google.golang.org/appengine/taskqueue"
)

// Task represents an AppEngine task
type Task struct {
	aeTask *taskqueue.Task
	queue  string
}

// TaskQueue represents a wrapper over the AppEngine taskqueue domain
type TaskQueue interface {
	// NewTask creates a new POSTTask for a particular queue and returns it.
	NewTask(queue, path string, params url.Values) *Task

	// Queue adds a Task to the AppEngine queue
	Queue(hostname string, task *Task) error

	// SetDelay sets the Delay value of the task
	SetDelay(time.Duration, *Task)

	// SetEta sets the ETA value of the task
	SetEta(time.Time, *Task)
}

// taskQueueImpl is an implementation of TaskQueue
type taskQueueImpl struct {
	context Context
}

// NewTaskQueue returns a new instance of TaskQueue
func NewTaskQueue(context Context) TaskQueue {
	return &taskQueueImpl{
		context: context,
	}
}

func (t *taskQueueImpl) NewTask(queueName, path string, params url.Values) *Task {
	return &Task{
		aeTask: taskqueue.NewPOSTTask(path, params),
		queue:  queueName,
	}
}

func (t *taskQueueImpl) Queue(hostname string, task *Task) error {
	context := t.context.Context()

	if hostname != "" {
		task.aeTask.Header.Add("Host", hostname)
	}

	_, err = taskqueue.Add(context, task.aeTask, task.queue)
	return err
}

func (t *taskQueueImpl) SetDelay(delay time.Duration, task *Task) {
	task.aeTask.Delay = delay
}

func (t *taskQueueImpl) SetEta(eta time.Time, task *Task) {
	task.aeTask.ETA = eta
}
