package ae

import (
	"net/url"
	"time"

	"github.com/clavoie/erru"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/taskqueue"
)

// Task represents an AppEngine task
type Task struct {
	aeTask *taskqueue.Task
	queue  string
}

// ITask represents a wrapper over the AppEngine taskqueue domain
type ITask interface {
	// New creates a new Task value and returns it
	New(queue TaskQueue, path string, params url.Values) (*Task, error)
	// Queue adds a Task to the AppEngine queue
	Queue(*Task) error
	// SetDelay sets the Delay value of the task
	SetDelay(time.Duration, *Task)
	// SetEta sets the ETA value of the task
	//
	// IDK maybe this should be a field on the struct
	SetEta(time.Time, *Task)
}

// taskImpl is an implementation of ITask
type taskImpl struct {
	context Context
}

// newTask returns a new instance of ITask
func newTask(context Context) ITask {
	return &taskImpl{
		context: context,
	}
}

func (t *taskImpl) New(queue TaskQueue, path string, params url.Values) (*Task, error) {
	queueName, hasQueue := taskQueueToName[queue]

	if hasQueue == false {
		return nil, erru.Errorf("ae:task: unknown task queue: %v", queue)
	}

	return &Task{
		aeTask: taskqueue.NewPOSTTask(path, params),
		queue:  queueName,
	}, nil
}

func (t *taskImpl) Queue(task *Task) error {
	context := t.context.Context()
	hostName, err := appengine.ModuleHostname(context, "cron", "", "")

	if err != nil {
		log.Errorf(context, "ae:task:Queue cannot get cron hostname: %v", err)
		return err
	}

	task.aeTask.Header.Add("Host", hostName)
	_, err = taskqueue.Add(context, task.aeTask, task.queue)
	return err
}

func (t *taskImpl) SetDelay(delay time.Duration, task *Task) {
	task.aeTask.Delay = delay
}

func (t *taskImpl) SetEta(eta time.Time, task *Task) {
	task.aeTask.ETA = eta
}
