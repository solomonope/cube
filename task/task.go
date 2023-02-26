package task

import (
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
	"github.com/solomonope/cube/job"
)

type Task struct {
	ID            uuid.UUID
	Name          string
	State         job.State
	Image         string
	Memory        int
	Disk          int
	ExposedPorts  nat.PortSet
	PortBindings  map[string]string
	RestartPolicy string
	StartTime     time.Time
	FinishTime    time.Time
	Timestamp     time.Time
}

type TaskEvent struct {
	ID        uuid.UUID
	State     job.State
	Timestamp time.Time
	Task      Task
}

type Config struct {
	Name          string
	AttachStdin   bool
	AttachStdout  bool
	AttachStderr  bool
	Cmd           []string
	Image         string
	Memory        int64
	Disk          int64
	Env           []string
	RestartPolicy string
}

type Docker struct {
	Client *client.Client
	Config Config
}

type DockerResult struct {
	Error       error
	Action      string
	ContainerId string
	Result      string
}
