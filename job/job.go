package job

type State int

const (
	Pending State = iota
	Scheduled
	Completed
	Running
	Failed
)
