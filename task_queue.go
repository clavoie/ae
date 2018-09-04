package ae

// TaskQueue represents an AppEngine queue that processes tasks
type TaskQueue byte

const (
	MatchmakingTaskQueue TaskQueue = iota
	MatchTurnTaskQueue
)

var taskQueueToName = map[TaskQueue]string{
	MatchmakingTaskQueue: "matchmaking",
	MatchTurnTaskQueue:   "match-turn",
}
